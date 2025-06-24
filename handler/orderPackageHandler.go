package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/textures1245/payso-check-slip-backend/model"
)

type OrderPackageHandler interface {
}

type orderPackageHandler struct {
}

func NewOrderPackageHandler() orderPackageHandler {
	return orderPackageHandler{}
}

func CheckPaymentTransaction(ref string) bool {

	var response []model.OrderDetailPost
	url := os.Getenv("SERVER_URL") + `/order/orderdetailpost`
	method := "POST"
	log.Printf("CheckPaymentTransaction URL: %s", url)
	payload := strings.NewReader(`{
		"merchantId": "` + os.Getenv("PUBLIC_ORDER_MID") + `",
		"refno": "` + ref + `",
		"productDetail": "Sure Sure Package"
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Errorf("Handler CheckPaymentTransaction Error: %#v", err)
		return false
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("ngrok-skip-browser-warning", "true")
	req.Header.Add("apikey", os.Getenv("PUBLIC_ORDER_APIKEY"))
	req.Header.Add("merchantSecretKey", os.Getenv("PUBLIC_ORDER_SECRETKEY"))
	req.Header.Add("merchantID", os.Getenv("PUBLIC_ORDER_MID"))

	res, err := client.Do(req)
	if err != nil {
		log.Errorf("Handler CheckPaymentTransaction Error: %#v", err)
		return false
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	if err != nil {
		log.Errorf("Handler CheckPaymentTransaction Error: %#v", err)
		return false
	}
	if len(response) == 0 {
		return false
	}
	return true
}
