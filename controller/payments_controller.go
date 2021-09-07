package controller

import (
	"errors"
	"log"

	"app/controller/model"
	"app/service"
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

	paymentRequest, ok := request.(model.CreatePaymentRequest)
	if !ok {
		log.Println("Bad request")
		encodeResponse(w, http.StatusBadRequest, nil)
		return
	}

	err = validateRequest(paymentRequest)
	if err != nil {
		log.Println("Error: " + err.Error())
		encodeResponse(w, http.StatusBadRequest, nil)
		return
	}

	paymentIntentID, err := c.service.ProcessPayment(
		paymentRequest.Amount,
		paymentRequest.StripeToken,
	)
	if err != nil {
		log.Println("Error: " + err.Error())
		encodeResponse(w, http.StatusInternalServerError, nil)
		return
	}

	response := struct {
		PaymentIntentID string `json:"payment_intent_id"`
	}{
		PaymentIntentID: paymentIntentID,
	}

	encodeResponse(w, http.StatusOK, response)
}

func validateRequest(c model.CreatePaymentRequest) error {
	if c.Amount <= 0 {
		return errors.New("no amount given")
	}

	if c.StripeToken == "" {
		return errors.New("no stripe token")
	}

	return nil
}
