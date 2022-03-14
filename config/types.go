package config

type IRuntimeConfig interface {
	GetEnv() string
}

type RuntimeEnv string

const (
	RuntimeEnvLocal  RuntimeEnv = "LOCAL"
	RuntimeEnvDevlab RuntimeEnv = "DEVLAB"
	RuntimeEnvDev    RuntimeEnv = "DEV"
	RuntimeEnvTest   RuntimeEnv = "TEST"
	RuntimeEnvProd   RuntimeEnv = "PROD"
)

// IsValid  是否有效的运行环境
func (r RuntimeEnv) IsValid() bool {
	return r == RuntimeEnvLocal || r == RuntimeEnvDevlab || r == RuntimeEnvDev || r == RuntimeEnvTest || r == RuntimeEnvProd
}

func (r RuntimeEnv) String() string {
	return string(r)
}
