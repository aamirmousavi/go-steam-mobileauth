package mobileauth

import (
	"fmt"
)

func (a *SteamGuardAccount) ReLogin(user, pass string) (*SessionData, error) {
	ul := NewUserLogin(user, pass)
	for {
		if err := ul.DoLogin(); err != nil {
			switch err {
			case ErrNeedEmail:
				fmt.Print("Please enter your email code: ")
				var code string
				fmt.Scanln(&code)
				ul.EmailCode = code
			case ErrNeedCaptcha:
				fmt.Println("https://steamcommunity.com/public/captcha.php?gid=" + ul.CaptchaGID)
				fmt.Println("Please follow link to get captcha text.")
				fmt.Print("Please enter captcha text: ")
				var captchaText string
				fmt.Scanln(&captchaText)
				ul.CaptchaText = captchaText
			case ErrNeed2FA:
				fmt.Print("Please enter your mobile authenticator code: ")
				var code string
				fmt.Scanln(&code)
				ul.TwoFactorCode = code
			default:
				fmt.Printf("Failed to login: %v\n", err)
				return nil, err
			}
		} else {
			break
		}
	}
	return ul.Session, nil
}
