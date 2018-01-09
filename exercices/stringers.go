package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ipa IPAddr) String() (string){
	// res = string(ipa[0]) + "." + string(ipa[1]) + "." + string(ipa[2]) + "." + string(ipa[3])
	return fmt.Sprintf("%v.%v.%v.%v", ipa[0], ipa[1], ipa[2], ipa[3])
	// fmt.Println(res)
	// return
} 


//func ipstringer() {
func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}


