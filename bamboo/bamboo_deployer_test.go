package bamboo

import (
	"os"
	"testing"

	"bytes"
	"fmt"
	. "github.com/bborbe/assert"
	"github.com/bborbe/http/rest"
	"github.com/golang/glog"
	"io"
	"io/ioutil"
	"net/http"
)

func TestMain(m *testing.M) {
	exit := m.Run()
	glog.Flush()
	os.Exit(exit)
}

func TestImplementsDeployer(t *testing.T) {
	c := NewDeployer(nil, "http://example.com", "user", "pass")
	var i *Deployer
	if err := AssertThat(c, Implements(i)); err != nil {
		t.Fatal(err)
	}
}

func TestCreateAuth(t *testing.T) {
	c := NewDeployer(nil, "http://example.com", "user", "pass")
	if err := AssertThat(c.header(), NotNilValue()); err != nil {
		t.Fatal(err)
	}
}

const listProjectsJson = `[
  {
    "id": 2588673,
    "oid": 1154328879490400300,
    "key": {
      "key": "2588673"
    },
    "name": "Deploy Develop",
    "planKey": {
      "key": "TELU-TELUB"
    },
    "description": "",
    "environments": [
      {
        "id": 2719745,
        "key": {
          "key": "2588673-2719745"
        },
        "name": "Staging",
        "description": "",
        "deploymentProjectId": 2588673,
        "operations": {
          "canView": true,
          "canEdit": false,
          "canDelete": false,
          "allowedToExecute": false,
          "canExecute": false,
          "allowedToCreateVersion": false,
          "allowedToSetVersionStatus": false
        },
        "position": 0,
        "configurationState": "TASKED"
      }
    ],
    "operations": {
      "canView": true,
      "canEdit": false,
      "canDelete": false,
      "allowedToExecute": false,
      "canExecute": false,
      "allowedToCreateVersion": false,
      "allowedToSetVersionStatus": false
    }
  }
]`

func TestListProjectsFailed(t *testing.T) {
	deployer := NewDeployer(rest.New(func(req *http.Request) (resp *http.Response, err error) {
		return nil, fmt.Errorf("request failed")
	}), "http://example.com", "user", "pass")
	list, err := deployer.listProjects()
	if err := AssertThat(err, NotNilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(len(list), Is(0)); err != nil {
		t.Fatal(err)
	}
}

func TestListProjectsSuccess(t *testing.T) {
	deployer := NewDeployer(rest.New(func(req *http.Request) (resp *http.Response, err error) {
		return &http.Response{
			StatusCode: 200,
			Body:       createBody(listProjectsJson),
		}, nil
	}), "http://example.com", "user", "pass")
	list, err := deployer.listProjects()
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(len(list), Is(1)); err != nil {
		t.Fatal(err)
	}
}

func createBody(body string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewBufferString(body))
}
