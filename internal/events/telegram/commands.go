package telegram

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"html"
	"log/slog"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"tweets-tg-bot/internal/clients/telegram"
	"tweets-tg-bot/internal/commands"
	"tweets-tg-bot/internal/downloader"
	"tweets-tg-bot/internal/events/telegram/tgTypes"
)

var AllCommands = []commands.Cmd{
	commands.RndCmd, commands.HelpCmd, commands.StartCmd, commands.StatsCmd, commands.LeaveChat, commands.ChatInfo,
}

var ErrorUnknownCommand = errors.New("unknown command")
var ErrApiResponse = errors.New("api error")

func isChatIdInTestGroup(chatId int, userId int) bool {
	testGroupIds := []int{-4020168327, -1001441929255}
	users := []int{114927545}

	for _, groupId := range testGroupIds {
		if chatId == groupId {
			return true
		}
	}

	for _, uId := range users {
		if userId == uId {
			return true
		}
	}
	return false
}

func (p *Processor) checkPermissions(ctx context.Context, chatId int) error {
	chatInfo, err := p.tg.GetChat(ctx, chatId)
	if err != nil {
		return errors.Wrap(err, "failed to get chat info")
	}

	permissions := chatInfo.Result.Permissions

	if !permissions.CanSendMessages {
		return errors.New("not allowed to send messages")
	}

	if !permissions.CanSendPhotos {
		return errors.New("not allowed to send photos")
	}

	if !permissions.CanSendVideos {
		return errors.New("not allowed to send videos")
	}

	if !permissions.CanSendMediaMessages {
		return errors.New("not allowed to send media messages")
	}

	return nil
}

func (p *Processor) doCmd(ctx context.Context, text string, chatId int, chatname, username string, userId int) error {
	defer p.recoverPanic(text, chatId, username)

	text = strings.TrimSpace(text)

	cmd, parsed, err := p.parseCmd(text)
	if errors.Is(err, ErrorUnknownCommand) {
		return nil
	}
	if err != nil {
		return err
	}

	defer func() {
		p.users.Command(cmd, username)
	}()

	switch cmd {
	case commands.TweetCmd:
		fallthrough
	case commands.TikTokCmd:
		fallthrough
	case commands.InstagramCmd:
		slog.InfoContext(ctx, "got new command", "cmd", cmd, "command", text, slog.Group("chat", "id", chatId, "name", chatname), slog.Group("user", "id", userId, "name", username))
		return p.sendContentOrHandleError(ctx, chatId, cmd, parsed, username)
	case commands.StartCmd:
		return p.sendStart(chatId, username)
	case commands.HelpCmd:
		return p.sendHelp(chatId, username)
	case commands.RndCmd:
		return p.sendRandom(chatId, username)
	case commands.StatsCmd:
		return p.sendStats(chatId, userId)
	case commands.LeaveChat:
		return p.leaveChat(ctx, userId, text)
	case commands.ChatInfo:
		return p.chatInfo(ctx, text, chatId, userId)
	default:
		return nil
	}
}

func generateHeader(tweet tgTypes.TweetThread) string {
	result := ""

	action := "tweeted"
	if tweet.Source != "twitter" {
		action = "posted"
	}

	if tweet.UserName != "" || tweet.UserId != "" {
		result += fmt.Sprintf("<b>%s</b>(<i>%s</i>) %s:\n", tweet.UserName, tweet.UserId, action)
	}

	twTime := tweet.Time.Format("15:04 · 2 Jan 2006")

	result += fmt.Sprintf("%s", twTime)

	if tweet.Views != "" && tweet.Views != "0" {
		views, err := strconv.Atoi(tweet.Views)
		if err == nil {
			tweet.Views = shortNumber(views)
		}
		result += fmt.Sprintf(" %s Views\n", tweet.Views)
	}

	addedLine := false
	if tweet.Retweets != 0 {
		addedLine = true
		result += fmt.Sprintf(" %s Retweets", shortNumber(tweet.Retweets))
	}

	if tweet.Replies != 0 {
		addedLine = true
		result += fmt.Sprintf(" %s Replies", shortNumber(tweet.Replies))
	}

	if tweet.Quotes != 0 {
		addedLine = true
		result += fmt.Sprintf(" %s Quotes", shortNumber(tweet.Quotes))
	}

	if tweet.Likes != 0 {
		addedLine = true
		result += fmt.Sprintf(" %s Likes", shortNumber(tweet.Likes))
	}

	if addedLine {
		result += "\n"
	}

	result += "\n"

	if tweet.UserNote.Text != "" {
		result += fmt.Sprintf("<span class=\"tg-spoiler\"><b>%s:</b>\n<i>%s</i>\n\n</span>", tweet.UserNote.Title, tweet.UserNote.Text)
	}

	return result
}

