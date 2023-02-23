package sms

import (
	"bufio"
	"encoding/csv"
	"io"
	check "itog/internal/check"
	"log"
	"os"
	"sort"
	"strings"
)

type SMSData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

func Sms() [][]SMSData {
	var SmsItogCountry []SMSData
	var SmsItogProvider []SMSData
	csvFile, err := os.Open("sms.data")
	if err != nil {
		log.Println("Не удается прочитать входной файл ", err)
	}
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal("err!=nil ", error)
		}
		sl := strings.Split(line[0], ";")
		if !check.CheckСountry(sl[0]) {
			continue
		}
		if !check.CheckProviders(sl[3]) {
			continue
		}
		if len(sl) != 4 {
			continue
		}
		var s SMSData = SMSData{sl[0], sl[1], sl[2], sl[3]}
		SmsItogCountry = append(SmsItogCountry, s)
		SmsItogProvider = append(SmsItogProvider, s)
	}
	for i, c := range SmsItogCountry {
		SmsItogCountry[i].Country = check.Countries[c.Country]
		SmsItogProvider[i].Country = check.Countries[c.Country]
	}
	sort.Slice(SmsItogCountry, func(i int, j int) bool {
		return SmsItogCountry[i].Country < SmsItogCountry[j].Country
	})
	sort.Slice(SmsItogProvider, func(i int, j int) bool {
		return SmsItogProvider[i].Provider < SmsItogProvider[j].Provider
	})
	result := [][]SMSData{SmsItogProvider, SmsItogCountry}
	return result
}
