package email

import (
	"bufio"
	"encoding/csv"
	"io"
	check "itog/internal/check"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

func Email() map[string][][]EmailData {
	var EmailItog []EmailData
	csvFile, err := os.Open("email.data")
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
		if !check.CheckProvidersEmail(sl[1]) {
			continue
		}
		if len(sl) != 3 {
			continue
		}
		c, er := strconv.Atoi(sl[2])
		if err != nil {
			log.Println(er)
		}
		var s EmailData = EmailData{sl[0], sl[1], c}
		EmailItog = append(EmailItog, s)
	}
	countrysMap := make(map[string][]EmailData)
	countrysMaptwo := make(map[string][]EmailData)
	for i := 0; i < len(EmailItog); i++ {
		cercleCountry := EmailItog[i].Country
		countrysMap[cercleCountry] = append(countrysMap[cercleCountry], EmailItog[i])
		countrysMaptwo[cercleCountry] = append(countrysMaptwo[cercleCountry], EmailItog[i])
	}
	newMap := make(map[string][][]EmailData)
	for k, z := range countrysMap {
		sort.SliceStable(z, func(i, j int) bool {
			return z[i].DeliveryTime > z[j].DeliveryTime
		})
		sort.SliceStable(countrysMaptwo[k], func(i, j int) bool {
			return countrysMaptwo[k][i].DeliveryTime < countrysMaptwo[k][j].DeliveryTime
		})
		newMap[k] = [][]EmailData{z[:3], countrysMaptwo[k][:3]}
	}
	return newMap
}
