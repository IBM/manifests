package metadata

import (
	"github.com/kubeflow/manifests/tests"
	"testing"
)

func TestKustomize(t *testing.T) {
	testCase := &tests.KustomizeTestCase{
		Package:  "../../../../../stacks/openshift/components/metadata",
		Expected: "test_data/expected",
	}

	tests.RunTestCase(t, testCase)
}
