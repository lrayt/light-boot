package license_provider

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

const (
	LICENSE_V1 = "v1"
	LICENSE_V2 = "v2"
)

type LicenseInfo struct {
	LicenseExpiresDate int64    `json:"license_expires_date"`
	ServiceExpiresDate int64    `json:"service_expires_date"`
	MacAddress         []string `json:"mac_address"`
	Version            string   `json:"version"`
	CreateTime         int64    `json:"create_time"`
	Sign               string   `json:"sign"`
	LicenseActivated   bool     `json:"license_activated"`
}

func (info LicenseInfo) GenSign() string {
	var licenseMap = map[string]interface{}{
		"license_expires_date": info.LicenseExpiresDate,
		"service_expires_date": info.ServiceExpiresDate,
		"mac_address":          info.MacAddress,
		"version":              LICENSE_V1,
		"create_time":          info.CreateTime,
	}

	data, _ := json.Marshal(licenseMap)
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
