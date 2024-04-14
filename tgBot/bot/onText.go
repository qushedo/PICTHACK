package bot

import (
	"errors"
	tele "gopkg.in/telebot.v3"
	"tgBot/entities"
	"tgBot/states"
)

func (b *Bot) OnText(c tele.Context) error {
	var user entities.User

	userId := c.Sender().ID
	if b.db.Where("id = ?", userId).First(&user).Error != nil {
		return errors.New("пользователь не найден")
	}
	state := b.inputStates.Map[userId]
	msg := c.Message().Text

	switch state {
	case states.WaitingForUniversity:
		var course entities.Course
		if b.db.Where("university = ?", msg).First(&course).Error == nil {
			user.University = course.University
			b.db.Save(&user)

			return b.chooseFaculty(c)

		} else {
			return b.selectUniversity(c)
		}
	case states.WaitingForFaculty:
		var courses []entities.Course
		b.db.Where("megafaculty = ? OR megafaculty_id = ? OR faculty = ? OR faculty_id = ?", msg, msg, msg, msg).Find(&courses)
		if len(courses) > 0 {
			user.Faculty = courses[0].Faculty
			b.db.Save(&user)

			return b.pickItems(c, courses)
		} else {
			return b.chooseFaculty(c)
		}

	default:
		return c.Send(b.Layout.Text(c, "unknownCmd"))
	}
}
