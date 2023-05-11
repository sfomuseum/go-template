package funcs

import (
	"encoding/base64"
	"fmt"
	_ "log"
	
	"github.com/skip2/go-qrcode"
)

// Encodes 'uri' as a QR code and returns it as base-64-encoded data URI.
func QRCode(uri string) string {

	qr_png, err := qrcode.Encode(uri, qrcode.Medium, 256)
	
	if err != nil {
		// log.Printf("Failed to encode '%s' as QR code, %w", uri, err)
		return "#"
	}
	
	qr_b64 := base64.StdEncoding.EncodeToString(qr_png)
	qr_uri := fmt.Sprintf("data:image/%s;base64,%s", "image/png", qr_b64)

	return qr_uri
}
