package message_handler

import (
	"fmt"

	"github.com/bborbe/bot_agent/api"
	"github.com/bborbe/bot_agent/response"
	"github.com/golang/glog"
)

type bambooAgent struct {
}

func New() *bambooAgent {
	return new(bambooAgent)
}

func (h *bambooAgent) HandleMessage(request *api.Request) ([]*api.Response, error) {
	glog.V(2).Infof("handle message for token: %v", request.Id)
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
