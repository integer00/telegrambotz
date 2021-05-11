# telegrambotz
test bot

#### amazon config:
```bash
$ export AWS_ACCESS_KEY_ID="anaccesskey"
$ export AWS_SECRET_ACCESS_KEY="asecretkey"
```

https://core.telegram.org/bots/api#getting-updates

Get token: [https://core.telegram.org/bots#3-how-do-i-create-a-bot](https://core.telegram.org/bots#3-how-do-i-create-a-bot)

```bash
$ export TELEGRAM_TOKEN=<TOKEN>
```
### Examples cli
Check [updates](https://core.telegram.org/bots/api#getting-updates):
```bash
curl https://api.telegram.org/bot$TELEGRAM_TOKEN/getUpdates | jq .
```

send [message](https://core.telegram.org/bots/api#sendmessage):
```bash
curl -X POST \
     -H 'Content-Type: application/json' \
     -d '{"chat_id": "123456789", "text": "This is a test from curl", "disable_notification": true}' \
     https://api.telegram.org/bot$TELEGRAM_TOKEN/sendMessage

```