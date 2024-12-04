package main

import (
	"fmt"

	"github.com/sekzerd/winhid/hidapi"
)

func main() {
	r := hidapi.HidInit()
	if r != nil {
		panic(r)
	}

	s, err := hidapi.HidEnumerate(0, 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("len:%d", len(s))

	for _, v := range s {
		fmt.Printf("Path:%s\n", v.Path)
		fmt.Printf("VendorId:%d\n", v.VendorId)
		fmt.Printf("ProductId:%d\n", v.ProductId)
		fmt.Printf("SerialNumber:%s\n", v.SerialNumber)
		fmt.Printf("ReleaseNumber:%d\n", v.ReleaseNumber)
		fmt.Printf("ManufacturerString:%s\n", v.ManufacturerString)
		fmt.Printf("ProductString:%s\n", v.ProductString)
		fmt.Printf("UsagePage:%d\n", v.UsagePage)
		fmt.Printf("Usage:%d\n", v.Usage)
		fmt.Printf("InterfaceNumber:%d\n", v.InterfaceNumber)
		fmt.Printf("OutputReportLength:%d\n", v.OutputReportLength)
		fmt.Printf("InputReportLength:%d\n", v.InputReportLength)
		fmt.Printf("FeatureReportLength:%d\n", v.FeatureReportLength)
		fmt.Printf("DeviceVersionNumber:%d\n", v.DeviceVersionNumber)
	}
}
