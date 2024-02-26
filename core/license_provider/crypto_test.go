package license_provider

import (
	"testing"
)

func TestGenKey(t *testing.T) {
	p := NewLicenseProvider("D:\\workspace\\golang\\src\\git.gtmap.cn\\license")
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
	var privateKey = "-----BEGIN PRIVATE KEY-----\nMIIEpAIBAAKCAQEAvKCuiq0bBktYnwOe+V+krqgVeRPF0R/D2WZGj/I2Uc2MR8vO\nKdi6rDARhXyD/sF2Ek4Ny5uTDPvkKI6MkY70LUWAdXCxjN4v7Auo3gGycBiUE5qX\nLJtDH1i5P77GTdTVNpmmsvT+2icWpc0wiXGLxLzdH0N1f/uKt72/ak9bbS3vvAxz\nPcgN4dIDa7fvagZ42VFRIWz2r+ok2ScEB4r2E4c0vPzjb/CwToRBzaB7OXEtq5BB\ntWLRq0AqCMx3Zlapq0fMheXCYqf7+Q4MGxdgBDknsmAOlkvP74NhyHkSJFsrgKh8\nEUbr2mx/4n/5X3yLlbju0qh/R9lF/UA7hLvlRwIDAQABAoIBACLaGYXb+CtQjp5K\n3/u5lUcEHXuSkLFUCi7H+++q1CiHLw0w9fW7arpX41TjrcvLWRKGw3vEUgZLIFvy\nArz+SdV3iwWn/dZGU1psyRXEAIE/uVRp8ta2FU/cZ9YEXMWStpZxC2DzbYdHQ6L8\n56Pti+Kc9spI/5PXGUlb0NtGsGDakzWhemBdNQsTiCd1TH7KZ8diX5TEPB/8Q5uY\nWMgD0T/NZsbCnFncc6pDcbvGQBnKuL8FZGkmxJ/bS/NEA9zzu9QIGFMhqPMaGI8B\n/sNPMXqmXp5ck9GesFJ1lMqUfd4BdLD9vzzbfH7F8cyjPu/rzkQCTQoYg8XZ0CrO\nlBLx3XECgYEA9y580k43NRADE0B3rurcLxl/yes0j44Xis0khteR6NQph+Pr3z24\n0G4Bn1gnUff+q836dNk/sF0CAv4RHFiOm03kAhMHMv4/BTIROBjKTukxKqgXixqb\nClxK3hQK1lxUzLlT75tC/XqIhKm2sR7E8E974FPfGr6QdActbbf9DO8CgYEAw1tr\npWD5/sMfxxIKfZfcPlZiMuW32Ev9x+fGRUg6/GKCmMgKHznVUuRFaNK07iTIj+ip\nB5mfd4GzztQafSJhsdAq5Fz3ocpJ+l5+34ZJf5flGYCFvCiobBSxLwfOT6+Zj/cJ\ncy8ypV5A3TXDGMl8V/F4lulj+w3522X1ZhyeXSkCgYA/GXTv0tuxBgdi8MAcvQO8\nWkwO5aYjR2inHDtI2Nr8jryTXhGmiYEWZB6x6LUQ7bfrb8eR8KubgAK1dNo1XD37\nU2TzSpw53kGEOAXBRkLO5iSQ+RwZfI87k9fg3Uju898J6/2LWx08y/zxsMlcabho\niEHdQEvh7ee+Vt639ZH5rQKBgQCFqkKQCDdk+cr/YTeuUT6PpHlSC91rcCs3IT2G\n3/IEB78FnDFYxjBQpAFdxHWpT5+u32jv20vB8AvAI13waprfl+gREg8ZMxjE5uz2\nMaKOW7aUaVyXrlX2hX3qCPQ/j42JdGSR3CRSA5hsAsH3brzWS1DQKOaQuzJOzIdI\nk6JDaQKBgQD1G804ylRmmZ2duPoF95aCZr0re6xl5yEmyN6tKrBDDig8Fg+0u599\niC1s2Fn1yVpx4hAAULQN12kH+D2ujY1EltgZoGnmHIOsMtXkduLMTvku5crp2/PQ\ndTo9k4Po9uzXrsjcQvyaTvTaJFoU4lxr/8COHRBwap5SoCWNw7yZVg==\n-----END PRIVATE KEY-----\n"
	err3 := p.CheckByPrivateKey(privateKey)
	//t.Log(info)
	t.Log(err3)
}
