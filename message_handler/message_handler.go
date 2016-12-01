package message_handler

import (
	"fmt"

	"github.com/bborbe/bot_agent/api"
	"github.com/bborbe/bot_agent/response"
	"github.com/bborbe/bot_agent_bamboo/bamboo"
	"github.com/golang/glog"
)

type bambooAgent struct {
	deployer bamboo.Deployer
}

func New(deployer bamboo.Deployer) *bambooAgent {
	d := new(bambooAgent)
	d.deployer = deployer
	return d
}

func (h *bambooAgent) HandleMessage(request *api.Request) ([]*api.Response, error) {
	glog.V(2).Infof("handle message for token: %v", request.Id)

	if glog.V(4) {
		glog.Infof("request %+v", request)
	}

	if request.Message != fmt.Sprintf("bamboo %s", request.Bot) {
		glog.V(2).Infof("message contains no bamboo => skip")
		return nil, nil
	}
	if request.From == nil {
		glog.V(2).Infof("from is empty => skip")
		return nil, nil
	}
	glog.V(2).Infof("return response")
	return response.CreateReponseMessage(fmt.Sprintf("bamboo %s", request.From.UserName)), nil
}
