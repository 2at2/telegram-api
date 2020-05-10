# telegram-api
Telegram Bot API

```
bot, err := api.NewBot("111111111:XXXXXXXXXXXXXXXXXXXXXX")
if err != nil {
    return err
}

// Bot listener
messages := make(chan api.Message, 100)
callbacks := make(chan api.Callback, 100)

bot.Listen(
    messages,
    nil,
    callbacks,
    5*time.Second,
)

// Reading messages
go func(list chan api.Message) {
    for entity := range list {           
        go processIncomingMessage(entity)
    }
}(messages)

go func(list chan api.Callback) {
    for entity := range list {
        go processIncomingCallback(entity)
    }
}(callbacks)
```
