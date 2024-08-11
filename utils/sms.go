package utils

import (
	"UserLoginSystem/config"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type SmsRequest struct {
	ApiKey  string `json:"api_key"`
	To      string `json:"to"`
	Message string `json:"message"`
}

func SendSms(to string, message string) error {
	smsRequest := SmsRequest{
		ApiKey:  config.SmsApiKey,
		To:      to,
		Message: message,
	}
	jsonValue, err := json.Marshal(smsRequest)
	if err != nil {
		return err
	}
	resp, err := http.Post(config.SmsApiUrl, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send SMS, status code: %d", resp.StatusCode)
	}
	return nil
}
