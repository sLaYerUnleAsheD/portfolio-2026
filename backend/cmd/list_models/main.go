package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/genai"
)

func main() {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY env var not set")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	iter := client.Models.List(ctx, nil)
	for {
		model, err := iter.Next()
		if err != nil {
			if err.Error() == "iterator done" { // Simple check, real err is iterator.Done
				break
			}
			// Just try to keep going or break on real error. The SDK usually returns an iterator error.
			// Actually the correct way to break is when err != nil. Let's just print err and break.
			log.Println("Done or err:", err)
			break
		}
		fmt.Println("Model:", model.Name)
	}
}
