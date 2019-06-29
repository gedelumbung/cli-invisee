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
		Url       string
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
		}
	}
)

func Init(env string) *Invisee {
	url := "https://devmcw.invisee.com"
	agentCode := "Your Agent Code Here"
	if env == "production" {
		url = "https://api.invisee.com"
		agentCode = "Your Agent Code Here"
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

	r, err := req.Post(inv.Url+"/customer/login", req.BodyJSON(map[string]interface{}{
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

	r, err := req.Post(inv.Url+"/investment/list", req.BodyJSON(map[string]interface{}{
		"token": login.Data.Token,
	}))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(r)
	return nil
}
