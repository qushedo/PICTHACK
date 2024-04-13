package bot

import (
	tele "gopkg.in/telebot.v3"
	"tgBot/entities"
)

func (b *Bot) Authorized(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		var (
			user entities.User
		)
		if b.db.Where("id = ?", c.Sender().ID).First(&user).Error != nil {
			user.Id = c.Sender().ID
			user.FirstName = c.Sender().FirstName
			user.UserName = c.Sender().Username
			b.db.Save(&user)
		}
		return next(c)
	}
}

func (b *Bot) Localisation(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		var user entities.User
		if b.db.Where("id = ?", c.Sender().ID).First(&user).Error == nil {
			if user.Localisation == "" {
				return c.Send(
					"<b>🌍 Выберите предпочтительный язык</b>\n"+
						"<b>🌍 Choose your preferred language</b>",
					b.Layout.MarkupLocale("en", "pickLanguage"),
				)
			}

			//b.Use(b.Layout.Middleware("en", func(r tele.Recipient) string {
			//	return user.Localisation
			//}))
		}

		return next(c)
	}
}
