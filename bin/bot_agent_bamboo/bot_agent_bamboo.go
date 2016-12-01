package main

import (
	"runtime"

	auth_model "github.com/bborbe/auth/model"
	"github.com/bborbe/bot_agent/api"
	"github.com/bborbe/bot_agent/message_handler/restrict_to_tokens"
	"github.com/bborbe/bot_agent/request_consumer"
	"github.com/bborbe/bot_agent/sender"
	"github.com/bborbe/bot_agent_bamboo/bamboo"
	"github.com/bborbe/bot_agent_bamboo/message_handler"
	"github.com/bborbe/bot_agent_bamboo/model"
	flag "github.com/bborbe/flagenv"
	"github.com/bborbe/nsq_utils"
	"github.com/bborbe/nsq_utils/producer"
	"github.com/golang/glog"
)

const (
	defaultBotName            = "bamboo"
	parameterNsqLookupd       = "nsq-lookupd-address"
	parameterNsqd             = "nsqd-address"
	parameterBotName          = "bot-name"
	parameterRestrictToTokens = "restrict-to-tokens"
	parameterBambooUrl        = "bamboo-url"
	parameterBambooUsername   = "bamboo-username"
	parameterBambooPassword   = "bamboo-password"
)

var (
	nsqLookupdAddressPtr = flag.String(parameterNsqLookupd, "", "nsq lookupd address")
	nsqdAddressPtr       = flag.String(parameterNsqd, "", "nsqd address")
	botNamePtr           = flag.String(parameterBotName, defaultBotName, "bot name")
	restrictToTokensPtr  = flag.String(parameterRestrictToTokens, "", "restrict to tokens")
	bambooUrlPtr         = flag.String(parameterBambooUrl, "", "bamboo url")
	bambooUsernamePtr    = flag.String(parameterBambooUsername, "", "bamboo username")
	bambooPasswordPtr    = flag.String(parameterBambooPassword, "", "bamboo password")
)

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())

	if err := do(); err != nil {
		glog.Exit(err)
	}
}

func do() error {
	config := createConfig()
	if err := config.Validate(); err != nil {
		glog.V(2).Infof("validate config failed: %v", err)
		return err
	}

	requestConsumer, err := createRequestConsumer(config)
	if err != nil {
		glog.V(2).Infof("create request consumer failed: %v", err)
		return err
	}
	return requestConsumer.Run()
}

func createRequestConsumer(config model.Config) (request_consumer.RequestConsumer, error) {
	producer, err := producer.New(config.NsqdAddress)
	if err != nil {
		return nil, err
	}
	sender := sender.New(producer)

	bambooDeployer := bamboo.NewDeployer(config.BambooUrl, config.BambooUsername, config.BambooPassword)

	var messageHandler api.MessageHandler = message_handler.New(bambooDeployer)

	tokens := auth_model.ParseTokens(config.RestrictToTokens)
	if len(tokens) > 0 {
		messageHandler = restrict_to_tokens.New(
			messageHandler,
			tokens,
		)
	}

	return request_consumer.New(sender.Send, config.NsqdAddress, config.NsqLookupdAddress, config.BotName, messageHandler), nil
}

func createConfig() model.Config {
	return model.Config{
		NsqLookupdAddress: nsq_utils.NsqLookupdAddress(*nsqLookupdAddressPtr),
		NsqdAddress:       nsq_utils.NsqdAddress(*nsqdAddressPtr),
		BotName:           nsq_utils.NsqChannel(*botNamePtr),
		RestrictToTokens:  *restrictToTokensPtr,
		BambooUrl:         model.BambooUrl(*bambooUrlPtr),
		BambooUsername:    model.BambooUsername(*bambooUsernamePtr),
		BambooPassword:    model.BambooPassword(*bambooPasswordPtr),
	}
}
