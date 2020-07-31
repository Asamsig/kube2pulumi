package yaml2pcl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t *testing.T) {
	testNamespace(t)
	testNamespaceComments(t)
	test1PodArray(t)
}

func testNamespace(t *testing.T) {
	assertion := assert.New(t)

	testData := `
apiVersion: v1
kind: Namespace
metadata:
  name: foo

`

	expected := `resource foo "kubernetes:core/v1:Namespace" {
apiVersion = "v1"
kind = "Namespace"
metadata = {
name = "foo"
}
}
`
	result, err := Convert([]byte(testData))
	if err != nil {
		assertion.Error(err)
	} else {
		assertion.Equal(expected, result, "Single resource conversion was incorrect")
	}
}

func testNamespaceComments(t *testing.T) {
	assertion := assert.New(t)

	testData := `
apiVersion: v1
kind: Namespace
# this is a test comment
metadata:
  name: foo

`

	expected := `resource foo "kubernetes:core/v1:Namespace" {
apiVersion = "v1"
kind = "Namespace"
# this is a test comment
metadata = {
name = "foo"
}
}
`
	result, err := Convert([]byte(testData))
	if err != nil {
		assertion.Error(err)
	} else {
		assertion.Equal(expected, result, "Comments are converted incorrectly")
	}
}

func test1PodArray(t *testing.T) {
	assertion := assert.New(t)

	testData := `
apiVersion: v1
kind: Pod
metadata:
  namespace: foo
  name: bar
spec:
  containers:
    - name: nginx
      image: nginx:1.14-alpine
      resources:
        limits:
          memory: 20Mi
          cpu: 0.2
`

	expected := `resource bar "kubernetes:core/v1:Pod" {
apiVersion = "v1"
kind = "Pod"
metadata = {
namespace = "foo"
name = "bar"
}
spec = {
containers = [
{
name = "nginx"
image = "nginx:1.14-alpine"
resources = {
limits = {
memory = "20Mi"
cpu = 0.2
}
}
}
]
}
}
`
	result, err := Convert([]byte(testData))
	if err != nil {
		assertion.Error(err)
	} else {
		assertion.Equal(expected, result, "Nested array is converted incorrectly")
	}
}