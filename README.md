# Terratest with GCP example

This example uses Terratest to apply Terraform resources and then asserts that the creation was correct using Google's client-go libraries.


Run tests

- Env var for Google Project must be set (Any of [GOOGLE_PROJECT GOOGLE_CLOUD_PROJECT GOOGLE_CLOUD_PROJECT_ID GCLOUD_PROJECT CLOUDSDK_CORE_PROJECT])

```
GOOGLE_PROJECT=my-gcp-existing-project go test -v
```