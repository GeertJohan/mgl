package main

import (
	"fmt"
	"go.hid"
)

func main() {
	fmt.Println("--start--")

	leds, err := hid.Open(0x1770, 0xff00, nil)
	if err != nil {
		panic(err)
	}

	// //Gaming mode = only left area on 1 color with a intensity level
	// sendActivateArea(handle, AREA_LEFT, arguments[kColor1], arguments[kLevel])
	// commit(handle, MODE_GAMING)

	leds.Close()

	fmt.Println("--stop--")
}

func sendActivateArea(leds *hid.Device, area uint8, color uint8, level uint8) {

	// // Will send a 8 bytes array
	// unsigned char data[8];
	// memset(&data, 0x00, 8);
	// data[0] = 0x01; // Fixed report value.
	// data[1] = 0x02; // Fixed report value?

	// data[2] = 0x42; // 42 = set color inputs / 41 = confirm
	// data[3] = area; // 1 = left / 2 = middle / 3 = right
	// data[4] = color; // see color constants
	// data[5] = level; // see level constants
	// data[6] = 0x00; // empty 
	// data[7] = 0xec; // EOR

	// if (hid_send_feature_report(handle, data, 9) < 0) {
	// 	printf("Unable to send a feature report.\n");
	// }

}

func commit(leds *hid.Device, mode uint8) {
	// //CONFIRMATION. This needs to be sent for confirmate all the led operations
	// unsigned char data[8];
	// data[0] = 0x01;
	// data[1] = 0x02;

	// data[2] = 0x41; // commit byte
	// data[3] = mode; // current mode
	// data[4] = 0x00;  
	// data[5] = 0x00; 
	// data[6] = 0x00;
	// data[7] = 0xec;

	// if (hid_send_feature_report(handle, data, 9) < 0) {
	// 	printf("Unable to send a feature report.\n");
	// }
}
