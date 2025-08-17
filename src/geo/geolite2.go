package geo

import (
	"fmt"
	"log"
	"net"

	"github.com/oschwald/geoip2-golang"
)

// GeoLite2DB wraps the MaxMind GeoLite2 database reader.
// This struct lets us easily manage the database connection.
type GeoLite2DB struct {
	DB *geoip2.Reader
}

// NewGeoLite2DB opens the GeoLite2 database file and returns a GeoLite2DB instance.
// dbPath: path to the .mmdb file (e.g., "maxmind/GeoLite2-City.mmdb").
func NewGeoLite2DB(dbPath string) (*GeoLite2DB, error) {
	db, err := geoip2.Open(dbPath)
	if err != nil {
		return nil, err // Return error if the database can't be opened.
	}
	fmt.Println("GeoLite2 database opened successfully at:", dbPath)
	return &GeoLite2DB{DB: db}, nil
}

// GetLocation looks up geolocation info for a given IP address.
// Returns a City record (with country, city, coordinates, etc.) or an error.
func (g *GeoLite2DB) GetLocation(ip string) (*geoip2.City, error) {
	fmt.Println("Looking up location for IP:", ip)
	ipAddr := net.ParseIP(ip)
	if ipAddr == nil {
		return nil, fmt.Errorf("invalid IP address: %s", ip)
	}

	record, err := g.DB.City(ipAddr)
	if err != nil {
		return nil, err // Return error if lookup fails.
	}
	fmt.Printf("Lookup result for %s: %+v\n", ip, record)
	return record, nil
}

// Close safely closes the GeoLite2 database connection.
// Always call this when you're done using the database!
func (g *GeoLite2DB) Close() {
	fmt.Println("Closing GeoLite2 database.")
	if err := g.DB.Close(); err != nil {
		log.Printf("Error closing GeoLite2 database: %v", err)
	}
}
