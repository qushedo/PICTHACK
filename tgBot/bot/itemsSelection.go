package bot

import "C"
import (
	tele "gopkg.in/telebot.v3"
	"tgBot/states"
)

func (b *Bot) selectUniversity(c tele.Context) error {
	b.inputStates.Mx.RLock()
	b.inputStates.Map[c.Sender().ID] = states.WaitingForUniversity
	b.inputStates.Mx.RUnlock()
	return c.Send(b.Layout.Text(c, "university"))
}

func (b *Bot) chooseFaculty(c tele.Context) error {
	b.inputStates.Mx.RLock()
	b.inputStates.Map[c.Sender().ID] = states.WaitingForFaculty
	b.inputStates.Mx.RUnlock()
	return c.Send(b.Layout.Text(c, "faculty"))
}
