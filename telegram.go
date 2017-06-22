package caddy_telegram

import (
	"github.com/mholt/caddy"
	"github.com/mholt/caddy/caddyhttp/httpserver"
	"net/http"
	"fmt"
	"time"
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
	w.Header().Set("Test", time.Now().Format(time.RFC1123))
	w.Write([]byte("test"))

	return h.Next.ServeHTTP(w, r)
}

func setup(c *caddy.Controller) error {
	cnf := httpserver.GetConfig(c)
	for c.Next() {
		if !c.NextArg() {       // expect at least one value
			return c.ArgErr()   // otherwise it's an error
		}
		value := c.Val()        // use the value
		fmt.Println(value)
	}
	mid := func(next httpserver.Handler) httpserver.Handler {
		return &TelegramHandler{
			Next: next,
		}
	}

	cnf.AddMiddleware(mid)
	return nil
}