package bamboo

import (
	"fmt"
	"net/http"

	"github.com/bborbe/bot_agent_bamboo/model"
	"github.com/bborbe/http/header"
	"github.com/bborbe/http/rest"
	"github.com/golang/glog"
)

type executeRequest func(req *http.Request) (resp *http.Response, err error)

// Deployer can trigger a bamboo deployment
type Deployer interface {
	Deploy(number int) error
}

type deployer struct {
	rest           rest.Rest
	bambooUrl      model.BambooUrl
	bambooUsername model.BambooUsername
	bambooPassword model.BambooPassword
}

// NewDeployer returns a new instance of Deployer
func NewDeployer(
	rest rest.Rest,
	bambooUrl model.BambooUrl,
	bambooUsername model.BambooUsername,
	bambooPassword model.BambooPassword,
) *deployer {
	d := new(deployer)
	d.rest = rest
	d.bambooUrl = bambooUrl
	d.bambooUsername = bambooUsername
	d.bambooPassword = bambooPassword
	return d
}

func (d *deployer) header() http.Header {
	h := make(http.Header)
	h.Add("Authorization", fmt.Sprintf("Basic %s", header.CreateAuthorizationToken(d.bambooUsername.String(), d.bambooPassword.String())))
	h.Add("Accepts", "application/json")
	return h
}

func (d *deployer) Deploy(number int) error {
	glog.V(4).Infof("deploy to url: %v with user: %v and pw-length: %d", d.bambooUrl, d.bambooUsername, len(d.bambooPassword))
	err := d.rest.Call(d.bambooUrl.String(), nil, http.MethodGet, nil, nil, d.header())
	if err != nil {
		glog.V(1).Infof("call bamboo failed: %v", err)
		return err
	}
	return nil
}

type project struct {
	Id int `json:"id"`
}

func (d *deployer) listProjects() ([]project, error) {
	var data []project
	url := fmt.Sprintf("%s/rest/api/latest/deploy/project/all", d.bambooUrl)
	err := d.rest.Call(url, nil, http.MethodGet, nil, &data, d.header())
	if err != nil {
		glog.V(1).Infof("call bamboo failed: %v", err)
		return nil, err
	}
	return data, nil
}
