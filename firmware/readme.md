# HelmetHUD firmware

## Planned hardware

| Part | Model | Firmware connection |
| --- | --- | --- |
| Controller | [SparkFun Pro Micro ESP32-C3, DEV-23484](https://docs.sparkfun.com/SparkFun_Pro_Micro-ESP32C3/) | USB-C for bench power and flashing. Onboard Bluetooth LE for the iPhone link. |
| Display | [SparkFun 0.91″ Qwiic OLED, LCD-24606](https://www.sparkfun.com/sparkfun-qwiic-oled-display-0-91-in-128x32-lcd-24606.html) | SSD1306, 128 × 32 pixels, I²C address `0x3C`, powered by 3.3 V through Qwiic. |
| Display cable | [SparkFun 100 mm flexible Qwiic cable, PRT-17259](https://www.sparkfun.com/flexible-qwiic-cable-100mm.html) | Carries power, ground, SDA, and SCL between the boards. |
| Bench optics | [38 mm double-convex lens with a 50 mm focal length](https://www.homesciencetools.com/product/lens-double-convex-50mm-f-l/) and the existing helmet visor | The lens and visor are passive. |

The SparkFun board’s Qwiic port uses GPIO5 for SDA and GPIO6 for SCL. When you initialize the OLED, pass `qwiicSDA` and `qwiicSCL` from `board.go` to `machine.I2C0.Configure`. The generic ESP32-C3 target does not provide the SparkFun board’s pin mapping. [ESP32-C3 Pro Micro pin definitions](https://github.com/espressif/arduino-esp32/blob/master/variants/sparkfun_pro_micro_esp32c3/pins_arduino.h)

## Editor and build

This setup was checked with TinyGo 0.44.0 and Go 1.26.1. The `.tinygo-target` file lives beside `go.mod`, where the Neovim configuration looks for it. Neovim uses `tinygo info` for that target to tell `gopls` which standard library, architecture, and build tags to use. The TinyGo CLI takes the target explicitly.

From the repo root, build the starter firmware with:

```sh
cd firmware
mkdir -p build
tinygo build -target="$(cat .tinygo-target)" -o build/helmet-hud.elf .
```

On Apple Silicon, if TinyGo or `gopls` reports `bad CPU type in executable`, run `type -a go` and put the ARM64 Go installation first on your `PATH` before starting Neovim or building.

`main.go` is the starting point for your first exercise. Plug the OLED into the Qwiic port, configure I²C with the board pin constants, and draw a fixed `55` before adding Bluetooth.

## Simulate in Wokwi

[diagram.json](diagram.json) wires Wokwi's ESP32-C3-DevKitM-1 to an SSD1306 at `0x3C` using the same GPIO5 (SDA) and GPIO6 (SCL) as the SparkFun board. USB serial is enabled so TinyGo's `println` output can appear in the serial monitor.

1. From the repo root, build a binary for the simulator:

   ```sh
   cd firmware
   mkdir -p build
   tinygo build -target="$(cat .tinygo-target)" -o build/helmet-hud.bin .
   ```

2. Open a [new Wokwi ESP32-C3 project](https://wokwi.com/projects/new/esp32-c3). Replace the contents of its **diagram.json** tab with the contents of [our diagram.json](diagram.json).
3. Focus the code editor and press **F1** (**Fn–F1** on some Macs). Choose **Upload Firmware and Start Simulation…** and select `firmware/build/helmet-hud.bin`. After a code change, rebuild the binary and upload it again. [Wokwi's custom firmware instructions](https://docs.wokwi.com/guides/esp32#custom-application-firmware)

For the first run, add a `println("Hello, HUD!")` inside `main` and check the serial monitor. An empty `main` gives you no serial output or display image. Once serial works, initialize the OLED and draw a fixed `55` within a **128 × 32** logical drawing area. Wokwi's display model is **128 × 64**, so it can help you test drawing over I²C. Final screen placement and Bluetooth still need the physical hardware. [Wokwi SSD1306 model](https://docs.wokwi.com/parts/board-ssd1306), [Wokwi ESP32 features](https://docs.wokwi.com/guides/esp32#simulation-features)
