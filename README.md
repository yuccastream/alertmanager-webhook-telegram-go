# Alertmanager webhook for Telegram

Go version 1.16.4

## BUILD

* make build
* make docker

## Alertmanager configuration example

```yaml
	receivers:
	- name: 'telegram-webhook'
	  webhook_configs:
	  - url: http://ipGoAlert:8080/alert
	    send_resolved: true
```

## Running on docker

```bash
git clone https://github.com/yuccastream/alertmanager-webhook-telegram-go.git
cd ./alertmanager-webhook-telegram-go
docker build -t awt-go .

docker run -d --name awt -e "BOT_TOKEN=telegramBot:Token" -e "CHAT_ID=32" -p 8080:8080 yuccastream/awt-go:latest
```
