# Alertmanager webhook for Telegram

![Go](https://github.com/nopp/alertmanager-webhook-telegram-go/workflows/Go/badge.svg)
![Docker Image CI](https://github.com/nopp/alertmanager-webhook-telegram-go/workflows/Docker%20Image%20CI/badge.svg)
![Code scanning - action](https://github.com/nopp/alertmanager-webhook-telegram-go/workflows/Code%20scanning%20-%20action/badge.svg)

Python Version (https://github.com/nopp/alertmanager-webhook-telegram-python) 

Go version 1.13.9

## BUILD

* make build
* make docker

Alertmanager configuration example
==================================

	receivers:
	- name: 'telegram-webhook'
	  webhook_configs:
	  - url: http://ipGoAlert:8080/alert
	    send_resolved: true

Running on docker
=================
    git clone https://github.com/yuccastream/alertmanager-webhook-telegram-go.git
    cd ./alertmanager-webhook-telegram-go
    docker build -t awt-go .

    docker run -d --name awt \
    	-e "BOT_TOKEN=telegramBot:Token" \
    	-e "CHAT_ID=32" \
    	-p 8080:8080 yuccastream/awt-go:latest
