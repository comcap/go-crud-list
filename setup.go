package main

import (
	"context"
	"io"
	"log"
	validatorService "meeting-room/service/validator"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"

	"meeting-room/app"
	"meeting-room/service/util"

	jaegerConf "github.com/uber/jaeger-client-go/config"
	jaegerLog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"

	"meeting-room/config"
	eventRepo "meeting-room/repository/event"
	eventService "meeting-room/service/event/implement"
)

func setupJaeger(appConfig *config.Config) io.Closer {
	cfg, err := jaegerConf.FromEnv()
	panicIfErr(err)

	cfg.ServiceName = appConfig.AppName
	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1
	cfg.Reporter = &jaegerConf.ReporterConfig{LogSpans: true}

	jLogger := jaegerLog.StdLogger
	jMetricsFactory := metrics.NullFactory

	tracer, closer, err := cfg.NewTracer(
		jaegerConf.Logger(jLogger),
		jaegerConf.Metrics(jMetricsFactory),
	)
	panicIfErr(err)
	opentracing.SetGlobalTracer(tracer)

	return closer
}

func newApp(appConfig *config.Config) *app.App {
	ctx := context.Background()

	eRepo, err := eventRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBEventTableName)
	panicIfErr(err)

	validator := validatorService.New(eRepo)
	generateID, err := util.NewUUID()
	panicIfErr(err)

	event := eventService.New(validator, eRepo, generateID)

	return app.New(event, appConfig)
}

func setupLog() *logrus.Logger {
	lr := logrus.New()
	lr.SetFormatter(&logrus.JSONFormatter{})

	return lr
}

func panicIfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
