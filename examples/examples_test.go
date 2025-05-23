// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getProjectId(t *testing.T) string {
	name := os.Getenv("OVH_CLOUD_PROJECT_SERVICE")
	if name == "" {
		t.Skipf("Skipping test due to missing OVH_CLOUD_PROJECT_SERVICE environment variable")
	}

	return name
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions(t *testing.T) integration.ProgramTestOptions {
	project_id := getProjectId(t)
	return integration.ProgramTestOptions{
		Config: map[string]string{
			"project_id": project_id,
		},
	}
}
