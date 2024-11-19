package utility

import "github.com/godbus/dbus/v5"

func ParseOptions(options []map[string]dbus.Variant) map[string]dbus.Variant {
	opts := make(map[string]dbus.Variant)
	for _, option := range options {
		for k, v := range option {
			opts[k] = v
		}
	}
	return opts
}
