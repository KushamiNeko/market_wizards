package headerutils

//////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"cipherutils"
	"config"
	"net/http"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	CookieName     = "Gozila"
	CookiePathRoot = "/"

	CookieMaxAge = 2 * 60 * 60
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

func SetCookie(w http.ResponseWriter, cookieName, cookieValue, cookiePath string) {
	cookie := new(http.Cookie)
	cookie.Name = cookieName
	cookie.Value = cipherutils.AesCipherInstanceEncode(cookieValue)
	cookie.Path = cookiePath

	cookie.MaxAge = CookieMaxAge

	/////// VERY IMPORTANT ///////
	cookie.HttpOnly = true
	cookie.Secure = config.TLSConnection
	/////// VERY IMPORTANT ///////

	http.SetCookie(w, cookie)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////

func GetCookie(r *http.Request, cookieName string) (string, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return "", err
	}

	value, err := cipherutils.AesCipherInstanceDecode(cookie.Value)
	if err != nil {
		return "", err
	}

	return value, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////

func DeleteCookie(w http.ResponseWriter, cookieName, cookiePath string) {
	cookie := new(http.Cookie)
	cookie.Name = cookieName
	cookie.Value = "DELETED"
	cookie.Path = cookiePath

	cookie.MaxAge = -1

	/////// VERY IMPORTANT ///////
	cookie.HttpOnly = true
	cookie.Secure = config.TLSConnection
	/////// VERY IMPORTANT ///////

	http.SetCookie(w, cookie)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
