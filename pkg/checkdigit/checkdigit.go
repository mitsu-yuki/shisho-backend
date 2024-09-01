package checkdigit

import "github.com/osamingo/checkdigit"

func ISBN13IsValid(s string) bool {
	return checkdigit.NewISBN13().Verify(s)
}
