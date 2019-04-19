package bhtestdata

import (
	"os"
	"path/filepath"
	"runtime"
)

// basepath is the root directory of this package.
var basepath string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(currentFile)

	// host
	os.Setenv("BhServerGatewayHost", "")
	os.Setenv("BhServerGatewayPort", "8099")
	os.Setenv("BhServerMode", "release")

	// ssl
	os.Setenv("BhServerSSLEnable", "true")
	os.Setenv("BhServerSSLCaFile", Path("ca.pem"))
	os.Setenv("BhServerSSLCertFile", Path("server1.pem"))
	os.Setenv("BhServerSSLKeyFile", Path("server1.key"))
	os.Setenv("BhServerSSLServerName", "x.test.youtube.com")
}

// Path returns the absolute path the given relative file or directory path,
// relative to the google.golang.org/grpc/directory in the user's GOPATH.
// If rel is already absolute, it is returned unmodified.
func Path(rel string) string {
	if filepath.IsAbs(rel) {
		return rel
	}
	return filepath.Join(basepath, rel)
}
