package message_handler

import (
	"fmt"
	"strings"

	"flag"

	"github.com/bborbe/bot_agent/api"
	"github.com/bborbe/bot_agent/response"
	"github.com/golang/glog"
)

func init() {
	flag.Set("v", "4")
	flag.Set("logtostderr", "true")
}

type message []string

func NewMessage(message_string string) message {
	words := strings.Split(message_string, " ")
	m := message(words)
	return m
}

func (msg message) isRelevant() bool {
	return len(msg) > 0 && msg[0] == "/bamboo"
}
func (msg message) isValid() bool {
	return len(msg) == 3
}

type bambooAgent struct {
}

func New() *bambooAgent {
	return new(bambooAgent)
}

func (h *bambooAgent) HandleMessage(request *api.Request) ([]*api.Response, error) {
	glog.V(2).Infof("handle message: %s", request.Message)
	msg := NewMessage(request.Message)

	if !msg.isRelevant() {
		glog.V(4).Infof("message is not relevant for this handler => skip")
		return nil, nil
	}

	if !msg.isValid() {
		glog.V(4).Infof("message is not valid => reporting usage")
		return response.CreateReponseMessage("usage: /deploy ProjectName EnvironmentName"), nil
	}

	glog.V(2).Infof(fmt.Sprintf("looking up env %s in proj %s", msg[1], msg[2]))

	glog.V(2).Infof("return response")
	return response.CreateReponseMessage("deploy triggered"), nil
}
