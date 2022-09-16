package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

const namespace = "mirth"

var (
	up = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "up"),
		"Was the last Mirth query successful",
		nil, nil,
	)

	messagesReceived = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "message_recvieved_total"),
		"how many message have been received (per channel)",
		[]string{"channel"}, nil,
	)
)

type MexportCollector struct {
	username string
}

func NewMexportCollector(username string) *MexportCollector {
	return &MexportCollector{
		username,
	}
}

func (m *MexportCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- up
	ch <- messagesReceived
}

func (m *MexportCollector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(
		up,
		prometheus.GaugeValue,
		1,
	)
}
