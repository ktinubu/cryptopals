package set2

import (
	set1 "cryptopals/set_1"
	"fmt"
	"strings"
)

type CookieStore struct {
	users map[string]map[string]string
}

func RandomKey() []byte {
	return RandBytes(set1.AesBlockLen)
}

func NewCookieStore() *CookieStore {
	return &CookieStore{users: make(map[string]map[string]string)}
}

func (c *CookieStore) AddProfile(email string, profile map[string]string) {
	if email != "" {
		c.users[email] = profile
	}
}

func ParseCookie(cookie string) (email string, profile map[string]string) {
	tuples := strings.Split(cookie, "&")
	profile = make(map[string]string)
	for _, t := range tuples {
		tokens := strings.SplitN(t, "=", -1)
		k, v := tokens[0], tokens[1]
		profile[k] = v
		if k == "email" {
			email = v
		}
	}
	return email, profile
}

func (c *CookieStore) GetProfile(email string) string {
	email = strings.ReplaceAll(email, "&", "")
	email = strings.ReplaceAll(email, "=", "")
	profile, ok := c.users[email]
	if !ok {
		return ""
	}
	tuples := []string{}
	for k, v := range profile {
		tuples = append(tuples, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(tuples, "&")
}

func EncryptCookie(cookie string, key []byte) []byte {
	ecbEnc := set1.AesEcbEncrypter(key)
	src := Pkcs7Pad([]byte(cookie), set1.AesBlockLen)
	dst := make([]byte, len(src))
	ecbEnc.EncryptBlocks(dst, src)
	return dst
}

func DecryptCookie(cookie, key []byte) map[string]string {
	ecbDec := set1.AesEcbDecrypter(key)
	dst := make([]byte, len(cookie))
	ecbDec.DecryptBlocks(dst, cookie)
	_, profile := ParseCookie(string(dst))
	return profile
}
