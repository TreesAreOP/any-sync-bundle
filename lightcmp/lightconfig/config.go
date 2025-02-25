package lightconfig

import (
	"github.com/anyproto/any-sync-consensusnode/config"

	commonaccount "github.com/anyproto/any-sync/accountservice"
	"github.com/anyproto/any-sync/app"
	"github.com/anyproto/any-sync/app/logger"
	"github.com/anyproto/any-sync/net/rpc"
	"github.com/anyproto/any-sync/net/transport/quic"
	"github.com/anyproto/any-sync/net/transport/yamux"
	"github.com/anyproto/any-sync/nodeconf"
)

var log = logger.NewNamed("light." + config.CName)

type LightConfig struct {
	// Global
	Account       commonaccount.Config
	Network       nodeconf.Configuration
	ListenTCPAddr []string
	ListenUDPAddr []string
	// Consensus
	ConsensusDBPath string
}

//
// App Component
//

func (c *LightConfig) Init(a *app.App) error {
	log.Info("call Init")

	return nil
}

func (c *LightConfig) Name() (name string) {
	return config.CName
}

//
// Component
//

func (c *LightConfig) GetDrpc() rpc.Config {
	log.Info("call GetDrpc")

	return rpc.Config{
		Stream: rpc.StreamConfig{
			MaxMsgSizeMb: 256,
		},
	}
}

func (c *LightConfig) GetAccount() commonaccount.Config {
	log.Info("call GetAccount")

	return c.Account
}

func (c *LightConfig) GetNodeConf() nodeconf.Configuration {
	log.Info("call GetNodeConf")

	return c.Network
}

func (c *LightConfig) GetYamux() yamux.Config {
	log.Info("call GetYamux")

	return yamux.Config{
		ListenAddrs: c.ListenTCPAddr,
	}
}

func (c *LightConfig) GetQuic() quic.Config {
	log.Info("call GetQuic")

	return quic.Config{
		ListenAddrs: c.ListenUDPAddr,
	}
}

func (c *LightConfig) GetConsensusDBPath() string {
	log.Info("call GetConsensusDBPath")

	return c.ConsensusDBPath
}
