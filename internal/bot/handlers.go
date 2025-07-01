package bot

import (
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
)

func (b *Bot) StartHandler() func(c telebot.Context) error {
	return func(c telebot.Context) error {
		chat := c.Chat()
		user := c.Sender()
		
		b.logger.WithFields(logrus.Fields{
			"user_id": user.ID,
			"username": user.Username,
		}).Info("User started the bot")


		if chat.Type == telebot.ChatPrivate {
			b.logger.WithFields(logrus.Fields{
				"source": "StartHandler",
				"user_id": user.ID,
				"username": user.Username,
				"type_chat": chat.Type,
			}).Infof("User (%d | %s) clicked /start in private chat wit bot", user.ID, user.Username)

			startMsg := "Оу, привіт, зіронько! 🌟 Хочеш створити гру для своїх найкращих подруг? Натискай кнопку нижче і вперед до пригод!"

			creatorID := fmt.Sprintf("%d", c.Sender().ID)
			deepLink := "https://t.me/bestie_game_bot?startgroup=" + creatorID

			menu := &telebot.ReplyMarkup{}
			btnDeepLink := menu.URL("➕ Створити гру", deepLink)
			btnHelp := menu.Data("❓ Help Me", "help_me")

			menu.Inline(
				menu.Row(btnDeepLink),
				menu.Row(btnHelp),
			)

			// b.Handle(&btnHelp, HelpMeHandler(bot))

			return c.Send(startMsg, menu)
		}

		payload := c.Message().Payload
		if payload == "" {
			return c.Send("Щось пішло не так. 😔 Спробуй створити гру ще раз через особисте повідомлення боту.")
		}

		creatorID, err := strconv.ParseInt(payload, 10, 64)
		if err != nil {
		  b.logger.Errorf("Не вдалося розпізнати ID користувача: %v", err)
			return c.Send("Помилка при запуску гри. Спробуй ще раз.")
		}
    
		b.logger.WithFields(logrus.Fields{
			"source": "StartHandler",
			"group": chat.Title,
			"group_id": chat.ID,
			"admin_id:": creatorID,
			"admin": user.Username,
		}).Info("The bot was added to the group via a button in a private chat with the bot")
		
		return c.Send("🎉 Гру створено! Додайте своїх подруг і вперед до веселощів!")
	}
}

func (b *Bot) HelpMeHandler() func (c telebot.Context) error {
	return func(c telebot.Context) error {
		helpText := `
			Привіт, зіронько! 🌟 Я бот для ігор з подругами на відстані. Ось мої команди:

/start - Почати бота і створити нову гру або доєднатися до існуючої
/help - Показати це повідомлення

В грі ти можеш:
- Відповідати на завдання (текст, фото, відео, голосові повідомлення)
- Пропустити завдання (максимум 3 рази)
- Отримувати сповіщення про активність друзів

Якщо потрібна допомога, натисни кнопку "Хелп мі" в меню!
		`
		return c.Send(helpText)
	}
}