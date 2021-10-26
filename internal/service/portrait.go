package service

import (
	"bytes"
	"errors"
	"log"
	"net"
	"time"

	"github.com/xielizyh/ctid_service/pkg/app"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type PortraitAuthRequest struct {
	UserName   string `json:"user_name" binding:"required,min=2,max=100"`
	CertNumber string `json:"cert_number" binding:"required"`
	Photo      string `json:"photo" binding:"required"`
}

// Portrait 人像认证
// java -jar AuthenticationClientDemo2.0.jar 张三 511381199503624578 aGVsbG93b3JsZA==
func PortraitAuth(param *PortraitAuthRequest) (string, error) {
	// // windows: echo hello
	// // cmd := exec.Command("cmd", "/C", "echo hello")
	// // linux: echo hello
	// // cmd := exec.Command("bash", "-c", "echo hello")
	// // cmd := exec.Command("cmd", "/C", "java -jar -h")
	// var err error
	// authCmd := app.GetAuthCmd()
	// // param.Photo 参数太长，不能执行. The filename or extension is too long.
	// cmd := exec.Command("cmd", "/C", authCmd, param.UserName, param.CertNumber, param.Photo)
	// log.Printf("The auth cmd is %v\n", cmd)
	// output, err := cmd.CombinedOutput()
	// if err != nil {
	// 	return "", err
	// }
	// log.Println(string(output))

	// // 最后结果为：00XX\r\n
	// result := string(output[len(output)-6 : len(output)-2])

	// return result, nil

	addr := app.GetAuthServerAddr()
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	// 写身份信息
	var info bytes.Buffer
	// info.WriteString(param.UserName)
	userName, _ := simplifiedchinese.GBK.NewEncoder().String(param.UserName)
	info.WriteString(userName)
	info.WriteString(" ")
	info.WriteString(param.CertNumber)
	info.WriteString(" ")
	info.WriteString(param.Photo)
	info.WriteString("\n")
	// info := param.UserName + " " + param.CertNumber + " " + param.Photo + "\n"
	log.Println(info.String())
	_, err = conn.Write(info.Bytes())
	if err != nil {
		return "", err
	}

	// 读认证结果: 长度+认证结果码
	var result string
	recvBuf := make([]byte, 6)
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	recvLen, err := conn.Read(recvBuf)
	if err != nil {
		return "", err
	}

	if recvLen == 6 {
		result = string(recvBuf[2:])
	} else {
		return "", errors.New("auth server return result length error")
	}
	// log.Println(result)

	return result, nil
}
