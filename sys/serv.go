package sys

import (
	"github.com/rs/xid"
	"github/hootuu/utils/configure"
	"os"
	"strings"
)

var ServerID string
var RunMode Mode

func Exit(err error) {
	if err != nil {
		Error("Crash error: ", err.Error())
	}
	os.Exit(0)
}

func init() {
	ServerID = strings.ToUpper(xid.New().String())
	RunMode = ModeValueOf(configure.GetString("sys.mode", string(LOCAL)))
	Warn("# Server ID: ", ServerID)
	Warn("# Run Mode: ", strings.ToUpper(string(RunMode)))
}
