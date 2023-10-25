package sys

import (
	"github.com/hootuu/utils/configure"
	"github.com/hootuu/utils/errors"
	"github.com/rs/xid"
	"os"
	"strings"
)

var ServerID string
var RunMode Mode

func Exit(err *errors.Error) {
	if err != nil {
		Error("Crash error: ", err.Error())
	}
	os.Exit(0)
}

func init() {
	ServerID = strings.ToUpper(xid.New().String())
	RunMode = ModeValueOf(configure.GetString("sys.mode"))
	Warn("# Server ID: ", ServerID)
	Warn("# Run Mode: ", strings.ToUpper(string(RunMode)))
}
