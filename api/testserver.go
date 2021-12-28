package api

import (
	"git.jetbrains.space/yabbi/dsp/bidder/api"
	"github.com/nexcode/wenex"
)

func NewTestServer(
	wnx *wenex.Wenex,

) *TestServer {
	return &TestServer{
		wnx: wnx,
	}
}

type TestServer struct {
	wnx *wenex.Wenex
}

func (t *TestServer) Query(call api.Call, str *string) error {
	*str = "блять"
	return nil
}
