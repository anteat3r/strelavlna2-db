package main

import (
	"log"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"

	// "github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

const (
  ADMINS = "admins"
  CONFIG = "config"
  CONTESTS = "contests"
  PLAYERS = "players"
  PROBS = "probs"
  TEAMS = "teams"
  
  MATH = "Math"
  PHYSICS = "Physics"

  EASY = "Easy"
  MEDIUM = "Medium"
  HARD = "Hard"
)

var (
  app *pocketbase.PocketBase
  cache = make(map[string]any)
)

func main() {
  app = pocketbase.New()
  app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
    e.Router.GET("/ping", ping)
    e.Router.GET("/config/:key", config)
    e.Router.GET("/isadmin/:name", isAdmin)
    return nil
  })
  if err := app.Start(); err != nil {
    log.Fatal(err)
  }
}

func ping(c echo.Context) error {
  rec, err := app.Dao().FindFirstRecordByData(
    CONFIG, "key", "test",
  )
  if err != nil {
    return c.String(404, "No test config data :(")
  }
  res := rec.GetString("value")
  return c.String(200, res)
}

func config(c echo.Context) error {
  rec, err := app.Dao().FindFirstRecordByData(
    CONFIG, "key", c.PathParam("key"),
  )
  if err != nil { return err }
  res := rec.GetString("value")
  return c.String(200, res)
}

func isAdmin(c echo.Context) error {
  _, err := app.Dao().FindFirstRecordByData(
    ADMINS, "username", c.PathParam("user"),
  )
  if err != nil { return c.String(200, "false") }
  return c.String(200, "true")
}
