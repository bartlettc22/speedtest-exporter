# Speedtest Export Helm Chart

## Installation
From this directory
```bash
git clone https://github.com/bartlettc22/speedtest-exporter.git
cd speedtest-exporter/deploy/chart

helm upgrade --install speedtest-exporter ./
```

## Service Monitor
Prometheus `ServiceMonitor` can be enabled by setting `serviceMonitor.enabled=true` in the values file.
