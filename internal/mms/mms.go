package mms

import (
	"encoding/json"
	"io"
	check "itog/internal/check"
	"log"
	"net/http"
	"sort"
)

type MMSData struct {
	Country string `json:"country"`

	Provider string `json:"provider"`

	Bandwidth string `json:"bandwidth"`

	ResponseTime string `json:"response_time"`
}

func Mms() [][]MMSData {
	resp, err := http.Get("http://127.0.0.1:8383/mms")
	if err != nil {
		log.Println("Выполнить запрос не вышло")
		log.Println(err)

		return [][]MMSData{}
	}
	if resp.StatusCode != 200 {
		var mmsData []MMSData
		log.Println(mmsData)
		return [][]MMSData{}
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return [][]MMSData{}
	}

	data := make([]MMSData, 0)
	mmsCountry := make([]MMSData, 0)
	mmsProvider := make([]MMSData, 0)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return [][]MMSData{}
	}
	for _, x := range data {
		if !check.CheckСountry(x.Country) {
			continue
		}
		if !check.CheckProviders(x.Provider) {
			continue
		}
		mmsCountry = append(mmsCountry, x)
		mmsProvider = append(mmsProvider, x)
	}
	for i, c := range mmsCountry {
		mmsCountry[i].Country = check.Countries[c.Country]
		mmsProvider[i].Country = check.Countries[c.Country]
	}
	sort.Slice(mmsCountry, func(i int, j int) bool {
		return mmsCountry[i].Country < mmsCountry[j].Country
	})
	sort.Slice(mmsProvider, func(i int, j int) bool {
		return mmsProvider[i].Provider < mmsProvider[j].Provider
	})
	result := [][]MMSData{mmsCountry, mmsProvider}
	return result
}
