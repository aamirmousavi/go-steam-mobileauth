package mobileauth

import (
	"encoding/json"
	"fmt"
	"net/http/cookiejar"
)

func (a *SteamGuardAccount) ConfirmationsDetail(ConfirmationId string) (string, error) {
	qp, err := a.GenerateConfirmationQueryParams("details")
	if err != nil {
		return "", err
	}
	cookies, _ := cookiejar.New(nil)
	a.Session.AddCookies(cookies)
	respBody, err := WebRequest(UrlConfirmationService+"/details/"+ConfirmationId, "GET", &qp, cookies, nil, nil)
	if err != nil {
		return "", err
	}
	respString := string(respBody)
	var _m map[string]interface{}
	err = json.Unmarshal([]byte(respString), &_m)
	if err != nil {
		fmt.Println("json 1 error: ", err)
		return "", err
	}
	if ok, okType := _m["success"].(bool); !ok || !okType {
		return "", fmt.Errorf("success is false")
	}
	html, ok := _m["html"].(string)
	if !ok {
		return "", fmt.Errorf("'html' key type is %T not string", _m["html"])
	}
	return html, nil
}
