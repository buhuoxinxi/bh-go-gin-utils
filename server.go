package bhginutils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	// init config
	DefaultConfigFn()
}

// config server config
var (
	config *Config // config
)

// NewServer new server
func NewServer() *gin.Engine {

	// mode
	if config.ServerMode != gin.ReleaseMode && config.ServerMode != gin.DebugMode && config.ServerMode != gin.TestMode {
		config.ServerMode = gin.ReleaseMode
	}
	gin.SetMode(config.ServerMode)

	// server
	if config.ServerMode == gin.ReleaseMode {
		return gin.New()
	}
	return gin.Default()
}

// RunServer run server
func RunServer(engine *gin.Engine) {
	// info
	addr := ":" + config.ServerPort
	logrus.Infof("server addr %s, isHTTPS(%v)", addr, config.SSLEnable)

	// ssl
	if config.SSLEnable {
		if err := engine.RunTLS(addr, config.SSLCertFile, config.SSLKeyFile); err != nil {
			logrus.Fatalf("RunServer RunTLS(engine, addr, certFile, keyFile) error : %v", err)
		}
	}

	// start
	if err := engine.Run(addr); err != nil {
		logrus.Fatalf("RunServer Run(engine, addr) error : %v", err)
	}
}
