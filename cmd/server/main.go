package main

import (
	"fmt"
	"net/http"

	"github.com/ezhdanovskiy/twirp-experiments/internal/butlerserver"
	"github.com/ezhdanovskiy/twirp-experiments/internal/haberdasherserver"
	"github.com/ezhdanovskiy/twirp-experiments/rpc/butler"
	"github.com/ezhdanovskiy/twirp-experiments/rpc/haberdasher"
)

func main() {
	mux := http.NewServeMux()

	haberdasherSrv := &haberdasherserver.Server{} // implements Haberdasher interface
	haberdasherSrvHandler := haberdasher.NewHaberdasherServer(haberdasherSrv)
	fmt.Printf("haberdasherSrvHandler.PathPrefix(): %#v\n", haberdasherSrvHandler.PathPrefix())
	mux.Handle(haberdasherSrvHandler.PathPrefix(), haberdasherSrvHandler)

	butlerSrv := &butlerserver.Server{}
	butlerSrvHandler := butler.NewButlerServer(butlerSrv)
	fmt.Printf("butlerSrvHandler.PathPrefix(): %#v\n", butlerSrvHandler.PathPrefix())
	mux.Handle(butlerSrvHandler.PathPrefix(), butlerSrvHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