func generateText(tweet tgTypes.TweetContent) string {
	result := ""

	result += fmt.Sprintf("%s \n", tweet.Text)

	return result
}

func shortNumber(n int) string {
	str := fmt.Sprintf("%d", n)
	if n >= 1000 {
		str = fmt.Sprintf("%.1fK", float64(n)/1000)
	}
	if n >= 10000 {
		str = fmt.Sprintf("%dK", n/1000)
	}
	return str
}

func timeTrack(ctx context.Context, start time.Time, name string) {
	elapsed := time.Since(start)
	slog.InfoContext(ctx, "time", "func", name, "elapsed", elapsed.String())
}

func (p *Processor) sendContentOrHandleError(ctx context.Context, chatId int, cmd commands.Cmd, cmdUrl commands.ParsedCmdUrl, username string) error {
	err := p.send(ctx, chatId, cmd, cmdUrl)
	var err2 error
	if errors.Is(err, telegram.ErrNoEnoughRightToSendPhoto) {
		err2 = p.tg.SendMessage(chatId, "<i>Sorry, the bot doesn't have enough right to send photo contained in the provided link. Please allow sending photos in the chat settings.</i>")
	}
	if errors.Is(err, telegram.ErrNoEnoughRightToSendVideo) {
		err2 = p.tg.SendMessage(chatId, "<i>Sorry, the bot doesn't have enough right to send video contained in the provided link. Please allow sending video in the chat settings.</i>")
	}
	if err2 != nil {
		p.sendErrorToAdmin(cmdUrl.StrippedUrl, chatId, username, err2)
	}
	if err != nil {
		p.sendErrorToAdmin(cmdUrl.StrippedUrl, chatId, username, err)
		return err
	}
	return nil
}

func (p *Processor) send(ctx context.Context, chatId int, cmd commands.Cmd, cmdUrl commands.ParsedCmdUrl) error {
	defer timeTrack(ctx, time.Now(), string(cmd))

	content, err := p.contentManager.GetContent(ctx, cmd, cmdUrl)
	if err != nil {
		return errors.Wrap(err, "contentManager.GetContent")
	}

	// simple stupid retry
	attempt := 0
	for {
		attempt++
		err = p.sendContentAsMessage(chatId, content)
		if err == nil || attempt > 3 {
			break
		}
		slog.ErrorContext(ctx, "sendContentAsMessage failed, retrying in 1 sec...", "err", err)
		time.Sleep(1 * time.Second)
	}

	if err != nil {
		return errors.Wrap(err, "sendContentAsMessage")
	}
	return nil
}

