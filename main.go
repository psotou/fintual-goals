package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const (
	baseURL = "https://fintual.cl/api/goals"
)
const (
	hoursPerDay  = 24
	daysPerMonth = 30
)

type goal struct {
	Data []struct {
		Attributes struct {
			Name      string    `json:"name"`
			Nav       float64   `json:"nav"`
			Deposited float64   `json:"deposited"`
			Profit    float64   `json:"profit"`
			CreatedAt time.Time `json:"created_at"`
		} `json:"attributes"`
	} `json:"data"`
}

func main() {
	u, _ := url.Parse(baseURL)

	if err := godotenv.Load(); err != nil {
		log.Fatal(errors.New("Couldn't load any .env file"))
	}

	q := u.Query()
	q.Add("user_email", os.Getenv("USER_EMAIL"))
	q.Add("user_token", os.Getenv("USER_TOKEN"))
	u.RawQuery = q.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var g goal
	json.Unmarshal(body, &g)

	p := message.NewPrinter(language.Spanish)
	fmt.Printf("\033[1m%22s %10s %10s %10s %8s\033[0m\n", "", "Deposited", "Balance", "Profit", "Months")
	for _, v := range g.Data {
		durationInHours := time.Since(v.Attributes.CreatedAt).Hours()
		durationInMonths := durationInHours / (hoursPerDay * daysPerMonth)
		p.Printf("%22s %10.f %10.f %10.f %8.1f\n", v.Attributes.Name, v.Attributes.Deposited, v.Attributes.Nav, v.Attributes.Profit, durationInMonths)
	}
}
