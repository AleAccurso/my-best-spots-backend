package dtos

type SpotReqCreateDTO struct {
	Admin       bool                 `bson:"admin, omitempty" json:"admin"`
	Category    string               `bson:"category, omitempty" json:"category"`
	Everyone    bool                 `bson:"everyone, omitempty" json:"everyone"`
	Location    LocationReqCreateDTO `bson:"location, omitempty" json:"location"`
	LoggedUsers bool                 `bson:"logged_users, omitempty" json:"logged_users"`
	Name        string               `bson:"name, omitempty" json:"name"`
	Note        string               `bson:"note, omitempty" json:"note"`
	Phone       string               `bson:"phone, omitempty" json:"phone"`
}
