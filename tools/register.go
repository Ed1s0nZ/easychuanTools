package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type SuccessResponse struct {
	Token                  string  `json:"token"`
	UserID                 string  `json:"user_id"`
	Email                  string  `json:"email"`
	NickName               string  `json:"nick_name"`
	Phone                  *string `json:"phone"`
	IsVIP                  bool    `json:"is_vip"`
	MySpaceSize            int64   `json:"myspace_offline_size"`
	MySpaceValidity        string  `json:"myspace_offline_file_validity"`
	InlinkValidity         string  `json:"inlink_offline_file_validity"`
	FetchcodeSize          int64   `json:"fetchcode_offline_size"`
	FetchcodeValidity      string  `json:"fetchcode_offline_file_validity"`
	FetchcodeDownloadLimit int     `json:"fetchcode_offline_file_download_limit"`
}

type ErrorResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// 注册
func SendRegRequest(email string, password string) (string, error) {
	url := "https://easychuan.cn/users"
	payload := UserPayload{Email: email, Password: password}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Device", "3hm0tgb_kh2g")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode == http.StatusOK {
		var successResp SuccessResponse
		if err := json.Unmarshal(body, &successResp); err != nil {
			return "", fmt.Errorf("failed to unmarshal success response: %v", err)
		}
		return successResp.Token, nil
	}

	var errorResp ErrorResponse
	if err := json.Unmarshal(body, &errorResp); err != nil {
		return "", fmt.Errorf("failed to unmarshal error response: %v", err)
	}
	return "", fmt.Errorf("registration failed: %s", errorResp.Error.Message)
}
