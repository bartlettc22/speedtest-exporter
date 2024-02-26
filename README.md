# Speed Test Exporter
Runs a [Ookla Speedtest®](https://www.speedtest.net/) at regular intervals and exports the results in Prometheus format.

## Config
Configuration is through the following environment variables

|Env Var|Description|Type|Default|
|-|-|-|-|
|`SPEEDTEST_DEBUG`|Enables debug logs|bool|`false`|
|`SPEEDTEST_LISTEN_PORT`|Listen port for metrics|int|`8080`|
|`SPEEDTEST_INTERVAL_SECONDS`|Time period between tests, in seconds|int|`300`|

## Metrics
```
# HELP speedtest_download_bandwidth The download bandwidth of last successful speed test
# TYPE speedtest_download_bandwidth gauge
speedtest_download_bandwidth 9.067908e+07
# HELP speedtest_download_bytes The downloaded bytes of last successful speed test
# TYPE speedtest_download_bytes counter
speedtest_download_bytes 8.19117752e+08
# HELP speedtest_download_elapsed The elapsed download time of last successful speed test
# TYPE speedtest_download_elapsed counter
speedtest_download_elapsed 8914
# HELP speedtest_download_latency_high_seconds The max download latency of last successful speed test
# TYPE speedtest_download_latency_high_seconds gauge
speedtest_download_latency_high_seconds 0.06931400299072266
# HELP speedtest_download_latency_iqm_seconds The download interquartile mean latency of last successful speed test
# TYPE speedtest_download_latency_iqm_seconds gauge
speedtest_download_latency_iqm_seconds 0.004543000366538763
# HELP speedtest_download_latency_jitter_seconds The download latency jitter of last successful speed test
# TYPE speedtest_download_latency_jitter_seconds gauge
speedtest_download_latency_jitter_seconds 0.003743999870494008
# HELP speedtest_download_latency_low_seconds The min download latency of last successful speed test
# TYPE speedtest_download_latency_low_seconds gauge
speedtest_download_latency_low_seconds 0.0019859999883919954
# HELP speedtest_isp_info Internet service provider info of last successful speed test
# TYPE speedtest_isp_info gauge
speedtest_isp_info{external_ip="x.x.x.x",isp="Comcast"} 1
# HELP speedtest_last_error_timestamp The timestamp of the last unsuccessful speed test
# TYPE speedtest_last_error_timestamp gauge
speedtest_last_error_timestamp 0
# HELP speedtest_last_run_timestamp The timestamp of the last successful speed test
# TYPE speedtest_last_run_timestamp gauge
speedtest_last_run_timestamp 1.708907441e+09
# HELP speedtest_packet_loss The packet loss percentage of last successful speed test
# TYPE speedtest_packet_loss gauge
speedtest_packet_loss 0
# HELP speedtest_ping_high_seconds The ping min of last successful speed test
# TYPE speedtest_ping_high_seconds gauge
speedtest_ping_high_seconds 0.0029159998521208763
# HELP speedtest_ping_jitter_seconds The ping jitter of last successful speed test
# TYPE speedtest_ping_jitter_seconds gauge
speedtest_ping_jitter_seconds 0.00033500001882202923
# HELP speedtest_ping_latency_seconds The ping latency of last successful speed test
# TYPE speedtest_ping_latency_seconds gauge
speedtest_ping_latency_seconds 0.0020680001471191645
# HELP speedtest_ping_low_seconds The ping max of last successful speed test
# TYPE speedtest_ping_low_seconds gauge
speedtest_ping_low_seconds 0.0018469999777153134
# HELP speedtest_run_duration_seconds The run duration of successful speed tests
# TYPE speedtest_run_duration_seconds histogram
speedtest_run_duration_seconds_bucket{le="0"} 0
speedtest_run_duration_seconds_bucket{le="5"} 0
speedtest_run_duration_seconds_bucket{le="10"} 0
speedtest_run_duration_seconds_bucket{le="15"} 0
speedtest_run_duration_seconds_bucket{le="20"} 1
speedtest_run_duration_seconds_bucket{le="25"} 1
speedtest_run_duration_seconds_bucket{le="30"} 1
speedtest_run_duration_seconds_bucket{le="35"} 1
speedtest_run_duration_seconds_bucket{le="+Inf"} 1
speedtest_run_duration_seconds_sum 17.879575701
speedtest_run_duration_seconds_count 1
# HELP speedtest_runs The number of total successful speed test runs
# TYPE speedtest_runs counter
speedtest_runs{status="succeeded"} 1
# HELP speedtest_server_info Server info labels of last successful speed test
# TYPE speedtest_server_info gauge
speedtest_server_info{country="United States",host="speedtest-dnvr.mynextlight.net",id="47683",ip="66.186.203.166",location="Denver, CO",name="NextLight",port="8080"} 1
# HELP speedtest_upload_bandwidth The upload bandwidth of last successful speed test
# TYPE speedtest_upload_bandwidth gauge
speedtest_upload_bandwidth 1.15750378e+08
# HELP speedtest_upload_bytes The uploaded bytes of last successful speed test
# TYPE speedtest_upload_bytes counter
speedtest_upload_bytes 9.19615914e+08
# HELP speedtest_upload_elapsed The elapsed upload time of last successful speed test
# TYPE speedtest_upload_elapsed counter
speedtest_upload_elapsed 8008
# HELP speedtest_upload_latency_high_seconds The max upload latency of last successful speed test
# TYPE speedtest_upload_latency_high_seconds gauge
speedtest_upload_latency_high_seconds 0.16744199395179749
# HELP speedtest_upload_latency_iqm_seconds The upload interquartile mean latency of last successful speed test
# TYPE speedtest_upload_latency_iqm_seconds gauge
speedtest_upload_latency_iqm_seconds 0.0967710018157959
# HELP speedtest_upload_latency_jitter_seconds The upload latency jitter of last successful speed test
# TYPE speedtest_upload_latency_jitter_seconds gauge
speedtest_upload_latency_jitter_seconds 0.026027999818325043
# HELP speedtest_upload_latency_low_seconds The min upload latency of last successful speed test
# TYPE speedtest_upload_latency_low_seconds gauge
speedtest_upload_latency_low_seconds 0.0032790000550448895
# HELP speedtest_vpn_status Indicates VPN status of last successful speed test
# TYPE speedtest_vpn_status gauge
speedtest_vpn_status 0
```
