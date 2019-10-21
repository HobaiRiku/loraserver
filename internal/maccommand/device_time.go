package maccommand

import (
	"github.com/brocaar/loraserver/internal/gps"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/brocaar/loraserver/internal/models"
	"github.com/brocaar/loraserver/internal/storage"
	"github.com/brocaar/lorawan"
)

func handleDeviceTimeReq(ds *storage.DeviceSession, rxPacket models.RXPacket) ([]storage.MACCommandBlock, error) {
	if len(rxPacket.RXInfoSet) == 0 {
		return nil, errors.New("rx info-set contains zero items")
	}

	var timeSinceGPSEpoch time.Duration
	var timeField time.Time

	for _, rxInfo := range rxPacket.RXInfoSet {
		if rxInfo.TimeSinceGPSEpoch != nil {
			timeSinceGPSEpoch = time.Duration(*rxInfo.TimeSinceGPSEpoch)
		} else if rxInfo.Time != nil {
			timeField = *rxInfo.Time
		}
	}

	log.WithFields(log.Fields{
		"dev_eui": ds.DevEUI,
	}).Info("device_time_req received")

	// fallback on time field when time since GPS epoch is not available
	if timeSinceGPSEpoch == 0 {
		// fallback on current server time when time field is not available
		if timeField.IsZero() {
			timeField = time.Now()
		}

		timeSinceGPSEpoch = gps.Time(timeField).TimeSinceGPSEpoch()
	}
	return []storage.MACCommandBlock{
		{
			CID: lorawan.DeviceTimeAns,
			MACCommands: storage.MACCommands{
				{
					CID: lorawan.DeviceTimeAns,
					Payload: &lorawan.DeviceTimeAnsPayload{
						TimeSinceGPSEpoch: timeSinceGPSEpoch,
					},
				},
			},
		},
	}, nil
}
