package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type GoalData struct {
	Metadata []Data `json:"data"`
}

type Data struct {
	Attribute Attrb `json:"attributes"`
}

type Attrb struct {
	Name      string  `json:"name"`
	Amount    float64 `json:"nav"`
	Deposited float64 `json:"deposited"`
	Profit    float64 `json:"profit"`
}

func main() {
	requestUrl := fmt.Sprintf("https://fintual.cl/api/goals?user_email=pascualsu%sgmail.com&user_token=%s", "%40", os.Getenv("FINTUAL_TOKEN"))

	req, resData := GetReq(requestUrl)
	defer req.Body.Close()

	var resObject GoalData
	json.Unmarshal(resData, &resObject)

	fmt.Printf("%24s %10s %10s %10s\n", "GOAL NAME", "BALANCE", "DEPOSITED", "PROFIT")
	fmt.Printf("%24s %10.f %10.f %10.f\n", resObject.Metadata[0].Attribute.Name, resObject.Metadata[0].Attribute.Amount, resObject.Metadata[0].Attribute.Deposited, resObject.Metadata[0].Attribute.Profit)
	fmt.Printf("%24s %10.f %10.f %10.f\n", resObject.Metadata[1].Attribute.Name, resObject.Metadata[1].Attribute.Amount, resObject.Metadata[1].Attribute.Deposited, resObject.Metadata[1].Attribute.Profit)
	fmt.Printf("%24s %10s %10s %10.f\n", "TOTAL PROFIT", "", "", resObject.Metadata[0].Attribute.Profit+resObject.Metadata[1].Attribute.Profit)
}

func GetReq(url string) (*http.Response, []byte) {
	req, err := http.Get(url)
	if err != nil {
		log.Fatal("Error en la conexión", err.Error())
	}

	req.Header.Add("Accept", "application/json")
	resData, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	return req, resData
}
