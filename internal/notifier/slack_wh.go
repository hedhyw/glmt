package notifier

import (
	"context"

	"gitlab.com/gitlab-merge-tool/glmt/internal/templating"

	"github.com/slack-go/slack"
)

func NewSlackWebHookNotifier(url, messageTemplate string) *SlackWebHookNotifier {
	return &SlackWebHookNotifier{
		url:             url,
		messageTemplate: messageTemplate,
	}
}

type SlackWebHookNotifier struct {
	url             string
	messageTemplate string
}

func (sn *SlackWebHookNotifier) Send(ctx context.Context, args map[string]string, add string) error {
	templ := sn.messageTemplate
	if templ == "" {
		templ = "<!here>\n{{.Description}}\n{{.MergeRequestURL}}"
	}

	m := templating.CreateText("slack_wh_message", templ, args)

	if add != "" {
		m = "\n" + add
	}

	msg := &slack.WebhookMessage{
		Text: m,
	}
	return slack.PostWebhookContext(ctx, sn.url, msg)
}
