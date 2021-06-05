package notify

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"


	"github.com/koenverburg/vaccine-notifier/src/utils"
	"github.com/slack-go/slack"
)

func NotifyViaSlack(namespace string, msgText string, attachmentText string, color string) {
	attachment := slack.Attachment{
		Text:   attachmentText,
		Color:  color,
		Footer: "Vaccine Notifier - cronjob",
		Ts:     json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
	}
	msg := slack.WebhookMessage{
		Text:        msgText,
		Attachments: []slack.Attachment{attachment},
		Username:    fmt.Sprintf("Vaccine Notifier - %s", namespace),

		Channel: utils.Env("SLACK_CHANNEL"),
		IconURL: "https://github.com/koenverburg.png?size=48",
	}

	err := slack.PostWebhook(utils.Env("SLACK_WEBHOOK_URL"), &msg)
	utils.CheckIfErr(err)
}
