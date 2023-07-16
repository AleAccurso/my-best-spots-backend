package enums

import "strings"

type SpotAccessRight string

const (
	ADMIN        SpotAccessRight = "ADMIN"
	EVERYONE     SpotAccessRight = "EVERYONE"
	LOGGED_USERS SpotAccessRight = "LOGGED_USERS"
)

var (
    spotAccessRightsMap = map[string]SpotAccessRight{
        "ADMIN":   ADMIN,
        "EVERYONE": EVERYONE,
        "LOGGED_USERS": LOGGED_USERS,
    }
)

func ParseToSpotAccessRight(str string) (SpotAccessRight, bool) {
	spotAccessRight, ok := spotAccessRightsMap[strings.ToUpper(str)]
    return spotAccessRight, ok
}
