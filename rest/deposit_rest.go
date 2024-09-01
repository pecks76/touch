package rest

import (
	"encoding/json"
	"net/http"
	"restservice/service/receipt"
	"strconv"
)

type DepositRestService interface {
	GET(w http.ResponseWriter, r *http.Request)
}

type depositRestService struct {
	depositService receipt.DepositService
}

func NewDepositRestService(depositService receipt.DepositService) DepositRestService {
	return &depositRestService{depositService: depositService}
}

func (drs depositRestService) GET(w http.ResponseWriter, r *http.Request) {

	// todo: check that this is a GET

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid deposit Id", http.StatusBadRequest)
		return
	}

	depositMsg := drs.depositService.GetDepositReport(id)
	if depositMsg.Id != 0  {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(depositMsg)
	} else {
		http.Error(w, "Failed to find deposit", http.StatusNotFound)
	}
}

