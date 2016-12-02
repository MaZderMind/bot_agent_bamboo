package bamboo

import (
	"os"
	"regexp"
	"testing"

	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	. "github.com/bborbe/assert"
	"github.com/bborbe/http/rest"
	"github.com/golang/glog"
)

func createBody(body string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewBufferString(body))
}

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

const versionsJson = `{
  "size": 15,
  "versions": [
    {
      "id": 4685835,
      "name": "release-15",
      "creationDate": 1470752775471,
      "items": [
        {
          "id": 4751371,
          "name": "WAR",
          "planResultKey": {
            "key": "TELU-TELUB-34",
            "entityKey": {
              "key": "TELU-TELUB"
            },
            "resultNumber": 34
          },
          "type": "BAM_ARTIFACT",
          "label": "WAR",
          "location": "",
          "copyPattern": "**/*.war",
          "size": 20735428,
          "artifact": {
            "id": 4194331,
            "label": "WAR",
            "size": 20735428,
            "isSharedArtifact": true,
            "isGloballyStored": false,
            "linkType": "com.atlassian.bamboo.plugin.artifact.handler.local:ServerLocalArtifactHandler",
            "planResultKey": {
              "key": "TELU-TELUB-34",
              "entityKey": {
                "key": "TELU-TELUB"
              },
              "resultNumber": 34
            },
            "archiverType": "NONE"
          }
        }
      ],
      "operations": {
        "canView": false,
        "canEdit": false,
        "canDelete": false,
        "allowedToExecute": false,
        "canExecute": false,
        "allowedToCreateVersion": false,
        "allowedToSetVersionStatus": false
      },
      "planBranchName": "develop",
      "ageZeroPoint": 1470752779087
    },
    {
      "id": 4685834,
      "name": "release-14",
      "creationDate": 1467811308708,
      "items": [
        {
          "id": 4751370,
          "name": "WAR",
          "planResultKey": {
            "key": "TELU-TELUB-33",
            "entityKey": {
              "key": "TELU-TELUB"
            },
            "resultNumber": 33
          },
          "type": "BAM_ARTIFACT",
          "label": "WAR",
          "location": "",
          "copyPattern": "**/*.war",
          "size": 20735426,
          "artifact": {
            "id": 4194330,
            "label": "WAR",
            "size": 20735426,
            "isSharedArtifact": true,
            "isGloballyStored": false,
            "linkType": "com.atlassian.bamboo.plugin.artifact.handler.local:ServerLocalArtifactHandler",
            "planResultKey": {
              "key": "TELU-TELUB-33",
              "entityKey": {
                "key": "TELU-TELUB"
              },
              "resultNumber": 33
            },
            "archiverType": "NONE"
          }
        }
      ],
      "operations": {
        "canView": false,
        "canEdit": false,
        "canDelete": false,
        "allowedToExecute": false,
        "canExecute": false,
        "allowedToCreateVersion": false,
        "allowedToSetVersionStatus": false
      },
      "planBranchName": "develop",
      "ageZeroPoint": 1467811310157
    },
    {
      "id": 4685833,
      "name": "release-13",
      "creationDate": 1465549061666,
      "items": [
        {
          "id": 4751369,
          "name": "WAR",
          "planResultKey": {
            "key": "TELU-TELUB-32",
            "entityKey": {
              "key": "TELU-TELUB"
            },
            "resultNumber": 32
          },
          "type": "BAM_ARTIFACT",
          "label": "WAR",
          "location": "",
          "copyPattern": "**/*.war",
          "size": 20735428,
          "artifact": {
            "id": 4194329,
            "label": "WAR",
            "size": 20735428,
            "isSharedArtifact": true,
            "isGloballyStored": false,
            "linkType": "com.atlassian.bamboo.plugin.artifact.handler.local:ServerLocalArtifactHandler",
            "planResultKey": {
              "key": "TELU-TELUB-32",
              "entityKey": {
                "key": "TELU-TELUB"
              },
              "resultNumber": 32
            },
            "archiverType": "NONE"
          }
        }
      ],
      "operations": {
        "canView": false,
        "canEdit": false,
        "canDelete": false,
        "allowedToExecute": false,
        "canExecute": false,
        "allowedToCreateVersion": false,
        "allowedToSetVersionStatus": false
      },
      "planBranchName": "develop",
      "ageZeroPoint": 1465549061764
    },
    {
      "id": 4685832,
      "name": "release-12",
      "creationDate": 1465548634218,
      "items": [
        {
          "id": 4751368,
          "name": "WAR",
          "planResultKey": {
            "key": "TELU-TELUB-31",
            "entityKey": {
              "key": "TELU-TELUB"
            },
            "resultNumber": 31
          },
          "type": "BAM_ARTIFACT",
          "label": "WAR",
          "location": "",
          "copyPattern": "**/*.war",
          "size": 20735425,
          "artifact": {
            "id": 4194328,
            "label": "WAR",
            "size": 20735425,
            "isSharedArtifact": true,
            "isGloballyStored": false,
            "linkType": "com.atlassian.bamboo.plugin.artifact.handler.local:ServerLocalArtifactHandler",
            "planResultKey": {
              "key": "TELU-TELUB-31",
              "entityKey": {
                "key": "TELU-TELUB"
              },
              "resultNumber": 31
            },
            "archiverType": "NONE"
          }
        }
      ],
      "operations": {
        "canView": false,
        "canEdit": false,
        "canDelete": false,
        "allowedToExecute": false,
        "canExecute": false,
        "allowedToCreateVersion": false,
        "allowedToSetVersionStatus": false
      },
      "planBranchName": "develop",
      "ageZeroPoint": 1465548634351
    },
    {
      "id": 4685831,
      "name": "release-11",
      "creationDate": 1465488263703,
      "items": [
        {
          "id": 4751367,
          "name": "WAR",
          "planResultKey": {
            "key": "TELU-TELUB-30",
            "entityKey": {
              "key": "TELU-TELUB"
            },
            "resultNumber": 30
          },
          "type": "BAM_ARTIFACT",
          "label": "WAR",
          "location": "",
          "copyPattern": "**/*.war",
          "size": 20735429,
          "artifact": {
            "id": 4194327,
            "label": "WAR",
            "size": 20735429,
            "isSharedArtifact": true,
            "isGloballyStored": false,
            "linkType": "com.atlassian.bamboo.plugin.artifact.handler.local:ServerLocalArtifactHandler",
            "planResultKey": {
              "key": "TELU-TELUB-30",
              "entityKey": {
                "key": "TELU-TELUB"
              },
              "resultNumber": 30
            },
            "archiverType": "NONE"
          }
        }
      ],
      "operations": {
        "canView": false,
        "canEdit": false,
        "canDelete": false,
        "allowedToExecute": false,
        "canExecute": false,
        "allowedToCreateVersion": false,
        "allowedToSetVersionStatus": false
      },
      "planBranchName": "develop",
      "ageZeroPoint": 1465488263825
    },
    {
      "id": 4685830,
      "name": "release-10",
      "creationDate": 1465292609403,
      "items": [
        {
          "id": 4751366,
          "name": "WAR",
          "planResultKey": {
            "key": "TELU-TELUB-29",
            "entityKey": {
              "key": "TELU-TELUB"
            },
            "resultNumber": 29
          },
          "type": "BAM_ARTIFACT",
          "label": "WAR",
          "location": "",
          "copyPattern": "**/*.war",
          "size": 20735425,
          "artifact": {
            "id": 4194326,
            "label": "WAR",
            "size": 20735425,
            "isSharedArtifact": true,
            "isGloballyStored": false,
            "linkType": "com.atlassian.bamboo.plugin.artifact.handler.local:ServerLocalArtifactHandler",
            "planResultKey": {
              "key": "TELU-TELUB-29",
              "entityKey": {
                "key": "TELU-TELUB"
              },
              "resultNumber": 29
            },
            "archiverType": "NONE"
          }
        }
      ],
      "operations": {
        "canView": false,
        "canEdit": false,
        "canDelete": false,
        "allowedToExecute": false,
        "canExecute": false,
        "allowedToCreateVersion": false,
        "allowedToSetVersionStatus": false
      },
      "planBranchName": "develop",
      "ageZeroPoint": 1465292609570
    },
    {
      "id": 4685829,
      "name": "release-9",
      "creationDate": 1464776573957,
      "items": [
        {
          "id": 4751365,
          "name": "WAR",
          "planResultKey": {
            "key": "TELU-TELUB-28",
            "entityKey": {
              "key": "TELU-TELUB"
            },
            "resultNumber": 28
          },
          "type": "BAM_ARTIFACT",
          "label": "WAR",
          "location": "",
          "copyPattern": "**/*.war",
          "size": 20735423,
          "artifact": {
            "id": 4194325,
            "label": "WAR",
            "size": 20735423,
            "isSharedArtifact": true,
            "isGloballyStored": false,
            "linkType": "com.atlassian.bamboo.plugin.artifact.handler.local:ServerLocalArtifactHandler",
            "planResultKey": {
              "key": "TELU-TELUB-28",
              "entityKey": {
                "key": "TELU-TELUB"
              },
              "resultNumber": 28
            },
            "archiverType": "NONE"
          }
        }
      ],
      "operations": {
        "canView": false,
        "canEdit": false,
        "canDelete": false,
        "allowedToExecute": false,
        "canExecute": false,
        "allowedToCreateVersion": false,
        "allowedToSetVersionStatus": false
      },
      "planBranchName": "develop",
      "ageZeroPoint": 1464776574104
    },
    {
      "id": 4685828,
      "name": "release-8",
      "creationDate": 1464776196346,
      "items": [
        {
          "id": 4751364,
          "name": "WAR",
          "planResultKey": {
            "key": "TELU-TELUB-25",
            "entityKey": {
              "key": "TELU-TELUB"
            },
            "resultNumber": 25
          },
          "type": "BAM_ARTIFACT",
          "label": "WAR",
          "location": "",
          "copyPattern": "**/*.war",
          "size": 20735423,
          "artifact": {
            "id": 4194324,
            "label": "WAR",
            "size": 20735423,
            "isSharedArtifact": true,
            "isGloballyStored": false,
            "linkType": "com.atlassian.bamboo.plugin.artifact.handler.local:ServerLocalArtifactHandler",
            "planResultKey": {
              "key": "TELU-TELUB-25",
              "entityKey": {
                "key": "TELU-TELUB"
              },
              "resultNumber": 25
            },
            "archiverType": "NONE"
          }
        }
      ],
      "operations": {
        "canView": false,
        "canEdit": false,
        "canDelete": false,
        "allowedToExecute": false,
        "canExecute": false,
        "allowedToCreateVersion": false,
        "allowedToSetVersionStatus": false
      },
      "planBranchName": "develop",
      "ageZeroPoint": 1464776196479
    },
    {
      "id": 4685827,
      "name": "release-7",
      "creationDate": 1463734142113,
      "items": [
        {
          "id": 4751363,
          "name": "WAR",
          "planResultKey": {
            "key": "TELU-TELUB-23",
            "entityKey": {
              "key": "TELU-TELUB"
            },
            "resultNumber": 23
          },
          "type": "BAM_ARTIFACT",
          "label": "WAR",
          "location": "",
          "copyPattern": "**/*.war",
          "size": 20735422,
          "artifact": {
            "id": 4194322,
            "label": "WAR",
            "size": 20735422,
            "isSharedArtifact": true,
            "isGloballyStored": false,
            "linkType": "com.atlassian.bamboo.plugin.artifact.handler.local:ServerLocalArtifactHandler",
            "planResultKey": {
              "key": "TELU-TELUB-23",
              "entityKey": {
                "key": "TELU-TELUB"
              },
              "resultNumber": 23
            },
            "archiverType": "NONE"
          }
        }
      ],
      "operations": {
        "canView": false,
        "canEdit": false,
        "canDelete": false,
        "allowedToExecute": false,
        "canExecute": false,
        "allowedToCreateVersion": false,
        "allowedToSetVersionStatus": false
      },
      "planBranchName": "develop",
      "ageZeroPoint": 1463734142259
    },
    {
      "id": 4685826,
      "name": "release-6",
      "creationDate": 1463733473227,
      "items": [
        {
          "id": 4751362,
          "name": "WAR",
          "planResultKey": {
            "key": "TELU-TELUB-22",
            "entityKey": {
              "key": "TELU-TELUB"
            },
            "resultNumber": 22
          },
          "type": "BAM_ARTIFACT",
          "label": "WAR",
          "location": "",
          "copyPattern": "**/*.war",
          "size": 20735069,
          "artifact": {
            "id": 4194320,
            "label": "WAR",
            "size": 20735069,
            "isSharedArtifact": true,
            "isGloballyStored": false,
            "linkType": "com.atlassian.bamboo.plugin.artifact.handler.local:ServerLocalArtifactHandler",
            "planResultKey": {
              "key": "TELU-TELUB-22",
              "entityKey": {
                "key": "TELU-TELUB"
              },
              "resultNumber": 22
            },
            "archiverType": "NONE"
          }
        }
      ],
      "operations": {
        "canView": false,
        "canEdit": false,
        "canDelete": false,
        "allowedToExecute": false,
        "canExecute": false,
        "allowedToCreateVersion": false,
        "allowedToSetVersionStatus": false
      },
      "planBranchName": "develop",
      "ageZeroPoint": 1463733473329
    },
    {
      "id": 4685825,
      "name": "release-5",
      "creationDate": 1463732870082,
      "items": [
        {
          "id": 4751361,
          "name": "WAR",
          "planResultKey": {
            "key": "TELU-TELUB-21",
            "entityKey": {
              "key": "TELU-TELUB"
            },
            "resultNumber": 21
          },
          "type": "BAM_ARTIFACT",
          "label": "WAR",
          "location": "",
          "copyPattern": "**/*.war",
          "size": 20735063,
          "artifact": {
            "id": 4194318,
            "label": "WAR",
            "size": 20735063,
            "isSharedArtifact": true,
            "isGloballyStored": false,
            "linkType": "com.atlassian.bamboo.plugin.artifact.handler.local:ServerLocalArtifactHandler",
            "planResultKey": {
              "key": "TELU-TELUB-21",
              "entityKey": {
                "key": "TELU-TELUB"
              },
              "resultNumber": 21
            },
            "archiverType": "NONE"
          }
        }
      ],
      "operations": {
        "canView": false,
        "canEdit": false,
        "canDelete": false,
        "allowedToExecute": false,
        "canExecute": false,
        "allowedToCreateVersion": false,
        "allowedToSetVersionStatus": false
      },
      "planBranchName": "develop",
      "ageZeroPoint": 1463732870744
    },
    {
      "id": 2785284,
      "name": "release-4",
      "creationDate": 1461332084411,
      "items": [
        {
          "id": 2850820,
          "name": "WAR",
          "planResultKey": {
            "key": "TELU-TELUB-20",
            "entityKey": {
              "key": "TELU-TELUB"
            },
            "resultNumber": 20
          },
          "type": "BAM_ARTIFACT",
          "label": "WAR",
          "location": "",
          "copyPattern": "**/*.war",
          "size": 3150062,
          "artifact": {
            "id": 2031631,
            "label": "WAR",
            "size": 3150062,
            "isSharedArtifact": true,
            "isGloballyStored": false,
            "linkType": "com.atlassian.bamboo.plugin.artifact.handler.local:ServerLocalArtifactHandler",
            "planResultKey": {
              "key": "TELU-TELUB-20",
              "entityKey": {
                "key": "TELU-TELUB"
              },
              "resultNumber": 20
            },
            "archiverType": "NONE"
          }
        }
      ],
      "operations": {
        "canView": false,
        "canEdit": false,
        "canDelete": false,
        "allowedToExecute": false,
        "canExecute": false,
        "allowedToCreateVersion": false,
        "allowedToSetVersionStatus": false
      },
      "planBranchName": "develop",
      "ageZeroPoint": 1461332084554
    },
    {
      "id": 2785283,
      "name": "release-3",
      "creationDate": 1461325244768,
      "items": [
        {
          "id": 2850819,
          "name": "WAR",
          "planResultKey": {
            "key": "TELU-TELUB-19",
            "entityKey": {
              "key": "TELU-TELUB"
            },
            "resultNumber": 19
          },
          "type": "BAM_ARTIFACT",
          "label": "WAR",
          "location": "",
          "copyPattern": "**/*.war",
          "size": 1060320,
          "artifact": {
            "id": 2031628,
            "label": "WAR",
            "size": 1060320,
            "isSharedArtifact": true,
            "isGloballyStored": false,
            "linkType": "com.atlassian.bamboo.plugin.artifact.handler.local:ServerLocalArtifactHandler",
            "planResultKey": {
              "key": "TELU-TELUB-19",
              "entityKey": {
                "key": "TELU-TELUB"
              },
              "resultNumber": 19
            },
            "archiverType": "NONE"
          }
        }
      ],
      "operations": {
        "canView": false,
        "canEdit": false,
        "canDelete": false,
        "allowedToExecute": false,
        "canExecute": false,
        "allowedToCreateVersion": false,
        "allowedToSetVersionStatus": false
      },
      "planBranchName": "develop",
      "ageZeroPoint": 1461325244901
    },
    {
      "id": 2785282,
      "name": "release-2",
      "creationDate": 1461324703433,
      "items": [
        {
          "id": 2850818,
          "name": "WAR",
          "planResultKey": {
            "key": "TELU-TELUB-18",
            "entityKey": {
              "key": "TELU-TELUB"
            },
            "resultNumber": 18
          },
          "type": "BAM_ARTIFACT",
          "label": "WAR",
          "location": "",
          "copyPattern": "**/*.war",
          "size": 1060323,
          "artifact": {
            "id": 2031627,
            "label": "WAR",
            "size": 1060323,
            "isSharedArtifact": true,
            "isGloballyStored": false,
            "linkType": "com.atlassian.bamboo.plugin.artifact.handler.local:ServerLocalArtifactHandler",
            "planResultKey": {
              "key": "TELU-TELUB-18",
              "entityKey": {
                "key": "TELU-TELUB"
              },
              "resultNumber": 18
            },
            "archiverType": "NONE"
          }
        }
      ],
      "operations": {
        "canView": false,
        "canEdit": false,
        "canDelete": false,
        "allowedToExecute": false,
        "canExecute": false,
        "allowedToCreateVersion": false,
        "allowedToSetVersionStatus": false
      },
      "planBranchName": "develop",
      "ageZeroPoint": 1461325037434
    },
    {
      "id": 2785281,
      "name": "release-1",
      "creationDate": 1461324483571,
      "items": [
        {
          "id": 2850817,
          "name": "WAR",
          "planResultKey": {
            "key": "TELU-TELUB-17",
            "entityKey": {
              "key": "TELU-TELUB"
            },
            "resultNumber": 17
          },
          "type": "BAM_ARTIFACT",
          "label": "WAR",
          "location": "",
          "copyPattern": "**/*.war",
          "size": 1060276,
          "artifact": {
            "id": 2031626,
            "label": "WAR",
            "size": 1060276,
            "isSharedArtifact": true,
            "isGloballyStored": false,
            "linkType": "com.atlassian.bamboo.plugin.artifact.handler.local:ServerLocalArtifactHandler",
            "planResultKey": {
              "key": "TELU-TELUB-17",
              "entityKey": {
                "key": "TELU-TELUB"
              },
              "resultNumber": 17
            },
            "archiverType": "NONE"
          }
        }
      ],
      "operations": {
        "canView": false,
        "canEdit": false,
        "canDelete": false,
        "allowedToExecute": false,
        "canExecute": false,
        "allowedToCreateVersion": false,
        "allowedToSetVersionStatus": false
      },
      "planBranchName": "develop",
      "ageZeroPoint": 1461324484413
    }
  ],
  "start-index": 0,
  "max-result": 15
}`

