package fetch

import (
	b64 "encoding/base64"
	"fmt"
	"net/http"

	config2 "cli/app/lib/config"
)

func FetchWithAuth(url string) (*http.Response, error) {
	config, err := config2.GetConfig()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Basic "+b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", config.Authorization.Username, config.Authorization.Password))))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
