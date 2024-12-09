package logger

import (
	"gin-frame-base/internal/config"
	"gin-frame-base/internal/tool"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
)

var (
	options []zap.Option
	conf    config.Logger
)

func New(LoggerConf config.Logger) (logger *zap.Logger, err error) {
	conf = LoggerConf

	rootDir, _ := tool.GetRootDir()
	logDir := rootDir + conf.FilePath
	err = os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		return
	}

	loggerConf := genConfig()

	loggerConf.EncoderConfig = genEncodeConfig()

	writer, err := genWriteSyncer()
	if err != nil {
		return nil, err
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(loggerConf.EncoderConfig),
		writer,
		loggerConf.Level,
	)

	// 添加调用位置
	options = append(options, zap.AddCaller())

	logger = zap.New(core, options...)
	return
}

// 生成WriteSyncer
func genWriteSyncer() (writeSyncer zapcore.WriteSyncer, err error) {
	// 创建日志存放目录
	rootDir, _ := tool.GetRootDir()
	logDir := rootDir + conf.FilePath
	err = os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		return
	}

	fileName := formatLogFileName(conf.FileName)
	lumberJack := &lumberjack.Logger{
		Filename:   logDir + fileName,
		MaxSize:    conf.MaxSize, // megabytes
		MaxBackups: conf.MaxBackups,
		MaxAge:     conf.MaxAge, //days
		Compress:   true,
	}

	writeSyncer = zapcore.AddSync(lumberJack)
	return
}

func formatLogFileName(baseFileName string) string {
	// 当前日期
	currentDate := time.Now().Format("2006-01-02")
	return strings.Replace(baseFileName, "{date}", currentDate, 1)
}

// 生成配置
func genConfig() (config zap.Config) {
	config = zap.NewProductionConfig()

	config.EncoderConfig = genEncodeConfig()
	config.Level = zap.NewAtomicLevelAt(GetLevel())

	return

}

// 生成编码配置
func genEncodeConfig() (c zapcore.EncoderConfig) {
	c = zap.NewProductionEncoderConfig()

	c.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("2006-01-02 15:04:05.000"))
	}

	c.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(strings.ToUpper(l.String()))
	}

	c.TimeKey = "time"

	return
}

// GetLevel 转换为zapcore的level
func GetLevel() (level zapcore.Level) {
	switch conf.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	return
}
