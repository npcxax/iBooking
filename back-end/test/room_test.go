package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type room struct {
	ID         int64  `json:"id"`
	RoomNumber string `json:"room_number"`
	Location   string `json:"location"`
	OpenTime   string `json:"open_time"`
	CloseTime  string `json:"close_time"`
	Overnight  bool   `json:"overnight"`
}

var roomID string = ""

func TestCreateRoom(t *testing.T) {
	// 设置路由
	apiEndpoint := "/room/auth/createRoom"

	// 设置请求内容
	Room := room{
		RoomNumber: "101",
		Location:   "南楼",
		OpenTime:   "7",
		CloseTime:  "22",
		Overnight:  false,
	}
	jsonData, err := json.Marshal(Room)
	assert.NoError(t, err)

	// 构建测试请求
	req, err := http.NewRequest(http.MethodPost, serverURL+apiEndpoint, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	assert.NoError(t, err)

	// 发送请求并获取响应
	resp, err := client.Do(req)
	assert.NoError(t, err)

	// 验证响应状态码
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// 验证其他响应属性或内容
}

func TestGetRoom(t *testing.T) {
	// 设置路由
	apiEndpoint := "/room/"

	// 设置请求内容

	// 构建测试请求
	req, err := http.NewRequest(http.MethodGet, serverURL+apiEndpoint, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	assert.NoError(t, err)

	// 发送请求并获取响应
	resp, err := client.Do(req)
	assert.NoError(t, err)

	// 获取roomID
	body, err := io.ReadAll(resp.Body)
	bodyStr := string(body)
	// fmt.Println(bodyStr)
	roomID = bodyStr[16:35]
	// fmt.Printf("%v\n", roomID)

	// 验证响应状态码
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// 验证其他响应属性或内容
}

func TestGetRoomByID(t *testing.T) {
	// 设置路由
	apiEndpoint := "/room/" + roomID

	// 设置请求内容

	// 构建测试请求
	req, err := http.NewRequest(http.MethodGet, serverURL+apiEndpoint, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	assert.NoError(t, err)

	// 发送请求并获取响应
	resp, err := client.Do(req)
	assert.NoError(t, err)

	// 验证响应状态码
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// 验证其他响应属性或内容
}

func TestUpdateRoom(t *testing.T) {
	// 设置路由
	apiEndpoint := "/room/auth/updateRoom"

	// 设置请求内容
	roomData := map[string]string{
		"room_id":     roomID,
		"room_number": "102",
	}
	jsonData, err := json.Marshal(roomData)
	assert.NoError(t, err)

	// 构建测试请求
	req, err := http.NewRequest(http.MethodPost, serverURL+apiEndpoint, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	assert.NoError(t, err)

	// 发送请求并获取响应
	resp, err := client.Do(req)
	assert.NoError(t, err)

	// 验证响应状态码
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// 验证其他响应属性或内容
}
