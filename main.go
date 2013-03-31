package main

import (
	"fmt"
	"github.com/GeertJohan/go.hid"
	"log"
)

func main() {
	fmt.Println("--start--")

	leds, err := hid.Open(0x1770, 0xff00, "")
	if err != nil {
		log.Fatalf("Could not open leds device: %s", err)
	}
	defer leds.Close()

	// 1 normal
	// 2 gaming
	// 10 vaag witte kleur

	//Gaming mode = only left area on 1 color with a intensity level
	sendActivateArea(leds, areaLeft, colorSky, level1)
	sendActivateArea(leds, areaMiddle, colorBlue, level1)
	sendActivateArea(leds, areaRight, colorPurple, level1)
	commit(leds, modeNormal)

	fmt.Println("--stop--")
}

// send info on a specific area
func sendActivateArea(leds *hid.Device, area uint8, color uint8, level uint8) {
	data := make([]byte, 9)
	data[0] = 0x01  // Fixed report value.
	data[1] = 0x02  // Fixed report value?
	data[2] = 0x42  // 42 = set color inputs / 41 = confirm
	data[3] = area  // 1 = left / 2 = middle / 3 = right
	data[4] = color // see color constants
	data[5] = level // see level constants
	data[6] = 0x00  // empty 
	data[7] = 0x00  // empty 
	data[8] = 0xec  // EOR

	_, err := leds.SendFeatureReport(data)
	if err != nil {
		log.Fatalf("Could not send feature report to activate area. %s\n", err)
	}
}

// send confirmation to leds device
func commit(leds *hid.Device, mode uint8) {
	data := make([]byte, 9)
	data[0] = 0x01 // Fixed report value.
	data[1] = 0x02 // Fixed report value?
	data[2] = 0x41 // commit byte
	data[3] = mode // current mode
	data[4] = 0x00 // empty 
	data[5] = 0x00 // empty 
	data[6] = 0x00 // empty 
	data[7] = 0x00 // empty 
	data[8] = 0xec // EOR 

	_, err := leds.SendFeatureReport(data)
	if err != nil {
		log.Fatalf("Could not send feature report to commit. %s\n", err)
	}
}
