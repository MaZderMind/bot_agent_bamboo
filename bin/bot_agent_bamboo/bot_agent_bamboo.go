package main

import (
	"fmt"

	"runtime"

	auth_model "github.com/bborbe/auth/model"
	"github.com/bborbe/bot_agent/api"
	"github.com/bborbe/bot_agent/message_handler/restrict_to_tokens"
	"github.com/bborbe/bot_agent/request_consumer"
	"github.com/bborbe/bot_agent/sender"
	"github.com/bborbe/bot_agent_bamboo/message_handler"
	flag "github.com/bborbe/flagenv"
	"github.com/bborbe/nsq_utils"
	"github.com/bborbe/nsq_utils/producer"
	"github.com/golang/glog"
)

const (
	parameterNsqLookupd       = "nsq-lookupd-address"
	parameterNsqd             = "nsqd-address"
	defaultBotName            = "bamboo"
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
	requestConsumer, err := createRequestConsumer()
	if err != nil {
		return err
	}
	return requestConsumer.Run()
}

func createRequestConsumer() (request_consumer.RequestConsumer, error) {
	nsqdAddress := nsq_utils.NsqdAddress(*nsqdAddressPtr)
	nsqLookupdAddress := nsq_utils.NsqLookupdAddress(*nsqLookupdAddressPtr)
	botname := *botNamePtr
	restrictToTokens := *restrictToTokensPtr

	if len(nsqLookupdAddress) == 0 {
		return nil, fmt.Errorf("parameter %s missing", parameterNsqLookupd)
	}
	if len(nsqdAddress) == 0 {
		return nil, fmt.Errorf("parameter %s missing", parameterNsqd)
	}
	if len(botname) == 0 {
		return nil, fmt.Errorf("parameter %s missing", parameterBotName)
	}
	producer, err := producer.New(nsqdAddress)
	if err != nil {
		return nil, err
	}
	sender := sender.New(producer)
	var messageHandler api.MessageHandler = message_handler.New()

	tokens := auth_model.ParseTokens(restrictToTokens)
	if len(tokens) > 0 {
		messageHandler = restrict_to_tokens.New(
			messageHandler,
			tokens,
		)
	}

	return request_consumer.New(sender.Send, nsqdAddress, nsqLookupdAddress, nsq_utils.NsqChannel(botname), messageHandler), nil
}
