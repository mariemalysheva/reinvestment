package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/handler"
	v3 "github.com/swaggest/swgui/v3"
	"log"
	"net/http"
)

const apiName = "Tokenized Reinvestment"

// Swagger struct
type Swagger struct {
	JSON string
}

func NewAPI(impl *handler.Implementation) *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/client/salt/{client_id}", impl.GetClientSalt).Methods(http.MethodGet)
	r.HandleFunc("/verify/salt/{client_id}", impl.VerifyClientSalt).Methods(http.MethodGet)
	r.HandleFunc("/clients", impl.GetClients).Methods(http.MethodGet)
	r.HandleFunc("/client/records/{client_id}", impl.GetClientRecords).Methods(http.MethodGet)
	r.HandleFunc("/clients/add", impl.PostAddClient).Methods(http.MethodPost)
	r.HandleFunc("/clients/transfer", impl.PostTransferClients).Methods(http.MethodPost)
	r.HandleFunc("/price", impl.PostSetPrice).Methods(http.MethodPost)
	r.HandleFunc("/rate", impl.PostSetRate).Methods(http.MethodPost)
	r.HandleFunc("/reschedule", impl.PostRescheduleReinvestment).Methods(http.MethodPost)
	r.HandleFunc("/reinvestments", impl.GetReinvestmentPeriods).Methods(http.MethodGet)

	fs := http.FileServer(http.Dir("./internal/router/json"))
	http.Handle("/swagger.json", fs)
	http.Handle("/", v3.NewHandler(apiName, "/swagger.json", "/docs"))

	go func() {
		appPort := 3001
		log.Printf("starting swagger at port %d", appPort)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", appPort), nil); err != nil {
			log.Fatalf("failed to launch swagger: %v", err)
		}
	}()

	return r
}
