// Code generated by GoTool. DO NOT EDIT.

package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

type LoggingService struct {
	next   Interface
	logger log.Logger
}

func (s *LoggingService) Resize(ctx context.Context, data []byte, width int, height int) (err error) {
	defer func(begin time.Time) {
		s.logger.Log("ctx", ctx, "data", data, "width", width, "height", height, "err", err, "took", time.Since(begin))
	}(time.Now())
	return s.next.Resize(ctx, data, width, height)
}

func (s *LoggingService) ResizeByURL(ctx context.Context, url string, width int, height int) (err error) {
	defer func(begin time.Time) {
		s.logger.Log("ctx", ctx, "url", url, "width", width, "height", height, "err", err, "took", time.Since(begin))
	}(time.Now())
	return s.next.ResizeByURL(ctx, url, width, height)
}

func NewLoggingService(next Interface, logger log.Logger) *LoggingService {
	return &LoggingService{next: next, logger: logger}
}
