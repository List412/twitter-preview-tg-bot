package graphql

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

func NewClient(host string) *Client {
	return &Client{host}
}

type Client struct {
	host string
}

func (c *Client) GetVideo(ctx context.Context, id string) (*ParsedPost, error) {

	resp, err := gqlRequest(ctx, id)
	if err != nil {
		return nil, err
	}

	log.Printf("GetPost done %s", id)

	return resp, nil
}

func gqlRequest(ctx context.Context, code string) (*ParsedPost, error) {
	url := "https://www.instagram.com/api/graphql"
	method := "POST"

	payload := strings.NewReader("av=0&__d=www&__user=0&__a=1&__req=3&__hs=19702.HYP%3Ainstagram_web_pkg.2.1..0.0&dpr=2&__ccg=UNKNOWN&__rev=1010329543&__s=ytibog%3Adaumdy%3A3lk1qh&__hsi=7311321296052053265&__dyn=7xeUjG1mxu1syUbFp60DU98nwgU29zEdEc8co2qwJw5ux609vCwjE1xoswIwuo2awlU-cw5Mx62G3i1ywOwv89k2C1Fwc60AEC7U2czXwae4UaEW2G1NwwwNwKwHw8Xxm16wUwtEvw4JwJwSyES1Twoob82ZwrUdUbGwmk1xwmo6O1FwlE6OFA6fxy4Ujw&__csr=gjhXlMxdaWXDamZbmF8ytrmBqGHXRBx2vyV4iQpGvKbCGiU-eLFoSHzqDyqzaKRKFm-ahuiqimXl7ypGjx2OeuqhuBDhHDyWDAgCGGdzEOciihElzUargG4FU01cGpE2W805eiw1S606EE25G44md40dbw1aCrc1txC0uG3VzE8Q2q0nK089w0adG&__comet_req=7&lsd=AVpQxgXKVKs&jazoest=21006&__spin_r=1010329543&__spin_b=trunk&__spin_t=1702299643&fb_api_caller_class=RelayModern&fb_api_req_friendly_name=PolarisPostActionLoadPostQueryQuery&variables=%7B%22shortcode%22%3A%22" + code + "%22%2C%22fetch_comment_count%22%3A40%2C%22fetch_related_profile_media_count%22%3A3%2C%22parent_comment_count%22%3A24%2C%22child_comment_count%22%3A3%2C%22fetch_like_count%22%3A10%2C%22fetch_tagged_user_count%22%3Anull%2C%22fetch_preview_comment_count%22%3A2%2C%22has_threaded_comments%22%3Atrue%2C%22hoisted_comment_id%22%3Anull%2C%22hoisted_reply_id%22%3Anull%7D&server_timestamps=true&doc_id=10015901848480474")

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, method, url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("authority", "www.instagram.com")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "ru")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cookie", "csrftoken=4uEFj2FLgdivDIhwhQvWmf; mid=ZXbNoAAEAAEyW7_iU-3KhQ1e_P8O; ig_did=8447FB01-50CB-40B4-BF71-96FD60599770; datr=sAF3ZbqSNWefZVtcruJTpACc; csrftoken=FxH3VTv4mRviA8kqGgpU2B")
	req.Header.Add("dpr", "2")
	req.Header.Add("origin", "https://www.instagram.com")
	req.Header.Add("referer", "https://www.instagram.com/p/CzBjgFiISfF/")
	req.Header.Add("sec-ch-prefers-color-scheme", "dark")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"119\", \"Chromium\";v=\"119\", \"Not?A_Brand\";v=\"24\"")
	req.Header.Add("sec-ch-ua-full-version-list", "\"Google Chrome\";v=\"119.0.6045.199\", \"Chromium\";v=\"119.0.6045.199\", \"Not?A_Brand\";v=\"24.0.0.0\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-model", "\"\"")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-ch-ua-platform-version", "\"13.6.1\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")
	req.Header.Add("viewport-width", "853")
	req.Header.Add("x-asbd-id", "129477")
	req.Header.Add("x-csrftoken", "4uEFj2FLgdivDIhwhQvWmf")
	req.Header.Add("x-fb-friendly-name", "PolarisPostActionLoadPostQueryQuery")
	req.Header.Add("x-fb-lsd", "AVpQxgXKVKs")
	req.Header.Add("x-ig-app-id", "936619743392459")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	gqlResp := new(ParsedPost)
	if err := json.Unmarshal(body, gqlResp); err != nil {
		return nil, err
	}

	if gqlResp.Data.XdtShortcodeMedia.ID == "" {
		return nil, errors.New("got empty result")
	}

	return gqlResp, nil
}
