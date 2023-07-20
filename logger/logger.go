package logger
import{
  "go.uber.org/zap"
  "go.uber.org/zap/zapcore"
  "gopkg.in/natefinch/lumberjack.v2"
  }

var l *zap.Logger

func InitLogger(logPath,logLevel string) error{
  hook:= lumberjack.Logger{
    Filename: logPath,
    MaxSize: 1024,
    MaxBackups:3,
    MaxAge: 7,
    Compress: true,
  }
  w := zapcore.AddSync(&hook)
  var level zapcore.Level
  switch logLevel{
    case "debug":
      level = zap.DebugLevel
    case "info":
      level = zap.InfoLevel
    case "error":
      level = zap.ErrorLevel
    default:
      leve = zap.DebugLevel
  }
}
