package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/araddon/dateparse"
	"github.com/emersion/go-webdav"
	"github.com/emersion/go-webdav/carddav"
)

type Payload struct {
	Text     string `json:"text"`
	Username string `json:"username"`
	IconUrl  string `json:"icon_url"`
}

func main() {

	btvDebug := os.Getenv("BTV_DEBUG")
	server := os.Getenv("WEBDAV_SERVER")
	addressBook := os.Getenv("WEBDAV_ADRESSBOOK")
	username := os.Getenv("WEBDAV_USERNAME")
	password := os.Getenv("WEBDAV_PASSWORD")
	webhookUrl := os.Getenv("WEBHOOK_URL")
	botName := os.Getenv("BOT_NAME")
	iconUrl := os.Getenv("ICON_URL")

	if btvDebug != "" {
		fmt.Println("WEBDAV_SERVER", server)
		fmt.Println("WEBDAV_ADDRESSBOOK", addressBook)
		fmt.Println("WEBDAV_USERNAME", username)
		fmt.Println("WEBDAV_PASSWORD", password)
		fmt.Println("WEEB_HOOK", webhookUrl)
		fmt.Println("BOT_NAME", botName)
		fmt.Println("ICON_URL", iconUrl)
	}

	var webDavClient webdav.HTTPClient = http.DefaultClient
	if username != "" {
		webDavClient = webdav.HTTPClientWithBasicAuth(webDavClient, username, password)
	}

	cardDavClient, err := carddav.NewClient(webDavClient, server)
	if err != nil {
		log.Fatal(err.Error())
	}

	res, err := cardDavClient.QueryAddressBook(addressBook, &carddav.AddressBookQuery{
		DataRequest: carddav.AddressDataRequest{
			AllProp: true,
		},
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, v := range res {

		name := v.Card.Value("FN")
		birthDay := v.Card.Value("BDAY")
		if birthDay == "" {
			continue
		}

		t, err := dateparse.ParseLocal(birthDay)
		if err != nil {
			fmt.Println("Error while parsing date :", err)
		}

		now := time.Now()
		if now.Day() == t.Day() && now.Month() == t.Month() {

			message := ""
			var defaultYear int = 1604

			if now.Year() == defaultYear {
				message = name
			} else {

				diffYears := fmt.Sprintf("%d", now.Year()-t.Year())
				message = name + " (" + diffYears + ")"
			}

			payload := &Payload{
				Text:     message,
				Username: botName,
				IconUrl:  iconUrl,
			}

			if btvDebug != "" {
				fmt.Printf("%#v\n", payload)
				fmt.Printf("%s\n", message)
			}

			dataBytes := new(bytes.Buffer)
			err := json.NewEncoder(dataBytes).Encode(payload)
			if err != nil {
				log.Fatalln(err)
			}

			payloadMarshal, err := json.Marshal(payload)
			if err != nil {
				log.Fatalln(err)
			}

			resp, err := http.Post(webhookUrl, "application/json", bytes.NewBuffer(payloadMarshal))
			if err != nil {
				log.Fatalln(err)
			}

			s, _ := ioutil.ReadAll(resp.Body)
			if btvDebug != "" {
				fmt.Println("STATUS CODE", resp.StatusCode)
				fmt.Println("STATUS CODE", string(s))
			}
		}
	}

}
