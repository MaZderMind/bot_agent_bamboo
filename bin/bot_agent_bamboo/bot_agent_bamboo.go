package main

import (
	"runtime"

	auth_model "github.com/bborbe/auth/model"
	"github.com/bborbe/bot_agent_bamboo/factory"
	"github.com/bborbe/bot_agent_bamboo/model"
	flag "github.com/bborbe/flagenv"
	"github.com/bborbe/nsq_utils"
	"github.com/bborbe/nsq_utils/producer"
	"github.com/golang/glog"
)

const (
	defaultBotName                   = "bamboo"
	parameterNsqLookupd              = "nsq-lookupd-address"
	parameterNsqd                    = "nsqd-address"
	parameterBotName                 = "bot-name"
	parameterRestrictToTokens        = "restrict-to-tokens"
	parameterRequiredGroups          = "required-groups"
	parameterBambooUrl               = "bamboo-url"
	parameterBambooUsername          = "bamboo-username"
	parameterBambooPassword          = "bamboo-password"
	parameterAuthUrl                 = "auth-url"
	parameterAuthApplicationName     = "auth-application-name"
	parameterAuthApplicationPassword = "auth-application-password"
	parameterPrefix                  = "prefix"
)

var (
	nsqLookupdAddressPtr       = flag.String(parameterNsqLookupd, "", "nsq lookupd address")
	nsqdAddressPtr             = flag.String(parameterNsqd, "", "nsqd address")
	botNamePtr                 = flag.String(parameterBotName, defaultBotName, "bot name")
	requiredGroupsPtr          = flag.String(parameterRequiredGroups, "", "required groups reperated by comma")
	restrictToTokensPtr        = flag.String(parameterRestrictToTokens, "", "restrict to tokens")
	bambooUrlPtr               = flag.String(parameterBambooUrl, "", "bamboo url")
	bambooUsernamePtr          = flag.String(parameterBambooUsername, "", "bamboo username")
	bambooPasswordPtr          = flag.String(parameterBambooPassword, "", "bamboo password")
	authUrlPtr                 = flag.String(parameterAuthUrl, "", "auth url")
	authApplicationNamePtr     = flag.String(parameterAuthApplicationName, "", "auth application name")
	authApplicationPasswordPtr = flag.String(parameterAuthApplicationPassword, "", "auth application password")
	prefixPtr                  = flag.String(parameterPrefix, "/deploy", "prefix commands start with")
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
	producer, err := producer.New(config.NsqdAddress)
	if err != nil {
		return err
	}
	factory := factory.New(config, producer)
	return factory.RequestConsumer().Run()
}

func createConfig() model.Config {
	return model.Config{
		Prefix:                  model.Prefix(*prefixPtr),
		NsqLookupdAddress:       nsq_utils.NsqLookupdAddress(*nsqLookupdAddressPtr),
		NsqdAddress:             nsq_utils.NsqdAddress(*nsqdAddressPtr),
		BotName:                 nsq_utils.NsqChannel(*botNamePtr),
		RequiredGroupNames:      auth_model.ParseGroupNames(*requiredGroupsPtr),
		RestrictToTokens:        auth_model.ParseTokens(*restrictToTokensPtr),
		BambooUrl:               model.BambooUrl(*bambooUrlPtr),
		BambooUsername:          model.BambooUsername(*bambooUsernamePtr),
		BambooPassword:          model.BambooPassword(*bambooPasswordPtr),
		AuthUrl:                 auth_model.Url(*authUrlPtr),
		AuthApplicationName:     auth_model.ApplicationName(*authApplicationNamePtr),
		AuthApplicationPassword: auth_model.ApplicationPassword(*authApplicationPasswordPtr),
	}
}
