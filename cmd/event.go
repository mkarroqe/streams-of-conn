package cmd

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Event struct {
	ID        string   `json:"id"`
	Timestamp int64    `json:"timestamp"`
	User      User     `json:"user"`
	Message   string   `json:"message"`
	Tags      []string `json:"tags"`
}

type User struct {
	ID       int    `json:"id"`
	ImageURL string `json:"image_url"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

func extractEventData(jsonEvent []byte) Event {
	var currEvent Event
	if err := json.Unmarshal(jsonEvent, &currEvent); err != nil {
		fmt.Println("error extracting websocket data", err)
	}

	// Add precending hashtag to tags & remove spaces
	for i, tag := range currEvent.Tags {
		tag = strings.ReplaceAll(tag, " ", "")
		currEvent.Tags[i] = "#" + tag
	}

	return currEvent
}

func calculateEventsPerMin(eventCount int, startTime time.Time) float64 {
	elapsedTime := time.Since(startTime)
	elapsedMinutes := elapsedTime.Minutes()

	return float64(eventCount) / elapsedMinutes
}

func displayEvent(eventCount int, eventRate float64, socialFeed bool, eventBuffer []Event) {
	fmt.Print("\033[2J\033[H") // Move cursor to the top-left corner of the screen and clear

	// TODO: place this logic elsewhere?
	if socialFeed {
		for _, event := range eventBuffer {
			fmt.Printf("\n\t%v | @%s \n", time.Unix(event.Timestamp, 0).Format(time.RFC822), event.User.Username)
			fmt.Printf("\t%s\n", event.Message)
			fmt.Printf("\t%s\n", strings.Join(event.Tags, ", "))
		}
	} else {
		for _, event := range eventBuffer {
			fmt.Printf("\n\tEvent ID: %s\n", event.ID)
			fmt.Printf("\tEvent Timestamp: %v\n", time.Unix(event.Timestamp, 0).Format(time.RFC822))
			fmt.Printf("\tUser ID: %d\n", event.User.ID)
			fmt.Printf("\tUser Image URL: %s\n", event.User.ImageURL)
			fmt.Printf("\tUser Name: %s\n", event.User.Name)
			fmt.Printf("\tUser Username: %s\n", event.User.Username)
			fmt.Printf("\tMessage: %s\n", event.Message)
			fmt.Printf("\tTags: %s\n", strings.Join(event.Tags, ", "))
		}
	}

	fmt.Println("\n--------------------------------------------------------------------------------")
	fmt.Printf(" 📲 Streaming from %s \n", url)
	fmt.Printf(" 🖊️  Total Events: %d \n", eventCount)
	// TODO(fixme): make whitespace dynamically generated
	if eventCount > 0 {
		// TODO: calculate eventRate here instead?
		fmt.Printf(" ⏱️  Event Rate (per minute): %.2f \n", eventRate)
	}
	fmt.Println("\n--------------------------------------------------------------------------------")
	fmt.Println(" 💡 Press 'Ctrl+C' to quit.        ")
	fmt.Println("--------------------------------------------------------------------------------")
}
