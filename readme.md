## Adding Webhook
````
curl -F "url=<WEBHOOK_DOMAIN>/webhook" https://api.telegram.org/bot<TELEGRAM_BOT_TOKEN>/setWebhook
````

## Pass enviroment variable locally

````
TELEGRAM_BOT_TOKEN=<TELEGRAM_BOT_TOKEN> go run main.go  handler.go   
````