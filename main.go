package main

import (
	"fmt"

	"github.com/sekzerd/go-winhid/hidapi"
)

func main() {
	fmt.Println("start hid_init")
	r := hidapi.Hid_init()
	fmt.Println("end hid_init")

	if r != nil {
		panic(r)
	}
	s, err := hidapi.Hid_enumerate(0, 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("len:%d", len(s))

	for _, v := range s {
		fmt.Printf("path:%s\n", v.path)
		fmt.Printf("vendor_id:%d\n", v.vendor_id)
		fmt.Printf("product_id:%d\n", v.product_id)
		fmt.Printf("serial_number:%s\n", v.serial_number)
		fmt.Printf("release_number:%d\n", v.release_number)
		fmt.Printf("manufacturer_string:%s\n", v.manufacturer_string)
		fmt.Printf("product_string:%s\n", v.product_string)
		fmt.Printf("usage_page:%d\n", v.usage_page)
		fmt.Printf("usage:%d\n", v.usage)
		fmt.Printf("interface_number:%d\n", v.interface_number)
		fmt.Printf("output_report_length:%d\n", v.output_report_length)
		fmt.Printf("input_report_length:%d\n", v.input_report_length)
		fmt.Printf("feature_report_length:%d\n", v.feature_report_length)
		fmt.Printf("device_version_number:%d\n", v.device_version_number)
	}
}
