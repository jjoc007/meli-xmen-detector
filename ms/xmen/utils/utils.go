package utils

import (
	b64 "encoding/base64"
)

func Base64Encode(text string) string {
	return b64.StdEncoding.EncodeToString([]byte(text))
}