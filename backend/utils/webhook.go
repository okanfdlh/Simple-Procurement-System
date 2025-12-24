package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type WebhookPayload struct {
	PurchasingID uint    `json:"purchasing_id"`
	GrandTotal   float64 `json:"grand_total"`
}

func SendWebhook(purchasingID uint, total float64) {
	payload := WebhookPayload{
		PurchasingID: purchasingID,
		GrandTotal:   total,
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(
		"POST",
		"https://example.com/webhook",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Webhook failed:", err)
		return
	}
	defer resp.Body.Close()
}
