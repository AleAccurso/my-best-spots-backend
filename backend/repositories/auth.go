package repositories

import (
	"gorm.io/gorm"
)

type AuthRepository struct {
	database *gorm.DB
}

func InitialiseAuthRepository(db *gorm.DB) AuthRepository {
	return AuthRepository{
		database: db,
	}
}

// func (repository AuthRepository) AddCategory(context context.Context, Category models.Category) (*primitive.ObjectID, error) {

// 	now := time.Now().UTC()
// 	Category.CreatedAt, Category.UpdatedAt = now, now

// 	result, err := repository.database.Categories.InsertOne(context, Category)
// 	if err != nil {
// 		return nil, errors.New("reposiotry/unable-to-register")
// 	}

// 	newID, ok := result.InsertedID.(primitive.ObjectID)
// 	if !ok {
// 		return nil, errors.New("reposiotry/unable-to-convert-id")
// 	}

// 	return &newID, nil
// }
