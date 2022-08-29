package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sandeepchugh/banking/domain"
	"github.com/sandeepchugh/banking/service"
	"log"
	"net/http"
	"os"
)

func Start() {

	router := mux.NewRouter()

	// wiring
	//ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	//routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
