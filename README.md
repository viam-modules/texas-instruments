# [`texas-instruments` module](https://github.com/viam-modules/texas-instruments)

This [texas-instruments module](https://app.viam.com/module/viam/texas-instruments) implements a [Texas Instruments TDA4VM board](https://devices.amazonaws.com/detail/a3G8a00000E2QErEAN/TI-TDA4VM-Starter-Kit-for-Edge-AI-vision-systems) using the [`rdk:component:board` API](https://docs.viam.com/appendix/apis/components/board/), and a [Texas Instruments INA219](https://www.ti.com/product/INA219) power sensor and a [Texas Instruments INA226](https://www.ti.com/product/INA226) power sensor using the [`rdk:component:power_sensor` API](https://docs.viam.com/appendix/apis/components/power-sensor/).

See [Configure your ti board](#Configure-your-ti-board) or [Configure your ina power sensor](#Configure-your-ina-power-sensor) for more information on configuring these components with Viam.

## Setup your ti board

Follow the [setup guide](https://docs.viam.com/installation/prepare/sk-tda4vm/) to prepare your TDA4VM for running `viam-server` before configuring a `ti` board.

> [!NOTE]
> Before configuring your board, you must [create a machine](https://docs.viam.com/cloud/machines/#add-a-new-machine).

Navigate to the [**CONFIGURE** tab](https://docs.viam.com/configure/) of your [machine](https://docs.viam.com/fleet/machines/) in the [Viam app](https://app.viam.com/).
[Add board / texas-instruments:ti to your machine](https://docs.viam.com/configure/#components).

## Configure your ti board

### Example configuration
```json
  {}
```

### Example configuration with optional digital interrupts
```json
{
  "digital_interrupts": [
    {
      "name": "your-interrupt-1",
      "pin": "15"
    },
    {
      "name": "your-interrupt-2",
      "pin": "16"
    }
  ]
}
```

### Attributes

The following attributes are available for `viam:texas-instruments:ti` boards:

| Attribute | Type | Required? | Description |
| --------- | ---- | --------- | ----------  |
| `digital_interrupts` | object | Optional | Any digital interrupts's pin number and name. |

For instructions on implementing digital interrupts, see [Digital interrupt configuration](#Digital-interrupt-configuration).


### Digital interrupt configuration
[Interrupts](https://en.wikipedia.org/wiki/Interrupt) are a method of signaling precise state changes.
Configuring digital interrupts to monitor GPIO pins on your board is useful when your application needs to know precisely when there is a change in GPIO value between high and low.

- When an interrupt configured on your board processes a change in the state of the GPIO pin it is configured to monitor, it ticks to record the state change.
  You can stream these ticks with the board API's [`StreamTicks()`](https://docs.viam.com/appendix/apis/components/board/#streamticks), or get the current value of the digital interrupt with [`Value()`](https://docs.viam.com/appendix/apis/components/board/#value).
- Calling [`GetGPIO()`](https://docs.viam.com/appendix/apis/components/board/#getgpio) on a GPIO pin, which you can do without configuring interrupts, is useful when you want to know a pin's value at specific points in your program, but is less precise and convenient than using an interrupt.

Integrate `digital_interrupts` into your machine in the `attributes` of your board by adding the following to your board's `attributes` configuration:

```json {class="line-numbers linkable-line-numbers"}
{
  "digital_interrupts": [
    {
      "name": "<your-digital-interrupt-name>",
      "pin": "<your-digital-interrupt-pin-number>"
    }
  ]
}
```

#### Attributes

The following attributes are available for `digital_interrupts`:

| Name | Type | Required? | Description |
| ---- | ---- | --------- | ----------- |
|`name` | string | **Required** | Your name for the digital interrupt. |
|`pin`| string | **Required** | The pin number of the board's GPIO pin that you wish to configure the digital interrupt for. |


## Set up your ina power sensor
> [!NOTE]
> Before configuring your power sensor, you must [create a machine](https://docs.viam.com/cloud/machines/#add-a-new-machine).

Navigate to the [**CONFIGURE** tab](https://docs.viam.com/configure/) of your [machine](https://docs.viam.com/fleet/machines/) in the [Viam app](https://app.viam.com/).
To use the ina219 model, [add power_sensor / texas-instruments:ina219 to your machine](https://docs.viam.com/configure/#components). To use the ina226 model, [add power_sensor / texas-instruments:ina226 to your machine](https://docs.viam.com/configure/#components)

## Configure your ina power sensor

On the new component panel, copy and paste the following attribute template into your power sensor's attributes field:

```json
{
  "i2c_bus": <string>,
  "i2c-addr": <int>,
  "max_current_amps": <float>,
  "shunt_resistance": <float>
}
```

### Attributes

The following attributes are available for `viam:texas-instruments:ina219` and `viam:texas-instruments:ina226` power sensors:

| Attribute | Type | Required? | Description |
| ---- | ---- | --------- | ----------- |
| `i2c_bus` | string | **Required** | The index of the i2c bus that the sensor is connected to. |
| `i2c_addr` | integer | Optional | The sensor's unique [i2c address](https://learn.adafruit.com/i2c-addresses/overview). Default: `0x40` |
| `max_current_amps` | float | Optional | Default: 3.2A. The maximum current that the sensor can measure in amperes (A). |
| `shunt_resistance` | float | Optional | Default: 0.1Ω. The shunt resistance value of the sensor in Ohms (Ω). |

Refer to your power sensor data sheet for specifics.

### Example configuration

For ina219 model, refer to the below example configuration. For ina226 model, you can still refer to the below example configuration, but change the `"model"` field to `"viam:texas-instruments:ina226"`

### `viam:texas-instruments:ina219`
```json
  {
      "name": "<your-texas-instruments-ina-power-sensor-name>",
      "model": "viam:texas-instruments:ina219",
      "type": "power_sensor",
      "namespace": "rdk",
      "attributes": {
        "i2c_bus": 1
      }
      "depends_on": []
  }
```

## Next Steps
- To test your board or power sensor, expand the **TEST** section of its configuration pane or go to the [**CONTROL** tab](https://docs.viam.com/fleet/control/).
- To write code against your board or power sensor, use one of the [available SDKs](https://docs.viam.com/sdks/).
- To view examples using a board or power sensor component, explore [these tutorials](https://docs.viam.com/tutorials/).
