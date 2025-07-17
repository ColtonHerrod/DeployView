package utilities

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func GetAccountDeployments(account string) ([]string) {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	stsSvc := sts.NewFromConfig(cfg)
	identity, err := stsSvc.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Account: %s, Arn: %s", aws.ToString(identity.Account), aws.ToString(identity.Arn))
	
	if aws.ToString(identity.Account) != account {
		roleArn := "arn:aws:iam::" + account + ":role/DeployViewRole"
		creds := stscreds.NewAssumeRoleProvider(stsSvc, roleArn, func(o *stscreds.AssumeRoleOptions) {
			o.RoleSessionName = "DeployViewSession"
		})
		cfg.Credentials = aws.NewCredentialsCache(creds)
	}
	
	cdSvc := codedeploy.NewFromConfig(cfg)

	deployments, err := cdSvc.ListDeployments(ctx, &codedeploy.ListDeploymentsInput{})

	var deploymentInfo []string

	if err != nil {
		log.Fatal(err)
	}
	for _, deployment := range deployments.Deployments {
		deployment, err := cdSvc.GetDeployment(ctx, &codedeploy.GetDeploymentInput{
			DeploymentId: aws.String(deployment),
		})
		if err != nil {
			log.Printf("Error getting deployment %s: %v", deployment, err)
			continue
		}
		deploymentInfo = append(deploymentInfo, fmt.Sprintf("Deployment ID: %s, Status: %s\n", deployment, deployment.DeploymentInfo.Status))
	}

	return deploymentInfo
}