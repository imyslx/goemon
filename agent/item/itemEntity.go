package item

import "time"

// MonitoringSetting : setting construct.
type MonitoringSetting struct {
	ManagerConf     ManagerConfig
	MonitoringItems []MonitoringItem
}

// ManagerConfig : Some monitoring settings. Its not a monitor items.
type ManagerConfig struct {
	ManagerName string
	ManagerPort int
	AgentName   string
}

// MonitoringItem : for monitoring construct.
type MonitoringItem struct {
	Key       string
	Command   string
	Type      string
	Duration  int
	DataType  string
	Tick      *time.Ticker
	exitChan  chan bool
	IsWorking bool
}
