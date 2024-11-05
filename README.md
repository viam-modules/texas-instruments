# [`texas-instruments` module](https://github.com/viam-modules/texas-instruments)

This [texas-instruments module](https://app.viam.com/module/viam/texas-instruments) implements a [Texas Instruments TDA4VM board](https://devices.amazonaws.com/detail/a3G8a00000E2QErEAN/TI-TDA4VM-Starter-Kit-for-Edge-AI-vision-systems) using the [`rdk:component:board` API](https://docs.viam.com/appendix/apis/components/board/).

## Setup

Follow the [setup guide](https://docs.viam.com/installation/prepare/sk-tda4vm/) to prepare your TDA4VM for running `viam-server` before configuring a `ti` board.

## Configure your ti board

> [!NOTE]
> Before configuring your board, you must [create a machine](https://docs.viam.com/cloud/machines/#add-a-new-machine).

Navigate to the [**CONFIGURE** tab](https://docs.viam.com/configure/) of your [machine](https://docs.viam.com/fleet/machines/) in the [Viam app](https://app.viam.com/).
[Add board / texas-instruments:ti to your machine](https://docs.viam.com/configure/#components).

### Attributes

The following attributes are available for `viam:texas-instruments:ti` boards:

| Attribute | Type | Required? | Description |
| --------- | ---- | --------- | ----------  |
| `digital_interrupts` | object | Optional | Any digital interrupts's pin number and name. |

For instructions on implementing digital interrupts, see [Digital interrupt configuration](#Digital-interrupt-configuration).

## Example configuration

### `viam:texas-instruments:ti`
```json
  {
     "name": "<your-texas-instruments-ti-board-name>",
      "model": "viam:texas-instruments:ti",
      "type": "board",
      "namespace": "rdk",
      "attributes": {},
      "depends_on": []
  }
```

### Next Steps
- To test your board, expand the **TEST** section of its configuration pane or go to the [**CONTROL** tab](https://docs.viam.com/fleet/control/).
- To write code against your board, use one of the [available SDKs](https://docs.viam.com/sdks/).
- To view examples using a board component, explore [these tutorials](https://docs.viam.com/tutorials/).

## Digital interrupt configuration
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

### Attributes

The following attributes are available for `digital_interrupts`:

<!-- prettier-ignore -->
| Name | Type | Required? | Description |
| ---- | ---- | --------- | ----------- |
|`name` | string | **Required** | Your name for the digital interrupt. |
|`pin`| string | **Required** | The pin number of the board's GPIO pin that you wish to configure the digital interrupt for. |

### Example configuration

```json {class="line-numbers linkable-line-numbers"}
{
  "components": [
    {
      "name": "<your-texas-instruments-board-name>",
      "model": "viam:texas-instruments:ti",
      "type": "board",
      "namespace": "rdk",
      "attributes": {
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
    }
  ]
}
```
