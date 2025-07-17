package utilities

type RequestBody struct {
	AccountID       string `json:"accountId"`
	RoleArn         string `json:"roleArn"`
	ApplicationName string `json:"applicationName"`
}