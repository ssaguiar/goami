// Copyright 2014 Helton Marques
//
//	Use of this source code is governed by a LGPL
//	license that can be found in the LICENSE file.
//

package main

import (
	"fmt"
	"github.com/heltonmarx/goami/ami"
)

func main() {
	socket, err := ami.NewSocket("127.0.0.1:5038")
	if err != nil {
		fmt.Printf("socket error: %v\n", err)
		return
	}
	if _, err := ami.Connect(socket); err != nil {
		return
	}
	var ret bool

	//Login
	uuid, _ := ami.GetUUID()
	ret, err = ami.Login(socket, "admin", "admin", "Off", uuid)
	if err != nil || ret == false {
		fmt.Printf("login error (%v)\n", err)
		return
	}
	fmt.Printf("login ok!\n")

	data := ami.KhompSMSData{
		Device:       "b0",
		Destination:  "4899893791",
		Confirmation: true,
		Message:      "hey ho, let's go",
	}

	//SendSMS
	//	func KSendSMS(socket *Socket, actionID string, data KhompSMSData) (map[string]string, error) {
	//
	s, err := ami.KSendSMS(socket, uuid, data)
	if err != nil {
		fmt.Printf("sms sms error\n", err)
	}
	fmt.Printf("response: [%v]\n", s)

	//Logoff
	fmt.Printf("logoff\n")
	ret, err = ami.Logoff(socket, uuid)
	if err != nil || ret == false {
		fmt.Printf("logoff error: (%v)\n", err)
		return
	}
	fmt.Printf("goodbye !\n")
}
