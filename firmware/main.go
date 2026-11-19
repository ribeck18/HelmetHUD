package main

import (
	"fmt"
	"image/color"
	"machine"

	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont"
)

func main() {
	fmt.Println("Hello HUD!")

	if err := machine.I2C0.Configure(machine.I2CConfig{
		SDA:       qwiicSDA,
		SCL:       qwiicSCL,
		Frequency: 100_000,
	}); err != nil {
		panic(err)
	}

	oled := ssd1306.NewI2C(machine.I2C0)
	oled.Configure(ssd1306.Config{
		Width:   128,
		Height:  64,
		Address: 0x3c,
	})

	//clear the in memory buffer for what will be sent to display
	oled.ClearBuffer()

	//write the display in the buffer
	tinyfont.WriteLine(oled, &tinyfont.Picopixel, 2, 12, "Hello HelmetHUD!", color.RGBA{R: 255, B: 255, G: 255, A: 255})

	//send to the screen
	if err := oled.Display(); err != nil {
		fmt.Println(err)
	}
}
