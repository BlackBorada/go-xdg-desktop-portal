package main

import (
	"log"

	screencast "github.com/BlackBorada/go-xdg-desktop-portal/screen_cast"
	"github.com/godbus/dbus/v5"
)

func main() {

	options := map[string]dbus.Variant{
		"handle_token":         dbus.MakeVariant("test"),
		"session_handle_token": dbus.MakeVariant("test"),
		"multiple":             dbus.MakeVariant(false),
		"types":                dbus.MakeVariant(screencast.SourceTypeWindow | screencast.SourceTypeMonitor),
		"cursor_mode":          dbus.MakeVariant(screencast.CursorModeEmbedded),
		"persist_mode":         dbus.MakeVariant(screencast.PersistModePermanent),
	}

	windowOptions := map[string]dbus.Variant{
		"show_details": dbus.MakeVariant(true),
		"show_labels":  dbus.MakeVariant(true),
		"allow_clear":  dbus.MakeVariant(true),
	}
	options["window_options"] = dbus.MakeVariant(windowOptions)

	conn, err := dbus.SessionBus()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	sc := screencast.NewScreencast(conn)

	sessionHandle, err := sc.CreateSession(options)
	if err != nil {
		log.Fatal(err)
	}

	err = sc.SelectSourece(sessionHandle, options)
	if err != nil {
		log.Fatal(err)
	}

	err = sc.Start(sessionHandle, options)
	if err != nil {
		log.Fatal(err)
	}

}
