package config

import (
	"go/build"
	"os"
	"time"
)

var (
	LogPath             = "./"
	HmyURL              = "http://localhost:9500/"
	HarmonyPath         = "./"
	BlockchainInterval  = 3000 * time.Millisecond
	SystemStatsInterval = 500 * time.Millisecond
	WidgetInterval      = 1000 * time.Millisecond
)

func SetConfig(env string) {
	if env == "local" {
		gopath := os.Getenv("GOPATH")
		if gopath == "" {
			gopath = build.Default.GOPATH
		}
		LogPath = gopath + "/src/github.com/harmony-one/harmony/tmp_log/"
		HarmonyPath = gopath + "/src/github.com/harmony-one/harmony/bin/"

		BlockchainInterval = 3000 * time.Millisecond
		SystemStatsInterval = 250 * time.Millisecond
	} else if env == "ec2" {
		LogPath = "./latest/"
		HarmonyPath = "./"
		BlockchainInterval = 5000 * time.Millisecond
		SystemStatsInterval = 500 * time.Millisecond
	}
}
