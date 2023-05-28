package utils

import (
	"my-best-spots-backend/enums"

	"golang.org/x/exp/slices"
)

func IsAllowedLanguage(language string) bool {
	if !slices.Contains(enums.AllowedLanguages, language) {
		return false
	} else {
		return true
	}

}
