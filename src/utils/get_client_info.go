package utils

import (
	"log"

	"github.com/ip2location/ip2location-go/v9"
	ua "github.com/mileusna/useragent"
)

func GetCountry(ip string) string {
	db, err := ip2location.OpenDB("./src/public/IP2LOCATION-LITE-DB1.IPV6.BIN")

	// fallback
	defer func() string {
		var r string
		if rv := recover(); rv != nil {
			r = ""
		}
		return r
	}()

	if err != nil {
		log.Panicln("[GetCountry] OpenDB: ", err)
	}

	results, err := db.Get_all(ip)

	if err != nil {
		log.Panicln("[GetCountry] Get_all: ", err)
	}

	country := results.Country_short

	// Invalid IP address
	if country == "Invalid IP address." {
		log.Panicln("[GetCountry] Invalid IP address")
	}

	return results.Country_short
}

func GetBrowser(userAgent string) string {
	ua := ua.Parse(userAgent)
	browserName := ua.Name
	return browserName
}

func GetOs(userAgent string) string {
	ua := ua.Parse(userAgent)
	osName := ua.OS
	return osName
}

func GetDevice(userAgent string) string {
	ua := ua.Parse(userAgent)
	deviceName := ""
	if ua.Mobile {
		deviceName = "mobile"
	}
	if ua.Tablet {
		deviceName = "tablet"
	}
	if ua.Desktop {
		deviceName = "desktop"
	}
	return deviceName
}
