package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// A Response struct to map the Entire Response
type Response struct {
	Data Data `json:"data"`
}

type Data struct {
	Catalog Catalog `json:"Catalog"`
}

type Catalog struct {
	SearchStore SearchStore `json:"searchStore"`
}

type SearchStore struct {
	Element []Element `json:"elements"`
}

type Element struct {
	Title      string     `json:"title"`
	Promotions Promotions `json:"promotions"`
}

type Promotions struct {
	PromotionalOffers []PromotionalOffers `json:"promotionalOffers"`
}

type PromotionalOffers struct {
	MainPromotionalOffers []MainPromotionalOffers `json:"promotionalOffers"`
}

type MainPromotionalOffers struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

func main() {
	response, err := http.Get("https://store-site-backend-static-ipv4.ak.epicgames.com/freeGamesPromotions?locale=en-US&country=PL&allowCountries=PL")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	for i := 0; i < len(responseObject.Data.Catalog.SearchStore.Element); i++ {
		for x := 0; x < len(responseObject.Data.Catalog.SearchStore.Element[i].Promotions.PromotionalOffers); x++ {
			fmt.Printf("Hurry \"%v\" is now free\n", responseObject.Data.Catalog.SearchStore.Element[i].Title)
		}
	}

}