func (p *Processor) sendContentAsMessage(chatId int, tweet tgTypes.TweetThread) error {
	for i, tw := range tweet.Tweets {

		message := html.EscapeString(generateText(tw))

		if i == 0 {
			message = generateHeader(tweet) + message
		}

		messages, err := chunkString(message, 1024)
		if err != nil {
			return err
		}

		for _, m := range messages {
			if len(tweet.Tweets[i].Media.Videos) >= 2 || len(tweet.Tweets[i].Media.Photos) > 0 ||
				(len(tweet.Tweets[i].Media.Videos) == 1 && len(tweet.Tweets[i].Media.Photos) >= 1) {

				var mediasForEncoding []telegram.MediaForEncoding
				if len(tweet.Tweets[i].Media.Photos) > 0 {
					downloadedPhotos, err := downloader.Download(tweet.Tweets[i].Media.Photos)
					if err != nil {
						return errors.Wrap(err, "downloading photo files")
					}
					tweet.Tweets[i].Media.Photos = downloadedPhotos

					mediasForEncoding = append(mediasForEncoding, telegram.MediaForEncoding{
						Media:     tweet.Tweets[i].Media.Photos,
						MediaType: telegram.MediaTypePhoto,
					})
				}

				if len(tweet.Tweets[i].Media.Videos) > 0 {
					downloadedVideos, err := downloader.Download(tweet.Tweets[i].Media.Videos)
					if err != nil {
						return errors.Wrap(err, "failed to download videos")
					}
					tweet.Tweets[i].Media.Videos = downloadedVideos

					mediasForEncoding = append(mediasForEncoding, telegram.MediaForEncoding{
						Media:     tweet.Tweets[i].Media.Videos,
						MediaType: telegram.MediaTypeVideo,
					})
				}

				allMedia := append(tweet.Tweets[i].Media.Videos, tweet.Tweets[i].Media.Photos...)

				err = p.tg.SendMedia(chatId, m, mediasForEncoding, allMedia)
				if err != nil {
					return errors.Wrap(err, "SendMedia")
				}
				tweet.Tweets[i].Media.Videos = nil
				tweet.Tweets[i].Media.Photos = nil
				continue
			}

			if len(tweet.Tweets[i].Media.Videos) == 1 {
				downloadedVideos, err := downloader.Download(tweet.Tweets[i].Media.Videos)
				if err != nil {
					return errors.Wrap(err, "failed to download video")
				}
				tweet.Tweets[i].Media.Videos = downloadedVideos

				mediasForEncoding := []telegram.MediaForEncoding{
					{
						Media:     tweet.Tweets[i].Media.Videos,
						MediaType: telegram.MediaTypeVideo,
					},
				}

				//err = p.tg.SendVideo(chatId, m, video(tw))
				err = p.tg.SendMedia(chatId, m, mediasForEncoding, tweet.Tweets[i].Media.Videos)
				if err != nil {
					return errors.Wrap(err, "SendVideo error")
				}
				tweet.Tweets[i].Media.Videos = nil
				continue
			}

			if len(tweet.Tweets[i].Media.Photos) == 1 {
				err = p.tg.SendPhoto(chatId, m, photos(tw)[0].Url)
				if err != nil {
					return errors.Wrap(err, "SendPhoto")
				}
				tweet.Tweets[i].Media.Photos = nil
				continue
			}

			if len(tweet.Tweets[i].Media.Photos) >= 2 {
				mediaForEncoding := []telegram.MediaForEncoding{
					{
						Media:     photos(tw),
						MediaType: telegram.MediaTypePhoto,
					},
				}

				err = p.tg.SendPhotos(chatId, m, mediaForEncoding)
				if err != nil {
					return errors.Wrap(err, "SendPhotos")
				}
				tweet.Tweets[i].Media.Photos = nil
				continue
			}

			if err := p.tg.SendMessage(chatId, m); err != nil {
				return errors.Wrap(err, "SendMessage")
			}
			continue
		}
	}
	return nil
}

func chunkString(s string, chunkSize int) ([]string, error) {
	//regex, err := regexp.Compile(".{1,25}\\b|.{1,25}")
	//if err != nil {
	//	return nil, err
	//}
	//
	//chunks := regex.FindAllString(s, -1)
	var chunks []string
	runes := []rune(s)

	if len(runes) == 0 {
		return []string{s}, nil
	}

	for i := 0; i < len(runes); i += chunkSize {
		nn := i + chunkSize
		if nn > len(runes) {
			nn = len(runes)
		}
		chunks = append(chunks, string(runes[i:nn]))
	}
	return chunks, nil
}

func (p *Processor) sendRandom(chatId int, username string) error {
	n := rand.Intn(100)
	if err := p.tg.SendMessage(chatId, fmt.Sprintf("random %d", n)); err != nil {
		return err
	}

	return nil
}

