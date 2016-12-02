package message_handler

import (
	"fmt"
	"strconv"

	auth_model "github.com/bborbe/auth/model"
	"github.com/bborbe/bot_agent/api"
	"github.com/bborbe/bot_agent/command"
	"github.com/bborbe/bot_agent/response"
	"github.com/bborbe/bot_agent_bamboo/bamboo"
	"github.com/bborbe/bot_agent_bamboo/model"
	"github.com/golang/glog"
)

type hasRequiredGroups func(authToken auth_model.AuthToken) bool

type handler struct {
	deployer          bamboo.Deployer
	hasRequiredGroups hasRequiredGroups
	command           command.Command
}

func New(
	prefix model.Prefix,
	deployer bamboo.Deployer,
	hasRequiredGroups hasRequiredGroups,
) *handler {
	d := new(handler)
	d.deployer = deployer
	d.hasRequiredGroups = hasRequiredGroups
	d.command = command.New(prefix.String(), "[ID]")
	return d
}

func (h *handler) Match(request *api.Request) bool {
	return h.command.MatchRequest(request) && h.hasRequiredGroups(request.AuthToken)
}

func (h *handler) Help(request *api.Request) []string {
	return []string{h.command.Help()}
}

func (h *handler) HandleMessage(request *api.Request) ([]*api.Response, error) {
	glog.V(3).Infof("handle deploy command")
	id, err := h.command.Parameter(request, "[ID]")
	if err != nil {
		glog.V(3).Infof("parse command failed: %v", err)
		return nil, err
	}
	if err := h.deploy(id); err != nil {
		return response.CreateReponseMessage(fmt.Sprintf("trigger deployment failed: %v", err)), nil
	}
	glog.V(2).Infof("return response")
	return response.CreateReponseMessage("deployment triggered succcesful"), nil
}

func (h *handler) deploy(idString string) error {
	id, err := strconv.Atoi(idString)
	if err != nil {
		glog.V(1).Infof("parse deployment id failed: %v", err)
		return err
	}
	if err = h.deployer.Deploy(id); err != nil {
		glog.V(1).Infof("deploy failed: %v", err)
		return err
	}
	glog.V(2).Infof("success")
	return nil
}
