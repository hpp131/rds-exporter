package main

import "github.com/prometheus/client_golang/prometheus"

//collect prometheus metrics

type  RdsMetrisCollector struct{
	CpuUsage  *prometheus.Desc
	IOPSUsage *prometheus.Desc
	MemoryUsage *prometheus.Desc
	DiskUsage *prometheus.Desc
	ConnectionUsage *prometheus.Desc
}

// 多个指标共用相同的variableLabel,如何优化？
func NewRdsMetrisCollector() prometheus.Collector{
	c := &RdsMetrisCollector{
		CpuUsage: prometheus.NewDesc(
			"CpuUsage",
			"CpuUsage",
			[]string{"InstanceId", "BusinessUnitTag", "InstanceName"},
			nil,
		),
		IOPSUsage: prometheus.NewDesc(
			"IOPSUsage",
			"IOPSUsage",
			[]string{"InstanceId", "BusinessUnitTag", "InstanceName"},
			nil,	
		),
		MemoryUsage: prometheus.NewDesc(
			"MemoryUsage",
			"MemoryUsage",
			[]string{"InstanceId", "BusinessUnitTag", "InstanceName"},
			nil,
		),
		DiskUsage: prometheus.NewDesc(
			"DiskUsage",
			"DiskUsage",
			[]string{"InstanceId", "BusinessUnitTag", "InstanceName"},
			nil,
		),
		ConnectionUsage: prometheus.NewDesc(
			"ConnectionUsage",
			"ConnectionUsage",
			[]string{"InstanceId", "BusinessUnitTag", "InstanceName"},
			nil,
		),
	}
	
	return c
}


func(r *RdsMetrisCollector)  Describe(ch chan<- *prometheus.Desc) {
	ch <- r.CpuUsage
	ch <- r.IOPSUsage
	ch <- r.DiskUsage
	ch <- r.ConnectionUsage
	ch <- r.MemoryUsage
}

func (c *RdsMetrisCollector) Collect(ch chan<- prometheus.Metric) {
	c.collectFunc(ch)
}

// collect metrics data
func (c *RdsMetrisCollector) collectFunc(ch chan <- prometheus.Metric){
	// var resultList [][]*RdsResponse
	// wg.Add(5)
	for _, v := range  metricsList{
		m := getMetrics(v)
		for _,r := range m{
			switch r.MetricName{
				case "CpuUsage":
					ch <- prometheus.MustNewConstMetric(c.CpuUsage, prometheus.GaugeValue, float64(r.Maximum), r.InstanceId, r.Tag["BusinessUnit"], r.InstanceName)
				case "IOPSUsage":
					ch <- prometheus.MustNewConstMetric(c.IOPSUsage, prometheus.GaugeValue, float64(r.Maximum), r.InstanceId, r.Tag["BusinessUnit"], r.InstanceName)
				case "DiskUsage":
					ch <- prometheus.MustNewConstMetric(c.DiskUsage, prometheus.GaugeValue, float64(r.Maximum), r.InstanceId, r.Tag["BusinessUnit"], r.InstanceName)
				case "ConnectionUsage":
					ch <- prometheus.MustNewConstMetric(c.ConnectionUsage, prometheus.GaugeValue, float64(r.Maximum), r.InstanceId, r.Tag["BusinessUnit"], r.InstanceName)
				case "MemoryUsage":
					ch <- prometheus.MustNewConstMetric(c.MemoryUsage, prometheus.GaugeValue, float64(r.Maximum), r.InstanceId, r.Tag["BusinessUnit"], r.InstanceName)
			}
		} 	
	} 
}