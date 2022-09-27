package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type tokenInfo struct {
	Aud           string `json:"aud"`
	Azp           string `json:"azp"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Exp           string `json:"exp"`
	Iat           string `json:"iat"`
	Iss           string `json:"iss"`
	Sub           string `json:"sub"`
	Alg           string `json:"alg"`
	Kid           string `json:"kid"`
	Typ           string `json:"typ"`
}

func IsAuthorized(jwt string) (bool, error) {
	resp, err := http.Get(fmt.Sprintf("https://oauth2.googleapis.com/tokeninfo?id_token=%s", jwt))
	if err != nil {
		return false, err
	}

	// If the response status code is 200, then the token is valid.
	if resp.StatusCode == 200 {
		var tokenInfo tokenInfo
		err = json.NewDecoder(resp.Body).Decode(&tokenInfo)
		if err != nil {
			fmt.Println("something went wrong decoding the token info:", err)
			return false, err
		}

		// Check that the token is not expired.
		exp, err := strconv.ParseInt(tokenInfo.Exp, 10, 64)
		if err != nil {
			fmt.Println("something went wrong parsing the token expiration time:", err)
			return false, err
		}
		if exp < time.Now().UTC().Unix() {
			fmt.Println("the token is expired.")
			return false, nil
		}

		// If the now time is more than issued at time by 5 seconds, then consider the token invalid.
		// This is to make sure a token is not used maliciously.
		// I could implement some database to block used JWTs, but I am lazy :).
		iat, err := strconv.ParseInt(tokenInfo.Iat, 10, 64)
		if err != nil {
			fmt.Println("something went wrong parsing the token issued at time:", err)
			return false, err
		}
		if time.Now().UTC().Unix()-iat > 5 {
			fmt.Println("the token is used maliciously after more than 5 seconds.")
			return false, nil
		}

		return true, nil
	}

	// Any other status code means the token is invalid.
	// The API endpoint above returns 400 in case it is invalid, but this is just to be sure.
	return false, nil
}
