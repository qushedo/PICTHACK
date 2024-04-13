package bot

import (
	tele "gopkg.in/telebot.v3"
	"tgBot/entities"
)

func (b *Bot) OnLocalisation(c tele.Context) error {
	var user entities.User
	err := b.db.Where("id = ?", c.Sender().ID).First(&user).Error
	if err == nil {
		user.Localisation = c.Callback().Data
		b.db.Save(&user)
		return c.Delete()
	}
	return err
}
