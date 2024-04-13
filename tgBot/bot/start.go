package bot

import tele "gopkg.in/telebot.v3"

func (b *Bot) onStart(c tele.Context) error {
	return c.Send(
		b.Layout.Text(c, "start"),
		b.Layout.Markup(c, "replyOpenMenu", b.Layout.Text(c, "openMenu")),
	)
}
