//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=config.yaml bbclient.json
package bbclient

import (
	"context"
	b64 "encoding/base64"
	"log"
	"net/http"
)

var BbClient *ClientWithResponses

func init() {
	var err error
	hc := http.Client{}
	BbClient, err = NewClientWithResponses("https://api.bitbucket.org/2.0", WithHTTPClient(&hc), WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Add("Authorization", "Basic "+b64.StdEncoding.EncodeToString([]byte("user:password")))
		return nil
	}))
	if err != nil {
		log.Fatal(err)
	}

}
