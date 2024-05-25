//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=config.yaml bbclient.json
package bbclient

import (
	"context"
	b64 "encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"

	config2 "cli/app/lib/config"
)

var BbClient *ClientWithResponses

func init() {
	var err error
	hc := http.Client{}
	config, err := config2.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	BbClient, err = NewClientWithResponses("https://api.bitbucket.org/2.0", WithHTTPClient(&hc), WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		if config.Authorization.Username == "" || config.Authorization.Password == "" {
			log.Fatal("please run 'bbc init' to initialize your Bitbucket Cloud CLI")
		}
		req.Header.Add("Authorization", "Basic "+b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", config.Authorization.Username, config.Authorization.Password))))
		return nil
	}),
		WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			if req.Body != nil {
				b, _ := io.ReadAll(req.Body)
				fmt.Println(string(b))
			}

			return nil
		}))
	if err != nil {
		log.Fatal(err)
	}
}
