package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/syncship/moby-dick/internal/commands"
)

func main() {
	client, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Open()
	if err != nil {
		log.Fatal("Error opening the Discord connection: ", err)
	}

	commands.Router.SetPrefix(os.Getenv("DISCORD_PREFIX"))

	client.AddHandler(commands.Router.OnMessageCreateHandler)

	client.Identify.Intents = discordgo.IntentsAll

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	client.Close()
}
