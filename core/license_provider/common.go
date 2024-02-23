package license_provider

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
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

func (info LicenseInfo) toMap() map[string]interface{} {
	return map[string]interface{}{
		"license_expires_date": info.LicenseExpiresDate,
		"service_expires_date": info.ServiceExpiresDate,
		"mac_address":          info.MacAddress,
		"version":              "v1",
		"create_time":          info.CreateTime,
	}
}

func (info *LicenseInfo) GenSign() {
	data, _ := json.Marshal(info.toMap())
	h := md5.New()
	h.Write(data)
	info.Sign = hex.EncodeToString(h.Sum(nil))
}

func (info LicenseInfo) IsLegal() bool {
	data, _ := json.Marshal(info.toMap())
	h := md5.New()
	h.Write(data)
	return info.Sign == hex.EncodeToString(h.Sum(nil))
}
