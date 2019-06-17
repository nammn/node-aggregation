package database

import (
	"time"
)

// Corresponds to a Node usage in a specific time slice
type NodeStat struct {
	Timestamp time.Time
	TimeSlice float64
	Cpu       float64
	Mem       float64
	NodeName  string
}

// Corresponds to a Node process usage in a specific time slice
type NodeProcessStat struct {
	Timestamp time.Time
	TimeSlice float64
	CpuUsed   float64
	MemUsed   float64
	NodeName  string
	Name      string
	URL       string
}
