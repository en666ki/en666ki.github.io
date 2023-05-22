package geoloc

import (
	"log"
	"net"
	"strings"

	"github.com/oschwald/maxminddb-golang"
)

// This example shows how to decode to a struct.
func Lookup_struct(ipString string) string {
	log.Println(ipString)
	db, err := maxminddb.Open("geoloc/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ip := net.ParseIP(strings.Split(ipString, ":")[0])
	log.Println(ip)

	var record struct {
		Country struct {
			ISOCode string `maxminddb:"iso_code"`
		} `maxminddb:"country"`
	} // Or any appropriate struct

	err = db.Lookup(ip, &record)
	if err != nil {
		log.Panic(err)
	}
	return record.Country.ISOCode
}
