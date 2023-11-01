package dtos

type RegionResDTO struct {
	Name string `bson:"name,omitempty" json:"name"`
	Key  string `bson:"key,omitempty" json:"key"`
}

type RegionListResDTO struct {
	Regions []RegionResDTO `bson:"regions,omitempty" json:"regions"`
}
