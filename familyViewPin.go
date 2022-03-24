package mobileauth

import (
	"net/http/cookiejar"
	"net/url"
)

func (a *SteamGuardAccount) FamilyPin(pin string) error {
	qp := url.Values{}
	qp.Set("pin", pin)
	qp.Set("sessionid", a.Session.SessionID)
	cookies, _ := cookiejar.New(nil)
	a.Session.AddCookies(cookies)
	resp, err := WebRequestRaw("https://steamcommunity.com"+"/parental"+"/ajaxunlock", "POST", &qp, cookies, nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	for _, cook := range resp.Cookies() {
		if cook.Name == "steamparental" {
			a.Session.SteamParental = cook.Value
		}
	}
	return nil
}
