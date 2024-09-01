package cli

const (
	RED    = "\033[31m"
	GREEN  = "\033[32m"
	YELLOW = "\033[33m"
	BLUE   = "\033[34m"
	BOLD   = "\033[1m"
	RESET  = "\033[0m"
)

func Red(s string) string {
	return RED + s + RESET
}

func Green(s string) string {
	return GREEN + s + RESET
}

func Yellow(s string) string {
	return YELLOW + s + RESET
}

func Blue(s string) string {
	return BLUE + s + RESET
}

func Bold(s string) string {
	return BOLD + s + RESET
}
