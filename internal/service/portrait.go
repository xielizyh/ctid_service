package service

import (
	"log"
	"os/exec"

	"github.com/xielizyh/ctid_service/pkg/app"
)

type PortraitAuthRequest struct {
	UserName   string `json:"user_name" binding:"required,min=2,max=100"`
	CertNumber string `json:"cert_number" binding:"required"`
	Photo      string `json:"photo" binding:"required"`
}

// Portrait 人像认证
// java -jar AuthenticationClientDemo2.0.jar 张三 511381199503624578 aGVsbG93b3JsZA==
func PortraitAuth(param *PortraitAuthRequest) (string, error) {
	// windows: echo hello
	// cmd := exec.Command("cmd", "/C", "echo hello")
	// linux: echo hello
	// cmd := exec.Command("bash", "-c", "echo hello")
	// cmd := exec.Command("cmd", "/C", "java -jar -h")
	var err error
	authCmd := app.GetAuthCmd()
	cmd := exec.Command("cmd", "/C", authCmd, param.UserName, param.CertNumber, param.Photo)
	log.Printf("The auth cmd is %v\n", cmd)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	log.Println(string(output))

	// 最后结果为：00XX\r\n
	result := string(output[len(output)-6 : len(output)-2])

	return result, nil
}
