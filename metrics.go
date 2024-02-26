package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (

	// Run Info
	metricRuns = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "speedtest_runs",
		Help: "The number of total successful speed test runs",
	}, []string{"status"})
	metricRunDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "speedtest_run_duration_seconds",
		Help:    "The run duration of successful speed tests",
		Buckets: prometheus.LinearBuckets(0, 5, 8),
	})
	metricLastRun = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "speedtest_last_run_timestamp",
		Help: "The timestamp of the last successful speed test",
	})

	// Ping
	metricPingJitter = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "speedtest_ping_jitter_seconds",
		Help: "The ping jitter of last successful speed test",
	})
	metricPingLatency = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "speedtest_ping_latency_seconds",
		Help: "The ping latency of last successful speed test",
	})
	metricPingLow = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "speedtest_ping_low_seconds",
		Help: "The ping max of last successful speed test",
	})
	metricPingHigh = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "speedtest_ping_high_seconds",
		Help: "The ping min of last successful speed test",
	})

	// Download
	metricDownloadBandwidth = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "speedtest_download_bandwidth",
		Help: "The download bandwidth of last successful speed test",
	})
	metricDownloadBytes = promauto.NewCounter(prometheus.CounterOpts{
		Name: "speedtest_download_bytes",
		Help: "The downloaded bytes of last successful speed test",
	})
	metricDownloadElapsed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "speedtest_download_elapsed",
		Help: "The elapsed download time of last successful speed test",
	})
	metricDownloadLatencyIQM = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "speedtest_download_latency_iqm_seconds",
		Help: "The download interquartile mean latency of last successful speed test",
	})
	metricDownloadLatencyLow = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "speedtest_download_latency_low_seconds",
		Help: "The min download latency of last successful speed test",
	})
	metricDownloadLatencyHigh = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "speedtest_download_latency_high_seconds",
		Help: "The max download latency of last successful speed test",
	})
	metricDownloadLatencyJitter = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "speedtest_download_latency_jitter_seconds",
		Help: "The download latency jitter of last successful speed test",
	})

	// Upload
	metricUploadBandwidth = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "speedtest_upload_bandwidth",
		Help: "The upload bandwidth of last successful speed test",
	})
	metricUploadBytes = promauto.NewCounter(prometheus.CounterOpts{
		Name: "speedtest_upload_bytes",
		Help: "The uploaded bytes of last successful speed test",
	})
	metricUploadElapsed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "speedtest_upload_elapsed",
		Help: "The elapsed upload time of last successful speed test",
	})
	metricUploadLatencyIQM = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "speedtest_upload_latency_iqm_seconds",
		Help: "The upload interquartile mean latency of last successful speed test",
	})
	metricUploadLatencyLow = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "speedtest_upload_latency_low_seconds",
		Help: "The min upload latency of last successful speed test",
	})
	metricUploadLatencyHigh = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "speedtest_upload_latency_high_seconds",
		Help: "The max upload latency of last successful speed test",
	})
	metricUploadLatencyJitter = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "speedtest_upload_latency_jitter_seconds",
		Help: "The upload latency jitter of last successful speed test",
	})

	// General Info
	metricPacketLoss = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "speedtest_packet_loss",
		Help: "The packet loss percentage of last successful speed test",
	},
	)
	metricInfo = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "speedtest_isp_info",
		Help: "Internet service provider info of last successful speed test",
	}, []string{"isp", "external_ip"})
	metricVPNStatus = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "speedtest_vpn_status",
		Help: "Indicates VPN status of last successful speed test",
	})
	metricServerInfo = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "speedtest_server_info",
		Help: "Server info labels of last successful speed test",
	}, []string{"id", "host", "port", "name", "location", "country", "ip"})

	// Errors
	metricErrorTimestamp = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "speedtest_last_error_timestamp",
		Help: "The timestamp of the last unsuccessful speed test",
	})
	metricErrors = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "speedtest_errors",
		Help: "Errors encountered during speed test execution",
	}, []string{"reason"})
)
