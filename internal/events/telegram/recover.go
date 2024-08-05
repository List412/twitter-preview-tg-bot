package telegram

import "fmt"

func (p *Processor) recoverPanic(text string, chatId int, username string) {
	if r := recover(); r != nil {
		p.sendErrorToAdmin(text, chatId, username, r)
	}
}

func (p *Processor) sendErrorToAdmin(text string, chatId int, username string, err interface{}) {
	adminMessage := fmt.Sprintf("error occurred with message: %s \nfrom user: %s in chat: %d\n error: %v", text, username, chatId, err)

	_ = p.tg.SendMessage(p.users.GetAdminId(), adminMessage)
}
