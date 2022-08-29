package app

import (
	"github.com/gorilla/mux"
	"github.com/sandeepchugh/banking/domain"
	"github.com/sandeepchugh/banking/service"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	// wiring
	//ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	//routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
