package item

import (
	"fmt"

	zlog "github.com/rs/zerolog/log"
	"github.com/shirou/gopsutil/host"
)

// WorkerProcess :
func (mi *MonitoringItem) WorkerProcess() {
	if mi.Type == "Agent" {
		mi.SystemWorkerProcess()
	} else {
		mi.CommandWorkerProcess()
	}
}

// SystemWorkerProcess :
func (mi *MonitoringItem) SystemWorkerProcess() {
	switch mi.Command {
	case "HostInfo":
		v, err := host.Info()
		if err != nil {
			zlog.Error().Err(err).Msg("Get an error in HostInfo")
		}
		fmt.Print(v)
	default:
		zlog.Info().Msg("No Monitoring Preset definied.")
	}
}

// CommandWorkerProcess :
func (mi *MonitoringItem) CommandWorkerProcess() {

	switch mi.Command {
	case "HostInfo":
		v, err := host.Info()
		if err != nil {
			zlog.Error().Err(err).Msg("Get an error in HostInfo")
		}
		fmt.Print(v)
	default:
		zlog.Info().Msg("No Monitoring Preset definied.")
	}

}
