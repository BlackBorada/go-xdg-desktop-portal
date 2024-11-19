package main

import (
	"fmt"
	"log"

	"github.com/BlackBorada/go-xdg-desktop-portal/account"
	"github.com/godbus/dbus/v5"
)

func main() {
	conn, err := dbus.SessionBus()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	options := map[string]dbus.Variant{
		"handle_token":         dbus.MakeVariant("test"),
		"session_handle_token": dbus.MakeVariant("test"),
		"reason":               dbus.MakeVariant("test"),
	}

	userInfo, err := account.GetUserInformation(conn, "", options)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(userInfo)
}
