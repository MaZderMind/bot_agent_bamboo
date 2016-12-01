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
	PARAMETER_NSQ_LOOKUPD        = "nsq-lookupd-address"
	PARAMETER_NSQD               = "nsqd-address"
	DEFAULT_BOT_NAME             = "bamboo"
	PARAMETER_BOT_NAME           = "bot-name"
	PARAMETER_RESTRICT_TO_TOKENS = "restrict-to-tokens"
)

var (
	nsqLookupdAddressPtr = flag.String(PARAMETER_NSQ_LOOKUPD, "", "nsq lookupd address")
	nsqdAddressPtr       = flag.String(PARAMETER_NSQD, "", "nsqd address")
	botNamePtr           = flag.String(PARAMETER_BOT_NAME, DEFAULT_BOT_NAME, "bot name")
	restrictToTokensPtr  = flag.String(PARAMETER_RESTRICT_TO_TOKENS, "", "restrict to tokens")
)

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())

	err := do(
		nsq_utils.NsqdAddress(*nsqdAddressPtr),
		nsq_utils.NsqLookupdAddress(*nsqLookupdAddressPtr),
		*botNamePtr,
		*restrictToTokensPtr,
	)
	if err != nil {
		glog.Exit(err)
	}
}

func do(
	nsqdAddress nsq_utils.NsqdAddress,
	nsqLookupdAddress nsq_utils.NsqLookupdAddress,
	botname string,
	restrictToTokens string,
) error {
	requestConsumer, err := createRequestConsumer(nsqdAddress, nsqLookupdAddress, botname, restrictToTokens)
	if err != nil {
		return err
	}
	return requestConsumer.Run()
}

func createRequestConsumer(
	nsqdAddress nsq_utils.NsqdAddress,
	nsqLookupdAddress nsq_utils.NsqLookupdAddress,
	botname string,
	restrictToTokens string,
) (request_consumer.RequestConsumer, error) {
	if len(nsqLookupdAddress) == 0 {
		return nil, fmt.Errorf("parameter %s missing", PARAMETER_NSQ_LOOKUPD)
	}
	if len(nsqdAddress) == 0 {
		return nil, fmt.Errorf("parameter %s missing", PARAMETER_NSQD)
	}
	if len(botname) == 0 {
		return nil, fmt.Errorf("parameter %s missing", PARAMETER_BOT_NAME)
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
