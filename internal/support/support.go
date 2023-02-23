package support

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

func Support() []int {
	resp, err := http.Get("http://127.0.0.1:8383/support")
	if err != nil {
		log.Println("Выполнить запрос не вышло")
		log.Println(err)

		return []int{}
	}
	if resp.StatusCode != 200 {
		var DataSupport []SupportData
		log.Println(DataSupport)
		return []int{}
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []int{}
	}
	sup := make([]SupportData, 0)
	err = json.Unmarshal(body, &sup)
	if err != nil {
		return []int{}
	}
	itog := SupportItog(sup)
	return itog

}

func SupportItog(supports []SupportData) []int {
	Itog := make([]int, 0)
	amountActiveTickets := 0
	for _, elem := range supports {
		amountActiveTickets += elem.ActiveTickets
	}
	loading := 0
	if amountActiveTickets < 9 {
		loading = 1
	} else if amountActiveTickets >= 9 && amountActiveTickets <= 16 {
		loading = 2
	} else {
		loading = 3
	}
	Itog = append(Itog, loading)
	averageTime := amountActiveTickets * 60 / 18
	Itog = append(Itog, averageTime)
	return Itog
}
