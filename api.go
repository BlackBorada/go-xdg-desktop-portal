package api

const (
	ObjectName = "org.freedesktop.portal.Desktop"

	ObjectPath = "/org/freedesktop/portal/desktop"

	AccountInterface = "org.freedesktop.portal.Account"

	ScreenCastInterface = "org.freedesktop.portal.ScreenCast"

	RequestInterface = "org.freedesktop.portal.Request"

	SessionInterface = "org.freedesktop.portal.Session"

	ResponseMember = "Response"

	PortalImpl = "org.freedesktop.impl.portal"
)

type SessionOptions struct {
	HandleToken        string
	SessionHandleToken string
}
