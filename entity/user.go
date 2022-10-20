package entity

import "github.com/go-pg/pg/v10"

// User : structure for user
type User struct {
	tableName struct{} `sql:"co_user"`
	ID        int      `json:"id,omitempty" xml:"id,attr" sql:"user_id,pk"`
	Username  string   `json:"username,omitempty" xml:"username,attr" sql:"user_username"`
	Password  string   `json:"password,omitempty" xml:"password,attr" sql:"user_password"`
	ACL       int      `json:"acl,omitempty" xml:"acl,attr" sql:"user_acl"`
}

// UserRepositoryFactory : structure knowing all parameters to create the repository.
type UserRepositoryFactory struct {
	Database *pg.DB
}

// Build : Create the repository
func (f *UserRepositoryFactory) Build() *UserRepository {
	return &UserRepository{
		Database: f.Database,
		Column:   "user.*",
	}
}

// UserRepository : structure knowing all parameters to perform exchange with database for user type
type UserRepository struct {
	Database *pg.DB
	Column   string
}

// GetUSer : Get a user by a given ID
func (r *UserRepository) GetUser(id int) (*User, error) {
	user := User{ID: id}
	err := r.Database.Model(&user).WherePK().Select()
	if err != nil {
		return &User{}, nil
	}

	return &user, nil
}

// PostUser : Post a user
func (r *UserRepository) PostUser(user User) (int, error) {
	_, err := r.Database.Model(&user).Insert()

	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

// UpdateUser : Update a user
func (r *UserRepository) UpdateUser(user User) (int, error) {
	_, err := r.Database.Model(&user).Update()

	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

// DeleteUser : Delete a user by a given id
func (r *UserRepository) DeleteUser(id int) error {
	user := User{ID: id}
	_, err := r.Database.Model(&user).Delete()

	if err != nil {
		return err
	}

	return nil
}
