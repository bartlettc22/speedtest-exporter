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
	ListenPort       int    `default:"8080" split_words:"true"`
	IntervalSeconds  int    `default:"3600" split_words:"true"`
	NameLabel        string `default:"default" split_words:"true"`
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

	// Set default status metrics labels with zero value
	metricRuns.WithLabelValues(c.NameLabel, "succeeded")
	metricRuns.WithLabelValues(c.NameLabel, "failed")
	metricRunDuration.WithLabelValues(c.NameLabel, "succeeded")
	metricRunDuration.WithLabelValues(c.NameLabel, "failed")
	metricLastRun.WithLabelValues(c.NameLabel, "succeeded")
	metricLastRun.WithLabelValues(c.NameLabel, "failed")

	// Initial test on startup
	log.Infof("running initial test")
	for {
		if updateSpeedTestResults(&c) {
			break
		}
		time.Sleep(10 * time.Second)
	}

	go func() {
		for {
			time.Sleep(c.intervalDuration)
			updateSpeedTestResults(&c)
		}
	}()

	log.Infof("listening on :%d", c.ListenPort)
	http.Handle("/metrics", promhttp.Handler())
	err = http.ListenAndServe(fmt.Sprintf(":%d", c.ListenPort), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func updateSpeedTestResults(c *Config) (success bool) {

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
		metricLastRun.WithLabelValues(c.NameLabel, "failed").Set(float64(time.Now().Unix()))
		metricRuns.WithLabelValues(c.NameLabel, "failed").Inc()
		metricRunDuration.WithLabelValues(c.NameLabel, "failed").Observe(time.Since(start).Seconds())
		return false
	}

	speedTestResults := &SpeedTestResults{}

	if outb.String() == "" {
		log.Warn("empty results")
		metricLastRun.WithLabelValues(c.NameLabel, "failed").Set(float64(time.Now().Unix()))
		metricRuns.WithLabelValues(c.NameLabel, "failed").Inc()
		metricRunDuration.WithLabelValues(c.NameLabel, "failed").Observe(time.Since(start).Seconds())
		return false
	}

	err = json.Unmarshal(outb.Bytes(), speedTestResults)
	if err != nil {
		log.Warnf("unmarshal error: %v", err)
		metricLastRun.WithLabelValues(c.NameLabel, "failed").Set(float64(time.Now().Unix()))
		metricRuns.WithLabelValues(c.NameLabel, "failed").Inc()
		metricRunDuration.WithLabelValues(c.NameLabel, "failed").Observe(time.Since(start).Seconds())
		return false
	}

	if speedTestResults.Type != "result" {
		log.Warnf("non-result response: %s", outb.Bytes())
		metricLastRun.WithLabelValues(c.NameLabel, "failed").Set(float64(time.Now().Unix()))
		metricRuns.WithLabelValues(c.NameLabel, "failed").Inc()
		metricRunDuration.WithLabelValues(c.NameLabel, "failed").Observe(time.Since(start).Seconds())
		return false
	}

	// Run Info
	runDuration := time.Since(start)
	metricRunDuration.WithLabelValues(c.NameLabel, "succeeded").Observe(runDuration.Seconds())
	metricLastRun.WithLabelValues(c.NameLabel, "succeeded").Set(float64(speedTestResults.Timestamp.Unix()))

	// Ping
	metricPingJitter.WithLabelValues(c.NameLabel).Set(float64(speedTestResults.Ping.Jitter / 1000))
	metricPingLatency.WithLabelValues(c.NameLabel).Set(float64(speedTestResults.Ping.Latency / 1000))
	metricPingLow.WithLabelValues(c.NameLabel).Set(float64(speedTestResults.Ping.Low / 1000))
	metricPingHigh.WithLabelValues(c.NameLabel).Set(float64(speedTestResults.Ping.High / 1000))

	// Download
	metricDownloadBandwidth.WithLabelValues(c.NameLabel).Set(float64(speedTestResults.Download.Bandwidth))
	metricDownloadBytes.WithLabelValues(c.NameLabel).Add(float64(speedTestResults.Download.Bytes))
	metricDownloadElapsed.WithLabelValues(c.NameLabel).Add(float64(speedTestResults.Download.Elapsed / 1000))
	metricDownloadLatencyIQM.WithLabelValues(c.NameLabel).Set(float64(speedTestResults.Download.Latency.IQM / 1000))
	metricDownloadLatencyLow.WithLabelValues(c.NameLabel).Set(float64(speedTestResults.Download.Latency.Low / 1000))
	metricDownloadLatencyHigh.WithLabelValues(c.NameLabel).Set(float64(speedTestResults.Download.Latency.High / 1000))
	metricDownloadLatencyJitter.WithLabelValues(c.NameLabel).Set(float64(speedTestResults.Download.Latency.Jitter / 1000))

	// Upload
	metricUploadBandwidth.WithLabelValues(c.NameLabel).Set(float64(speedTestResults.Upload.Bandwidth))
	metricUploadBytes.WithLabelValues(c.NameLabel).Add(float64(speedTestResults.Upload.Bytes))
	metricUploadElapsed.WithLabelValues(c.NameLabel).Add(float64(speedTestResults.Upload.Elapsed / 1000))
	metricUploadLatencyIQM.WithLabelValues(c.NameLabel).Set(float64(speedTestResults.Upload.Latency.IQM / 1000))
	metricUploadLatencyLow.WithLabelValues(c.NameLabel).Set(float64(speedTestResults.Upload.Latency.Low / 1000))
	metricUploadLatencyHigh.WithLabelValues(c.NameLabel).Set(float64(speedTestResults.Upload.Latency.High / 1000))
	metricUploadLatencyJitter.WithLabelValues(c.NameLabel).Set(float64(speedTestResults.Upload.Latency.Jitter / 1000))

	// General Info
	metricPacketLoss.WithLabelValues(c.NameLabel).Set(float64(speedTestResults.PacketLoss))
	metricInfo.Reset()
	metricInfo.WithLabelValues(c.NameLabel, speedTestResults.ISP, speedTestResults.Iface.ExternalIP).Set(1)
	isVPN := 0
	if speedTestResults.Iface.IsVPN {
		isVPN = 1
	}
	metricVPNStatus.WithLabelValues(c.NameLabel).Set(float64(isVPN))
	metricServerInfo.Reset()
	metricServerInfo.WithLabelValues(
		c.NameLabel,
		fmt.Sprintf("%d", speedTestResults.Server.ID),
		speedTestResults.Server.Host,
		fmt.Sprintf("%d", speedTestResults.Server.Port),
		speedTestResults.Server.Name,
		speedTestResults.Server.Location,
		speedTestResults.Server.Country,
		speedTestResults.Server.IP,
	).Set(1)
	metricRuns.WithLabelValues(c.NameLabel, "succeeded").Inc()

	log.Debugf("speed test completed in %s", runDuration.String())

	return true
}
