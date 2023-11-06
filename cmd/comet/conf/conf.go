package conf

import "os"

var (
	SocketPort    string // comet port
	PProfPort     string // comet pprof port
	RpcPort       string // comet grpc server
	WebsocketPort string // comet websocket port
	CurrPodIp     string // 本机真实IP（可变）
	SerENV        string // 当前环境(开发/测试/用户验收/线上/等等), 可能的值：dev/fat/uat/prod
)

func init() {
	SerENV = os.Getenv("GROUP")
	if len(SerENV) == 0 { // 默认
		SerENV = Dev
	}
}

const (
	Dev  = "dev"
	Fat  = "fat"
	Uat  = "uat"
	Prod = "prod"
)
