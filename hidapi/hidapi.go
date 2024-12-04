package hidapi

/*
#cgo windows LDFLAGS: -lsetupapi -lhid
#cgo windows CFLAGS: -w
#include "hidapi.h"
#include "hidapi.c"
*/
import "C"
import (
	"fmt"

	"golang.org/x/sys/windows"
)

type HidDevice struct {
	Path               string
	VendorId           uint16
	ProductId          uint16
	SerialNumber       string
	ReleaseNumber      uint16
	ManufacturerString string
	ProductString      string
	UsagePage          uint16
	Usage              uint16
	InterfaceNumber    int32

	OutputReportLength  uint16
	InputReportLength   uint16
	FeatureReportLength uint16
	DeviceVersionNumber uint16
}

func C_wchar_p_to_GoString(s *C.wchar_t) (ret string) {
	return windows.UTF16PtrToString((*uint16)(s))
}

func HidEnumerate(vendor_id uint16, product_id uint16) (ret []HidDevice, err error) {
	s := C.hid_enumerate(C.ushort(vendor_id), C.ushort(product_id))
	for s.next != nil {
		dev := HidDevice{
			Path:               C.GoString(s.path),
			VendorId:           uint16(s.vendor_id),
			ProductId:          uint16(s.product_id),
			SerialNumber:       C_wchar_p_to_GoString(s.serial_number),
			ReleaseNumber:      uint16(s.release_number),
			ManufacturerString: C_wchar_p_to_GoString(s.manufacturer_string),
			ProductString:      C_wchar_p_to_GoString(s.product_string),
			InterfaceNumber:    int32(s.interface_number),
		}
		info := C.hid_open_path(s.path)
		if info == nil {
			dev.InputReportLength = 0
			dev.OutputReportLength = 0
			dev.FeatureReportLength = 0
			ret = append(ret, dev)
			s = s.next
			continue
		}
		dev.InputReportLength = uint16(info.input_report_length)
		dev.OutputReportLength = uint16(info.output_report_length)
		dev.FeatureReportLength = uint16(info.feature_report_length)

		ret = append(ret, dev)
		s = s.next
	}
	return ret, err
}

func HidInit() (err error) {
	r := C.hid_init()
	if r != 0 {
		fmt.Printf("%d", r)
		err = fmt.Errorf("hid_init failed error code: %d", r)
	}
	return err
}
func HidExit() (err error) {
	r := C.hid_exit()
	if r != 0 {
		err = fmt.Errorf("hid_exit failed error code: %d", r)
	}
	return err
}
