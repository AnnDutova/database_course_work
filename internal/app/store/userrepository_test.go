package store_test

import (
	"testing"

	"github.com/AnnDutova/database_course_work/internal/app/model"
	"github.com/AnnDutova/database_course_work/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")
	u, err := s.User().Create(&model.User{
		Email: "user1@example.com",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepositoryFindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)
	s.User().Create(&model.User{
		Email: email,
	})
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
