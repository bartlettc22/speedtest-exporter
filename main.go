package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Debug            bool
	ListenPort       int `default:"8080" split_words:"true"`
	IntervalSeconds  int `default:"300" split_words:"true"`
	intervalDuration time.Duration
}

func main() {

	var c Config
	err := envconfig.Process("speedtest", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	if c.Debug {
		log.SetLevel(log.DebugLevel)
	}
	log.Infof("log level: %s", log.GetLevel().String())

	c.intervalDuration = time.Duration(time.Duration(c.IntervalSeconds) * time.Second)
	log.Infof("interval: %s", c.intervalDuration.String())

	// Initial test on startup
	log.Infof("running initial test")
	for {
		if updateSpeedTestResults() {
			break
		}
		time.Sleep(10 * time.Second)
	}

	go func() {
		for {
			time.Sleep(c.intervalDuration)
			updateSpeedTestResults()
		}
	}()

	log.Infof("listening on :%d", c.ListenPort)
	http.Handle("/metrics", promhttp.Handler())
	err = http.ListenAndServe(fmt.Sprintf(":%d", c.ListenPort), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func updateSpeedTestResults() (success bool) {

	log.Debug("running speed test")
	start := time.Now()

	cmd := exec.Command("./speedtest", "--accept-license", "--format", "json-pretty")
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		log.Warn(err)
		log.Warnf("stdout %s", outb.String())
		log.Warnf("stderr %s", errb.String())
		metricErrors.WithLabelValues("cmd_error").Inc()
		metricErrorTimestamp.Set(float64(time.Now().Unix()))
		metricRuns.WithLabelValues("failed").Inc()
		return false
	}

	speedTestResults := &SpeedTestResults{}

	failed := false
	if outb.String() == "" {
		log.Warn("empty results")
		metricErrors.WithLabelValues("empty_results").Inc()
		failed = true
	}

	if errb.String() != "" {
		log.Warnf("error fetching results: %v", errb.String())
		metricErrors.WithLabelValues("error").Inc()
		failed = true
	}

	if failed {
		metricErrorTimestamp.Set(float64(time.Now().Unix()))
		metricRuns.WithLabelValues("failed").Inc()
		return false
	}

	err = json.Unmarshal(outb.Bytes(), speedTestResults)
	if err != nil {
		metricErrors.WithLabelValues("unmarshal_error").Inc()
		metricErrorTimestamp.Set(float64(time.Now().Unix()))
		metricRuns.WithLabelValues("failed").Inc()
		return false
	}

	// Run Info
	runDuration := time.Since(start)
	metricRunDuration.Observe(runDuration.Seconds())
	metricLastRun.Set(float64(speedTestResults.Timestamp.Unix()))

	// Ping
	metricPingJitter.Set(float64(speedTestResults.Ping.Jitter / 1000))
	metricPingLatency.Set(float64(speedTestResults.Ping.Latency / 1000))
	metricPingLow.Set(float64(speedTestResults.Ping.Low / 1000))
	metricPingHigh.Set(float64(speedTestResults.Ping.High / 1000))

	// Download
	metricDownloadBandwidth.Set(float64(speedTestResults.Download.Bandwidth))
	metricDownloadBytes.Add(float64(speedTestResults.Download.Bytes))
	metricDownloadElapsed.Add(float64(speedTestResults.Download.Elapsed))
	metricDownloadLatencyIQM.Set(float64(speedTestResults.Download.Latency.IQM / 1000))
	metricDownloadLatencyLow.Set(float64(speedTestResults.Download.Latency.Low / 1000))
	metricDownloadLatencyHigh.Set(float64(speedTestResults.Download.Latency.High / 1000))
	metricDownloadLatencyJitter.Set(float64(speedTestResults.Download.Latency.Jitter / 1000))

	// Upload
	metricUploadBandwidth.Set(float64(speedTestResults.Upload.Bandwidth))
	metricUploadBytes.Add(float64(speedTestResults.Upload.Bytes))
	metricUploadElapsed.Add(float64(speedTestResults.Upload.Elapsed))
	metricUploadLatencyIQM.Set(float64(speedTestResults.Upload.Latency.IQM / 1000))
	metricUploadLatencyLow.Set(float64(speedTestResults.Upload.Latency.Low / 1000))
	metricUploadLatencyHigh.Set(float64(speedTestResults.Upload.Latency.High / 1000))
	metricUploadLatencyJitter.Set(float64(speedTestResults.Upload.Latency.Jitter / 1000))

	// General Info
	metricPacketLoss.Set(float64(speedTestResults.PacketLoss))
	metricInfo.Reset()
	metricInfo.WithLabelValues(speedTestResults.ISP, speedTestResults.Iface.ExternalIP).Set(1)
	isVPN := 0
	if speedTestResults.Iface.IsVPN {
		isVPN = 1
	}
	metricVPNStatus.Set(float64(isVPN))
	metricServerInfo.Reset()
	metricServerInfo.WithLabelValues(
		fmt.Sprintf("%d", speedTestResults.Server.ID),
		speedTestResults.Server.Host,
		fmt.Sprintf("%d", speedTestResults.Server.Port),
		speedTestResults.Server.Name,
		speedTestResults.Server.Location,
		speedTestResults.Server.Country,
		speedTestResults.Server.IP,
	).Set(1)
	metricRuns.WithLabelValues("succeeded").Inc()

	log.Debugf("speed test completed in %s", runDuration.String())

	return true
}
