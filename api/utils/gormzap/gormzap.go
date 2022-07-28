package gormzap

import (
	"context"
	"time"

	"go.uber.org/zap"
	gormLogger "gorm.io/gorm/logger"
)

type DbLogger struct {
	ZapLogger *zap.Logger
	LogLevel  gormLogger.LogLevel
}

func New(zapLogger *zap.Logger) DbLogger {
	return DbLogger{
		ZapLogger: zapLogger,
		LogLevel:  gormLogger.Warn,
	}
}

func (l DbLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	return DbLogger{
		ZapLogger: l.ZapLogger,
		LogLevel:  level,
	}
}

func (l DbLogger) Error(ctx context.Context, msg string, args ...interface{}) {
	if l.LogLevel < gormLogger.Error {
		return
	}
	l.ZapLogger.Sugar().Errorf(msg, args...)
}

func (l DbLogger) Warn(ctx context.Context, msg string, args ...interface{}) {
	if l.LogLevel < gormLogger.Warn {
		return
	}
	l.ZapLogger.Sugar().Warnf(msg, args...)
}

func (l DbLogger) Info(ctx context.Context, msg string, args ...interface{}) {
	if l.LogLevel < gormLogger.Info {
		return
	}
	l.ZapLogger.Sugar().Infof(msg, args...)
}

func (l DbLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= gormLogger.Error:
		sql, rows := fc()
		l.ZapLogger.Error("trace", zap.Error(err), zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	case l.LogLevel >= gormLogger.Warn:
		sql, rows := fc()
		l.ZapLogger.Warn("trace", zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	case l.LogLevel >= gormLogger.Info:
		sql, rows := fc()
		l.ZapLogger.Warn("trace", zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	}
}
