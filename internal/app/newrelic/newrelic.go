package newrelic

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/mochammadshenna/aplikasi-po/internal/state"
	"github.com/mochammadshenna/aplikasi-po/internal/util/authentication"
	"github.com/mochammadshenna/aplikasi-po/internal/util/common"
	"github.com/mochammadshenna/aplikasi-po/internal/util/helper"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var (
	App                *newrelic.Application
	newrelicCredential = os.Getenv("NEWRELIC_KEY")
)

func Init() {
	appName := "aplikasi_po_production"
	if appName == "" {
		return
	}
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(newrelicCredential),
		newrelic.ConfigDistributedTracerEnabled(true),
		func(c *newrelic.Config) {
			// set newrelic hostname based on environment and deployed time
			loc, err := time.LoadLocation("")
			if err != nil {
				panic(err)
			}
			c.HostDisplayName = fmt.Sprintf("%s-%s", state.App.Environment, time.Now().In(loc).Format("20060102_150405_9999Z0700"))
			c.CrossApplicationTracer.Enabled = true
		},
	)
	helper.PanicError(err)
	App = app

}

func StartSegmentWithFuncName(ctx context.Context) *newrelic.Segment {
	segmentName := common.FuncCallerName(3)
	return newrelic.FromContext(ctx).StartSegment(segmentName)
}

func LogUserID(r *http.Request) {
	txn := newrelic.FromContext(r.Context())
	userClaim := authentication.ExtractClaim(r)

	if userClaim.IDUser != 0 {
		txn.AddAttribute("user-ID", userClaim.IDUser)
	}
}
