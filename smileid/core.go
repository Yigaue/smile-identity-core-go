package smileid

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"time"
)

func GenerateSignature(config *Config) string {
	ts := time.Now().Format(time.RFC3339)
	message := generateMessage(ts, config.PartnerId)
	hmc := hmac.New(sha256.New, []byte(config.ApiKey))
	hmc.Write([]byte(message.String()))
	return base64.StdEncoding.EncodeToString(hmc.Sum(nil))
}

func generateMessage(timestamp string, partnerId string) strings.Builder {
	var message strings.Builder
  message.WriteString(timestamp)
  message.WriteString(partnerId)
  message.WriteString("sid_request")
	return message
}

func ConfirmSignature(receivedSignature string, receivedTimestamp string, partnerId string, apiKey string) bool {
	generatedSignature := GenerateSignature(&Config{PartnerId: partnerId, ApiKey: apiKey})
	return generatedSignature == receivedSignature
}

