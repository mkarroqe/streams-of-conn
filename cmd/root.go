package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)

var url string
var bufferSize int
var socialFeed bool

var rootCmd = &cobra.Command{
	Use:   "streams-of-conn",
	Short: "CLI app that streams events from a websocket.",
	Long:  "A simple CLI app for streaming events from a websocket. A stream of con(ciousness)nection if you will.",
	Run:   runWebSocket,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&url, "url", "u", "ws://beeps.gg/stream", "Sets the WebSocket URL.")
	rootCmd.PersistentFlags().BoolVar(&socialFeed, "socialFeed", false, "When true, only timestamp, username, message, and tags are displayed.")

	// TODO: consider buffer restrictions?
	rootCmd.PersistentFlags().IntVar(&bufferSize, "bufferSize", 7, "Sets the buffer flag to an int value.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runWebSocket(cmd *cobra.Command, args []string) {
	// Establish WebSocket connection
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		fmt.Println("Dial error:", err)
	}
	defer ws.Close()

	// Channel to receive interrupt signals (Ctrl+C)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Channel to receive WebSocket messages
	messageChan := make(chan []byte)

	// Initialize event metrics
	eventCount := 0
	startTime := time.Now()
	eventRate := -1.0
	eventBuffer := make([]Event, 0, bufferSize)

	// goroutine that continuously listens to the channel
	go func() {
		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				fmt.Println("Error reading message:", err)
				return
			}
			messageChan <- message
		}
	}()

	go func() {
		for {
			select {
			case message := <-messageChan:
				currEvent := extractEventData(message)

				// Update the event buffer
				eventBuffer = append(eventBuffer, currEvent)
				if len(eventBuffer) > bufferSize {
					eventBuffer = eventBuffer[1:]
				}

				eventCount++
				eventRate = calculateEventsPerMin(eventCount, startTime)
				displayEvent(eventCount, eventRate, socialFeed, eventBuffer)

			case <-interrupt:
				fmt.Println("/nReceived interrupt. Exiting.")
				os.Exit(0)
			}
		}
	}()

	// Keep the main goroutine running until an interrupt signal is received
	<-interrupt
}
