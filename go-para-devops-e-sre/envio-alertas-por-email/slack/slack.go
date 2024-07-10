package slack

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func SendSlack(textAlert string) {
	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		panic("SLACK_TOKEN not found")
	}

	channelId := os.Getenv("SLACK_CHANEL_ID")
	if channelId == "" {
		panic("SLACK_CHANEL_ID not found")
	}

	client := slack.New(token, slack.OptionDebug(true))
	attachment := slack.Attachment{
		Color:   "danger",
		Pretext: "Alerta de servidor down",
		Text:    textAlert,
	}
	_, timestamp, err := client.PostMessage(
		channelId,
		slack.MsgOptionAttachments(attachment),
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Mensagem enviada com sucesso %s as %s\n", channelId, timestamp)
}
