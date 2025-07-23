package internal

import (
    "context"
    "fmt"
    "google.golang.org/genai"
)

func GenerateMessage(message string) (string, error) {
    ctx := context.Background()
    client, err := genai.NewClient(ctx, &genai.ClientConfig{
        APIKey:  "api_key",
        Backend: genai.BackendGeminiAPI,
    })
    if err != nil {
		return "", err
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text(fmt.Sprintf("Here is some content extracted from a web page. Based on this, write a short, friendly reminder message to a user encouraging them to visit the website. First, understand what the page is about by analyzing the content and structure. Then, generate a 1–2 sentence message summarizing what the user can expect to find on the page and gently reminding them to check it out. and this site is chosen by the user so just think this as a reminder message abouthe site. and remember only send the final reminder do not add other thing just the reminder message Here is the content of the web page: %v",message)),
        nil,
    )
    if err != nil {
		return "", err
    }
	fmt.Println(result.Text())
	return result.Text(), nil
}
