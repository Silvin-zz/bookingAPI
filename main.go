package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Silvin/booking/models"

	"github.com/Silvin/booking"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var bookingClient = booking.Booking{}

func GetClientEndPoint(w http.ResponseWriter, req *http.Request) {
	clients, _ := bookingClient.GetAllClients()
	fmt.Println(clients)
	json.NewEncoder(w).Encode(clients)

	// values := req.URL.Query() // Returns a url.Values, which is a map[string][]string
	// hola, _ := values["hola"]
	// dos, _ := values["dos"]
	// //params := mux.Vars(req) // solo aplica para el caso de delete o post o put
	// fmt.Println(hola)
	// fmt.Println(dos)
	// fmt.Println("saludos")

}

// ****************GET ****************

func GetEventEndPoint(w http.ResponseWriter, req *http.Request) {
	events, _ := bookingClient.GetAllEvents()
	fmt.Println(events)
	json.NewEncoder(w).Encode(events)

}

func GetPaymentTypeEndPoint(w http.ResponseWriter, req *http.Request) {
	paymentTypes, _ := bookingClient.GetAllPaymentsType()
	fmt.Println(paymentTypes)
	json.NewEncoder(w).Encode(paymentTypes)

}

func GetComissionEndPoint(w http.ResponseWriter, req *http.Request) {
	comissions, _ := bookingClient.GetAllComission()
	fmt.Println(comissions)
	json.NewEncoder(w).Encode(comissions)

}

// ************** POST *********************

func NewClientEndPoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println("saludos .......................................................................................")
	decoder := json.NewDecoder(req.Body)

	tmpClient := models.Client{}
	_ = decoder.Decode(&tmpClient)

	fmt.Println(tmpClient.Name)
	fmt.Println(tmpClient.Username)

	defaultComission, _ := bookingClient.GetDefaultComission()
	client, _ := bookingClient.AddClient(tmpClient.Name, tmpClient.Username, tmpClient.Password, defaultComission)
	json.NewEncoder(w).Encode(client)

}

func NewComissionEndPoint(w http.ResponseWriter, req *http.Request) {

	fmt.Println("saludos .......................................................................................")
	decoder := json.NewDecoder(req.Body)
	tmpComission := models.Comission{}
	_ = decoder.Decode(&tmpComission)
	fmt.Println(tmpComission.Name)
	fmt.Println(tmpComission.Value)
	comission, _ := bookingClient.AddCommission(tmpComission.Name, tmpComission.Value, tmpComission.IsPercent, tmpComission.IsDefault)
	json.NewEncoder(w).Encode(comission)

}

func NewEventEndPoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println("saludos .......................................................................................")
	decoder := json.NewDecoder(req.Body)

	tmpEvent := models.Event{}
	_ = decoder.Decode(&tmpEvent)

	fmt.Println(tmpEvent.Name)
	fmt.Println(tmpEvent.Comission.Id)

	event, _ := bookingClient.AddEvent(tmpEvent.Name, tmpEvent.Client_id, tmpEvent.Comission)
	json.NewEncoder(w).Encode(event)

}

func main() {
	bookingClient.Init("127.0.0.1:27017", "test")
	//bookingClient.RemoveDB() //Remove the database if exists

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Accept", "Origin", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"PUT", "PATCH", "POST", "DELETE", "GET", "OPTIONS"})

	// *** Routing
	router := mux.NewRouter()

	router.HandleFunc("/client", GetClientEndPoint).Methods("GET")
	router.HandleFunc("/event", GetEventEndPoint).Methods("GET")
	router.HandleFunc("/paymenttype", GetPaymentTypeEndPoint).Methods("GET")
	router.HandleFunc("/comission", GetComissionEndPoint).Methods("GET")

	router.HandleFunc("/comission", NewComissionEndPoint).Methods("POST")
	router.HandleFunc("/client", NewClientEndPoint).Methods("POST")
	router.HandleFunc("/event", NewEventEndPoint).Methods("POST")

	log.Fatal(http.ListenAndServe(":12345", handlers.CORS(originsOk, headersOk, methodsOk)(router)))

}
