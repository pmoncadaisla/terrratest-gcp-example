package test

import (
	"context"
	"log"
	"testing"

	"cloud.google.com/go/bigquery"
	"github.com/gruntwork-io/terratest/modules/gcp"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformHelloWorldExample(t *testing.T) {

	t.Parallel()

	projectID := gcp.GetGoogleProjectIDFromEnvVar(t)

	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "examples/example1",
		EnvVars: map[string]string{
			"GOOGLE_CLOUD_PROJECT": projectID,
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("bigquery.NewClient: %v", err)
	}
	defer client.Close()

	dataset := client.Dataset("example_dataset")
	datasetMetadata, _ := dataset.Metadata(ctx)

	assert.Equal(t, "test", datasetMetadata.Name)
	assert.Equal(t, "example_dataset", dataset.DatasetID)
	assert.Equal(t, "test", datasetMetadata.Name)
	assert.Contains(t, datasetMetadata.Labels, "env")
	assert.Equal(t, "default", datasetMetadata.Labels["env"])

}
