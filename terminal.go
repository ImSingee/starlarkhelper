package starlarkhelper

const (
	bold      = "\x1b[1m"
	italic    = "\x1b[3m"
	underline = "\x1b[4m"

	blue = "\x1b[34m"
	cyan = "\x1b[36m"

	unformat = "\x1b[0m"
)

func terminalUnderlineString(s string) string {
	return underline + s + unformat
}

func terminalItalicString(s string) string {
	return italic + s + unformat
}

func terminalBoldString(s string) string {
	return bold + s + unformat
}

func terminalBlueString(s string) string {
	return blue + s + unformat
}

func terminalCyanString(s string) string {
	return cyan + s + unformat
}
