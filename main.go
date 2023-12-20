package main

import (
	"net/http"

	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const(
	AccessKey = "xxx"
	SecretKey = "xxx"
	CpuUsage = "CpuUsage"
	IOPSUsage = "IOPSUsage"
	MemoryUsage = "MemoryUsage"
	DiskUsage = "DiskUsage"
	ConnectionUsage = "ConnectionUsage"
)

var metricsList = []string{"CpuUsage", "IOPSUsage", "MemoryUsage", "DiskUsage", "ConnectionUsage"}
var instanceidName = map[string]string{}



func main()  {

	// initialize  intanceidName
	rdsName()
	// update intanceidName per 30s
	go func ()  {
		time.Sleep(30*time.Second)
		rdsName()
	}()
	reg := prometheus.NewRegistry()
	reg.Register(NewRdsMetrisCollector())
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	http.ListenAndServe(":9001", nil)

}

