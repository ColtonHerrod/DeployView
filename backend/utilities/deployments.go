package utilities

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codedeploy"
)

func GetAccountDeployments(account string) (*codedeploy.ListDeploymentsOutput, error) {
	ctx := context.TODO()
	cfg := GetCredentials(account)
	
	cdSvc := codedeploy.NewFromConfig(cfg)

	return cdSvc.ListDeployments(ctx, &codedeploy.ListDeploymentsInput{})
}

func CreateDeployment(account string, input codedeploy.CreateDeploymentInput) (*codedeploy.CreateDeploymentOutput, error) {
	ctx := context.TODO()
	cfg := GetCredentials(account)

	cdSvc := codedeploy.NewFromConfig(cfg)

	return cdSvc.CreateDeployment(ctx, &input)
}