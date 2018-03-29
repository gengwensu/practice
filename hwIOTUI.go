/*
hwIOTUI: app to rerieve information based on a time frame and device from hwIOTCollector.
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func keepResults(s string, n int) string {
	result := strings.Join(strings.Split(s, "\n")[:n], "\n")
	return strings.Replace(result, "\r", "", -1)
}

func main() {

	fmt.Scanf()
	fmt.Printf("Command is %s\n", *Istr)
	switch iCommand {
	case "deviceInfo":
		//Query all devices
		resp, err := http.Get("https://www.honeywell.com/device")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println("get:\n", keepResults(string(body), 3))
	case "deviceData":
		//Query a particular device with deviceId with a time frame
		resp, err := http.Get("https://www.honeywell.com/device/id/?start=starttime,end=endtime")
		//starttime and endtime are timestamps
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println("get:\n", keepResults(string(body), 3))
	default:
		fmt.Println("Next command please.")
	}
}
