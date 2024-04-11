package smileid

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"time"
)


func GenerateSignature(partnerid string) string {
	ts := time.Now().Format(time.RFC3339)
	var message strings.Builder
  message.WriteString(ts)
  message.WriteString(partnerid)
  message.WriteString("sid_request")
	hmc := hmac.New(sha256.New, []byte(partnerid))
	hmc.Write([]byte(message.String()))
	return base64.StdEncoding.EncodeToString(hmc.Sum(nil))
}

