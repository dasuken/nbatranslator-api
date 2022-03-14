package translator

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

var DefaultDeepLClient Client
var defaultBaseUrl = "https://api-free.deepl.com/v2/translate"

func init() {
	secret := os.Getenv("DEEPL_SECRET")
	DefaultDeepLClient = NewDeepLClient(secret, "JA")
}

type deepLClient struct {
	baseURL string
	secret  string
	target  string
}

func NewDeepLClient(secret, target string) Client {
	return &deepLClient{baseURL: defaultBaseUrl, secret: secret,target: target}
}

type deepL struct {
	Translations    []struct {
		DetectedSourceLanguage    string    `json:"detected_source_language" `
		Text    string    `json:"text" `
	}    `json:"translations" `
}

func (c *deepLClient) Do(text string) (string, error) {
	if len(c.secret) == 0 {
		return "", fmt.Errorf("invalid depeL secret key")
	}

	u := &url.Values{}
	u.Set("auth_key", c.secret)
	u.Set("target_lang", c.target)
	u.Set("text", text)

	targetUrl := c.baseURL + "?" + u.Encode()
	fmt.Println(targetUrl)

	res, err := http.Get(targetUrl)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var d deepL
	if err := json.Unmarshal(b, &d); err != nil {
		return "", err
	}

	if d.Translations == nil {
		return "", errors.New("Unmarshal error. translations is nil")
	}

	return d.Translations[0].Text, nil
}
