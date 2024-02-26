# Changelog

# v0.2.0
* Added global Prometheus label to all metrics `name` (default value `default`).  This is intended to be a label that can be customized when running multiple instances of this service.  For example, from different vpn/non-vpn networks.
* Added config option `SPEEDTEST_NAME_LABEL` to change the value of the Prometheus label `name`.
* Changed default interval from `300s` to `3600s`
* Changed previous `speedtest_server_info` label `name` to `server_name`
* Added `status` to `speedtest_last_run_timestamp`, `speedtest_run_duration_seconds`, and `speedtest_runs` to indicate success or failure
* Removed metrics `speedtest_last_error_timestamp`, `speedtest_errors`
* Updated metric name `speedtest_upload_elapsed` and `speedtest_download_elapsed` to `speedtest_upload_elapsed_seconds` and `speedtest_download_elapsed_seconds` respectively.  Also changed the value to `seconds` as it was previously `milliseconds`.
* Renamed `speedtest_runs` to `speedtest_run_total`
* Initialized the following metrics to zero properly: `speedtest_run_total`, `speedtest_run_duration_seconds`, `speedtest_last_run_timestamp`.
* Updated various metric HELP text

## v0.1.0
* Initial release