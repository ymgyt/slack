.. image:: https://godoc.org/github.com/ymgyt/slack/webhook?status.svg
.. image:: https://coveralls.io/repos/github/ymgyt/slack/badge.svg?branch=develop
:target: https://coveralls.io/github/ymgyt/slack?branch=develop

.. image:: https://goreportcard.com/badge/github.com/ymgyt/slack

=====
slack
=====

slack api client 


Webhook
=======

Usage
-----

.. code-block:: go

   import "github.com/ymgyt/slack/webhook"
   
   func Notify() {
   	wh, err := webhook.New(webhook.Config{
   		URL:       "https://hooks.slack.com/services/XXXXXXXXXXXX/XXXXXXXXXXXXXXXXXXX",
   		Channel:   "general",
   		Username:  "gopher",
   		IconEmoji: "+1",
   		Timeout:   0,
   		Dump:      true,
   	})
   	if err != nil {
   		panic(err)
   	}
   	err = wh.SendPayload(&webhook.Payload{
   		Text: "text content",
   		Attachments: []*webhook.Attachment{
   			{
   				Fallback: "fallback content",
   				Text:     "attachment text 1",
   				Pretext:  "pretext",
   				Color:    "good",
   				Fields: []*webhook.Field{
   					{
   						Title: "title 1",
   						Value: "field content. hello !",
   						Short: true,
   					},
   					{
   						Title: "title 2",
   						Value: "field content. hello !",
   						Short: true,
   					},
   				},
   			},
   			{
   				Fallback: "fallback content",
   				Text:     "attachment text 2",
   				Pretext:  "pretext",
   				Color:    "warning",
   				Fields: []*webhook.Field{
   					{
   						Title: "title 1",
   						Value: "field content. hello !",
   						Short: true,
   					},
   					{
   						Title: "title 2",
   						Value: "field content. hello !",
   						Short: true,
   					},
   				},
   			},
   		},
   	})
   	if err != nil {
   		panic(err)
   	}
   }
   
