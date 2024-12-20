package screenshot

import (
	"fmt"

	api "github.com/BlackBorada/go-xdg-desktop-portal"

	"github.com/BlackBorada/go-xdg-desktop-portal/pkg/utility"
	"github.com/BlackBorada/go-xdg-desktop-portal/request"
	"github.com/godbus/dbus/v5"
)

type Screenshot struct {
	conn *dbus.Conn
}

func NewScreenshot(conn *dbus.Conn) *Screenshot {
	return &Screenshot{
		conn: conn,
	}
}

func (s *Screenshot) GetScreenshot(window string, options ...map[string]dbus.Variant) (*Screenshot, error) {
	obj := s.conn.Object(api.ObjectName, api.ObjectPath)

	opts := utility.ParseOptions(options)

	call := obj.Call(api.ScreenshotInterface+".Screenshot", 0, window, opts)
	if call.Err != nil {
		return nil, fmt.Errorf("failed to call Screenshot: %w", call.Err)
	}

	response := request.Request{
		Conn: s.conn,
		Call: call,
	}

	res, err := response.Request()
	if err != nil {
		return nil, fmt.Errorf("failed to get response: %w", err)
	}

	fmt.Println(res)

	//TODO: response parsing
	return nil, nil
}
