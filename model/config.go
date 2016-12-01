package model

import (
	"fmt"

	auth_model "github.com/bborbe/auth/model"
	"github.com/bborbe/nsq_utils"
)

// Config of collector
type Config struct {
	NsqLookupdAddress       nsq_utils.NsqLookupdAddress
	NsqdAddress             nsq_utils.NsqdAddress
	BotName                 nsq_utils.NsqChannel
	RequiredGroupNames      []auth_model.GroupName
	RestrictToTokens        []auth_model.AuthToken
	BambooUrl               BambooUrl
	BambooUsername          BambooUsername
	BambooPassword          BambooPassword
	AuthUrl                 auth_model.Url
	AuthApplicationName     auth_model.ApplicationName
	AuthApplicationPassword auth_model.ApplicationPassword
	Prefix                  Prefix
}

// Validate the config values
func (config *Config) Validate() error {
	if len(config.NsqLookupdAddress) == 0 {
		return fmt.Errorf("parameter NsqLookupdAddress missing")
	}
	if len(config.NsqdAddress) == 0 {
		return fmt.Errorf("parameter NsqdAddress missing")
	}
	if len(config.BotName) == 0 {
		return fmt.Errorf("parameter BotName missing")
	}
	if len(config.BambooUrl) == 0 {
		return fmt.Errorf("parameter BambooUrl missing")
	}
	if len(config.BambooUsername) == 0 {
		return fmt.Errorf("parameter BambooUsername missing")
	}
	if len(config.BambooPassword) == 0 {
		return fmt.Errorf("parameter AuthUrl missing")
	}
	if len(config.BambooPassword) == 0 {
		return fmt.Errorf("parameter AuthApplicationName missing")
	}
	if len(config.BambooPassword) == 0 {
		return fmt.Errorf("parameter AuthApplicationPassword missing")
	}
	if len(config.BambooPassword) == 0 {
		return fmt.Errorf("parameter BambooPassword missing")
	}
	if len(config.Prefix) == 0 {
		return fmt.Errorf("parameter Prefix missing")
	}
	return nil
}

// Prefix commands start with
type Prefix string

func (p Prefix) String() string {
	return string(p)
}

// BambooUrl is the url of bamboo
type BambooUrl string

func (b BambooUrl) String() string {
	return string(b)
}

// BambooUsername is the username of bamboo
type BambooUsername string

func (b BambooUsername) String() string {
	return string(b)
}

// BambooPassword is the password of bamboo
type BambooPassword string

func (b BambooPassword) String() string {
	return string(b)
}

type ApiUrl string

func (a ApiUrl) String() string {
	return string(a)
}

type ApiUsername string

func (a ApiUsername) String() string {
	return string(a)
}

type ApiPassword string

func (a ApiPassword) String() string {
	return string(a)
}
