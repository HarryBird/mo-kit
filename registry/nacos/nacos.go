package nacos

import (
	"errors"
	"fmt"
	"strings"

	"github.com/HarryBird/mo-kit/config"
	"github.com/nacos-group/nacos-sdk-go/clients"
	client "github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

const (
	ClusterBeijing = "MO-CLUSTER-BJ"
	GroupDefault   = "MO-GROUP-DEFAULT"
)

// ErrInvalidConfig .
var ErrInvalidConfig = errors.New("mo-kit -> nacos: invalid config, miss address or port")

// DefaultServerConfig .
func DefaultServerConfig(conf IServerConfig) ([]constant.ServerConfig, error) {
	opts := make([]constant.ServerOption, 0, 2)

	if conf.GetAddress() == "" || conf.GetPort() == 0 {
		return nil, ErrInvalidConfig
	}

	if conf.GetScheme() != "" {
		opts = append(opts, constant.WithScheme(conf.GetScheme()))
	}

	if conf.GetContextPath() != "" {
		opts = append(opts, constant.WithContextPath(conf.GetContextPath()))
	}

	return []constant.ServerConfig{
		*constant.NewServerConfig(conf.GetAddress(), conf.GetPort(), opts...),
	}, nil
}

// DefaultClientConfig .
func DefaultClientConfig(conf config.IRuntimeConfig) (*constant.ClientConfig, error) {
	env := config.RuntimeEnv(strings.ToUpper(conf.GetEnv()))

	if !env.IsValid() {
		return nil, config.ErrInvalidEnv
	}

	logLevel := "info"
	if env == config.RuntimeEnvProd {
		logLevel = "warn"
	}

	return &constant.ClientConfig{
		NamespaceId:         fmt.Sprintf("MO-%s", env),
		TimeoutMs:           5000,
		BeatInterval:        5000,
		NotLoadCacheAtStart: true,
		LogDir:              "./tmp/nacos/log",
		CacheDir:            "./tmp/nacos/cache",
		LogRollingConfig: &constant.ClientLogRollingConfig{
			MaxSize:    100,
			MaxAge:     7,
			MaxBackups: 7,
		},
		LogLevel: logLevel,
	}, nil
}

// DefaultClient .
func DefaultClient(servCfg IServerConfig, runCfg config.IRuntimeConfig) (client.INamingClient, []constant.ServerConfig, *constant.ClientConfig, error) {
	var err error
	sc, err := DefaultServerConfig(servCfg)
	if err != nil {
		return nil, nil, nil, err
	}

	cc, err := DefaultClientConfig(runCfg)
	if err != nil {
		return nil, sc, nil, err
	}

	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ServerConfigs: sc,
			ClientConfig:  cc,
		},
	)
	if err != nil {
		return nil, sc, cc, err
	}

	return cli, sc, cc, err
}
