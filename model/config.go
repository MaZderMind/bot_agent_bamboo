package model

import (
	"fmt"

	"github.com/bborbe/nsq_utils"
)

// Config of collector
type Config struct {
	NsqLookupdAddress nsq_utils.NsqLookupdAddress `json:"nsq-lookupd-address"`
	NsqdAddress       nsq_utils.NsqdAddress       `json:"nsqd-address"`
	BotName           nsq_utils.NsqChannel        `json:"bot-name"`
	RestrictToTokens  string                      `json:"restrict-to-tokens"`
	BambooUrl         BambooUrl                   `json:"bamboo-url"`
	BambooUsername    BambooUsername              `json:"bamboo-username"`
	BambooPassword    BambooPassword              `json:"bamboo-password"`
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
		return fmt.Errorf("parameter BambooPassword missing")
	}
	return nil
}

// BambooUrl is the url of bamboo
type BambooUrl string

// BambooUsername is the username of bamboo
type BambooUsername string

// BambooPassword is the password of bamboo
type BambooPassword string
