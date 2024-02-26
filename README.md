# Speed Test Exporter
Runs a [Ookla Speedtest®](https://www.speedtest.net/) at regular intervals and exports the results in Prometheus format.

## Config
Configuration is through the following environment variables

|Env Var|Description|Type|Default|
|-|-|-|-|
|`SPEEDTEST_DEBUG`|Enables debug logs|bool|`false`|
|`SPEEDTEST_LISTEN_PORT`|Listen port for metrics|int|`8080`|
|`SPEEDTEST_INTERVAL_SECONDS`|Time period between tests, in seconds|int|`300`|
|`SPEEDTEST_NAME_LABEL`|Value that is filled in for `name` label on all metrics.  Intended to be used to distinguish between tests in different environments.|string|`default`|

## Metrics
```
# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 0
go_gc_duration_seconds{quantile="0.25"} 0
go_gc_duration_seconds{quantile="0.5"} 0
go_gc_duration_seconds{quantile="0.75"} 0
go_gc_duration_seconds{quantile="1"} 0
go_gc_duration_seconds_sum 0
go_gc_duration_seconds_count 0
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 7
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.21.7"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 235456
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 235456
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 4282
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 0
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 2.477616e+06
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 235456
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 2.097152e+06
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 1.671168e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 603
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 2.097152e+06
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 3.76832e+06
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 0
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 0
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 603
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 9600
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 15600
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 56280
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 65184
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.194304e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 1.212822e+06
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 425984
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 425984
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 7.969808e+06
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 5
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 0.02
# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge
process_max_fds 524288
# HELP process_open_fds Number of open file descriptors.
# TYPE process_open_fds gauge
process_open_fds 9
# HELP process_resident_memory_bytes Resident memory size in bytes.
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 7.671808e+06
# HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1.70896529397e+09
# HELP process_virtual_memory_bytes Virtual memory size in bytes.
# TYPE process_virtual_memory_bytes gauge
process_virtual_memory_bytes 1.263747072e+09
# HELP process_virtual_memory_max_bytes Maximum amount of virtual memory available in bytes.
# TYPE process_virtual_memory_max_bytes gauge
process_virtual_memory_max_bytes 1.8446744073709552e+19
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 0
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
# HELP speedtest_download_bandwidth The download bandwidth, in bytes/s, of last successful speed test
# TYPE speedtest_download_bandwidth gauge
speedtest_download_bandwidth{name="default"} 1.17641899e+08
# HELP speedtest_download_bytes The downloaded bytes of last successful speed test
# TYPE speedtest_download_bytes counter
speedtest_download_bytes{name="default"} 4.2297528e+08
# HELP speedtest_download_elapsed_seconds The elapsed download time of last successful speed test
# TYPE speedtest_download_elapsed_seconds counter
speedtest_download_elapsed_seconds{name="default"} 3
# HELP speedtest_download_latency_high_seconds The max download latency of last successful speed test
# TYPE speedtest_download_latency_high_seconds gauge
speedtest_download_latency_high_seconds{name="default"} 0.012550000101327896
# HELP speedtest_download_latency_iqm_seconds The download interquartile mean latency of last successful speed test
# TYPE speedtest_download_latency_iqm_seconds gauge
speedtest_download_latency_iqm_seconds{name="default"} 0.0081619992852211
# HELP speedtest_download_latency_jitter_seconds The download latency jitter of last successful speed test
# TYPE speedtest_download_latency_jitter_seconds gauge
speedtest_download_latency_jitter_seconds{name="default"} 0.0007749999640509486
# HELP speedtest_download_latency_low_seconds The min download latency of last successful speed test
# TYPE speedtest_download_latency_low_seconds gauge
speedtest_download_latency_low_seconds{name="default"} 0.0017389999702572823
# HELP speedtest_isp_info Internet service provider info of last successful speed test
# TYPE speedtest_isp_info gauge
speedtest_isp_info{external_ip="x.x.x.x",isp="Comcast",name="default"} 1
# HELP speedtest_last_run_timestamp The timestamp of the last speed test
# TYPE speedtest_last_run_timestamp gauge
speedtest_last_run_timestamp{name="default",status="failed"} 0
speedtest_last_run_timestamp{name="default",status="succeeded"} 1.708965304e+09
# HELP speedtest_packet_loss The packet loss percentage of last successful speed test
# TYPE speedtest_packet_loss gauge
speedtest_packet_loss{name="default"} 0
# HELP speedtest_ping_high_seconds The idle ping min of last successful speed test
# TYPE speedtest_ping_high_seconds gauge
speedtest_ping_high_seconds{name="default"} 0.0020860000513494015
# HELP speedtest_ping_jitter_seconds The idle ping jitter of last successful speed test
# TYPE speedtest_ping_jitter_seconds gauge
speedtest_ping_jitter_seconds{name="default"} 9.200000204145908e-05
# HELP speedtest_ping_latency_seconds The idle ping latency of last successful speed test
# TYPE speedtest_ping_latency_seconds gauge
speedtest_ping_latency_seconds{name="default"} 0.0020129999611526728
# HELP speedtest_ping_low_seconds The idle ping max of last successful speed test
# TYPE speedtest_ping_low_seconds gauge
speedtest_ping_low_seconds{name="default"} 0.0019520000787451863
# HELP speedtest_run_duration_seconds The run duration of the speed tests
# TYPE speedtest_run_duration_seconds histogram
speedtest_run_duration_seconds_bucket{name="default",status="failed",le="0"} 0
speedtest_run_duration_seconds_bucket{name="default",status="failed",le="5"} 0
speedtest_run_duration_seconds_bucket{name="default",status="failed",le="10"} 0
speedtest_run_duration_seconds_bucket{name="default",status="failed",le="15"} 0
speedtest_run_duration_seconds_bucket{name="default",status="failed",le="20"} 0
speedtest_run_duration_seconds_bucket{name="default",status="failed",le="25"} 0
speedtest_run_duration_seconds_bucket{name="default",status="failed",le="30"} 0
speedtest_run_duration_seconds_bucket{name="default",status="failed",le="35"} 0
speedtest_run_duration_seconds_bucket{name="default",status="failed",le="+Inf"} 0
speedtest_run_duration_seconds_sum{name="default",status="failed"} 0
speedtest_run_duration_seconds_count{name="default",status="failed"} 0
speedtest_run_duration_seconds_bucket{name="default",status="succeeded",le="0"} 0
speedtest_run_duration_seconds_bucket{name="default",status="succeeded",le="5"} 0
speedtest_run_duration_seconds_bucket{name="default",status="succeeded",le="10"} 0
speedtest_run_duration_seconds_bucket{name="default",status="succeeded",le="15"} 1
speedtest_run_duration_seconds_bucket{name="default",status="succeeded",le="20"} 1
speedtest_run_duration_seconds_bucket{name="default",status="succeeded",le="25"} 1
speedtest_run_duration_seconds_bucket{name="default",status="succeeded",le="30"} 1
speedtest_run_duration_seconds_bucket{name="default",status="succeeded",le="35"} 1
speedtest_run_duration_seconds_bucket{name="default",status="succeeded",le="+Inf"} 1
speedtest_run_duration_seconds_sum{name="default",status="succeeded"} 10.214874283
speedtest_run_duration_seconds_count{name="default",status="succeeded"} 1
# HELP speedtest_run_total The number of total speed test runs
# TYPE speedtest_run_total counter
speedtest_run_total{name="default",status="failed"} 0
speedtest_run_total{name="default",status="succeeded"} 1
# HELP speedtest_server_info Server info labels of last successful speed test
# TYPE speedtest_server_info gauge
speedtest_server_info{country="United States",host="speedtest-dnvr.mynextlight.net",id="47683",ip="66.186.203.166",location="Denver, CO",name="default",port="8080",server_name="NextLight"} 1
# HELP speedtest_upload_bandwidth The upload bandwidth, in bytes/s, of last successful speed test
# TYPE speedtest_upload_bandwidth gauge
speedtest_upload_bandwidth{name="default"} 1.15625551e+08
# HELP speedtest_upload_bytes The uploaded bytes of last successful speed test
# TYPE speedtest_upload_bytes counter
speedtest_upload_bytes{name="default"} 6.57519155e+08
# HELP speedtest_upload_elapsed_seconds The elapsed upload time of last successful speed test
# TYPE speedtest_upload_elapsed_seconds counter
speedtest_upload_elapsed_seconds{name="default"} 5
# HELP speedtest_upload_latency_high_seconds The max upload latency of last successful speed test
# TYPE speedtest_upload_latency_high_seconds gauge
speedtest_upload_latency_high_seconds{name="default"} 0.09055399894714355
# HELP speedtest_upload_latency_iqm_seconds The upload interquartile mean latency of last successful speed test
# TYPE speedtest_upload_latency_iqm_seconds gauge
speedtest_upload_latency_iqm_seconds{name="default"} 0.0544780008494854
# HELP speedtest_upload_latency_jitter_seconds The upload latency jitter of last successful speed test
# TYPE speedtest_upload_latency_jitter_seconds gauge
speedtest_upload_latency_jitter_seconds{name="default"} 0.007607999723404646
# HELP speedtest_upload_latency_low_seconds The min upload latency of last successful speed test
# TYPE speedtest_upload_latency_low_seconds gauge
speedtest_upload_latency_low_seconds{name="default"} 0.002257999964058399
# HELP speedtest_vpn_status Indicates VPN status of last successful speed test
# TYPE speedtest_vpn_status gauge
speedtest_vpn_status{name="default"} 0
```
