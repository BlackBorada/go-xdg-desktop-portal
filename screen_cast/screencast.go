package screencast

import (
	"fmt"

	api "github.com/BlackBorada/go-xdg-desktop-portal"

	"github.com/BlackBorada/go-xdg-desktop-portal/pkg/utility"
	"github.com/BlackBorada/go-xdg-desktop-portal/request"
	"github.com/godbus/dbus/v5"
)

const (
	CursorModeHidden   uint32 = 0
	CursorModeEmbedded uint32 = 1
	CursorModeMetadata uint32 = 2
)

const (
	PersistModeNo        uint32 = 0
	PersistModeTemporary uint32 = 1
	PersistModePermanent uint32 = 2
)

const (
	SourceTypeMonitor uint32 = 1
	SourceTypeWindow  uint32 = 2
	SourceTypeVirtual uint32 = 4
)

type Screencast struct {
	conn *dbus.Conn
}

func NewScreencast(conn *dbus.Conn) *Screencast {
	return &Screencast{
		conn: conn,
	}
}

func (sc *Screencast) CreateSession(options ...map[string]dbus.Variant) (dbus.ObjectPath, error) {
	obj := sc.conn.Object(api.ObjectName, api.ObjectPath)

	opts := utility.ParseOptions(options)

	call := obj.Call(api.ScreenCastInterface+".CreateSession", 0, opts)
	if call.Err != nil {
		return "", fmt.Errorf("failed to call CreateSession: %w", call.Err)
	}

	response := request.Request{
		Conn: sc.conn,
		Call: call,
	}

	res, err := response.Request()
	if err != nil {
		return "", fmt.Errorf("failed to get response: %w", err)
	}

	sessionHandle, ok := res.Response["session_handle"]
	if !ok {
		return "", fmt.Errorf("missing required fields in response")
	}

	return dbus.ObjectPath(sessionHandle.Value().(string)), nil

}

func (sc *Screencast) SelectSourece(sessionHandle dbus.ObjectPath, options ...map[string]dbus.Variant) error {
	obj := sc.conn.Object(api.ObjectName, api.ObjectPath)

	opts := utility.ParseOptions(options)

	call := obj.Call(api.ScreenCastInterface+".SelectSources", 0, sessionHandle, opts)
	if call.Err != nil {
		return fmt.Errorf("failed to call SelectSources: %w", call.Err)
	}

	response := request.Request{
		Conn: sc.conn,
		Call: call,
	}

	source, err := response.Request()
	if err != nil {
		return fmt.Errorf("failed to get response: %w", err)
	}

	fmt.Println(source)
	return nil
}

func (sc *Screencast) Start(sessionHandle dbus.ObjectPath, options ...map[string]dbus.Variant) error {
	obj := sc.conn.Object(api.ObjectName, api.ObjectPath)

	opts := utility.ParseOptions(options)

	call := obj.Call(api.ScreenCastInterface+".Start", 0, sessionHandle, "/", opts)
	if call.Err != nil {
		return fmt.Errorf("failed to call Start: %w", call.Err)
	}

	response := request.Request{
		Conn: sc.conn,
		Call: call,
	}

	handle, err := response.Request()
	if err != nil {
		return fmt.Errorf("failed to get response: %w", err)
	}
	//TODO: Handle parse
	fmt.Println(handle)

	return nil
}

func (sc *Screencast) Close(sessionHandle dbus.ObjectPath) error {
	obj := sc.conn.Object(api.ObjectName, sessionHandle)

	call := obj.Call(api.ScreenCastInterface+".Close", 0)
	if call.Err != nil {
		return fmt.Errorf("failed to call Close: %w", call.Err)
	}
	return nil
}

// TODO: Implement other methods
// pipeWire
