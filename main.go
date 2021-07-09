package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	//Create a new discord session using the provided bot token

	dg, err := discordgo.New("Bot " + os.Getenv(("DISCORD_TOEKN")))
	if err != nil {
		fmt.Println("Error creating new bot session", err)
		return
	}
	//Add a handler function to our bot
	dg.AddHandler(messageCreate)
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

//This function is called anytime a message is sent in any channel.

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	//Ignore all messages sent by our bot
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "</>ping" {
		s.ChannelMessageSend(m.ChannelID, "SHEEESHHHH")
	}
}
