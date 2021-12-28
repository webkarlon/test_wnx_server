package main

import (
	"git.jetbrains.space/yabbi/dsp/components/extrarpc"
	"github.com/nexcode/wenex"
	"os"
	"path"
	"test_wnx_server/api"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	configPath := path.Base(wd) + "/bidder"

	wnx, err := wenex.New(configPath, nil, nil)
	if err != nil {
		panic(err)
	}

	rpcTestServer := api.NewTestServer(wnx)

	rpcServer := extrarpc.NewServer(wnx.Config.MustString("RPC-SERVER"))

	rpcServer.Register(rpcTestServer)

	go func() {
		panic(rpcServer.Run())
	}()

	if err = wnx.Run(); err != nil {
		panic(err)
	}
}
