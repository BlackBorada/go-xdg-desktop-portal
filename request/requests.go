package request

import (
	"fmt"

	api "github.com/BlackBorada/go-xdg-desktop-portal"
	"github.com/godbus/dbus/v5"
)

type Request struct {
	Conn *dbus.Conn
	Call *dbus.Call
}

type Response struct {
	Response map[string]dbus.Variant
}

func (r *Request) Request() (*Response, error) {
	var responsePath dbus.ObjectPath

	err := r.Call.Store(&responsePath)
	if err != nil {
		return nil, err
	}

	err = r.Conn.AddMatchSignal(
		dbus.WithMatchObjectPath(responsePath),
		dbus.WithMatchInterface(api.RequestInterface),
		dbus.WithMatchMember(api.ResponseMember),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to add signal match: %w", err)
	}

	signals := make(chan *dbus.Signal, 1)
	r.Conn.Signal(signals)
	defer r.Conn.RemoveSignal(signals)

	signal := <-signals

	if signal == nil {
		return nil, fmt.Errorf("failed to receive signal")
	}

	if len(signal.Body) < 2 {
		return nil, fmt.Errorf("invalid response format")
	}

	responseCode, ok := signal.Body[0].(uint32)
	if !ok {
		return nil, fmt.Errorf("invalid response code type")
	}

	// 0 = Success, 1 = Cancelled, 2 = Other
	if responseCode != 0 {
		return nil, fmt.Errorf("portal request failed with code: %d", responseCode)
	}

	// Получаем результат
	result, ok := signal.Body[1].(map[string]dbus.Variant)
	if !ok {
		return nil, fmt.Errorf("invalid result format")
	}

	return &Response{Response: result}, nil
}
