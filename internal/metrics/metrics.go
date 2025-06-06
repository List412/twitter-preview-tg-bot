package metrics

import (
	"context"
	"github.com/list412/tweets-tg-bot/internal/commands"
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	commandsMetric *prometheus.CounterVec
}

func (m *Metrics) Register(reg *prometheus.Registry) {
	counter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "tg_tweet_bot",
		Name:      "cmd_usage",
		Help:      "cmd usage counter",
	}, []string{"cmd"})

	reg.MustRegister(counter)
	m.commandsMetric = counter
}

func (m *Metrics) HandleCmd(ctx context.Context, cmd commands.Cmd) {
	m.commandsMetric.With(prometheus.Labels{"cmd": string(cmd)}).Inc()
}

func (m *Metrics) GetCmdStats(ctx context.Context, command string) (prometheus.Counter, error) {
	return m.commandsMetric.GetMetricWithLabelValues(command)
}
