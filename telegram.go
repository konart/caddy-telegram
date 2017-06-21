package caddy_telegram

import (
	"github.com/mholt/caddy"
	"github.com/mholt/caddy/caddyhttp/httpserver"
	"net/http"
	"fmt"
)

func init()  {
	caddy.RegisterPlugin("telegram", caddy.Plugin{
		ServerType: "http",
		Action: setup,
	})
}

type TelegramHandler struct {
	Next httpserver.Handler
}

func (h TelegramHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	return h.Next.ServeHTTP(w, r)
}

func setup(c *caddy.Controller) error {
	fmt.Println("test")
	return nil
}