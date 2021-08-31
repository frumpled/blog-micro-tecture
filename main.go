package main

import (
	"app/client"
	"app/controller"
	"app/repository"
	"app/service"
)

func main() {
	paymentClient := client.NewPaymentClient("sk_test_4eC39HqLyjWDarjtT1zdp7dc")
	repository := repository.NewTransactionRepository()
	service := service.NewPaymentService(paymentClient, repository)
	controllers := controller.NewControllers(service)

	controllers.Start()
}
