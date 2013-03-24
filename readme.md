
mgl is an application to manage the **M**SI **G**T **L**eds
This application is under development and doesn't work yet.

To compile this project you need libudev.h on your system.
Install this on debian/ubuntu with `sudo aptitude install libudev-dev`

HID
Because using the hidapi library by signal11 directly in this project wasn't too convenient, I started go.hid, which wraps the signal11/hidapi project.

History:
This project started as an attempt to rewrite https://github.com/PaNaVTEC/MSI_GT_GE_Led_Enabler in Go.