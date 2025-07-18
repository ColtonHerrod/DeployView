package utilities

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func GetCredentials(account string) aws.Config {
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

	return cfg
}