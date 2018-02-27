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
   
   func Example() {
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
   
