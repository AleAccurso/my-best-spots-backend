package helpers

// func IsLoggedCategoryOrAdmin(c *gin.Context, CategoryId primitive.ObjectID, Category models.Category) error {
// 	loggedCategoryEmail, ok := c.Get("Category_email")
// 	if !ok {
// 		return errors.New(constants.AUTH_UNVERIFIED_EMAIL)
// 	}

// 	loggedCategoryRole, ok := c.Get("Category_role")
// 	if !ok {
// 		return errors.New("cannot get logged Category role")
// 	}

// 	if Category.Email != loggedCategoryEmail && loggedCategoryRole != "admin" {
// 		return errors.New(constants.AUTH_UNAUTHORIZED)
// 	}

// 	return nil
// }
