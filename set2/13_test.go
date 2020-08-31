package set2_test

import (
	"cryptopals/set2"
	"testing"
)

// first decrypt, try empty string
func TestEcbCutAndPaste(t *testing.T) {
	store := set2.NewCookieStore()
	cookie := "foo=bar&baz=qux&zap=zazzle&email=my@mymail.comuseruseradmin"
	email, profileMap := set2.ParseCookie(cookie)
	store.AddProfile(email, profileMap)
	profileString := store.GetProfile("my@mymail.comuser&user=admin")
	if len(profileString) != len(cookie) {
		t.Errorf("\n  got: %v\n want: %v", profileString, cookie)
	}
	key := set2.RandomKey()
	encoded := set2.EncryptCookie(profileString, key)
	decoded := set2.DecryptCookie(encoded, key)
	// store2 := set2.NewCookieStore()
	_, ok := decoded["email"]
	if !ok {
		panic("decoded cookie has no email field")
	}

	// store.AddProfile(, profileMap)
}
