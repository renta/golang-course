package main

import (
	"fmt"
	"github.com/renta/golang-course/hw1-ntp/internal/ntp"
	"os"
	"time"
)

func main() {
	fmt.Printf("begin to ask for ntp time to a host %s \n", ntp.NtpHost)

	//ntpTime, err := internal.GetNtpTimeSimple()
	ntpTime, err := ntp.GetNtpTimeComplex()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error after query to ntp host %s, message: %s \n", ntp.NtpHost, err.Error())
		os.Exit(1)
	}

	fmt.Printf("current time is %s \n", ntpTime)
	fmt.Println("and formatted")
	fmt.Printf("current formatted (with RFC3339) time: %s \n", ntpTime.Format(time.RFC3339))
}