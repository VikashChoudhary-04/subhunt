package resolver

import (
	"context"
	"net"
	"time"
)

func Resolve(subs []string) []string {
	var alive []string

	for _, s := range subs {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		ips, err := net.DefaultResolver.LookupHost(ctx, s)
		if err == nil && len(ips) > 0 {
			alive = append(alive, s)
		}
	}

	return alive
}
