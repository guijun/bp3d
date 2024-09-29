package api

import (
	"bufio"
	"encoding/json"

	"github.com/gedex/bp3d/bp3d"
	"github.com/gedex/bp3d/users"
	"github.com/gedex/bp3d/vars"
	"github.com/gofiber/fiber/v3"
)

func IfHttpApi2(app *fiber.App) {

}

func IfHttpApi(app *fiber.App) {
	app.Get("/calc", func(c fiber.Ctx) error {
		us := (c.Locals(users.LOCAL_USESSION)).(*users.USession)

		session := vars.G_USessions.GetById(us.Id)

		var req bp3d.ReqPack
		if err := json.Unmarshal(c.Request().Body(), &req); err != nil {
			session.Packer.Pack()

			writer := bufio.NewWriterSize(nil, 1024)
			if binOut, _ := json.Marshal(session.Packer); len(binOut) > 0 {
				c.Response().WriteGzip(writer)
			}
		} else {
			return err
		}
		return nil

	})
}
