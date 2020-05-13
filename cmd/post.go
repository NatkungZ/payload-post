package main

import (
	"net/http"
	"fmt"
	"bytes"
	"io/ioutil"
	"bufio"
	"os"
	"flag"
)

func main() {
	webhookURL := flag.String("webhook-url", "", "a webhook URL")
	flag.Parse()

	if *webhookURL == "" {
		printUsage()
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	msgText, err := reader.ReadString('\n')

    fmt.Println("URL:>", *webhookURL)
    fmt.Println("msgText:>", msgText)

    var jsonStr = []byte(msgText)
    req, err := http.NewRequest("POST", *webhookURL, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}

func printUsage() {
	fmt.Println(`
Usage: discord STDIN --webhook-url=WEBHOOK_URL

  Send a message to Discord channel.

Options:
  --webhook-url A Discord Webhook URL of the channel you want to send message to`)
}
