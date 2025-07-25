// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/deployments": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deployments"
                ],
                "responses": {
                    "200": {
                        "description": "List of deployments",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request method",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/deployments/{account}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deployments"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Account ID",
                        "name": "account",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of deployments",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request method",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deployments"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Account ID",
                        "name": "account",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Deployment details",
                        "name": "deployment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/codedeploy.CreateDeploymentInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/codedeploy.CreateDeploymentOutput"
                        }
                    },
                    "400": {
                        "description": "Error decoding request body",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "codedeploy.CreateDeploymentInput": {
            "type": "object",
            "properties": {
                "applicationName": {
                    "description": "The name of an CodeDeploy application associated with the user or Amazon Web\nServices account.\n\nThis member is required.",
                    "type": "string"
                },
                "autoRollbackConfiguration": {
                    "description": "Configuration information for an automatic rollback that is added when a\ndeployment is created.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/types.AutoRollbackConfiguration"
                        }
                    ]
                },
                "deploymentConfigName": {
                    "description": "The name of a deployment configuration associated with the user or Amazon Web\nServices account.\n\nIf not specified, the value configured in the deployment group is used as the\ndefault. If the deployment group does not have a deployment configuration\nassociated with it, CodeDeployDefault . OneAtATime is used by default.",
                    "type": "string"
                },
                "deploymentGroupName": {
                    "description": "The name of the deployment group.",
                    "type": "string"
                },
                "description": {
                    "description": "A comment about the deployment.",
                    "type": "string"
                },
                "fileExistsBehavior": {
                    "description": "Information about how CodeDeploy handles files that already exist in a\ndeployment target location but weren't part of the previous successful\ndeployment.\n\nThe fileExistsBehavior parameter takes any of the following values:\n\n  - DISALLOW: The deployment fails. This is also the default behavior if no\n  option is specified.\n\n  - OVERWRITE: The version of the file from the application revision currently\n  being deployed replaces the version already on the instance.\n\n  - RETAIN: The version of the file already on the instance is kept and used as\n  part of the new deployment.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/types.FileExistsBehavior"
                        }
                    ]
                },
                "ignoreApplicationStopFailures": {
                    "description": "If true, then if an ApplicationStop , BeforeBlockTraffic , or AfterBlockTraffic\ndeployment lifecycle event to an instance fails, then the deployment continues\nto the next deployment lifecycle event. For example, if ApplicationStop fails,\nthe deployment continues with DownloadBundle . If BeforeBlockTraffic fails, the\ndeployment continues with BlockTraffic . If AfterBlockTraffic fails, the\ndeployment continues with ApplicationStop .\n\nIf false or not specified, then if a lifecycle event fails during a deployment\nto an instance, that deployment fails. If deployment to that instance is part of\nan overall deployment and the number of healthy hosts is not less than the\nminimum number of healthy hosts, then a deployment to the next instance is\nattempted.\n\nDuring a deployment, the CodeDeploy agent runs the scripts specified for\nApplicationStop , BeforeBlockTraffic , and AfterBlockTraffic in the AppSpec\nfile from the previous successful deployment. (All other scripts are run from\nthe AppSpec file in the current deployment.) If one of these scripts contains an\nerror and does not run successfully, the deployment can fail.\n\nIf the cause of the failure is a script from the last successful deployment\nthat will never run successfully, create a new deployment and use\nignoreApplicationStopFailures to specify that the ApplicationStop ,\nBeforeBlockTraffic , and AfterBlockTraffic failures should be ignored.",
                    "type": "boolean"
                },
                "overrideAlarmConfiguration": {
                    "description": "Allows you to specify information about alarms associated with a deployment.\nThe alarm configuration that you specify here will override the alarm\nconfiguration at the deployment group level. Consider overriding the alarm\nconfiguration if you have set up alarms at the deployment group level that are\ncausing deployment failures. In this case, you would call CreateDeployment to\ncreate a new deployment that uses a previous application revision that is known\nto work, and set its alarm configuration to turn off alarm polling. Turning off\nalarm polling ensures that the new deployment proceeds without being blocked by\nthe alarm that was generated by the previous, failed, deployment.\n\nIf you specify an overrideAlarmConfiguration , you need the\nUpdateDeploymentGroup IAM permission when calling CreateDeployment .",
                    "allOf": [
                        {
                            "$ref": "#/definitions/types.AlarmConfiguration"
                        }
                    ]
                },
                "revision": {
                    "description": "The type and location of the revision to deploy.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/types.RevisionLocation"
                        }
                    ]
                },
                "targetInstances": {
                    "description": "Information about the instances that belong to the replacement environment in\na blue/green deployment.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/types.TargetInstances"
                        }
                    ]
                },
                "updateOutdatedInstancesOnly": {
                    "description": "Indicates whether to deploy to all instances or only to instances that are not\nrunning the latest application revision.",
                    "type": "boolean"
                }
            }
        },
        "codedeploy.CreateDeploymentOutput": {
            "type": "object",
            "properties": {
                "deploymentId": {
                    "description": "The unique ID of a deployment.",
                    "type": "string"
                },
                "resultMetadata": {
                    "description": "Metadata pertaining to the operation's result.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/middleware.Metadata"
                        }
                    ]
                }
            }
        },
        "middleware.Metadata": {
            "type": "object"
        },
        "types.Alarm": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "The name of the alarm. Maximum length is 255 characters. Each alarm name can be\nused only once in a list of alarms.",
                    "type": "string"
                }
            }
        },
        "types.AlarmConfiguration": {
            "type": "object",
            "properties": {
                "alarms": {
                    "description": "A list of alarms configured for the deployment or deployment group. A maximum\nof 10 alarms can be added.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Alarm"
                    }
                },
                "enabled": {
                    "description": "Indicates whether the alarm configuration is enabled.",
                    "type": "boolean"
                },
                "ignorePollAlarmFailure": {
                    "description": "Indicates whether a deployment should continue if information about the current\nstate of alarms cannot be retrieved from Amazon CloudWatch. The default value is\nfalse.\n\n  - true : The deployment proceeds even if alarm status information can't be\n  retrieved from Amazon CloudWatch.\n\n  - false : The deployment stops if alarm status information can't be retrieved\n  from Amazon CloudWatch.",
                    "type": "boolean"
                }
            }
        },
        "types.AppSpecContent": {
            "type": "object",
            "properties": {
                "content": {
                    "description": "The YAML-formatted or JSON-formatted revision string.\n\nFor an Lambda deployment, the content includes a Lambda function name, the\nalias for its original version, and the alias for its replacement version. The\ndeployment shifts traffic from the original version of the Lambda function to\nthe replacement version.\n\nFor an Amazon ECS deployment, the content includes the task name, information\nabout the load balancer that serves traffic to the container, and more.\n\nFor both types of deployments, the content can specify Lambda functions that\nrun at specified hooks, such as BeforeInstall , during a deployment.",
                    "type": "string"
                },
                "sha256": {
                    "description": "The SHA256 hash value of the revision content.",
                    "type": "string"
                }
            }
        },
        "types.AutoRollbackConfiguration": {
            "type": "object",
            "properties": {
                "enabled": {
                    "description": "Indicates whether a defined automatic rollback configuration is currently\nenabled.",
                    "type": "boolean"
                },
                "events": {
                    "description": "The event type or types that trigger a rollback.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.AutoRollbackEvent"
                    }
                }
            }
        },
        "types.AutoRollbackEvent": {
            "type": "string",
            "enum": [
                "DEPLOYMENT_FAILURE",
                "DEPLOYMENT_STOP_ON_ALARM",
                "DEPLOYMENT_STOP_ON_REQUEST"
            ],
            "x-enum-varnames": [
                "AutoRollbackEventDeploymentFailure",
                "AutoRollbackEventDeploymentStopOnAlarm",
                "AutoRollbackEventDeploymentStopOnRequest"
            ]
        },
        "types.BundleType": {
            "type": "string",
            "enum": [
                "tar",
                "tgz",
                "zip",
                "YAML",
                "JSON"
            ],
            "x-enum-varnames": [
                "BundleTypeTar",
                "BundleTypeTarGZip",
                "BundleTypeZip",
                "BundleTypeYaml",
                "BundleTypeJson"
            ]
        },
        "types.EC2TagFilter": {
            "type": "object",
            "properties": {
                "key": {
                    "description": "The tag filter key.",
                    "type": "string"
                },
                "type": {
                    "description": "The tag filter type:\n\n  - KEY_ONLY : Key only.\n\n  - VALUE_ONLY : Value only.\n\n  - KEY_AND_VALUE : Key and value.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/types.EC2TagFilterType"
                        }
                    ]
                },
                "value": {
                    "description": "The tag filter value.",
                    "type": "string"
                }
            }
        },
        "types.EC2TagFilterType": {
            "type": "string",
            "enum": [
                "KEY_ONLY",
                "VALUE_ONLY",
                "KEY_AND_VALUE"
            ],
            "x-enum-varnames": [
                "EC2TagFilterTypeKeyOnly",
                "EC2TagFilterTypeValueOnly",
                "EC2TagFilterTypeKeyAndValue"
            ]
        },
        "types.EC2TagSet": {
            "type": "object",
            "properties": {
                "ec2TagSetList": {
                    "description": "A list that contains other lists of Amazon EC2 instance tag groups. For an\ninstance to be included in the deployment group, it must be identified by all of\nthe tag groups in the list.",
                    "type": "array",
                    "items": {
                        "type": "array",
                        "items": {
                            "$ref": "#/definitions/types.EC2TagFilter"
                        }
                    }
                }
            }
        },
        "types.FileExistsBehavior": {
            "type": "string",
            "enum": [
                "DISALLOW",
                "OVERWRITE",
                "RETAIN"
            ],
            "x-enum-varnames": [
                "FileExistsBehaviorDisallow",
                "FileExistsBehaviorOverwrite",
                "FileExistsBehaviorRetain"
            ]
        },
        "types.GitHubLocation": {
            "type": "object",
            "properties": {
                "commitId": {
                    "description": "The SHA1 commit ID of the GitHub commit that represents the bundled artifacts\nfor the application revision.",
                    "type": "string"
                },
                "repository": {
                    "description": "The GitHub account and repository pair that stores a reference to the commit\nthat represents the bundled artifacts for the application revision.\n\nSpecified as account/repository.",
                    "type": "string"
                }
            }
        },
        "types.RawString": {
            "type": "object",
            "properties": {
                "content": {
                    "description": "The YAML-formatted or JSON-formatted revision string. It includes information\nabout which Lambda function to update and optional Lambda functions that\nvalidate deployment lifecycle events.",
                    "type": "string"
                },
                "sha256": {
                    "description": "The SHA256 hash value of the revision content.",
                    "type": "string"
                }
            }
        },
        "types.RevisionLocation": {
            "type": "object",
            "properties": {
                "appSpecContent": {
                    "description": "The content of an AppSpec file for an Lambda or Amazon ECS deployment. The\ncontent is formatted as JSON or YAML and stored as a RawString.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/types.AppSpecContent"
                        }
                    ]
                },
                "gitHubLocation": {
                    "description": "Information about the location of application artifacts stored in GitHub.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/types.GitHubLocation"
                        }
                    ]
                },
                "revisionType": {
                    "description": "The type of application revision:\n\n  - S3: An application revision stored in Amazon S3.\n\n  - GitHub: An application revision stored in GitHub (EC2/On-premises\n  deployments only).\n\n  - String: A YAML-formatted or JSON-formatted string (Lambda deployments only).\n\n  - AppSpecContent: An AppSpecContent object that contains the contents of an\n  AppSpec file for an Lambda or Amazon ECS deployment. The content is formatted as\n  JSON or YAML stored as a RawString.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/types.RevisionLocationType"
                        }
                    ]
                },
                "s3Location": {
                    "description": "Information about the location of a revision stored in Amazon S3.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/types.S3Location"
                        }
                    ]
                },
                "string_": {
                    "description": "Information about the location of an Lambda deployment revision stored as a\nRawString.\n\nDeprecated: RawString and String revision type are deprecated, use\nAppSpecContent type instead.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/types.RawString"
                        }
                    ]
                }
            }
        },
        "types.RevisionLocationType": {
            "type": "string",
            "enum": [
                "S3",
                "GitHub",
                "String",
                "AppSpecContent"
            ],
            "x-enum-varnames": [
                "RevisionLocationTypeS3",
                "RevisionLocationTypeGitHub",
                "RevisionLocationTypeString",
                "RevisionLocationTypeAppSpecContent"
            ]
        },
        "types.S3Location": {
            "type": "object",
            "properties": {
                "bucket": {
                    "description": "The name of the Amazon S3 bucket where the application revision is stored.",
                    "type": "string"
                },
                "bundleType": {
                    "description": "The file type of the application revision. Must be one of the following:\n\n  - tar : A tar archive file.\n\n  - tgz : A compressed tar archive file.\n\n  - zip : A zip archive file.\n\n  - YAML : A YAML-formatted file.\n\n  - JSON : A JSON-formatted file.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/types.BundleType"
                        }
                    ]
                },
                "etag": {
                    "description": "The ETag of the Amazon S3 object that represents the bundled artifacts for the\napplication revision.\n\nIf the ETag is not specified as an input parameter, ETag validation of the\nobject is skipped.",
                    "type": "string"
                },
                "key": {
                    "description": "The name of the Amazon S3 object that represents the bundled artifacts for the\napplication revision.",
                    "type": "string"
                },
                "version": {
                    "description": "A specific version of the Amazon S3 object that represents the bundled\nartifacts for the application revision.\n\nIf the version is not specified, the system uses the most recent version by\ndefault.",
                    "type": "string"
                }
            }
        },
        "types.TargetInstances": {
            "type": "object",
            "properties": {
                "autoScalingGroups": {
                    "description": "The names of one or more Auto Scaling groups to identify a replacement\nenvironment for a blue/green deployment.",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "ec2TagSet": {
                    "description": "Information about the groups of Amazon EC2 instance tags that an instance must\nbe identified by in order for it to be included in the replacement environment\nfor a blue/green deployment. Cannot be used in the same call as tagFilters .",
                    "allOf": [
                        {
                            "$ref": "#/definitions/types.EC2TagSet"
                        }
                    ]
                },
                "tagFilters": {
                    "description": "The tag filter key, type, and value used to identify Amazon EC2 instances in a\nreplacement environment for a blue/green deployment. Cannot be used in the same\ncall as ec2TagSet .",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.EC2TagFilter"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "DeployView API",
	Description:      "This is the API documentation for DeployView, a deployment management tool.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
