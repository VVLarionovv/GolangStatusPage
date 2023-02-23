package handler

import (
	"encoding/json"
	"itog/internal/billing"
	"itog/internal/email"
	"itog/internal/incident"
	"itog/internal/mms"
	"itog/internal/sms"
	"itog/internal/support"
	"itog/internal/voice"
	"net/http"

	"github.com/gorilla/mux"
)

type ResultT struct {
	Status bool       `json:"status"`
	Data   ResultSetT `json:"data"`
	Error  string     `json:"error"`
}

type ResultSetT struct {
	SMS       [][]sms.SMSData                `json:"sms"`
	MMS       [][]mms.MMSData                `json:"mms"`
	VoiceCall []voice.VoiceCallData          `json:"voice_call"`
	Email     map[string][][]email.EmailData `json:"email"`
	Billing   billing.BillingData            `json:"billing"`
	Support   []int                          `json:"support"`
	Incidents []incident.IncidentData        `json:"incident"`
}

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/api", handleConnection).Methods("GET")
	http.ListenAndServe("127.0.0.1:8282", router)
}
func handleConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	result := getResultData()
	var itogResult ResultT
	if result.CheckResultSetT() {
		itogResult.Status = true
		itogResult.Data = result
	} else {
		itogResult.Status = false
		itogResult.Error = "Ошибка при сборе данных"
	}
	byteResultAns, err := json.Marshal(itogResult)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(byteResultAns)
}
func getResultData() ResultSetT {

	smsValye := sms.Sms()

	mmsValye := mms.Mms()

	voiceValye := voice.Voice()

	emailValye := email.Email()

	billingValye := billing.Billing()

	supportValye := support.Support()

	incident := incident.Incident()

	var ResultSetTItog ResultSetT = ResultSetT{smsValye, mmsValye, voiceValye, emailValye,
		billingValye, supportValye, incident}
	return ResultSetTItog
}
func (r ResultSetT) CheckResultSetT() bool {
	var b bool
	if r.SMS == nil {
		return b
	} else if r.MMS == nil {
		return b
	} else if r.VoiceCall == nil {
		return b
	} else if r.Email == nil {
		return b
	} else if r.Support == nil {
		return b
	} else if r.Incidents == nil {
		return b
	}
	b = true
	return b
}
