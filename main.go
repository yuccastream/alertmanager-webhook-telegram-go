package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	botapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gorilla/mux"
)

var (
	BotToken      string
	ChatID        int64
	addressListen string
)

const (
	timeDateFormat = "2006-01-02 15:04:05"
)

func main() {

	s := os.Getenv("CHAT_ID")
	if s == "" {
		log.Fatal("Empty env CHAT_ID")
	}
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	ChatID = n
	BotToken = os.Getenv("BOT_TOKEN")
	if BotToken == "" {
		log.Fatal("Empty env BOT_TOKEN")
	}

	flag.StringVar(&addressListen, "l", ":8080", "Listen adderss")
	flag.Parse()

	if BotToken == "" || ChatID == 0 {
		log.Fatal("Empty env BOT_TOKEN or CHAT_ID")
	}

	router := mux.NewRouter()
	router.HandleFunc("/alert", ToTelegram).Methods("POST")

	log.Println("Listen", addressListen)
	log.Fatal(http.ListenAndServe(addressListen, router))
}

type alertmanagerAlert struct {
	Receiver string `json:"receiver"`
	Status   string `json:"status"`
	Alerts   []struct {
		Status string `json:"status"`
		Labels struct {
			Name      string `json:"name"`
			Instance  string `json:"instance"`
			Alertname string `json:"alertname"`
			Service   string `json:"service"`
			Severity  string `json:"severity"`
		} `json:"labels"`
		Annotations struct {
			Info        string `json:"info"`
			Description string `json:"description"`
			Summary     string `json:"summary"`
		} `json:"annotations"`
		StartsAt     time.Time `json:"startsAt"`
		EndsAt       time.Time `json:"endsAt"`
		GeneratorURL string    `json:"generatorURL"`
		Fingerprint  string    `json:"fingerprint"`
	} `json:"alerts"`
	GroupLabels struct {
		Alertname string `json:"alertname"`
	} `json:"groupLabels"`
	CommonLabels struct {
		Alertname string `json:"alertname"`
		Service   string `json:"service"`
		Severity  string `json:"severity"`
	} `json:"commonLabels"`
	CommonAnnotations struct {
		Summary string `json:"summary"`
	} `json:"commonAnnotations"`
	ExternalURL string `json:"externalURL"`
	Version     string `json:"version"`
	GroupKey    string `json:"groupKey"`
}

// ToTelegram function responsible to send msg to telegram
func ToTelegram(w http.ResponseWriter, r *http.Request) {

	var alerts alertmanagerAlert

	bot, err := botapi.NewBotAPI(BotToken)
	if err != nil {
		log.Panic(err)
	}

	_ = json.NewDecoder(r.Body).Decode(&alerts)

	for _, alert := range alerts.Alerts {
		var (
			status string
			mtime  string
		)
		switch alert.Status {
		case "firing":
			status = "🔥 **" + alert.Labels.Alertname + "**"
			mtime = alert.EndsAt.Format(timeDateFormat)
		case "resolved":
			status = "✅ **" + alert.Labels.Alertname + "**"
			mtime = alert.StartsAt.Format(timeDateFormat)
		}
		telegramMsg := status + "\n"

		if alert.Annotations.Description != "" {
			telegramMsg += "Description: " + alert.Annotations.Description + "\n"
		}
		telegramMsg += mtime + "\n"

		msg := botapi.NewMessage(-ChatID, telegramMsg)
		bot.Send(msg)

	}

	log.Println(alerts)
	json.NewEncoder(w).Encode(alerts)

}
