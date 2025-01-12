package model

import (
	"projectforge.dev/projectforge/app/lib/types"
)

func ColToString(c *Column, prefix string) string {
	return TypeToString(c.Type, prefix+c.Proper())
}

func TypeToString(t types.Type, prop string) string {
	switch t.Key() {
	case types.KeyUUID:
		return prop + ".String()"
	default:
		return prop
	}
}

func ColToViewString(c *Column, prefix string) string {
	return TypeToViewString(c.Type, prefix+c.Proper(), c.Nullable)
}

func TypeToViewString(t types.Type, prop string, nullable bool) string {
	ret := ToGoString(t, prop, false)
	switch t.Key() {
	case types.KeyDate:
		if nullable {
			return "{%%= components.DisplayTimestampDay(" + ret + ") %%}"
		}
		return "{%%= components.DisplayTimestampDay(&" + ret + ") %%}"
	case types.KeyTimestamp:
		if nullable {
			return "{%%= components.DisplayTimestamp(" + ret + ") %%}"
		}
		return "{%%= components.DisplayTimestamp(&" + ret + ") %%}"
	default:
		return "{%%s " + ret + " %%}"
	}
}
