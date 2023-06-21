package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var bookingID = ""

type seat struct {
	ID     string `json:"seat_id"`
	X      string `json:"x"`
	Y      string `json:"y"`
	Status string `json:"status"`
	Plug   bool   `json:"plug"`
}

type createSeatData struct {
	RoomID string          `json:"room_id"`
	Seats  map[string]seat `json:"seats"`
}

var seatID = ""

func TestCreateSeat(t *testing.T) {
	// 设置路由
	apiEndpoint := "/seat/auth/createSeat"

	// 设置请求内容
	mapData := createSeatData{
		RoomID: roomID,
		Seats: map[string]seat{
			"1": {
				X:      "1",
				Y:      "1",
				Status: "1",
				Plug:   false,
			},
			"2": {
				X:      "1",
				Y:      "2",
				Status: "1",
				Plug:   true,
			},
		},
	}
	jsonData, err := json.Marshal(mapData)
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

func TestGetSeat(t *testing.T) {
	// 设置路由
	apiEndpoint := "/seat/"

	// 设置请求内容

	// 构建测试请求
	req, err := http.NewRequest(http.MethodGet, serverURL+apiEndpoint, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	assert.NoError(t, err)

	// 发送请求并获取响应
	resp, err := client.Do(req)
	assert.NoError(t, err)

	// 获取seatID
	body, err := io.ReadAll(resp.Body)
	bodyStr := string(body)
	// fmt.Println(bodyStr)
	seatID = bodyStr[16:35]
	// fmt.Printf("%v\n", seatID)

	// 验证响应状态码
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// 验证其他响应属性或内容
}

func TestGetSeatByID(t *testing.T) {
	// 设置路由
	apiEndpoint := "/seat/" + seatID

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

func TestGetSeatByRoomID(t *testing.T) {
	// 设置路由
	apiEndpoint := "/seat/getSeatByRoomID/" + roomID

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

func TestUpdateSeat(t *testing.T) {
	// 设置路由
	apiEndpoint := "/seat/auth/updateSeat"

	// 设置请求内容
	mapData := map[string]map[string]string{
		"1": {
			"seat_id": seatID,
			"y":       "3",
		},
	}
	jsonData, err := json.Marshal(mapData)
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

// booking test
type bookSeatData struct {
	UserID      string    `json:"user_id"`
	SeatID      string    `json:"seat_id"`
	RoomID      string    `json:"room_id"`
	Duration    string    `json:"duration"`
	BookingTime time.Time `json:"booking_time"`
}

func roundToNearestHour(t time.Time) time.Time {
	rounded := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
	return rounded
}

func TestBookSeat(t *testing.T) {
	// 设置路由
	apiEndpoint := "/booking/"

	// 设置请求内容
	mapData := bookSeatData{
		UserID:      userID,
		SeatID:      seatID,
		RoomID:      roomID,
		Duration:    "3",
		BookingTime: roundToNearestHour(time.Now().Add(time.Hour * 2)),
	}
	jsonData, err := json.Marshal(mapData)
	assert.NoError(t, err)

	// 构建测试请求
	req, err := http.NewRequest(http.MethodPost, serverURL+apiEndpoint, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+userToken)
	assert.NoError(t, err)

	// 发送请求并获取响应
	resp, err := client.Do(req)
	assert.NoError(t, err)

	// 获取bookingID
	body, err := io.ReadAll(resp.Body)
	bodyStr := string(body)
	// fmt.Println(bodyStr)
	bookingID = bodyStr[17:36]
	// fmt.Printf("%v\n", bookingID)

	// 验证响应状态码
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// 验证其他响应属性或内容
}

func TestGetBookingByUserID(t *testing.T) {
	// 设置路由
	apiEndpoint := "/booking/getBookingByUserID/" + userID

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

func TestGetBookingByID(t *testing.T) {
	// 设置路由
	apiEndpoint := "/booking/getBookingByID/" + bookingID

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

func TestUpdateBooking(t *testing.T) {
	// 设置路由
	apiEndpoint := "/booking/updateBooking/"

	// 设置请求内容
	mapData := map[string]string{
		"booking_id": bookingID,
		"is_signed":  "2",
	}
	jsonData, err := json.Marshal(mapData)
	assert.NoError(t, err)

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

func TestDeleteBooking(t *testing.T) {
	// 设置路由
	apiEndpoint := "/booking/deleteBooking"

	// 设置请求内容
	mapData := map[string]string{
		"booking_id": bookingID,
	}
	jsonData, err := json.Marshal(mapData)
	assert.NoError(t, err)

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

func TestGetBookingHistory(t *testing.T) {
	// 设置路由
	apiEndpoint := "/booking/bookingHistory/" + userID

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

// delete testing
func TestDeleteSeat(t *testing.T) {
	// 设置路由
	apiEndpoint := "/seat/auth/deleteSeat"

	// 设置请求内容
	mapData := map[string]string{
		"seat_id": seatID,
	}
	jsonData, err := json.Marshal(mapData)
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

func TestDeleteRoom(t *testing.T) {
	// 设置路由
	apiEndpoint := "/room/auth/deleteRoom"

	// 设置请求内容
	mapData := map[string]string{
		"room_id": roomID,
	}
	jsonData, err := json.Marshal(mapData)
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

func TestDeleteUser(t *testing.T) {
	// 设置路由
	apiEndpoint := "/user/auth/deleteUser"

	// 设置请求内容
	mapData := map[string]string{
		"user_id": userID,
	}
	jsonData, err := json.Marshal(mapData)
	assert.NoError(t, err)

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
