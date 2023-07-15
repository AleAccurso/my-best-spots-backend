package enums

type SpotAccessRight string

const (
	ADMIN        SpotAccessRight = "ADMIN"
	EVERYONE     SpotAccessRight = "EVERYONE"
	LOGGED_USERS SpotAccessRight = "LOGGED_USERS"
)

var SpotAccessRights = []SpotAccessRight{
	ADMIN,
	EVERYONE,
	LOGGED_USERS,
}
