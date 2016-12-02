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
	projects, err := d.listProjects()
	if err != nil {
		glog.V(1).Infof("list projects failed: %v", err)
		return err
	}
	if len(projects) == 0 {
		glog.V(1).Infof("project list is empty")
		return fmt.Errorf("project list is empty")
	}
	project := projects[0]
	versions, err := d.listVersions(project.Id)
	if err != nil {
		glog.V(1).Infof("list versions failed: %v", err)
		return err
	}
	if len(versions) == 0 {
		glog.V(1).Infof("version list is empty")
		return fmt.Errorf("version list is empty")
	}
	version := versions[0]
	err = d.deploy(project.Id, version.Id)
	if err != nil {
		glog.V(1).Infof("deploy failed: %v", err)
		return err
	}
	glog.V(2).Infof("deploy completed")
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
		glog.V(1).Infof("list projects failed: %v", err)
		return nil, err
	}
	return data, nil
}

type versions struct {
	Versions []version `json:"versions"`
}

type version struct {
	Id int `json:"id"`
}

func (d *deployer) listVersions(projectId int) ([]version, error) {
	var data versions
	url := fmt.Sprintf("%s/rest/api/latest/deploy/project/%d/versions", d.bambooUrl, projectId)
	err := d.rest.Call(url, nil, http.MethodGet, nil, &data, d.header())
	if err != nil {
		glog.V(1).Infof("list versions failed: %v", err)
		return nil, err
	}
	return data.Versions, nil
}

func (d *deployer) deploy(projectId int, versionId int) error {
	url := fmt.Sprintf("%s/rest/api/latest/queue/deployment/?environmentId=%d&versionId=%d", d.bambooUrl, projectId, versionId)
	err := d.rest.Call(url, nil, http.MethodPost, nil, nil, d.header())
	if err != nil {
		glog.V(1).Infof("deploy failed: %v", err)
		return err
	}
	return nil
}
