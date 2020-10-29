package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/twitchtv/twirp"

	"github.com/ezhdanovskiy/twirp-experiments/rpc/haberdasher"
)

func main() {
	client := haberdasher.NewHaberdasherJSONClient("http://localhost:8080", &http.Client{})

	var (
		hat *haberdasher.Hat
		err error
	)
	for i := 0; i < 5; i++ {
		hat, err = client.MakeHat(context.Background(), &haberdasher.Size{Inches: -1})
		if err != nil {
			fmt.Printf("err = %#v\n", err)
			if twerr, ok := err.(twirp.Error); ok {
				if twerr.Meta("retryable") != "" {
					// Log the error and go again.
					log.Printf("got error %q, retrying", twerr)
					continue
				}
			}
			// This was some fatal error!
			log.Fatal(err)
		}
	}
	fmt.Printf("%+v", hat)
}
