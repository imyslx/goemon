package item

import (
	"encoding/json"
	"io/ioutil"
	"time"

	zlog "github.com/rs/zerolog/log"
)

///// MonitoringSetting

// ReadLocalFile :
func ReadLocalFile(file string) (*MonitoringSetting, error) {

	data := new(*MonitoringSetting)

	// Read config file.
	//     ([]bytes, err)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return *data, err
	}

	if err := json.Unmarshal(jsonBytes, data); err != nil {
		return *data, err
	}
	return *data, nil
}

// GetConfig : re-create monitoring instance.
func GetConfig() (*MonitoringSetting, error) {

	// Get Config file.
	ms, err := ReadLocalFile("./conf/client.json")
	if err != nil {
		zlog.Fatal().Err(err).Msg("Could not get local config. ")
	}
	zlog.Info().Msg("Successful to get config. Continue.")

	ms.ValidateConfigs()
	ms.CreateTickers()

	return ms, err
}

// CreateTickers : Create Ticker for all of MonitoringItems.
func (ms *MonitoringSetting) CreateTickers() {
	zlog.Debug().Msg("Start to create Tickers.")

	items := ms.MonitoringItems
	for i, item := range items {
		items[i].Tick = time.NewTicker(time.Second * time.Duration(item.Duration))
		items[i].exitChan = make(chan bool)
		items[i].IsWorking = false
	}
	ms.MonitoringItems = items
	zlog.Info().Msg("Completed to create Tickers.")

}

// RunAllWokers :
func (ms *MonitoringSetting) RunAllWokers() {
	zlog.Info().Msg("Running All Monitors... ")

	for _, v := range ms.MonitoringItems {
		v.Work()
	}
	zlog.Info().Msg("[ OK ]")
}

// ExitAllItems : Exit all of MonitoringItems.
func (ms *MonitoringSetting) ExitAllItems() {
	zlog.Info().Msg("Exiting All Monitoring Items...")

	for _, v := range ms.MonitoringItems {
		v.exitChan <- true
	}
}

///// MonitoringItems

// Work : work the monitoring items.
func (mi MonitoringItem) Work() {

	zlog.Debug().Msg("Starting to work: " + mi.Key)

	go func() {
	forLoop:
		for {
			select {
			case <-mi.Tick.C:
				zlog.Info().Msg(mi.Key + ": I'm working !")
				mi.WorkerProcess()
			case <-mi.exitChan:
				break forLoop
			}
		}
		zlog.Info().Msg("Exited: " + mi.Key)
	}()

}
