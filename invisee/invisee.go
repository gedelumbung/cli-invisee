package invisee

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/imroc/req"
)

type (
	Invisee struct {
		Env       string
		URL       string
		AgentCode string
	}

	LoginResponse struct {
		Code int    `json:"code"`
		Info string `json:"info"`
		Data struct {
			LastLogin           string `json:"last_login"`
			Token               string `json:"token"`
			CustomerRiskProfile struct {
				Code  string `json:"code"`
				Value string `json:"value"`
			} `json:"customer_risk_profile"`
			CustomerStatus       string `json:"customer_status"`
			CustomerStatusBefore string `json:"customer_status_before"`
		} `json:"data"`
	}

	InvestmentsResponse struct {
		Code int `json:"code"`
		Data struct {
			Investment interface{} `json:"investment"`
		} `json:"data"`
	}

	TransactionsResponse struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
	}

	OrderStatusResponse struct {
		Code int         `json:"code"`
		Info string      `json:"info"`
		Data interface{} `json:"data"`
	}

	RangeOfPartialResponse struct {
		Code int         `json:"code"`
		Info string      `json:"info"`
		Data interface{} `json:"data"`
	}
)

func Init(env string) *Invisee {
	url := "https://devmcw.invisee.com"
	agentCode := "FUNDTC"
	if env == "production" {
		url = "https://api.invisee.com"
		agentCode = "FUNDTC"
	}

	return &Invisee{env, url, agentCode}
}

func Signature(inv *Invisee, customerKey string) string {
	hashSha256 := sha256.Sum256([]byte(inv.AgentCode))
	sha256ToString := hex.EncodeToString(hashSha256[:])
	hashSha384 := sha512.Sum384([]byte(sha256ToString + customerKey))
	sha384ToString := hex.EncodeToString(hashSha384[:])

	return sha384ToString
}

func Login(inv *Invisee, customerCif string, customerKey string) *LoginResponse {
	signature := Signature(inv, customerKey)

	r, err := req.Post(inv.URL+"/customer/login", req.BodyJSON(map[string]interface{}{
		"customer_cif": customerCif,
		"signature":    signature,
	}))

	if err != nil {
		fmt.Println(err)
	}

	var loginResponse *LoginResponse
	r.ToJSON(&loginResponse)
	return loginResponse
}

func Investments(inv *Invisee, customerCif string, customerKey string) interface{} {
	login := Login(inv, customerCif, customerKey)

	if login.Code != 0 {
		return "Failed"
	}

	r, err := req.Post(inv.URL+"/investment/list", req.BodyJSON(map[string]interface{}{
		"token": login.Data.Token,
	}))

	if err != nil {
		fmt.Println(err)
	}

	var investmentsResponse *InvestmentsResponse
	r.ToJSON(&investmentsResponse)
	return investmentsResponse
}

func Transactions(inv *Invisee, customerCif string, customerKey string) interface{} {
	login := Login(inv, customerCif, customerKey)

	if login.Code != 0 {
		return "Failed"
	}

	r, err := req.Post(inv.URL+"/transaction/list", req.BodyJSON(map[string]interface{}{
		"token": login.Data.Token,
	}))

	if err != nil {
		fmt.Println(err)
	}

	var transactionsResponse *TransactionsResponse
	r.ToJSON(&transactionsResponse)
	return transactionsResponse
}

func OrderStatus(inv *Invisee, customerCif string, customerKey string, orderNumber string) interface{} {
	login := Login(inv, customerCif, customerKey)

	if login.Code != 0 {
		return "Failed"
	}

	r, err := req.Post(inv.URL+"/transaction/check_order", req.BodyJSON(map[string]interface{}{
		"token":        login.Data.Token,
		"order_number": orderNumber,
	}))

	if err != nil {
		fmt.Println(err)
	}

	var orderStatusResponse *OrderStatusResponse
	r.ToJSON(&orderStatusResponse)
	return orderStatusResponse
}

func RangeOfPartial(inv *Invisee, customerCif string, customerKey string, orderNumber string) interface{} {
	signature := Signature(inv, customerKey)
	login := Login(inv, customerCif, customerKey)

	if login.Code != 0 {
		return "Failed"
	}

	r, err := req.Post(inv.URL+"/transaction/rangeOfPartial", req.BodyJSON(map[string]interface{}{
		"token":     login.Data.Token,
		"signature": signature,
		"invNo":     orderNumber,
	}))

	if err != nil {
		fmt.Println(err)
	}

	var rangeOfPartialResponse *RangeOfPartialResponse
	r.ToJSON(&rangeOfPartialResponse)
	return rangeOfPartialResponse
}
