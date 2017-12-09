package item

import (
	"encoding/json"
	"strconv"

	zlog "github.com/rs/zerolog/log"
)

// ValidateConfigs :
func (ms MonitoringSetting) ValidateConfigs() {
	ms.ValidateBasicSettings()
	ms.ValidateMonitoringItem()
}

// ValidateBasicSettings : check the basic settings.
func (ms MonitoringSetting) ValidateBasicSettings() {
	if ms.ManagerConf.ManagerName == "" {
		zlog.Fatal().Msg("Hostname config is Empty. Please define ManagerName.")
	}
}

// ValidateMonitoringItem :
func (ms *MonitoringSetting) ValidateMonitoringItem() {
	errFlag := 0
	items := ms.MonitoringItems

	for i, item := range items {
		itemBytes, err := json.Marshal(item)
		if err != nil {
			zlog.Fatal().Err(err).Msg("JSON Marshal error.")
		}

		if item.Key == "" {
			zlog.Error().
				Str("Data", string(itemBytes)).
				Msg("Key is Empty in MonitoringItem [" + strconv.Itoa(i) + "].")
			unset(items, i)
			errFlag = 1
		}

		if item.Type == "System" {
			//err := SystemMonitorValidation(item)
		} else if item.Type == "Command" && item.Command == "" {
			zlog.Error().
				Str("Data", string(itemBytes)).
				Msg("key " + item.Key + " : Command is Empty.")
			unset(items, i)
			errFlag = 1
		}

		if item.Duration == 0 {
			zlog.Info().
				Str("Data", string(itemBytes)).
				Msg("key " + item.Key + " : Duration is Empty.\nSet to default 60.")

			items[i].Duration = 60
		}
		if item.DataType == "" {
			zlog.Error().
				Str("Data", string(itemBytes)).
				Msg("key " + item.Key + " : DataType is Empty.")
			unset(items, i)
			errFlag = 1
		}
	}
	ms.MonitoringItems = items

	if errFlag == 1 {
		zlog.Error().Msg("Detected config error(s) in Monitoring Items. Ignored these items.")
	}

	zlog.Info().Msg("Complete the validation.")
}

func unset(s []MonitoringItem, i int) []MonitoringItem {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}
