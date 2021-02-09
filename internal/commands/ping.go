package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/syncship/moby-dick/pkg/router"
	"github.com/syncship/moby-dick/pkg/typeguard"
)

func init() {
	Router.AddCommand(router.Command{
		Name: "ping",
		Args: router.Arguments{
			"times": {
				To:       typeguard.WantInt(),
				Required: false,
				Output: typeguard.Output{
					Default: 1,
				},
			},
		},
		Run: func(s *discordgo.Session, m *discordgo.MessageCreate, a router.Arguments) {
			times, err := a["times"].Output.ToInt()
			if err != nil {
				log.Println(err.Error())
			}

			for i := 0; i < times; i++ {
				s.ChannelMessageSend(m.ChannelID, "p-pong :flushed:")
			}
		},
		Permissions: []int{discordgo.PermissionSendMessages},
	})
}
