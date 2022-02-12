package twitch

import (
	"github.com/nicklaw5/helix/v2"
	"github.com/sirupsen/logrus"
)

type ApiClient interface {
	GetStreams(params *helix.StreamsParams) (*helix.StreamsResponse, error)
	GetUsers(params *helix.UsersParams) (*helix.UsersResponse, error)
}

var Log = logrus.StandardLogger()
