package peer

import (
	"context"
	"fmt"
	"github.com/hootuu/utils/configure"
	"github.com/hootuu/utils/errors"
	"github.com/hootuu/utils/logger"
	"github.com/hootuu/utils/sys"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"
)

type Service struct {
	Code     string
	Startup  func() *errors.Error
	Shutdown func(ctx context.Context) *errors.Error
}

var gServices map[string]*Service
var gLock sync.Mutex

func doRegisterService(service *Service) {
	gLock.Lock()
	defer gLock.Unlock()
	_, exists := gServices[service.Code]
	if exists {
		sys.Exit(errors.Sys(fmt.Sprintf("The same service has been register.[ code=%s ]", service.Code)))
	}
	gServices[service.Code] = service
}

func RegisterService(service ...*Service) {
	if len(service) == 0 {
		sys.Exit(errors.Sys("no any service"))
	}
	for _, s := range service {
		doRegisterService(s)
	}
}

func Running() {
	for code, service := range gServices {
		sys.Warn("# Start the service: ", code, " ...... #")
		err := service.Startup()
		if err != nil {
			logger.Logger.Error("start server failed", zap.String("code", code), zap.Error(err))
			sys.Error("# Start service exception: ", code, " #")
			return
		}
		sys.Success("# Start the service ", code, " [OK] #")
	}

	sys.Info("# ALL Used Configure Items #")
	configure.Dump(func(key string, val any) {
		if strings.Index(strings.ToLower(key), "password") > -1 {
			sys.Info("  # ", key, " # ==>> ", "**********")
		} else {
			sys.Info("  # ", key, " # ==>> ", val)
		}

	})

	sys.Success("# HOTU Super Node Is At Your Service ......#")

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	sys.Info("# Shutting down the system ...... #")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for code, service := range gServices {
		sys.Info(" * Stop service: ", code, " ......")
		err := service.Shutdown(ctx)
		if err != nil {
			logger.Logger.Error("shutdown server failed", zap.String("code", code), zap.Error(err))
			sys.Error(" * Stop service ", code, " exception")
		}
		sys.Success(" * Stop service ", code, " [OK]")
	}
	sys.Success("# Shutting down the system [OK] #")
}

func init() {
	gServices = make(map[string]*Service)
}
