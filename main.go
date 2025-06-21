package main

import (
	"fmt"
	"log"
	"os"

	"yt-transcript/yttranscript"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run main.go <video_id> [language_code]")
	}
	videoID := os.Args[1]

	client, err := yttranscript.New()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	if len(os.Args) == 2 {
		// If no language code is provided, list available transcripts.
		fmt.Println("Listing available transcripts...")
		tracks, err := client.ListTranscripts(videoID)
		if err != nil {
			log.Fatalf("Failed to list transcripts: %v", err)
		}
		if len(tracks) == 0 {
			fmt.Println("No transcripts found for this video.")
			return
		}
		fmt.Println("Available transcripts:")
		for _, track := range tracks {
			fmt.Printf("- Language: %s, Name: %s, Kind: %s\n", track.LanguageCode, track.Name.SimpleText, track.Kind)
		}
		return
	}

	languageCode := os.Args[2]
	transcript, err := client.GetTranscript(videoID, languageCode)
	if err != nil {
		log.Fatalf("Failed to get transcript: %v", err)
	}

	fmt.Printf("\nTranscript (%s):\n", languageCode)
	for _, text := range transcript.Texts {
		fmt.Println(text.Content)
	}
}
