package incident

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sort"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"`
}

func Incident() []IncidentData {
	resp, err := http.Get("http://127.0.0.1:8383/accendent")
	if err != nil {
		log.Println("Выполнить запрос не вышло")
		log.Println(err)
		return []IncidentData{}
	}

	if resp.StatusCode != 200 {
		var DataSupport []IncidentData
		log.Println(DataSupport)
		return []IncidentData{}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []IncidentData{}
	}

	incident := make([]IncidentData, 0)
	err = json.Unmarshal(body, &incident)
	if err != nil {
		return []IncidentData{}
	}
	sort.Slice(incident, func(i int, j int) bool {
		return incident[i].Status < incident[j].Status
	})
	return incident

}
