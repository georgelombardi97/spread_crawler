package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"io"
	"net/http"
	//"os"
)

type KucoinOrderBookL2 struct {
	Code string                `json:"code"`
	Data KucoinOrderBookL2Data `json:"data"`
}

type KucoinOrderBookL2Data struct {
	Time     uint64     `json:"time"`
	Sequence string     `json:"sequence"`
	Bids     [][]string `json:"bids"`
	Asks     [][]string `json:"asks"`
}

func main() {
	for {
		fetchKucoinAndAppendToFile("ATOM-USDT")
		time.Sleep(10 * time.Second)
	}
}

func fetchKucoinAndAppendToFile(pairSymbol string) {
	// url := "https://api.kucoin.com/api/v1/market/orderbook/level2_20?symbol=ATOM-USDT"
	url := fmt.Sprintf("https://api.kucoin.com/api/v1/market/orderbook/level2_20?symbol=%s", pairSymbol)
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	b, _ := io.ReadAll(response.Body)

	var orderbook KucoinOrderBookL2
	errUnmarshal := json.Unmarshal(b, &orderbook)
	if errUnmarshal != nil {
		fmt.Println("err ", errUnmarshal)
	}
	// fmt.Printf("orderbook: %+v\n", orderbook)
	// fmt.Println("code", orderbook.Code)

	fmt.Println("health:", time.Now())

	filename := fmt.Sprintf("%s.txt", pairSymbol)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if f != nil {
		defer f.Close()

		jsonOrderBook, _ := json.Marshal(orderbook)
		_, errF := f.WriteString(string(jsonOrderBook) + "\n")
		if errF != nil {
			return
		}
	}

}
