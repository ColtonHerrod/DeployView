package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"deployview-backend/utilities"

	"github.com/gorilla/mux"
)

func AllDeploymentsHandler(w http.ResponseWriter, r *http.Request) {
	accounts := strings.Split(os.Getenv("AWS_ACCOUNTS"), ",")
	for _, account := range accounts {
		deployments := utilities.GetAccountDeployments(account)
		fmt.Fprintf(w, "Deployments for account %s:\n", account)
		for _, deployment := range deployments {
			fmt.Fprintf(w, "%s\n", deployment)
		}
	}
}

func DeploymentsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	account := vars["account"]

	log.Printf("Handling deployments for account: %s", account)

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	deployments := utilities.GetAccountDeployments(account)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deployments for account %s:\n", account)
	for _, deployment := range deployments {
		fmt.Fprintf(w, "%s\n", deployment)
	}
	log.Printf("Handled deployments for account: %s", account)
}
