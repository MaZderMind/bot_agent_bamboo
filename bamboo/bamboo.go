package bamboo

import (
	"github.com/bborbe/bot_agent_bamboo/model"
	"github.com/golang/glog"
)

// Deployer can trigger a bamboo deployment
type Deployer interface {
	Deploy() error
}

type deployer struct {
	bambooUrl      model.BambooUrl
	bambooUsername model.BambooUsername
	bambooPassword model.BambooPassword
}

// NewDeployer returns a new instance of Deployer
func NewDeployer(
	bambooUrl model.BambooUrl,
	bambooUsername model.BambooUsername,
	bambooPassword model.BambooPassword,
) *deployer {
	d := new(deployer)
	d.bambooUrl = bambooUrl
	d.bambooUsername = bambooUsername
	d.bambooPassword = bambooPassword
	return d
}

func (d *deployer) Deploy() error {
	glog.V(4).Infof("deploy to url: %v with user: %v and pw-length: %d", d.bambooUrl, d.bambooUsername, len(d.bambooPassword))
	return nil
}
