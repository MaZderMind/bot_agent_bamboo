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
	Deploy(project, environment string) error
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

func filterProjects(vs []project, f func(project) bool) []project {
	vsf := make([]project, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func (d *deployer) selectProject(projectName string) (*project, error) {
	projects, err := d.listProjects()
	if err != nil {
		glog.V(1).Infof("list projects failed: %v", err)
		return nil, err
	}

	if len(projects) == 0 {
		glog.V(1).Infof("project list is empty")
		return nil, fmt.Errorf("project list is empty")
	}

	filtered := filterProjects(projects, func(theProject project) bool {
		return theProject.Name == projectName
	})

	if len(filtered) == 0 {
		return nil, fmt.Errorf("No Project named %s found (searched %d Projects)", projectName, len(projects))
	} else if len(filtered) > 1 {
		return nil, fmt.Errorf("More then 1 Project named %s found", projectName)
	}

	return &filtered[0], nil
}

func (d *deployer) Deploy(projectName, environmentName string) error {
	glog.V(4).Infof("deploy to url: %v with user: %v and pw-length: %d", d.bambooUrl, d.bambooUsername, len(d.bambooPassword))
	selectedProject, err := d.selectProject(projectName)
	if err != nil {
		glog.V(1).Infof("project selection failed: %v", err)
		return err
	}

	versions, err := d.listVersions(selectedProject.Id)
	if err != nil {
		glog.V(1).Infof("list versions failed: %v", err)
		return err
	}
	if len(versions) == 0 {
		glog.V(1).Infof("version list is empty")
		return fmt.Errorf("version list is empty")
	}
	version := versions[0]

	environments, err := d.listEnvironments(selectedProject.Id)
	if err != nil {
		glog.V(1).Infof("list environments failed: %v", err)
		return err
	}
	if len(environments) == 0 {
		glog.V(1).Infof("environment  list is empty")
		return fmt.Errorf("environment  list is empty")
	}
	environment := environments[0]

	err = d.deploy(environment.Id, version.Id)
	if err != nil {
		glog.V(1).Infof("deploy failed: %v", err)
		return err
	}
	glog.V(2).Infof("deploy completed")
	return nil
}

type project struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
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

type environments struct {
	Environments []environment `json:"environments"`
}

type environment struct {
	Id int `json:"id"`
}

func (d *deployer) listEnvironments(projectId int) ([]environment, error) {
	var data environments
	url := fmt.Sprintf("%s/rest/api/latest/deploy/project/%d", d.bambooUrl, projectId)
	err := d.rest.Call(url, nil, http.MethodGet, nil, &data, d.header())
	if err != nil {
		glog.V(1).Infof("list versions failed: %v", err)
		return nil, err
	}
	return data.Environments, nil
}
