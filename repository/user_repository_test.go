package repository

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/app/realworld/model/user"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type userRepositorySuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	db   *gorm.DB

	rep UserRepository
}

func (s *userRepositorySuite) SetupTest() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.db, err = gorm.Open("mysql", db)
	require.NoError(s.T(), err)

	s.rep = *NewUserRepository(s.db)
}

func (s *userRepositorySuite) TestGetId() {
	user := user.User{
		ID:       1,
		Username: "Joe",
		Email:    "example@sample.com",
		Password: "password",
		Bio:      "i am ...",
	}

	result := sqlmock.NewRows(GetDBFields(user)).
		AddRow(user.ID, user.Username, user.Email, user.Password, user.Bio)

	s.mock.ExpectQuery("SELECT \\* FROM `users`").
		WillReturnRows(result).
		WillReturnError(nil)

	actual, err := s.rep.GetById(1)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), user.Username, actual.Username)
	assert.Equal(s.T(), user.Email, actual.Email)
	assert.Equal(s.T(), user.Password, actual.Password)
	assert.Equal(s.T(), user.Bio, actual.Bio)
}

func (s *userRepositorySuite) TestCreate() {
	user := user.User{
		ID:       1,
		Username: "Joe",
		Email:    "example@sample.com",
		Password: "password",
		Bio:      "i am ...",
	}

	// Insert 系のテストの書き方はここを参照。
	//https://godoc.org/github.com/DATA-DOG/go-sqlmock#NewResult
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO `users`").
		WithArgs(user.ID, user.Username, user.Email, user.Password, user.Bio).
		// Result は sqlの結果の影響を表現する構造体。
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	actual, err := s.rep.Create(user)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), user.Username, actual.Username)
	assert.Equal(s.T(), user.Email, actual.Email)
	assert.Equal(s.T(), user.Password, actual.Password)
	assert.Equal(s.T(), user.Bio, actual.Bio)
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(userRepositorySuite))
}
