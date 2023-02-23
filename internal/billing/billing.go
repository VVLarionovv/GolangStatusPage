package billing

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
)

type BillingData struct {
	CreateCustomer bool `json:"create_customer"`
	Purchase       bool `json:"purchase"`
	Payout         bool `json:"payout"`
	Recurring      bool `json:"recurring"`
	FraudControl   bool `json:"fraud_control"`
	CheckoutPage   bool `json:"checkout_page"`
}

func Billing() BillingData {
	csvFile, err := os.Open("billing.data")
	if err != nil {
		log.Println("Не удается прочитать входной файл ", err)
	}
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))

	line, error := reader.Read()
	if error != nil {
		log.Println(error, "nen")
	}
	bb := make([]bool, len(line[0]))
	for i, r := range line[0] {
		if r == '1' {
			bb[i] = true
		}
	}
	var BillingItog BillingData = BillingData{bb[5], bb[4], bb[3], bb[2], bb[1], bb[0]}
	return BillingItog
}
