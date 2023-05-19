package global

const (
	ConfigEnv         = "GD_CONFIG" //电脑环境变量，用于控制gin模式：release test debug
	ConfigDefaultFile = "./config/config.yaml"
	ConfigTestFile    = "./config/config.test.yaml"
	ConfigDebugFile   = "./config/config.debug.yaml"
	ConfigReleaseFile = "./config/config.release.yaml"
)

type (
	Mode string
)

const (
	ModeDev  Mode = "dev"     //开发模式
	ModeTest Mode = "test"    //测试模式
	ModeProd Mode = "prod"    //生产模式
)

func (e Mode) String() string {
	return string(e)
}