package license_provider

import (
	"testing"
)

func TestGenKey(t *testing.T) {
	p := NewLicenseProvider("D:\\workspace\\golang\\src\\github.com\\light-boot\\core\\license_provider")
	//err := p.GenKey()
	//t.Log(err)
	//var info = &LicenseInfo{
	//	LicenseExpiresDate: 1000,
	//	ServiceExpiresDate: 2000,
	//	MacAddress:         []string{"12.12.12.12", "34.34.34.34"},
	//	Version:            "v1",
	//	CreateTime:         3000,
	//	Sign:               "1",
	//	LicenseActivated:   false,
	//}
	//err2 := p.GenLicense(info)
	//t.Log(err2)

	//info, err3 := p.LoadLicenseInfo()
	//t.Log(info)
	//t.Log(err3)
	var privateKey = "-----BEGIN PRIVATE KEY-----\nMIIEpAIBAAKCAQEAvUKPQ+yKnasQQhcXfnx+HTFfV2i96a67TKXXQtxdbuGiaEfl\nx2bKMXRXxOhAVNF5SDsF0dfQkf6Ez1WcLjZmBXgcNnTgPDVI/p8I3tb/gewJUp5a\nk4IF5BimLd+8DLRuneQOx+/Y8RnuzTlmbQvHKypyDCJfwk9e/27ciXQZ7xK1tPNT\nqBz1XUdofXRy2FSrO1UH7DRk6B+wyDFvrHZu9ZDkXJGrCL+KEo4ru3mOMa183ofH\nWMzdazGP64BS4PmfJqh6vI1nf2G+0SRQXPnQ6hhhUfiE4dhsB32xXQwvae+fSnKL\n9UU5KLBmvBTOy+58YMcb6MkQm2Wbsfpjb6H3yQIDAQABAoIBAHH8V2rSWP3y3Dzh\nyX2dnxsX9FMhu1e1rArSyx3yuLciX/0saEC7MMWuS3pFYxFyg/gzzCKBN5XmU7jH\n1+OhUbqzh1Jo+/BnK0pnICGOHZ7zOP4wb17t7XxeOB/i41BHj2O6yggy9VbohoHF\ns2GzuUXPL5cNUN0njpCLBXonmcL5WWg48u5Dyd7/qGKW0Gx9/BgOzEhHZZsCnPgJ\nL5LwrC+uQiLqcKA9e0d6ant6eNmQ4uQlJdU2bY2I67jEI33oCtyKHPo7sxo5YUKW\n2lkW59Yr3zrMJEMibzmG5qjdZ44BP3VZgVCbWgCY111K/APpGjjas5qJX58u7Cdy\n4w6YRvECgYEA5UaXYad4/QqDw9S8wA69502ypqPKgcHDWuzfwW/bu5DjnHamwC3L\nU4cLbeyIiHZIvu/vz2eY1bAOFp09XxZxvzLEGNQKPcUUvFnZZDeC9ZRDN0J63luV\n5bIkCOw8T2HrEFQp2ulBUH6aSvHJNoYZKsfLMZCV6/PZOjyDYTmeRo0CgYEA01Hu\nAQQD0og1xwGQnXKAdPuu5TWiBb4JHCys9RshLrYkTW+TEGdMhsfeipWbiAFJM0q5\nSuY/QIpHNXQdCOiVFnzg2dCj66PRaDcLOd9LBRjAZO1zeVLDyK03+40H/MEQkS0i\nZtlUw1bPXlOwC2O0u7zRhXuSxl+oEPfNf+ieFS0CgYEA0epewW8W6/6Lg2msYt9c\naYEO11lRGJox5XbIqo9ijAltC5zqsTt8VH5pLXyJyP8bPY3qb2d3W/Yz/+p4/S1M\neXTea6j/s2xUvto4mnPkuDnFjLuWSS1rQYFnUrAPy9Jn6GxWYfJ79VnkpEdZJU2D\nj1KbrEanrP923MX4t7SMQCUCgYEAjYBTCsMRfCxRNkzJ9WPyuGK/niPntkil+QPB\nXO8prQmyDVgk9dr111qGne3IE64owfvLT2gV+rWdMSHrP0febQx50+BtoyAy3MHe\n9oWaynLkgbCunhLZnsq+real7o2o9k1ut8fXhmXevZ70Ruwg3YV2FxfEFKmA9Lu6\nE2hEUpECgYBvHnfP+ONUj5I/ynI+yLNP0XMrW4fx1snqHrT/xmSNcRaqfkftcoM4\nXs2AidZFu+ludLuRgF5dbkZg3Rwbn5zpF1QywcxH19fs+Np0Dn7RoGsIFFohu0dU\nMYSjHvMsL82G8oJc/1yOHnFbJu9pjoKusu52oZsHzTCh/qR8qrACjw==\n-----END PRIVATE KEY-----\n"
	err3 := p.CheckByPrivateKey(privateKey)
	//t.Log(info)
	t.Log(err3)
}
