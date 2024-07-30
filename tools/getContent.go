package tools

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GetContentSuccessResponse struct {
	TextContent string `json:"text_content"`
}

func GetContentRequest(token string) (string, error) {
	url := "https://easychuan.cn/myspace/texts"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Sec-Ch-Ua", `"Chromium";v="91", " Not;A Brand";v="99"`)
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Device", "xxxx_x123")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://easychuan.cn/myspace")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Connection", "close")

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
		var successResp GetContentSuccessResponse
		if err := json.Unmarshal(body, &successResp); err != nil {
			return "", fmt.Errorf("failed to unmarshal success response: %v", err)
		}
		fmt.Printf("文本内容: %s\n", successResp.TextContent)
		return successResp.TextContent, nil
	} else {
		var errorResp ErrorResponse
		if err := json.Unmarshal(body, &errorResp); err != nil {
			return "", fmt.Errorf("failed to unmarshal error response: %v", err)
		}
		return "", fmt.Errorf("获取文本信息失败，错误信息: %s", errorResp.Error.Message)
	}

}
