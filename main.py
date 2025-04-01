from django.http import JsonResponse, HttpResponseRedirect
from django.views.decorators.csrf import csrf_exempt
from django.urls import path
import boto3
from workos import client as workos
import json

workos.api_key = "your_workos_api_key"
redirect_uri = "https://deployview.herrod.net"

# SSO Login view
def auth(request):
    authorization_url = workos.client.sso.get_authorization_url(
        client_id="your_workos_client_id",
        redirect_uri=redirect_uri,
        provider="GitHubOAuth",
    )
    return HttpResponseRedirect(authorization_url)

# SSO Callback view
def callback(request):
    try:
        code = request.GET.get("code")
        profile_and_token = workos.client.sso.get_profile_and_token(code=code)

        profile = profile_and_token["profile"]
        org_id = "org_test_idp"

        if profile["organization_id"] != org_id:
            return JsonResponse({"detail": "Unauthorized"}, status=403)

        # Use the profile for further business logic
        return HttpResponseRedirect("/")
    except Exception as e:
        print("Error during SSO callback:", e)
        return JsonResponse({"detail": "Internal Server Error"}, status=500)

# CodeDeploy view
@csrf_exempt
def codedeploy(request):
    if request.method != "POST":
        return JsonResponse({"detail": "Method not allowed"}, status=405)

    try:
        data = json.loads(request.body)
        account_id = data.get("accountId")
        role_arn = data.get("roleArn")
        application_name = data.get("applicationName")

        if not account_id or not role_arn or not application_name:
            return JsonResponse({"detail": "Missing required parameters"}, status=400)

        sts_client = boto3.client("sts", region_name="us-east-1")
        assumed_role = sts_client.assume_role(
            RoleArn=role_arn,
            RoleSessionName="DeployViewSession",
        )

        credentials = assumed_role["Credentials"]
        codedeploy_client = boto3.client(
            "codedeploy",
            region_name="us-east-1",
            aws_access_key_id=credentials["AccessKeyId"],
            aws_secret_access_key=credentials["SecretAccessKey"],
            aws_session_token=credentials["SessionToken"],
        )

        app_details = codedeploy_client.get_application(
            applicationName=application_name
        )
        return JsonResponse(app_details)
    except Exception as e:
        print("Error fetching CodeDeploy application:", e)
        return JsonResponse({"detail": "Error fetching CodeDeploy information"}, status=500)

# URL configuration
urlpatterns = [
    path("auth/", auth),
    path("callback/", callback),
    path("codedeploy/", codedeploy),
]
