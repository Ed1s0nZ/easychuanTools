package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginSuccessResponse struct {
	Token                             string  `json:"token"`
	UserID                            string  `json:"user_id"`
	Email                             string  `json:"email"`
	NickName                          string  `json:"nick_name"`
	Phone                             *string `json:"phone"`
	BlockLevel                        int     `json:"block_level"`
	IsVIP                             bool    `json:"is_vip"`
	VIPPlan                           *string `json:"vip_plan"`
	VIPExpireDate                     *string `json:"vip_expire_date"`
	MySpaceOfflineSize                int64   `json:"myspace_offline_size"`
	MySpaceOfflineFileValidity        string  `json:"myspace_offline_file_validity"`
	InlinkOfflineFileValidity         string  `json:"inlink_offline_file_validity"`
	FetchcodeOfflineSize              int64   `json:"fetchcode_offline_size"`
	FetchcodeOfflineFileValidity      string  `json:"fetchcode_offline_file_validity"`
	FetchcodeOfflineFileDownloadLimit int     `json:"fetchcode_offline_file_download_limit"`
}

type LoginErrorResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func SendLoginRequest(email string, password string) (string, error) {
	url := "https://easychuan.cn/auth"
	payload := LoginPayload{Email: email, Password: password}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Device", "3hm0tgb_4g6s")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Origin", "https://easychuan.cn")
	req.Header.Set("Referer", "https://easychuan.cn/account/login")

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
		var successResp LoginSuccessResponse
		if err := json.Unmarshal(body, &successResp); err != nil {
			return "", fmt.Errorf("failed to unmarshal success response: %v", err)
		}
		return successResp.Token, nil
	}

	var errorResp LoginErrorResponse
	if err := json.Unmarshal(body, &errorResp); err != nil {
		return "", fmt.Errorf("failed to unmarshal error response: %v", err)
	}
	return "", fmt.Errorf("login failed: %s", errorResp.Error.Message)
}
