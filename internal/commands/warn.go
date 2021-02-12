package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/syncship/moby-dick/data/models"
	"github.com/syncship/moby-dick/pkg/router"
)

func init() {
	Router.AddCommand(router.Command{
		Name: "warn",
		Run: func(s *discordgo.Session, m *discordgo.MessageCreate, a router.Arguments) {
			users := m.Mentions

			if len(users) < 1 {
				s.ChannelMessageSend(m.ChannelID, "Você não mencionou nenhum usuário.")
				return
			}

			for user := range users {
				var userModel models.UserModel

				if err := db.Conn.One("ID", user, &userModel); err != nil {
					userModel = models.UserModel{
						ID:       user,
						Warnings: 1,
					}

					db.Conn.Save(&userModel)
					continue
				}

				userModel.Warnings++

				if err := db.Conn.Save(&userModel); err != nil {
					fmt.Println(err.Error())
				}
			}

			s.ChannelMessageSend(m.ChannelID, "Usuário(s) receberam um warning.")
		},
	})
}
