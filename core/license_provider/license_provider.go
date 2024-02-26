package license_provider

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/lrayt/light-boot/pkg/date"
	"github.com/lrayt/light-boot/pkg/ip"
	"os"
	"path/filepath"
)

type LicenseProvider struct {
	PrivateKeyPath string
	PublicKeyPath  string
	LicensePath    string
}

func NewLicenseProvider(workDir string) *LicenseProvider {
	return &LicenseProvider{
		LicensePath:    filepath.Join(workDir, "license"),
		PrivateKeyPath: filepath.Join(workDir, "private.pem"),
		PublicKeyPath:  filepath.Join(workDir, "public.pem"),
	}
}

// SaveFile 保存文件
func (p LicenseProvider) SaveFile(path string, data []byte) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	if n, err1 := file.Write(data); err != nil || n != len(data) {
		return errors.New(fmt.Sprintf("文件保存失败,err:%v\n", err1))
	}

	return file.Close()
}

// GenKey 生成密钥（公钥、私钥）
func (p LicenseProvider) GenKey() error {
	// 私钥
	privateKey, err1 := rsa.GenerateKey(rand.Reader, 2048)
	if err1 != nil {
		return err1
	}
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privateKeyBytes})
	if err := p.SaveFile(p.PrivateKeyPath, privateKeyPEM); err != nil {
		return err
	}

	// 公钥
	publicKey := &privateKey.PublicKey
	publicKeyBytes, err2 := x509.MarshalPKIXPublicKey(publicKey)
	if err2 != nil {
		return err2
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: publicKeyBytes})

	return p.SaveFile(p.PublicKeyPath, publicKeyPEM)
}

// GenLicense 生成证书（使用公钥对授权信息加密）
func (p LicenseProvider) GenLicense(info *LicenseInfo) error {
	// 读取公钥
	publicKeyPEM, err1 := os.ReadFile(p.PublicKeyPath)
	if err1 != nil {
		return err1
	}
	publicKeyBlock, _ := pem.Decode(publicKeyPEM)
	publicKeyParsed, err2 := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err2 != nil {
		return err2
	}
	publicKey, ok := publicKeyParsed.(*rsa.PublicKey)
	if !ok {
		return errors.New("无法解析为RSA公钥")
	}

	plaintext, err1 := json.Marshal(info)
	if err1 != nil {
		return err1
	}

	ciphertext, err3 := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plaintext)
	if err3 != nil {
		return err3
	}
	return p.SaveFile(p.LicensePath, ciphertext)
}

// LoadLicenseInfo 加载证书解密（使用私钥对证书解密）
func (p LicenseProvider) LoadLicenseInfo() (*LicenseInfo, error) {
	privateKeyPEM, err1 := os.ReadFile(p.PrivateKeyPath)
	if err1 != nil {
		return nil, err1
	}
	privateKeyBlock, _ := pem.Decode(privateKeyPEM)
	privateKey, err2 := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err2 != nil {
		return nil, err2
	}

	ciphertext, err3 := os.ReadFile(p.LicensePath)
	if err3 != nil {
		return nil, err3
	}

	decryptedText, err4 := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err4 != nil {
		return nil, err4
	}

	var info = new(LicenseInfo)
	if err := json.Unmarshal(decryptedText, info); err != nil {
		return nil, err
	}
	return info, nil
}

// DecryptLicense 证书解密
func (p LicenseProvider) DecryptLicense(privateKeyPEM string) (*LicenseInfo, error) {
	privateKeyBlock, _ := pem.Decode([]byte(privateKeyPEM))
	privateKey, err2 := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err2 != nil {
		return nil, err2
	}

	ciphertext, err3 := os.ReadFile(p.LicensePath)
	if err3 != nil {
		return nil, err3
	}

	decryptedText, err4 := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err4 != nil {
		return nil, err4
	}

	var info = new(LicenseInfo)
	if err := json.Unmarshal(decryptedText, info); err != nil {
		return nil, err
	}
	return info, nil
}

func (p LicenseProvider) CheckByPrivateKey(privateKeyPEM string) error {
	info, err1 := p.DecryptLicense(privateKeyPEM)
	if err1 != nil {
		return errors.New("授权证书不合法！")
	}

	var sign = info.GenSign()
	if sign != info.Sign {
		return errors.New("授权证书内容已被修改，证书失效！")
	}

	if !info.LicenseActivated && info.LicenseExpiresDate <= date.NowTime().Unix() {
		return errors.New(fmt.Sprintf("授权证书超过可激活时间[%s]，证书已失效", date.FormatDate(info.LicenseExpiresDate)))
	}

	if info.ServiceExpiresDate <= date.NowTime().Unix() {
		return errors.New(fmt.Sprintf("已超过授权截至时间[%s]，证书已失效", date.FormatDate(info.ServiceExpiresDate)))
	}

	if len(info.MacAddress) > 0 {
		macList, err2 := ip.GetMacAddressByNet()
		if err2 != nil || len(macList) <= 0 {
			return errors.New(fmt.Sprintf("获取主机标识失败，err:%v", err2))
		}
		var isContain = false
		for _, s := range macList {
			for _, t := range info.MacAddress {
				isContain = s == t
			}
		}
		if !isContain {
			return errors.New("此设备不在授权列表内！")
		}
	}

	if !info.LicenseActivated {
		info.LicenseActivated = true
		if err := p.GenLicense(info); err != nil {
			return err
		}
	}
	return nil
}
