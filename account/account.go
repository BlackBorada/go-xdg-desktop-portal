package account

import (
	"fmt"

	api "github.com/BlackBorada/go-xdg-desktop-portal"
	"github.com/BlackBorada/go-xdg-desktop-portal/request"
	"github.com/godbus/dbus/v5"
)

// AccountInterface определяет методы для работы с аккаунтом
// type AccountInterface interface {
// 	GetUserInformation(conn *dbus.Conn, window string, options ...map[string]dbus.Variant) (*UserInformation, error)
// }

// UserInformation содержит информацию о пользователе
type UserInformation struct {
	ID    string
	Name  string
	Image string
}

func GetUserInformation(conn *dbus.Conn, options ...map[string]dbus.Variant) (*UserInformation, error) {
	obj := conn.Object(api.ObjectName, api.ObjectPath)

	optionsV := make(map[string]dbus.Variant)
	for _, option := range options {
		for k, v := range option {
			optionsV[k] = v
		}
	}

	call := obj.Call(api.AccountInterface+".GetUserInformation", 0, "D-Bus", map[string]dbus.Variant{})
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

// parseUserInformation извлекает данные пользователя из ответа D-Bus
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
