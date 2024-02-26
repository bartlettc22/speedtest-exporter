package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (

	// Run Info
	metricRuns = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "speedtest_run_total",
		Help: "The number of total speed test runs",
	}, []string{"name", "status"})
	metricRunDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "speedtest_run_duration_seconds",
		Help:    "The run duration of the speed tests",
		Buckets: prometheus.LinearBuckets(0, 5, 8),
	}, []string{"name", "status"})
	metricLastRun = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "speedtest_last_run_timestamp",
		Help: "The timestamp of the last speed test",
	}, []string{"name", "status"})

	// Ping
	metricPingJitter = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "speedtest_ping_jitter_seconds",
		Help: "The idle ping jitter of last successful speed test",
	}, []string{"name"})
	metricPingLatency = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "speedtest_ping_latency_seconds",
		Help: "The idle ping latency of last successful speed test",
	}, []string{"name"})
	metricPingLow = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "speedtest_ping_low_seconds",
		Help: "The idle ping max of last successful speed test",
	}, []string{"name"})
	metricPingHigh = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "speedtest_ping_high_seconds",
		Help: "The idle ping min of last successful speed test",
	}, []string{"name"})

	// Download
	metricDownloadBandwidth = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "speedtest_download_bandwidth",
		Help: "The download bandwidth, in bytes/s, of last successful speed test",
	}, []string{"name"})
	metricDownloadBytes = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "speedtest_download_bytes",
		Help: "The downloaded bytes of last successful speed test",
	}, []string{"name"})
	metricDownloadElapsed = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "speedtest_download_elapsed_seconds",
		Help: "The elapsed download time of last successful speed test",
	}, []string{"name"})
	metricDownloadLatencyIQM = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "speedtest_download_latency_iqm_seconds",
		Help: "The download interquartile mean latency of last successful speed test",
	}, []string{"name"})
	metricDownloadLatencyLow = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "speedtest_download_latency_low_seconds",
		Help: "The min download latency of last successful speed test",
	}, []string{"name"})
	metricDownloadLatencyHigh = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "speedtest_download_latency_high_seconds",
		Help: "The max download latency of last successful speed test",
	}, []string{"name"})
	metricDownloadLatencyJitter = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "speedtest_download_latency_jitter_seconds",
		Help: "The download latency jitter of last successful speed test",
	}, []string{"name"})

	// Upload
	metricUploadBandwidth = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "speedtest_upload_bandwidth",
		Help: "The upload bandwidth, in bytes/s, of last successful speed test",
	}, []string{"name"})
	metricUploadBytes = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "speedtest_upload_bytes",
		Help: "The uploaded bytes of last successful speed test",
	}, []string{"name"})
	metricUploadElapsed = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "speedtest_upload_elapsed_seconds",
		Help: "The elapsed upload time of last successful speed test",
	}, []string{"name"})
	metricUploadLatencyIQM = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "speedtest_upload_latency_iqm_seconds",
		Help: "The upload interquartile mean latency of last successful speed test",
	}, []string{"name"})
	metricUploadLatencyLow = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "speedtest_upload_latency_low_seconds",
		Help: "The min upload latency of last successful speed test",
	}, []string{"name"})
	metricUploadLatencyHigh = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "speedtest_upload_latency_high_seconds",
		Help: "The max upload latency of last successful speed test",
	}, []string{"name"})
	metricUploadLatencyJitter = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "speedtest_upload_latency_jitter_seconds",
		Help: "The upload latency jitter of last successful speed test",
	}, []string{"name"})

	// General Info
	metricPacketLoss = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "speedtest_packet_loss",
		Help: "The packet loss percentage of last successful speed test",
	}, []string{"name"})
	metricInfo = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "speedtest_isp_info",
		Help: "Internet service provider info of last successful speed test",
	}, []string{"name", "isp", "external_ip"})
	metricVPNStatus = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "speedtest_vpn_status",
		Help: "Indicates VPN status of last successful speed test",
	}, []string{"name"})
	metricServerInfo = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "speedtest_server_info",
		Help: "Server info labels of last successful speed test",
	}, []string{"name", "id", "host", "port", "server_name", "location", "country", "ip"})
)
