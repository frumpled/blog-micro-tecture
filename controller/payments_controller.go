package controller

import (
	controller_model "app/controller/model"
	"app/service"

	"log"

	"net/http"
)

type pamentsController struct {
	service service.PaymentService
}

func (c pamentsController) processPayment(w http.ResponseWriter, r *http.Request) {
	log.Println("addEntry request received")

	request, err := decodeRequest(r)
	if err != nil {
		log.Println("Failed to decode response")
		encodeResponse(w, http.StatusBadRequest, nil)
		return
	}

	processPaymentRequest, ok := request.(controller_model.ProcessPaymentRequest)
	if !ok {
		log.Println("Bad request")
		encodeResponse(w, http.StatusBadRequest, nil)
		return
	}

	chargeID, err := c.service.ProcessPayment(processPaymentRequest)
	if err != nil {
		log.Println("Error: " + err.Error())
		encodeResponse(w, http.StatusInternalServerError, nil)
		return
	}

	response := struct {
		ChargeID string `json:"charge_id"`
	}{
		ChargeID: chargeID,
	}

	encodeResponse(w, http.StatusOK, response)
}
