package controller

import (
	"errors"
	"log"

	"app/service"
	"net/http"

	"github.com/stripe/stripe-go"
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

	stripeChargeParams, ok := request.(*stripe.ChargeParams)
	if !ok {
		log.Println("Bad request")
		encodeResponse(w, http.StatusBadRequest, nil)
		return
	}

	err = validateRequest(stripeChargeParams)
	if err != nil {
		log.Println("Error: " + err.Error())
		encodeResponse(w, http.StatusBadRequest, nil)
		return
	}

	chargeID, err := c.service.ProcessPayment(stripeChargeParams)
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

func validateRequest(c *stripe.ChargeParams) error {
	if c == nil {
		return errors.New("empty request")
	}

	if c.Amount == nil || *c.Amount < 0 {
		return errors.New("no amount given")
	}

	if c.Source.Token == nil || *c.Source.Token == "" {
		return errors.New("no source token")
	}

	if c.Description == nil || *c.Description == "" {
		return errors.New("description is empty")
	}

	return nil
}
