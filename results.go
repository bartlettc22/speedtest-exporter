package main

import "time"

type SpeedTestResults struct {
	Type       string             `json:"type"`
	Timestamp  time.Time          `json:"timestamp"`
	Ping       SpeedTestPing      `json:"ping"`
	Download   SpeedTestSpeed     `json:"download"`
	Upload     SpeedTestSpeed     `json:"upload"`
	PacketLoss int                `json:"packetLoss"`
	ISP        string             `json:"isp"`
	Iface      SpeedTestInterface `json:"interface"`
	Server     SpeedTestServer    `json:"server"`
	Result     SpeedTestResult    `json:"result"`
}

type SpeedTestPing struct {
	Jitter  float32 `json:"jitter"`
	Latency float32 `json:"latency"`
	Low     float32 `json:"low"`
	High    float32 `json:"high"`
}

type SpeedTestSpeed struct {
	Bandwidth int              `json:"bandwidth"`
	Bytes     int              `json:"bytes"`
	Elapsed   int              `json:"elapsed"`
	Latency   SpeedTestLatency `json:"latency"`
}

type SpeedTestInterface struct {
	InternalIP string `json:"internalIp"`
	Name       string `json:"name"`
	MACAddr    string `json:"macAddr"`
	IsVPN      bool   `json:"isVpn"`
	ExternalIP string `json:"externalIp"`
}

type SpeedTestServer struct {
	ID       int    `json:"id"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Country  string `json:"country"`
	IP       string `json:"ip"`
}

type SpeedTestResult struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	Persisted bool   `json:"persisted"`
}

type SpeedTestLatency struct {
	IQM    float32 `json:"iqm"`
	Low    float32 `json:"low"`
	High   float32 `json:"high"`
	Jitter float32 `json:"jitter"`
}
