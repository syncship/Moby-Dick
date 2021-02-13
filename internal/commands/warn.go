package commands

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/syncship/moby-dick/data/models"
	"github.com/syncship/moby-dick/pkg/router"
)

func init() {
	Router.AddCommand(router.Command{
		Name: "warn",
		Run: func(s *discordgo.Session, m *discordgo.MessageCreate, a router.Arguments) {
			users := m.Mentions
			howMany := len(users)

			if len(users) < 1 {
				s.ChannelMessageSend(m.ChannelID, "Você não mencionou nenhum usuário.")
				return
			}

			for i := range users {
				user := users[i]
				userID, _ := strconv.Atoi(user.ID)

				var userModel models.UserModel

				if err := db.Conn.One("ID", userID, &userModel); err != nil {
					userModel = models.NewUserModel(userID)
				}

				userModel.Warnings++

				if err := db.Conn.Save(&userModel); err == nil {
					users = append(users[:i], users[i+1:]...)
				}
			}

			if len(users) > 0 {
				s.ChannelMessageSend(m.ChannelID,
					fmt.Sprintf("Não foi possível dar warning pra esses usuários: %v", users))

				return
			}

			if howMany > 1 {
				s.ChannelMessageSend(m.ChannelID, "Usuário(s) receberam um warning.")

				return
			}

			s.ChannelMessageSend(m.ChannelID, "Usuário recebeu um warning.")
		},
	})
}
