package hibp

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	DefaultBaseURL = "https://api.pwnedpasswords.com/range/"
	DefaultTimeoutSeconds = 2 * time.Second
)

func IsPwned(pw string) (bool, error) {
	client := http.Client{
		Timeout:       DefaultTimeoutSeconds,
	}

	prefix, suffix := asHashes(pw)

	res, err := client.Get(DefaultBaseURL+prefix)
	if err != nil {
		return false, err
	}

	if res.StatusCode != http.StatusOK {
		return false, fmt.Errorf("bad status returned from HIBP. Code: %v, Body: %s", res.StatusCode, res.Body)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}

	if strings.Contains(strings.ToLower(string(body)), strings.ToLower(suffix)) {
		return true, nil
	}

	return false, nil
}

// asHashes returns first six characters of sha1 hash rest of hash for comparison
func asHashes(pw string) (string, string) {
	h := sha1.Sum([]byte(pw))
	str := fmt.Sprintf("%x", h)
	return str[0:5], str[5:]
}