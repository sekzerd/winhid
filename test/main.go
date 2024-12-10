package main

import (
	"fmt"
	"time"

	"github.com/sekzerd/hidapi/hidapi"
)

func print_info(v *hidapi.DeviceInfo) {
	fmt.Printf("Path:%s\n", v.Path)
	fmt.Printf("VendorId:%d\n", v.VendorID)
	fmt.Printf("ProductId:%d\n", v.ProductID)
	fmt.Printf("SerialNumber:%s\n", v.SerialNumber)
	fmt.Printf("ReleaseNumber:%d\n", v.ReleaseNumber)
	fmt.Printf("Manufacturer:%s\n", v.Manufacturer)
	fmt.Printf("Product:%s\n", v.Product)
	fmt.Printf("UsagePage:%d\n", v.UsagePage)
	fmt.Printf("Usage:%d\n", v.Usage)
	fmt.Printf("InterfaceNumber:%d\n", v.InterfaceNumber)
	// fmt.Printf("OutputReportLength:%d\n", v.OutputReportLength)
	// fmt.Printf("InputReportLength:%d\n", v.InputReportLength)
	// fmt.Printf("FeatureReportLength:%d\n", v.FeatureReportLength)
	// fmt.Printf("DeviceVersionNumber:%d\n", v.DeviceVersionNumber)
}
func list_device() {
	r := hidapi.Init()
	if r != nil {
		panic(r)
	}
	defer hidapi.Exit()

	s := hidapi.Enumerate(0, 0)

	for _, v := range s {
		if v.VendorID != 0x093a || v.UsagePage != 65376 || v.Usage != 97 {
			continue
		}
		do_write(v)
		time.Sleep(1 * time.Second)
		do_read(v)
		break
	}
	hidapi.Exit()
}

func do_write(v *hidapi.DeviceInfo) {
	dev, err := hidapi.OpenPath(v.Path)
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, 65)
	buffer[0] = 0x09
	buffer[1] = 0x82
	buffer[2] = 0x01
	buffer[3] = 0x00
	buffer[4] = 0x01
	buffer[5] = 0x00
	buffer[6] = 0x06
	buffer[63] = 0x6c
	dev.SetNonBlocking(true)
	err = dev.Write(buffer)
	println(err)
}

func do_read(v *hidapi.DeviceInfo) {
	dev, err := hidapi.OpenPath(v.Path)
	if err != nil {
		panic(err)
	}
	print_info(v)
	// buffer := make([]byte, 65)
	// dev.SetNonBlocking(true)

	buffer, err := dev.ReadTimeout(0x09, 65, 100)
	if err != nil {
		panic(err)
	}

	println(buffer)
}

func main() {
	list_device()
}
