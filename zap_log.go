package goutil

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var Logger *zap.Logger
var Log *zap.SugaredLogger

func InitZapLog(fileName string, maxSize int, maxBackups int, maxAge int, serviceName string) (*zap.Logger, *zap.SugaredLogger) {
	hook := lumberjack.Logger{
		//Filename:   "./logs/spikeProxy1.log", // 日志文件路径
		//MaxSize:    128,                      // 每个日志文件保存的最大尺寸 单位：M
		//MaxBackups: 30,                       // 日志文件最多保存多少个备份
		//MaxAge:     7,                        // 文件最多保存多少天

		Filename:   fileName,   // 日志文件路径
		MaxSize:    maxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: maxBackups, // 日志文件最多保存多少个备份
		MaxAge:     maxAge,     // 文件最多保存多少天

		Compress: true, // 是否压缩
	}
	encoderConfig := zapcore.EncoderConfig{
		//MessageKey：输入信息的key名
		//LevelKey：输出日志级别的key名
		//TimeKey：输出时间的key名
		//NameKey CallerKey StacktraceKey跟以上类似，看名字就知道
		//LineEnding：每行的分隔符。基本zapcore.DefaultLineEnding 即"\n"
		//EncodeLevel：基本zapcore.LowercaseLevelEncoder。将日志级别字符串转化为小写
		//EncodeTime：输出的时间格式
		//EncodeDuration：一般zapcore.SecondsDurationEncoder,执行消耗的时间转化成浮点型的秒
		//EncodeCaller：一般zapcore.ShortCallerEncoder，以包/文件:行号 格式化调用堆栈
		//EncodeName：可选值。
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line_num",
		MessageKey:     "message",
		StacktraceKey:  "trace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     formatEncodeTime,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("serviceName", serviceName))
	Logger = zap.New(core, caller, development, filed)
	Log = Logger.Sugar()
	return Logger, Log
}

func formatEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%d%02d%02d_%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
}

//func FormatLog(args []interface{}) *zap.SugaredLogger {
//	return Log.With(toJsonData(args)).Sugar()
//}
//
//func toJsonData(args []interface{}) zap.Field {
//	det := make([]string, 0)
//	if len(args) > 0 {
//		for _, v := range args {
//			det = append(det, fmt.Sprintf("%+v", v))
//		}
//	}
//	var zap = zap.Any("detail", det)
//	return zap
//}
//func DebugLogs(msg string, args ...interface{}) {
//	FormatLog(args).Debugf(msg)
//}
//
//func InfoLogs(msg string, args ...interface{}) {
//	FormatLog(args).Infof(msg)
//}
//
//func WarnLogs(msg string, args ...interface{}) {
//	FormatLog(args).Warnf(msg)
//}
//
//func ErrorLogs(msg string, args ...interface{}) {
//	FormatLog(args).Errorf(msg)
//}
