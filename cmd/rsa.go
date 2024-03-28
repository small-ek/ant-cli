package cmd

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

type Rsa struct {
}

// Action
func (r Rsa) Action(c *cli.Context) error {
	// 生成RSA私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("私钥生成失败:", err)
		return err
	}

	// 将RSA私钥编码为PEM格式
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	privateKeyFile, err := os.Create("private.pem")
	if err != nil {
		fmt.Println("私钥文件创建失败:", err)
		return err
	}
	defer privateKeyFile.Close()
	pem.Encode(privateKeyFile, privateKeyPEM)

	fmt.Println("私钥已保存到 private.pem")

	// 生成RSA公钥
	publicKey := &privateKey.PublicKey

	// 将RSA公钥编码为PEM格式
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		fmt.Println("公钥编码失败:", err)
		return err
	}
	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicKeyFile, err := os.Create("public.pem")
	if err != nil {
		fmt.Println("公钥文件创建失败:", err)
		return err
	}
	defer publicKeyFile.Close()
	pem.Encode(publicKeyFile, publicKeyPEM)

	fmt.Println("公钥已保存到 public.pem")
	return nil
}