func (p *Processor) chatInfo(ctx context.Context, text string, sendTo int, requestedBy int) error {
	isAdmin, err := p.users.IsAdmin(requestedBy)
	if err != nil {
		return errors.Wrap(err, "IsAdmin")
	}
	if !isAdmin {
		return nil
	}

	text = strings.TrimSpace(strings.Replace(text, string(commands.ChatInfo), "", 1))

	chatId, err := strconv.Atoi(text)
	if err != nil {
		return errors.Wrap(err, "failed to parse chat id")
	}

	chat, err := p.tg.GetChat(ctx, chatId)
	if err != nil {
		return errors.Wrap(err, "failed to get chat info")
	}

	chatJson, err := json.MarshalIndent(chat, "", "    ")
	if err != nil {
		return errors.Wrap(err, "failed to marshal chat info")
	}

	_ = p.tg.SendMessage(sendTo, fmt.Sprintf("chat info: \n <pre>%s</pre>", string(chatJson)))
	return nil
}

func (p *Processor) leaveChat(ctx context.Context, userId int, text string) error {
	slog.InfoContext(ctx, text)
	isAdmin, err := p.users.IsAdmin(userId)
	if err != nil {
		return errors.Wrap(err, "IsAdmin")
	}
	if !isAdmin {
		return nil
	}

	text = strings.TrimSpace(strings.Replace(text, string(commands.LeaveChat), "", 1))

	chatId, err := strconv.Atoi(text)
	if err != nil {
		return errors.Wrap(err, "failed to parse chat id")
	}

	err = p.tg.LeaveChat(ctx, chatId)
	if err != nil {
		return errors.Wrap(err, "failed to leave chat")
	}

	_ = p.tg.SendMessage(userId, fmt.Sprintf("bot left chat %d", chatId))

	return nil
}

func (p *Processor) sendStats(id int, userId int) error {
	isAdmin, err := p.users.IsAdmin(userId)
	if err != nil {
		return err
	}
	if !isAdmin {
		return nil
	}

	ctx := context.TODO()
	count, err := p.users.Count(ctx)
	if err != nil {
		return err
	}

	countShares, err := p.users.CountShare(ctx)
	if err != nil {
		return err
	}

	mau, dau, err := p.users.CountActiveUsers(ctx)
	if err != nil {
		return err
	}

	mpu, dpu, err := p.users.CountPassiveUsers(ctx)
	if err != nil {
		return err
	}

	comandsStat, err := p.users.CommandsStat(ctx)
	if err != nil {
		return err
	}

	message := fmt.Sprintf(
		`Users: %d 
Users who share tweets: %d

Sharing:
Monthly: %d 
Daily: %d 

Viewing:
Monthly: %d 
Daily: %d

`,
		count,
		countShares,
		mau,
		dau,
		mpu,
		dpu,
	)

	for k, v := range comandsStat {
		message += fmt.Sprintf("\n%s: %d", k, v)
	}

	if err := p.tg.SendMessage(id, message); err != nil {
		return err
	}

	return nil
}

func (p *Processor) parseCmd(text string) (commands.Cmd, commands.ParsedCmdUrl, error) {
	cmd, parsed, err := p.cmdParser.Parse(text)
	if err == nil {
		return cmd, parsed, err
	}

	for _, cmd := range AllCommands {
		if len(text) < len(cmd) {
			continue
		}

		cmdPart := text[:len(cmd)]

		if string(cmd) == cmdPart {
			return cmd, commands.ParsedCmdUrl{}, nil
		}
	}

	return "", commands.ParsedCmdUrl{}, ErrorUnknownCommand
}

func photos(tweet tgTypes.TweetContent) []tgTypes.MediaObject {
	return tweet.Media.Photos
}

func videos(tweet tgTypes.TweetContent) []tgTypes.MediaObject {
	return tweet.Media.Videos
}

func video(tweet tgTypes.TweetContent) tgTypes.MediaObject {
	if len(tweet.Media.Videos) > 0 {
		return tweet.Media.Videos[0]
	}
	return tgTypes.MediaObject{}
}

func (p *Processor) sendStart(chatId int, username string) error {
	return p.tg.SendMessage(chatId, msgHello)
}

func (p *Processor) sendHelp(chatId int, username string) error {
	return p.tg.SendMessage(chatId, msgHelp)
}
