package rest

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"restservice/msg"
	"restservice/service/receipt"
)

type RequestPayload struct {
	msg.Receipt `json:"receipt"`
}

type ReceiptRestService interface {
	POST(w http.ResponseWriter, r *http.Request)
}

type receiptRestService struct {
	receiptService receipt.ReceiptService
}

func NewReceiptRestService(receiptService receipt.ReceiptService) ReceiptRestService {
	return &receiptRestService{receiptService: receiptService}
}

func (rrs receiptRestService) POST(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var payload RequestPayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	log.Printf("Received receipt: %+v", payload.Receipt)
	err = rrs.receiptService.SaveReceipt(payload.Receipt)
	if err != nil {
		http.Error(w, "Failed to save receipt", http.StatusBadRequest)
	}
}
