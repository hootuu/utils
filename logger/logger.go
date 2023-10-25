package logger

import (
	"github.com/hootuu/utils/configure"
	"github.com/hootuu/utils/sys"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
)

type Level = string

const (
	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
)

func LevelOf(level Level) zapcore.Level {
	zapLevel := zapcore.InfoLevel
	switch level {
	case DebugLevel:
		zapLevel = zapcore.DebugLevel
	case InfoLevel:
		zapLevel = zapcore.InfoLevel
	case WarnLevel:
		zapLevel = zapcore.WarnLevel
	case ErrorLevel:
		zapLevel = zapcore.ErrorLevel
	default:
		zapLevel = zapcore.ErrorLevel
	}
	return zapLevel
}

func GetLogger(key string) *zap.Logger {
	rootPath := configure.GetString("logger."+key+".root", "./.logs/"+key+"/")
	logLevel := configure.GetString("logger."+key+".level", "debug")
	fileName := rootPath + configure.GetString("logger."+key+".file", key+".jsonl")
	maxSize := configure.GetInt("logger."+key+".size.max", 128)
	maxBackups := configure.GetInt("logger."+key+".backup.max", 30)
	maxAge := configure.GetInt("logger."+key+".age.max", 7)
	compress := configure.GetBool("logger."+key+".compress", false)
	sys.Warn("# Initialize the " + key + " log system ..... #")
	sys.Info(" * PATH: ", strings.ToUpper(rootPath), "      ${ logger."+key+".root }")
	sys.Info(" * LEVEL: ", strings.ToUpper(logLevel), "      ${ logger."+key+".level }")
	sys.Info(" * FILE: ", strings.ToUpper(fileName), "      ${ logger."+key+".file }")
	sys.Info(" * MAX SIZE: ", maxSize, "      ${ logger."+key+".size.max }")
	sys.Info(" * BACKUP MAX: ", maxBackups, "      ${ logger."+key+".backup.max }")
	sys.Info(" * AGE MAX: ", maxAge, "      ${ logger."+key+".age.max }")
	sys.Info(" * COMPRESS: ", compress, "      ${ logger."+key+".compress }")

	zapLogLevel := LevelOf(logLevel)

	hook := lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        key,
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zapLogLevel)

	writeSyncers := []zapcore.WriteSyncer{
		zapcore.AddSync(&hook),
	}

	if sys.RunMode.IsRd() {
		writeSyncers = append(writeSyncers, zapcore.AddSync(os.Stdout))
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writeSyncers...),
		atomicLevel,
	)

	sys.Success("# Initialize the " + key + " log system [OK] #")

	return zap.New(core)
}

var Logger *zap.Logger
var Console *zap.Logger

func init() {
	Logger = GetLogger("commons")
	Console = GetLogger("console")
	if !sys.RunMode.IsRd() {
		sys.ConsoleToLogger(func(msg string) {
			Console.Info(msg)
		})
	}
}