func TestListVersionsFailed(t *testing.T) {
	deployer := NewDeployer(rest.New(func(req *http.Request) (resp *http.Response, err error) {
		return nil, fmt.Errorf("request failed")
	}), "http://example.com", "user", "pass")
	list, err := deployer.listVersions(2588673)
	if err := AssertThat(err, NotNilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(len(list), Is(0)); err != nil {
		t.Fatal(err)
	}
}

func TestListVersionsSuccess(t *testing.T) {
	deployer := NewDeployer(rest.New(func(req *http.Request) (resp *http.Response, err error) {
		return &http.Response{
			StatusCode: 200,
			Body:       createBody(versionsJson),
		}, nil
	}), "http://example.com", "user", "pass")
	list, err := deployer.listVersions(2588673)
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(len(list), Is(15)); err != nil {
		t.Fatal(err)
	}
}

const environmentsJson = `{
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
}`

func TestListEnvironmentsFailed(t *testing.T) {
	deployer := NewDeployer(rest.New(func(req *http.Request) (resp *http.Response, err error) {
		return nil, fmt.Errorf("request failed")
	}), "http://example.com", "user", "pass")
	list, err := deployer.listEnvironments(2588673)
	if err := AssertThat(err, NotNilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(len(list), Is(0)); err != nil {
		t.Fatal(err)
	}
}

func TestListEnvironmentsSuccess(t *testing.T) {
	deployer := NewDeployer(rest.New(func(req *http.Request) (resp *http.Response, err error) {
		return &http.Response{
			StatusCode: 200,
			Body:       createBody(environmentsJson),
		}, nil
	}), "http://example.com", "user", "pass")
	list, err := deployer.listEnvironments(2588673)
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(len(list), Is(1)); err != nil {
		t.Fatal(err)
	}
}

func TestFilterProject(t *testing.T) {
	d := NewDeployer(rest.New(func(req *http.Request) (resp *http.Response, err error) {
		return &http.Response{
			StatusCode: 200,
			Body:       createBody(listProjectsJson),
		}, nil
	}), "http://example.com", "user", "pass")

	selected, err := d.selectProject("Deploy Develop")
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(selected, NotNilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(selected.Id, Is(2588673)); err != nil {
		t.Fatal(err)
	}
}

func TestEnqueeDeploy(t *testing.T) {
	listProjectsRequestCount := 0
	versionsRequestCount := 0
	environmentsRequestCount := 0
	queueDeploymentRequestCount := 0

	deployer := NewDeployer(rest.New(func(req *http.Request) (resp *http.Response, err error) {
		glog.Infof("!!!!! mock request for %s", req.URL.Path)

		body := ""

		if req.URL.Path == "/rest/api/latest/deploy/project/all" {
			body = listProjectsJson
			listProjectsRequestCount++
		} else if match, err := regexp.MatchString("^/rest/api/latest/deploy/project/[0-9]+/versions$", req.URL.Path); match {
			if err != nil {
				t.Fatal(err)
			}
			body = versionsJson
			versionsRequestCount++
		} else if match, err := regexp.MatchString("^/rest/api/latest/deploy/project/[0-9]+$", req.URL.Path); match {
			if err != nil {
				t.Fatal(err)
			}
			body = environmentsJson
			environmentsRequestCount++
		} else if req.URL.Path == "/rest/api/latest/queue/deployment/" {
			body = ""
			queueDeploymentRequestCount++
		} else {
			t.Fatal(fmt.Errorf("Unexpected Request-URL %s", req.URL))
		}

		return &http.Response{
			StatusCode: 200,
			Body:       createBody(body),
		}, nil
	}), "http://example.com", "user", "pass")

	err := deployer.Deploy("Deploy Develop", "Staging")
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(listProjectsRequestCount, Is(1)); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(versionsRequestCount, Is(1)); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(environmentsRequestCount, Is(1)); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(queueDeploymentRequestCount, Is(1)); err != nil {
		t.Fatal(err)
	}
}
