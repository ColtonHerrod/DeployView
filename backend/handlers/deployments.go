package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"deployview-backend/utilities"

	"github.com/aws/aws-sdk-go-v2/service/codedeploy"
	"encoding/json"
	"github.com/gorilla/mux"
)

// @Summary: AllDeploymentsHandler handles requests for all deployments across accounts.
// @Description: Retrieves and displays deployments for all accounts specified in the AWS_ACCOUNTS environment variable.
// @Tags deployments
// @Accept json
// @Produce json
// @Success 200 {object} map[string][]string "Map of account IDs to deployment IDs"
// @Failure 400 {string} string "Invalid request method"
// @Router /deployments [get]
func AllDeploymentsHandler(w http.ResponseWriter, r *http.Request) {
	accounts := strings.Split(os.Getenv("AWS_ACCOUNTS"), ",")
	deployments := make(map[string][]string)
	for _, account := range accounts {
		account_deployments, _ := utilities.GetAccountDeployments(account)
		if account_deployments != nil {
			deployments[account] = account_deployments.Deployments
		}
	}
	json.NewEncoder(w).Encode(deployments)
}

// @Summary: DeploymentsHandler handles requests for deployments of a specific account.
// @Description: Retrieves and displays deployments for a specific account passed as a URL parameter.
// @Tags deployments
// @Accept json
// @Produce json
// @Param account path string true "Account ID"
// @Success 200 {array} string "List of deployment IDs for the specified account"
// @Failure 400 {string} string "Invalid request method"
// @Router /deployments/{account} [get]
func AccountDeploymentsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	account := vars["account"]

	log.Printf("Handling deployments for account: %s", account)

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	accountDeployments, err := utilities.GetAccountDeployments(account)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving deployments: %v", err), http.StatusInternalServerError)
		return
	}
	deployments := accountDeployments.Deployments
	json.NewEncoder(w).Encode(deployments)
}

// @Summary: CreateDeploymentHandler handles requests to create a deployment for a specific account.
// @Description: Creates a deployment for a specific account passed as a URL parameter.
// @Tags deployments
// @Accept json
// @Produce json
// @Param account path string true "Account ID"
// @Param deployment body codedeploy.CreateDeploymentInput true "Deployment details"
// @Success 201 {object} codedeploy.CreateDeploymentOutput
// @Failure 400 {string} string "Error decoding request body"
// @Router /deployments/{account} [post]
func CreateDeploymentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	account := vars["account"]
	
	input := codedeploy.CreateDeploymentInput{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}
	utilities.CreateDeployment(account, input)
}