package bot

import (
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
	"gopkg.in/telebot.v3/middleware"
	"gorm.io/gorm"
	"tgBot/database"
	"tgBot/entities"
	"tgBot/states"
)

type Bot struct {
	*tele.Bot
	*layout.Layout
	db          *gorm.DB
	inputStates *states.InputMap
}

func New(path string) (*Bot, error) {
	lt, err := layout.New(path)
	if err != nil {
		return nil, err
	}

	b, err := tele.NewBot(lt.Settings())
	if err != nil {
		return nil, err
	}

	if cmds := lt.Commands(); cmds != nil {
		if err = b.SetCommands(cmds); err != nil {
			return nil, err
		}
	}

	db := database.Connect()

	bot := &Bot{
		Bot:         b,
		Layout:      lt,
		db:          db,
		inputStates: states.MakeInputMap(),
	}
	return bot, nil
}

func (b *Bot) Start() {
	b.Use(middleware.Logger())
	b.Use(middleware.AutoRespond())
	b.Use(b.Authorized)
	b.Handle(b.Layout.Callback("english"), b.OnLocalisation)
	b.Handle(b.Layout.Callback("russian"), b.OnLocalisation)
	b.Use(b.Localisation)

	b.Use(b.Layout.Middleware("ru", func(r tele.Recipient) string {
		var user entities.User
		b.db.Where("id = ?", r.Recipient()).First(&user)
		return user.Localisation
	}))

	b.Handle(tele.OnText, b.OnText)
	b.Use(b.ItemsMiddle)
	b.Handle("/start", b.onStart)

	b.Bot.Start()
}
