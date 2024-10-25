package eslog

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"github.com/wzhanjun/dtclients/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	logsChan = make(chan *LogRequest, 100000)
	workNum  = 5
	connLock sync.RWMutex
)

type EslogHandler struct {
	handler.NopFlushClose
	slog.LevelWithFormatter

	client LogClient
}

func NewEslogHandler() *EslogHandler {
	s := &EslogHandler{}
	s.Level = slog.InfoLevel
	go s.push()
	return s
}

func (s *EslogHandler) connect() LogClient {
	connLock.Lock()
	defer connLock.Unlock()

	if s.client == nil {
		conn, err := grpc.NewClient(
			config.Cfg.Slog.Address,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Printf("connect eslog client failed: %v", err)
			return nil
		}
		s.client = NewLogClient(conn)
	}
	return s.client
}

func (s *EslogHandler) Handle(r *slog.Record) error {
	logsChan <- &LogRequest{
		AppId:         config.Cfg.Slog.AppId,
		Label:         StrLabel(r),
		Level:         r.Level.String(),
		Content:       r.Message,
		Caller:        StrCaller(r),
		Datatime:      r.Time.Format("2006/01/02T15:04:05.000"),
		EsIndexPrefix: config.Cfg.Slog.Address,
	}
	return nil
}

func (s *EslogHandler) push() {
	if config.Cfg.Slog.Address == "" {
		for m := range logsChan {
			log.Println(m)
		}
		return
	}

	for i := 0; i < workNum; i++ {
		go func() {
			for m := range logsChan {
				if client := s.connect(); client != nil {
					ctx, cf := context.WithTimeout(context.Background(), time.Second*5)
					defer cf()
					switch m.Level {
					case slog.TraceLevel.Name(), slog.DebugLevel.Name():
						client.Debug(ctx, m)
					case slog.InfoLevel.Name():
						client.Info(ctx, m)
					case slog.WarnLevel.Name(), slog.NoticeLevel.Name():
						client.Warn(ctx, m)
					case slog.ErrorLevel.Name():
						client.Error(ctx, m)
					case slog.FatalLevel.Name(), slog.PanicLevel.Name():
						client.Fatal(ctx, m)
					}
				}
			}
		}()
	}
}
