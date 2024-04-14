package bot

import (
	tele "gopkg.in/telebot.v3"
	"tgBot/entities"
)

func (b *Bot) pickItems(c tele.Context, courses []entities.Course) error {
	var (
		rows []tele.Row
	)

	menu := &tele.ReplyMarkup{}
	for _, course := range courses {
		rows = append(rows, menu.Row(*b.Layout.Button(c, "subject", struct {
			Text string
			Data uint
			URL  string
		}{
			Text: course.Subject,
			Data: course.Id,
			URL:  course.Link,
		})))
	}

	menu.Inline(rows...)
	return c.Send(b.Layout.Text(c, "pickItems"), menu)
}
