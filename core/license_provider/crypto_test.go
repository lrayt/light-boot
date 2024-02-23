package license_provider

import (
	"testing"
)

func TestGenKey(t *testing.T) {
	p := NewLicenseProvider("D:\\tmp")
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
	var privateKey = "-----BEGIN PRIVATE KEY-----\nMIIEpQIBAAKCAQEAwuEzGT1yDzijb+qg6fPnhAF1BAzxyZ68O/F9n2YF6XgYt0bT\nUFqbtOKbNVgI6F1QGpg4nYWJdapQDsavILKeNaOyc66xTBtUhdEx8liHfXN4x2pP\nh6sPqM0sPZn4i10ajfli8O4xT1wVDhJuO5Xzxzz282wD+uc0kxygKed1ZdDDyapQ\n6EbKNZAF896CXUGlq1Rrxu62xrkXshK9+xILow3Xvo5B7Vln6wr0fvJ6Awqz9gxo\n60Sj3h44a3zo1GhNM7j8Ae4JI8dCGONTKMOx9DH+kjHRAOtVjWyYraAhvZNmVG/L\nUXLe6wYF3aKNBJ3ICg9m4I61Mo2k6PZtF8qKmwIDAQABAoIBAAMp9NJsFPX6TWz7\nujESLKgP/f5a13IWiafPe+KinWX09THEjJm5Xt3I8Awlqb9cBHjqD8E+8cRfR311\n7FHzlG36gmMuauJ5bx3dD5qHpWt7/HIAX9osBM0Qx68YWNiqYkv7yIdw1owbSpw3\n4GJHRPXvH3NERaxFJK1ewZrsjxNgEQmJHvougp4OjHX3dpQ3QsFLC0Xi6cptG6EQ\nFE6ZkJwPeRQbIteuZ63E1bmo844LnA/8np9/aXgFgO1zoDBprxCkZ+hw4HgRXTfc\nIp6rhtZx1vCJ+TV04UbItvAlnbEa6BnGViwgUr3OEWsJPdblMFu3Jxjdx7cWhY6T\nESshvAECgYEA7XAjshtm4IkrYyTizDvfTqrUPE4fssLDibOImEEVnryq93WRYvSd\ntUUnQSmFQaoqIA2LaUKZXhgrSEcmwBUmZ7MEZty4l5+LZbh5M+/SoE7E92XB1UHq\nt6CdIIo0qsxsB2F6sxOuySXGoX8x3yAw4I5Vnx7IloVCqiFo2Rdpo8ECgYEA0h1W\nYGljkfVBYjQN3aqQyG6+mwmbH5pcCZ9U4q0GPVna5x+NfCDaphhKzLfpFARi9TVV\nD/3VYsge+vrR6i5yJHY/CWSJ6o1qtZX7z6FT1YCxBpnumbfDU+aOkuDhZe0pAW6F\nd+OpaPzfVqErLOWfIQKUiDvJ1Fzf/1PqB+xulVsCgYEArTXORILq5MgMJ7Jj5PuU\nZu4GD8wqIZ0lSlH3RqLMI9WcnxNcMSUj23YMMeQZOxLo0iMvXWVhPpxBDuQg56VG\nNAYLIwLHgoy4A0e022eLbrinxZas9Wa11KlNlsxbqXGhKS06dmLYchKluXBxsETq\ngxYybfbl+7BRNo1S0HPc88ECgYEArHjwgkuANDuB5D6ecyqjliNxyonkD9kBW9Sy\nNS+aQ5oSmk3IMfA2Csk9/TWp2YiQQn/4xxuxoVhsNdDVpslhYJS/wdIJc8OJqGzp\nOtJOop70sVqCBSRW7fP/bImz+5rlYgt6+6KhWVDZc57wdfY23T3k9r89OVjTdh9H\nqCFn1KkCgYEAtyneA3q3kOhhGWi0VXf83tDbLp6S/lcBhkbcAXZ22ZuIkOMdXh4s\nx6QRM3ndeH+8X6HBeI5Zuce9WsZseDV/U+FGqrsJ9esM1yKF3OYJt+tPL/FzYkax\nGnjcSGbkqhxRSD/l/3v6cLHxVtaCXEF/z0dEv43bujdn9wsm1FYu1cU=\n-----END PRIVATE KEY-----\n"
	info, err3 := p.DecryptLicense(privateKey)
	t.Log(info)
	t.Log(err3)
}
