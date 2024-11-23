package middleware

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/mochammadshenna/aplikasi-po/internal/state"
	"github.com/mochammadshenna/aplikasi-po/internal/util/logger"
	"github.com/newrelic/go-agent/v3/integrations/logcontext"
	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/sirupsen/logrus"
)

func PostHandle(writer *responseWriter, request *http.Request) {
	startTimeHeader := writer.Header().Get(state.HttpHeaders().StartTime.String())
	startTime, _ := strconv.ParseInt(startTimeHeader, 10, 64)

	md := newrelic.FromContext(request.Context()).GetLinkingMetadata()

	enableRequestLog, _ := strconv.ParseBool(os.Getenv("ENABLE_REQUEST_LOG"))
	re := regexp.MustCompile(`^5\d{2}$`)
	if re.MatchString(strconv.Itoa(writer.Status())) || enableRequestLog {
		logger.Logger.WithFields(logrus.Fields{
			logger.LoggerField().RequestId:     writer.Header().Get(state.HttpHeaders().RequestId.String()),
			logger.LoggerField().Latency:       time.Since(time.Unix(0, startTime)).String(),
			logger.LoggerField().RequestMethod: request.Method,
			logger.LoggerField().Resource:      request.URL.Path,
			logger.LoggerField().UserAgent:     request.Header.Get(state.HttpHeaders().UserAgent.String()),
			logger.LoggerField().PlatformType:  request.Context().Value(state.HttpHeaders().PlatformType),
			logger.LoggerField().Platform:      request.Context().Value(state.HttpHeaders().Platform),
			logger.LoggerField().Version:       request.Context().Value(state.HttpHeaders().Version),
			logger.LoggerField().Status:        writer.Status(),
			logger.LoggerField().XForwardedFor: request.Header.Get(state.HttpHeaders().XForwardedFor.String()),
			logcontext.KeyTraceID:              md.TraceID,
			logcontext.KeySpanID:               md.SpanID,
			logcontext.KeyEntityName:           md.EntityName,
			logcontext.KeyEntityType:           md.EntityType,
			logcontext.KeyEntityGUID:           md.EntityGUID,
			logcontext.KeyHostname:             md.Hostname,
		}).Info(fmt.Sprintf("HTTP request, %s %s", request.Method, request.URL.Path))
	}
}
