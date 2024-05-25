package user

import (
	"context"

	"cli/app/bbclient"
)

func CurrentUser() (*bbclient.Account, error) {
	res, err := bbclient.BbClient.GetUserWithResponse(context.TODO())
	if err != nil {
		return nil, err
	}
	return res.JSON200, nil
}
