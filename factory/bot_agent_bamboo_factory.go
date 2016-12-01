package factory

import (
	"net/http"

	"github.com/bborbe/auth/client"
	auth_model "github.com/bborbe/auth/model"
	"github.com/bborbe/auth/service"
	"github.com/bborbe/bot_agent/api"
	"github.com/bborbe/bot_agent/message_handler/match"
	"github.com/bborbe/bot_agent/message_handler/restrict_to_tokens"
	"github.com/bborbe/bot_agent/request_consumer"
	"github.com/bborbe/bot_agent/sender"
	"github.com/bborbe/bot_agent_bamboo/bamboo"
	deploy_handler "github.com/bborbe/bot_agent_bamboo/deploy"
	"github.com/bborbe/bot_agent_bamboo/model"
	http_client_builder "github.com/bborbe/http/client_builder"
	http_rest "github.com/bborbe/http/rest"
	"github.com/golang/glog"
	"github.com/nsqio/go-nsq"
)

type botAgentBambooFactory struct {
	config   model.Config
	producer *nsq.Producer
}

func New(config model.Config, producer *nsq.Producer) *botAgentBambooFactory {
	b := new(botAgentBambooFactory)
	b.config = config
	b.producer = producer
	return b
}

func (b *botAgentBambooFactory) httpRest() http_rest.Rest {
	httpClient := http_client_builder.New().WithoutProxy().Build()
	return http_rest.New(httpClient.Do)
}

func (a *botAgentBambooFactory) httpClient() *http.Client {
	return http_client_builder.New().WithoutProxy().Build()
}

func (a *botAgentBambooFactory) authClient() client.Client {
	return client.New(a.httpClient().Do, auth_model.Url(a.config.AuthUrl), auth_model.ApplicationName(a.config.AuthApplicationName), auth_model.ApplicationPassword(a.config.AuthApplicationPassword))
}

func (a *botAgentBambooFactory) authService() service.AuthService {
	return a.authClient().AuthService()
}

func (b *botAgentBambooFactory) hasRequiredGroups(authToken auth_model.AuthToken) bool {
	result, err := b.authService().HasGroups(authToken, b.config.RequiredGroupNames)
	if err != nil {
		return false
	}
	return result
}

func (b *botAgentBambooFactory) RequestConsumer() request_consumer.RequestConsumer {

	bambooDeployer := bamboo.NewDeployer(b.httpRest(), b.config.BambooUrl, b.config.BambooUsername, b.config.BambooPassword)

	var messageHandler api.MessageHandler = match.New(
		b.config.Prefix.String(),
		deploy_handler.New(b.config.Prefix, bambooDeployer, b.hasRequiredGroups),
	)

	if len(b.config.RestrictToTokens) > 0 {
		glog.V(2).Infof("restrict to tokens: %v", b.config.RestrictToTokens)
		messageHandler = restrict_to_tokens.New(
			messageHandler,
			b.config.RestrictToTokens,
		)
	}

	sender := sender.New(b.producer)
	return request_consumer.New(sender.Send, b.config.NsqdAddress, b.config.NsqLookupdAddress, b.config.BotName, messageHandler)
}
