package main

import (
	"log"

	"github.com/BlackBorada/go-xdg-desktop-portal/screenshot"
	"github.com/godbus/dbus/v5"
)

func main() {

	options := map[string]dbus.Variant{
		"handle_token":         dbus.MakeVariant("test"),
		"session_handle_token": dbus.MakeVariant("test"),
		"interactive":          dbus.MakeVariant(true),
	}

	conn, err := dbus.SessionBus()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	s := screenshot.NewScreenshot(conn)

	_, err = s.GetScreenshot("1", options)
	if err != nil {
		log.Fatal(err)
	}

}
