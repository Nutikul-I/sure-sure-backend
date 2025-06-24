package util

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/gofiber/fiber/v2"
)

type APIResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Timestamp  string      `json:"timestamp"`
	Data       interface{} `json:"data"`
}

func JSONResponse(c *fiber.Ctx, statusCode int, code int, data interface{}) {
	statusMessages := map[int]string{
		2000: "Payment Request Created",
		2001: "3D Secure Redirect URI Provided",
		2002: "Payment Authorized",
		2003: "Payment Captured",
		2004: "Void Successful",
		2005: "Refund Successful",
		2006: "Success",
		4000: "Invalid Request",
		4001: "Unauthorized Merchant",
		4002: "Forbidden Operation",
		4003: "Transaction Not Found",
		4004: "Duplicate Request",
		4005: "Invalid Payment Amount",
		4006: "Payment Declined",
		4007: "Transaction Not Eligible",
		4008: "Refund Not Permitted",
		4009: "Invalid 3D Secure Status",
		4010: "Payment Expired",
		4011: "Missing Required Field",
		5000: "Internal Processing Error",
		5001: "Payment Reconciliation Error",
		5002: "Downstream Dependency Failure",
		5003: "Service Unavailable",
		5004: "Gateway Timeout",
		5101: "Cannot connect to V1",
	}
	var message string
	if message == "" {
		if msg, exists := statusMessages[code]; exists {
			message = msg
		} else {
			message = "Unknown Status Code"
		}
	}

	response := APIResponse{
		StatusCode: code,
		Message:    message,
		Timestamp:  time.Now().Format(time.RFC3339),
		Data:       data,
	}

	c.Status(statusCode).JSON(response)
}

func GenerateRandomText(l int) string {
	bytes := make([]byte, l) // ใช้ 16 ไบต์ (16 x 2 = 32 อักขระในรูปแบบ hex)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}
