package account

import (
	"fmt"

	api "github.com/BlackBorada/go-xdg-desktop-portal"

	"github.com/BlackBorada/go-xdg-desktop-portal/pkg/utility"
	"github.com/BlackBorada/go-xdg-desktop-portal/request"
	"github.com/godbus/dbus/v5"
)

type UserInformation struct {
	ID    string
	Name  string
	Image string
}

// GetUserInformation returns the user information for the given window.
// Options are key-value pairs.
// handle_token(s), session_handle_token(s), reason(s)
// https://flatpak.github.io/xdg-desktop-portal/docs/doc-org.freedesktop.portal.Account.html
func GetUserInformation(conn *dbus.Conn, window string, options ...map[string]dbus.Variant) (*UserInformation, error) {
	obj := conn.Object(api.ObjectName, api.ObjectPath)

	opts := utility.ParseOptions(options)

	call := obj.Call(api.AccountInterface+".GetUserInformation", 0, window, opts)
	if call.Err != nil {
		return nil, fmt.Errorf("failed to call GetUserInformation: %w", call.Err)
	}

	response := request.Request{
		Conn: conn,
		Call: call,
	}

	res, err := response.Request()
	if err != nil {
		return nil, fmt.Errorf("failed to get response: %w", err)
	}

	return parseUserInformation(res.Response)
}

func parseUserInformation(response map[string]dbus.Variant) (*UserInformation, error) {
	id, ok1 := response["id"]
	name, ok2 := response["name"]
	image, ok3 := response["image"]

	if !ok1 || !ok2 || !ok3 {
		return nil, fmt.Errorf("missing required fields in response")
	}

	return &UserInformation{
		ID:    id.Value().(string),
		Name:  name.Value().(string),
		Image: image.Value().(string),
	}, nil
}
