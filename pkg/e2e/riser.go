// +build e2e

package e2e

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

const deployTimeoutSeconds = 60

// deployOrFail calls riser deploy with --wait
func deployOrFail(t *testing.T, appDir, dockerTag, environment string) {
	deployArgsOrFail(t, appDir, dockerTag, environment)
}

func deployArgsOrFail(t *testing.T, appDir, dockerTag, environment string, args ...string) {
	// Wait one second less than the timeout so that we can get any possible error output from the deploy command before the command times out
	waitSeconds := deployTimeoutSeconds - 2
	joinedArgs := strings.Join(args, " ")
	riserOrFail(t, appDir, fmt.Sprintf("deploy %s %s --wait --wait-seconds=%d %s", dockerTag, environment, waitSeconds, joinedArgs))
}

func deleteDeploymentOrFail(t *testing.T, appDir, deploymentName, environment string) {
	riserOrFail(t, appDir, fmt.Sprintf("deployments delete %s %s --no-prompt", deploymentName, environment))
}

func riserOrFail(t *testing.T, appDir, command string) {
	shellOrFailTimeout(t, time.Duration(deployTimeoutSeconds)*time.Second, "cd %s && riser %s", appDir, command)
}
