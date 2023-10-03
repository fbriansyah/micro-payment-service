package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	dmbiller "github.com/fbriansyah/micro-payment-service/internal/application/domain/biller"
)

type PaymentResponse struct {
	Billing          dmbiller.Bill `json:"bill"`
	RefferenceNumber string        `json:"refference_number"`
	CreatedAt        time.Time     `json:"created_at"`
}

type InquiryResponse struct {
	Error   bool          `json:"error"`
	Message string        `json:"message"`
	Data    dmbiller.Bill `json:"data,omitempty"`
}

// Inquiry send request to biller endpoint using http
func (a *HttpAdapter) Inquiry(billNumber string) (InquiryResponse, error) {
	var requestData struct {
		BillNumber string `json:"bill_number"`
	}
	requestData.BillNumber = billNumber

	requestJson, err := json.Marshal(requestData)
	if err != nil {
		return InquiryResponse{}, err
	}

	request, err := http.NewRequest("POST", a.endpoint+"/inquiry", bytes.NewBuffer(requestJson))
	if err != nil {
		fmt.Println("NewRequest")
		return InquiryResponse{}, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := a.client.Do(request)
	if err != nil {
		return InquiryResponse{}, err
	}

	var responseJson InquiryResponse

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&responseJson)
	if err != nil {
		return InquiryResponse{}, err
	}

	if response.StatusCode != http.StatusOK {
		return InquiryResponse{}, errors.New(responseJson.Message)
	}

	if responseJson.Error {
		return InquiryResponse{}, errors.New(responseJson.Message)
	}

	return responseJson, nil
}
