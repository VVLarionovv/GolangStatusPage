package voice

import (
	"bufio"
	"encoding/csv"
	"io"
	check "itog/internal/check"
	"log"
	"os"
	"strconv"
	"strings"
)

type VoiceCallData struct {
	Country             string  `json:"country"`
	Bandwidth           string  `json:"bandwidth"`
	ResponseTime        string  `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoicePurity         int     `json:"voice_purity"`
	MedianOfCallsTime   int     `json:"median_of_calls_time"`
}

func Voice() []VoiceCallData {
	var voiceItog []VoiceCallData
	csvFile, err := os.Open("voice.data")
	if err != nil {
		log.Println("Не удается прочитать входной файл", err)
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
		if !check.CheckProvidersVoice(sl[3]) {
			continue
		}
		if len(sl) != 8 {
			continue
		}
		f, err := strconv.ParseFloat(sl[4], 32)
		if err != nil {
			log.Println(err)
		}
		k, er := strconv.Atoi(sl[5])
		if err != nil {
			log.Println(er)
		}
		z, er := strconv.Atoi(sl[6])
		if err != nil {
			log.Println(er)
		}
		c, er := strconv.Atoi(sl[7])
		if err != nil {
			log.Println(er)
		}
		var s VoiceCallData = VoiceCallData{sl[0], sl[1], sl[2], sl[3], float32(f), k, z, c}
		voiceItog = append(voiceItog, s)

	}
	return voiceItog

}
