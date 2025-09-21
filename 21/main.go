package main

import (
	"fmt"
)

type LengthProvider interface {
	GetLength() int
}

type MetricLength struct {
	meters int
}

func (m *MetricLength) GetLengthInMeters() int {
	return m.meters
}

type MetricToImperialAdapter struct {
	metric *MetricLength
}

func (a *MetricToImperialAdapter) GetLength() int {
	return a.metric.GetLengthInMeters() * 100
}

func main() {
	metric := &MetricLength{meters: 10}

	adapter := &MetricToImperialAdapter{metric: metric}

	var provider LengthProvider = adapter
	fmt.Println("Длина:", provider.GetLength())
}
