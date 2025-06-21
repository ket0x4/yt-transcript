# YouTube Transcript Downloader

A Go library and command-line tool to download transcripts (captions) for YouTube videos.

## Features

- List all available transcripts for a video.
- Download a transcript in a specific language.
- Can be used as a command-line tool or as a library in your own Go projects.

## Command-Line Usage

### Installation

Clone the repository:

```sh
git clone <repository-url>
cd yt-transcript
```

### Running the tool

You can run the tool using `go run`.

Or Build it with `go build -ldflags="-s -w" -o yt-transcript` and run the executable.

**List available transcripts:**

Provide a YouTube video ID to see all available caption tracks.

```sh
go run main.go <video_id>
```

**Example:**
```sh
go run main.go dQw4w9WgXcQ
```
**Output:**
```
Listing available transcripts...
Available transcripts:
- Language: en, Name: English, Kind: manual
- Language: es, Name: Spanish, Kind: asr
...
```

**Download a specific transcript:**

Provide the video ID and the desired language code.

```sh
go run main.go <video_id> <language_code>
```

**Example:**
```sh
go run main.go dQw4w9WgXcQ en
```
**Output:**
```
Transcript (en):
We're no strangers to love
You know the rules and so do I
...
```

## Library Usage

You can also use this project as a library in your own Go applications.

### Installation

```sh
go get <repository-url>/yttranscript
```

### Example

Here's a simple example of how to use the `yttranscript` library:

```go
package main

import (
	"fmt"
	"log"

	"<repository-url>/yttranscript"
)

func main() {
	videoID := "dQw4w9WgXcQ"

	// Create a new client
	client, err := yttranscript.New()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// List available transcripts
	tracks, err := client.ListTranscripts(videoID)
	if err != nil {
		log.Fatalf("Failed to list transcripts: %v", err)
	}

	fmt.Println("Available transcripts:")
	for _, track := range tracks {
		fmt.Printf("- Language: %s, Name: %s\n", track.LanguageCode, track.Name.SimpleText)
	}

	// Get the English transcript
	transcript, err := client.GetTranscript(videoID, "en")
	if err != nil {
		log.Fatalf("Failed to get transcript: %v", err)
	}

	fmt.Println("\nEnglish Transcript:")
	for _, text := range transcript.Texts {
		fmt.Println(text.Content)
	}
}
```

