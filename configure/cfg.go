package configure

import (
	"github.com/hootuu/utils/errors"
	"github.com/hootuu/utils/strs"
	"github.com/spf13/viper"
	"strings"
	"time"
)

func GetString(key string, def ...string) string {
	ok := exists(key)
	if ok {
		return viper.GetString(key)
	}
	if len(def) > 0 {
		return def[0]
	}
	return strs.EMPTY
}

func MustGetString(key string) (string, *errors.Error) {
	val := GetString(key)
	if len(val) == 0 {
		return val, errors.Sys("Miss configure: " + key + "! " +
			"This parameter should be set in an environment variable, startup parameter, or configuration file.")
	}
	return val, nil
}

func GetBool(key string, def ...bool) bool {
	ok := exists(key)
	if ok {
		return viper.GetBool(key)
	}
	if len(def) > 0 {
		return def[0]
	}
	return false
}

func GetInt(key string, def ...int) int {
	ok := exists(key)
	if ok {
		return viper.GetInt(key)
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

func GetUint64(key string, def ...uint64) uint64 {
	ok := exists(key)
	if ok {
		return viper.GetUint64(key)
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

func GetDuration(key string, def ...time.Duration) time.Duration {
	ok := exists(key)
	if ok {
		return viper.GetDuration(key)
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

var gConfigs = make(map[string]any)

func register(key string, val any) {
	gConfigs[key] = val
}

// Dump Used to output all used configurations
func Dump(call func(key string, val any)) {
	for k, v := range gConfigs {
		call(k, v)
	}
}

func exists(key string) bool {
	val := viper.Get(key)
	if val != nil {
		register(key, val)
	}
	return val != nil
}

func init() {
	viper.SetEnvPrefix("HOTU")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AddConfigPath("./_bin/")
	viper.AddConfigPath("./")
	viper.SetConfigName(GetString("sys.mode", "local"))
	viper.SetConfigType("yaml")
	_ = viper.ReadInConfig()
}
