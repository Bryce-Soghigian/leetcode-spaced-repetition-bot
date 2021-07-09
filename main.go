package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + "")
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Partitioned into sublists you can query
	var arrayQuestions = []string{"https://leetcode.com/problems/two-sum/", "https://leetcode.com/problems/best-time-to-buy-and-sell-stock/", "https://leetcode.com/problems/contains-duplicate/", "https://leetcode.com/problems/product-of-array-except-self/", "https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/", "https://leetcode.com/problems/3sum/", "https://leetcode.com/problems/container-with-most-water/"}
	var binaryQuestions = []string{"https://leetcode.com/problems/sum-of-two-integers/", "https://leetcode.com/problems/number-of-1-bits/", "https://leetcode.com/problems/counting-bits/", "https://leetcode.com/problems/missing-number/", "https://leetcode.com/problems/reverse-bits/"}
	var linkedListQuestions = []string{"https://leetcode.com/problems/reverse-linked-list/", "https://leetcode.com/problems/linked-list-cycle/", "https://leetcode.com/problems/merge-two-sorted-lists/", "https://leetcode.com/problems/merge-k-sorted-lists/", "https://leetcode.com/problems/remove-nth-node-from-end-of-list/", "https://leetcode.com/problems/reorder-list/"}
	var heapQuestions = []string{"https://leetcode.com/problems/merge-k-sorted-lists/", "https://leetcode.com/problems/top-k-frequent-elements/", "https://leetcode.com/problems/find-median-from-data-stream/"}

	//Blind 75 List
	var blindSeventyFive = []string{"https://leetcode.com/problems/two-sum/", "https://leetcode.com/problems/best-time-to-buy-and-sell-stock/", "https://leetcode.com/problems/contains-duplicate/", "https://leetcode.com/problems/product-of-array-except-self/", "https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/", "https://leetcode.com/problems/3sum/", "https://leetcode.com/problems/container-with-most-water/", "https://leetcode.com/problems/reverse-linked-list/", "https://leetcode.com/problems/linked-list-cycle/", "https://leetcode.com/problems/merge-two-sorted-lists/", "https://leetcode.com/problems/merge-k-sorted-lists/", "https://leetcode.com/problems/remove-nth-node-from-end-of-list/", "https://leetcode.com/problems/reorder-list/", "https://leetcode.com/problems/sum-of-two-integers/", "https://leetcode.com/problems/number-of-1-bits/", "https://leetcode.com/problems/counting-bits/", "https://leetcode.com/problems/missing-number/", "https://leetcode.com/problems/reverse-bits/", "https://leetcode.com/problems/merge-k-sorted-lists/", "https://leetcode.com/problems/top-k-frequent-elements/", "https://leetcode.com/problems/find-median-from-data-stream/"}
	var flag bool = false

	if flag == true {
		fmt.Println(linkedListQuestions, heapQuestions, binaryQuestions, blindSeventyFive)

	}
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "*yo" {
		s.ChannelMessageSend(m.ChannelID, "SHEEEEEEEEEEEEESSHHH")
	}
	//Learn a specific topic
	if m.Content == "*learn arrays" {
		//Set up spaced rep for arrays
		isSuccessful, scheduleLink := scheduleEvents(arrayQuestions, m.Author.ID)

		if isSuccessful {
			s.ChannelMessageSend(m.ChannelID, ("Newly Generated Spaced Repetition Calendar" + scheduleLink))
		} else {
			s.ChannelMessageSend(m.ChannelID, "Failed to generate a new spaced repetition for test user")
		}

	}
	// if m.Content == "*learn binary" {
	// 	for i := 0; i < len(binaryQuestions); i++ {
	// 		s.ChannelMessageSend(m.ChannelID, binaryQuestions[i])
	// 	}
	// }
}

/*
	1. Create a calendar with that users id


	We want to be reminded about each problem with a schedule

	dates:
		i days from today
		i days from today + 2 days
		i days from today + 30 days
		i days from today + 60 days


	For each item in the list we want to add all of these days to the new calendar we made



	then we respond with an invite to the calendar with the given spaced repetition schedule

*/
func scheduleEvents(questions []string, calendarId string) (bool, string) {
	var scheduledSuccessfully bool = true
	calendarLink := ""
	fmt.Println(calendarId)

	//Generate a calendar with the given id
	for i := 0; i < len(questions); i++ {
		//Add an event to the calendar
	}

	return (scheduledSuccessfully), (calendarLink)
}
