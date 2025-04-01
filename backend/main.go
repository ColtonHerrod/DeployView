package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type RequestBody struct {
	AccountID       string `json:"accountId"`
	RoleArn         string `json:"roleArn"`
	ApplicationName string `json:"applicationName"`
}

func main() {
	http.HandleFunc("/codedeploy", handleCodeDeploy)
	port := 3000
	fmt.Printf("Server running on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func handleCodeDeploy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var body RequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if body.AccountID == "" || body.RoleArn == "" || body.ApplicationName == "" {
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	cfg, err := config.LoadDefaultConfig(r.Context(), config.WithRegion("us-east-1"))
	if err != nil {
		http.Error(w, "Failed to load AWS config", http.StatusInternalServerError)
		log.Println("Error loading AWS config:", err)
		return
	}

	// Assume role for cross-account access
	stsClient := sts.NewFromConfig(cfg)
	creds := stscreds.NewAssumeRoleProvider(stsClient, body.RoleArn)
	assumedCfg := aws.Config{
		Credentials: aws.NewCredentialsCache(creds),
		Region:      "us-east-1",
	}

	// Fetch CodeDeploy application details
	codedeployClient := codedeploy.NewFromConfig(assumedCfg)
	appDetails, err := codedeployClient.GetApplication(r.Context(), &codedeploy.GetApplicationInput{
		ApplicationName: aws.String(body.ApplicationName),
	})
	if err != nil {
		http.Error(w, "Error fetching CodeDeploy information", http.StatusInternalServerError)
		log.Println("Error fetching CodeDeploy application:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(appDetails); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		log.Println("Error encoding response:", err)
	}
}
