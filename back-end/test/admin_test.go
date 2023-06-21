package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/npcxax/iBooking/pkg/models"
)

// 设置服务器地址、端口和协议
var serverURL = "http://10.177.88.168:8800"

// 创建一个HTTP客户端
var client = &http.Client{}

// 添加令牌到请求头
var (
	userToken  = ""
	adminToken = ""
	userID     = ""
	userName   = "shunzige"
)

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODczNTE5NzYsIm9yaWdfaWF0IjoxNjg3MjY1NTc2LCJ1c2VybmFtZSI6InNodW4ifQ.-3dIRADAlhUpKi7SmfoWIp3Bgsvu0ZrbDNGWE29B6X8"

//func TestCreateAdmin(t *testing.T) {
//	// 设置路由
//	apiEndpoint := "/admin"
//
//	// 设置请求内容
//	Admin := models.Administrator{
//		Username: "shun",
//		Password: "123456",
//	}
//
//	j, err := json.Marshal(Admin)
//	assert.NoError(t, err)
//	s := string(j)
//	body := strings.NewReader(s)
//
//	// 构建测试请求
//	req, err := http.NewRequest(http.MethodPost, serverURL+apiEndpoint, body)
//	assert.NoError(t, err)
//
//	// 发送请求并获取响应
//	resp, err := client.Do(req)
//	assert.NoError(t, err)
//
//	// 验证响应状态码
//	assert.Equal(t, http.StatusOK, resp.StatusCode)
//	// 验证其他响应属性或内容
//}

func TestAdminLogin(t *testing.T) {
	// 设置路由
	apiEndpoint := "/admin/login/"

	// 设置请求内容
	loginData := map[string]string{
		"username": "shun",
		"password": "123456",
	}

	jsonData, _ := json.Marshal(loginData)

	// 构建测试请求
	req, err := http.NewRequest(http.MethodPost, serverURL+apiEndpoint, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)

	// 发送请求并获取响应
	resp, err := client.Do(req)
	assert.NoError(t, err)

	// 获取admin token
	body, err := io.ReadAll(resp.Body)
	bodyStr := string(body)
	// fmt.Println(bodyStr)
	adminToken = bodyStr[53:212]
	// fmt.Printf("%v\n", adminToken)

	// 验证响应状态码
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// 验证其他响应属性或内容
}

func TestCreateUser(t *testing.T) {
	// 设置路由
	apiEndpoint := "/user/"

	// 设置请求内容
	Admin := models.User{
		Username: userName,
		Password: "123456",
	}

	j, err := json.Marshal(Admin)
	assert.NoError(t, err)
	s := string(j)
	body := strings.NewReader(s)

	// 构建测试请求
	req, err := http.NewRequest(http.MethodPost, serverURL+apiEndpoint, body)
	assert.NoError(t, err)

	// 发送请求并获取响应
	resp, err := client.Do(req)
	assert.NoError(t, err)

	// 验证响应状态码
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	// 验证其他响应属性或内容
}

func TestUserLogin(t *testing.T) {
	// 设置路由
	apiEndpoint := "/user/login"

	// 设置请求内容
	loginData := map[string]string{
		"username": userName,
		"password": "123456",
	}

	jsonData, _ := json.Marshal(loginData)

	// 构建测试请求
	req, err := http.NewRequest(http.MethodPost, serverURL+apiEndpoint, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)

	// 发送请求并获取响应
	resp, err := client.Do(req)
	assert.NoError(t, err)

	// 验证响应状态码
	if !assert.Equal(t, http.StatusOK, resp.StatusCode) {
		return
	}

	// 获取user token
	body, err := io.ReadAll(resp.Body)
	bodyStr := string(body)
	// fmt.Println(bodyStr)
	userToken = bodyStr[53:217]
	// fmt.Printf("%v\n", userToken)

	// 验证其他响应属性或内容
}

func TestUserLogout(t *testing.T) {
	// 设置路由
	apiEndpoint := "/user/auth/logout"

	// 设置请求内容

	// 构建测试请求
	req, err := http.NewRequest(http.MethodPost, serverURL+apiEndpoint, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+userToken)
	assert.NoError(t, err)

	// 发送请求并获取响应
	resp, err := client.Do(req)
	assert.NoError(t, err)

	// 验证响应状态码
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// 验证其他响应属性或内容
}

func TestUserRefreshToken(t *testing.T) {
	// 设置路由
	apiEndpoint := "/user/auth/refreshToken"

	// 设置请求内容

	// 构建测试请求
	req, err := http.NewRequest(http.MethodPost, serverURL+apiEndpoint, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+userToken)
	assert.NoError(t, err)

	// 发送请求并获取响应
	resp, err := client.Do(req)
	assert.NoError(t, err)

	// 验证响应状态码
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// 验证其他响应属性或内容
}

func TestGetUserByUsername(t *testing.T) {
	// 设置路由
	apiEndpoint := "/user/auth/getUserByUsername/" + userName

	// 设置请求内容

	// 构建测试请求
	req, err := http.NewRequest(http.MethodGet, serverURL+apiEndpoint, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+userToken)
	assert.NoError(t, err)

	// 发送请求并获取响应
	resp, err := client.Do(req)
	assert.NoError(t, err)

	// 验证响应状态码
	if !assert.Equal(t, http.StatusOK, resp.StatusCode) {
		return
	}

	// 获取userID
	body, err := io.ReadAll(resp.Body)
	bodyStr := string(body)
	// fmt.Println(bodyStr)
	userID = bodyStr[38:57]
	// fmt.Printf("%v\n", userID)

	// 验证其他响应属性或内容
}

func TestGetUserByID(t *testing.T) {
	// 设置路由
	apiEndpoint := "/user/auth/getUserByID/" + userID

	// 设置请求内容

	// 构建测试请求
	req, err := http.NewRequest(http.MethodGet, serverURL+apiEndpoint, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+userToken)
	assert.NoError(t, err)

	// 发送请求并获取响应
	resp, err := client.Do(req)
	assert.NoError(t, err)

	// 验证响应状态码
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// 验证其他响应属性或内容
}

func TestUpdatePassword(t *testing.T) {
	// 设置路由
	apiEndpoint := "/user/auth/password"

	// 设置请求内容
	loginData := map[string]string{
		"user_id":  userID,
		"password": "123",
	}

	jsonData, _ := json.Marshal(loginData)

	// 构建测试请求
	req, err := http.NewRequest(http.MethodPost, serverURL+apiEndpoint, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+userToken)
	assert.NoError(t, err)

	// 发送请求并获取响应
	resp, err := client.Do(req)
	assert.NoError(t, err)

	// 验证响应状态码
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// 验证其他响应属性或内容
}

func TestUpdateUser(t *testing.T) {
	// 设置路由
	apiEndpoint := "/user/auth/updateUser"

	// 设置请求内容
	updateData := map[string]string{
		"user_id": userID,
		"email":   "12345@qq.com",
	}

	jsonData, _ := json.Marshal(updateData)

	// 构建测试请求
	req, err := http.NewRequest(http.MethodPost, serverURL+apiEndpoint, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+userToken)
	assert.NoError(t, err)

	// 发送请求并获取响应
	resp, err := client.Do(req)
	assert.NoError(t, err)

	// 验证响应状态码
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// 验证其他响应属性或内容
}
