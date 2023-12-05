# streams-of-conn
beeps websocket coding challenge as per http://beeps.gg/interview

## Design Considerations
`streams-of-conn` is a CLI app that streams websocket data. While the instructions above were for a web-app implementation, I leaned into my backend skills and created a CLI app instead.

There are two modes available for displaying the data: default and social. In the default mode, all data is printed in the CLI as it is streamed. When the `socialFeed` flag is set to `true` however, only timestamp, username, message, and tags are displayed. In both modes, the requested stats are available at the bottom of the screen.

To account for different terminal size preferences, I added a `bufferSize` flag to allow users to specify how many events they wanted to view at once. A future improvement would be to update the implementation with a go framework like https://github.com/charmbracelet/bubbletea, which would allow for more UX-friendly features like scrolling.

## Usage
To run this app, run `go build` and `./streams-of-conn` to run with default settings.

```
./streams-of-conn --help           
A simple CLI app for streaming events from a websocket. A stream of con(ciousness)nection if you will.

Usage:
  streams-of-conn [flags]

Flags:
      --bufferSize int   Sets the buffer flag to an int value. (default 7)
  -h, --help             help for streams-of-conn
      --socialFeed       When true, only timestamp, username, message, and tags are displayed.
  -u, --url string       Sets the WebSocket URL. (default "ws://beeps.gg/stream")
```

## Example Output
#### Default Mode (all data)
```
	Event ID: aeFVsPhuc
	Event Timestamp: 04 Apr 98 00:36 EDT
	User ID: 3538
	User Image URL: http://loremflickr.com/200/204/
	User Name: Kate Beckinsale
	User Username: rau7274
	Message: Whichever inexpensive any have band as this your paint sleepy wisp few.
	Tags: #quinoa, #microdosing, #mumblecore, #gastropub, #jeanshorts

	Event ID: YmFddfU
	Event Timestamp: 04 Apr 98 01:26 EDT
	User ID: 13729
	User Image URL: http://loremflickr.com/205/207/
	User Name: Jessica Alba
	User Username: wisozk2630
	Message: His whoever been besides Laotian surprise that of substantial till being down.
	Tags: #chartreuse, #sartorial, #smallbatch, #stumptown, #nextlevel, #heirloom, #park, #fixie

--------------------------------------------------------------------------------
 📲  Streaming from ws://beeps.gg/stream 
 🖊️  Total Events: 17 
 ⏱️  Event Rate (per minute): 62.46 

--------------------------------------------------------------------------------
 💡 Press 'Ctrl+C' to quit.        
--------------------------------------------------------------------------------
```

#### Social Mode
```
	03 Apr 98 03:30 EDT | @carter8395 
	Hail we would laughter how sand this indeed kindness horde dive quarterly.
	#echo, #freegan, #helvetica, #chambray, #gentrify, #disrupt, #3wolfmoon, #offal, #deepv, #kogi

	03 Apr 98 04:20 EDT | @gleichner9009 
	This madly it consequently condemned adorable can soon such between first from.
	#cornhole, #venmo

	03 Apr 98 04:20 EDT | @wunsch2493 
	Well while Hindu waiter party out the its clearly yet where mine.
	#porkbelly, #distillery, #cred

--------------------------------------------------------------------------------
 📲  Streaming from ws://beeps.gg/stream 
 🖊️  Total Events: 34 
 ⏱️  Event Rate (per minute): 59.52 

--------------------------------------------------------------------------------
 💡 Press 'Ctrl+C' to quit.        
--------------------------------------------------------------------------------
```
