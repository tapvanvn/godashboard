package godashboard

import (
	"github.com/tapvanvn/gopubsubengine"
)

type DashboardReporter interface {
	Report(signal *Signal)
}

func NewPubsubDashboardReporter(hub gopubsubengine.Hub) *PubsubDashboardReporter {
	publisher, err := hub.PublishOn("dashboard")
	if err != nil {
		return nil
	}
	return &PubsubDashboardReporter{
		publisher: publisher,
	}
}

type PubsubDashboardReporter struct {
	publisher gopubsubengine.Publisher
}

func (dbr *PubsubDashboardReporter) Report(signal *Signal) {
	dbr.publisher.Publish(signal)
}
