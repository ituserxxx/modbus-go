// Copyright 2014 Quoc-Viet Nguyen. All rights reserved.
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package main

import (
	"collect/service"
	"collect/tools/modbus"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"log"
	"os"
)

const (
	rtuDevice = "\\\\.\\COM10"
)


func main1() {
	//var a interface{}
	//a = "ff"
	//if v, ok := a.(uint16); ok {
	//	println("yes")
	//}else{
	//	println("no")
	//	fmt.Printf("\n %T", v)
	//	println(uint16(0xff) == uint16(gconv.Uint16("ff")))
	//
	//}
	//println(uint16(0xff) == gconv.Uint16("0xff"))
	//println(uint16(0xa1) == gconv.Uint16("a1"))
	//println(uint16(0x00) == gconv.Uint16("0x0000"))
	//println(uint16(0x0009) == gconv.Uint16("0x09"))
	//println(uint16(0x0009) == gconv.Uint16("0x0009"))
	//println(uint16(0xff) == gconv.Uint16("0xff"))

}
func main() {
	handler := modbus.NewRTUClientHandler("/dev/sutpc_rs485s1")
	handler.BaudRate = 19200
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	//handler.SlaveId = byte(tools.GetUint16("ff"))
	handler.SlaveId = byte(service.GetUint16("6"))
	handler.Logger = log.New(os.Stdout, "rtu: ", log.LstdFlags)
	err := handler.Connect()
	if err != nil {
		print("error 1")
	}
	defer handler.Close()

	client := modbus.NewClient(handler)
	address := uint16(gconv.Byte(0))
	quantity := uint16(gconv.Byte(1))
	results, err := client.ReadInputRegisters(address, quantity)
	if err != nil {
		print(err.Error())
	}else {
		fmt.Printf("%#v",results)
		a1 := service.Hex2Dec(service.Dec2HexStr(results))
		fmt.Printf("%#v",a1)
	}
}
