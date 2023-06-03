package dtos

type CategoryPagingResDTO struct {
	Page      int8             `bson:"page,omitempty" json:"page"`
	Size      int8             `bson:"size,omitempty" json:"size"`
	NbPages   int8             `bson:"nb_pages,omitempty" json:"nb_pages"`
	NbResults int16            `bson:"nb_results,omitempty" json:"nb_results"`
	Data      []CategoryResDTO `bson:"data,omitempty" json:"data"`
}
