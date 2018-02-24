=====
slack
=====

slack api client 


Webhook
=======

Usage
-----

.. code-block:: go

   import (
       "github.com/ymgyt/slack/webhook"
   )
   
   func NotifySlack(msg string) {
   	webhook, err := New(&Config{
   		URL:     "https://hooks.slack.com/services/XXXXXXXXXXXXXXXXXXXX/AAAAAAAAAAAAAAAAAAAAAAAA",
   		Channel: "general",
   	})
   	if err != nil {
   		panic(err)
   	}
   	err = webhook.Send(msg)
   	if err != nil {
   		panic(err)
   	}
   }
