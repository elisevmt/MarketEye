package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/sjson"
)

func MakeClientInfo(c *fiber.Ctx, ipHeader string) (result []byte) {
	result, _ = sjson.SetBytes(result, "fingerprint", c.Get("Fingerprint"))
	result, _ = sjson.SetBytes(result, "ip", c.Get(ipHeader))
	result, _ = sjson.SetBytes(result, "userAgent", c.Get("User-Agent"))
	return
}

func MakeManualClientInfo(fingerprint, ip, userAgent string) (result []byte) {
	result, _ = sjson.SetBytes(result, "fingerprint", fingerprint)
	result, _ = sjson.SetBytes(result, "ip", ip)
	result, _ = sjson.SetBytes(result, "userAgent", userAgent)
	return
}
