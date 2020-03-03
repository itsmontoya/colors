package colors

const (
	escapeChar  = "\x1b"
	colorPrefix = "[%sm%s"
	colorSuffix = "[%dm"
	colorFormat = escapeChar + colorPrefix + escapeChar + colorSuffix
)
