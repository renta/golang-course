package ntp_adapter

import (
	"github.com/beevik/ntp"
	"time"
)

const NtpHost string  = "2.ru.pool.ntp.org"

func GetNtpTimeSimple() (time.Time, error) {
	ntpTime, err := ntp.Time(NtpHost)
	if err != nil {
		return time.Time{}, err
	}

	return ntpTime, err
}

func GetNtpTimeComplex() (time.Time, error) {
	options := ntp.QueryOptions{Version: 4}
	response, err := ntp.QueryWithOptions(NtpHost, options)
	if err != nil {
		return time.Time{}, err
	}

	err = response.Validate()
	if err != nil {
		return time.Time{}, err
	}
	ntpTime := time.Now().Add(response.ClockOffset)

	return ntpTime, err
}