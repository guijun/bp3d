package api

import (
	"github.com/nats-io/nats.go"
)

func IfNatsApi(app *nats.Conn) {
	app.Subscribe("", func(msg *nats.Msg) {

	})
}
