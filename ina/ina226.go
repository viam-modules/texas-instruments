//go:build linux

// Package ina implements ina power sensors to measure voltage, current, and power
// INA226 datasheet: https://www.ti.com/lit/ds/symlink/ina226.pdf

// The voltage, current and power can be read as
// 16 bit big endian integers from their given registers.
// This value is multiplied by the register LSB to get the reading in nanounits.

// Voltage LSB: 1.25 mV for INA226
// Current LSB: maximum expected current of the system / (1 << 15)
// Power LSB: 25*CurrentLSB for INA226

// The calibration register is programmed to measure current and power properly.
// The calibration register is set to: calibratescale / (currentLSB * senseResistor)

package ina

import (
	"context"

	"go.viam.com/rdk/components/powersensor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

// Model226 for viam supported texas-instruments ina226 power sensor.
var Model226 = resource.NewModel("viam", "texas-instruments", "ina226")

func init() {
	resource.RegisterComponent(
		powersensor.API,
		Model226,
		resource.Registration[powersensor.PowerSensor, *Config]{
			Constructor: func(
				ctx context.Context,
				deps resource.Dependencies,
				conf resource.Config,
				logger logging.Logger,
			) (powersensor.PowerSensor, error) {
				newConf, err := resource.NativeConfig[*Config](conf)
				if err != nil {
					return nil, err
				}
				return newINA(conf.ResourceName(), newConf, logger, modelName226)
			},
		})
}
