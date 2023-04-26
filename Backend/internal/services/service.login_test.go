package services

import (
	"errors"
	"fmt"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockup_RepoLogin struct {
	ExistedStudent string
	Password       string
	ExistedTutor   string
}

func (m mockup_RepoLogin) Login(l *schemas.SchemaLogin) (*ent.User, error) {
	uuuid, _ := uuid.NewUUID()
	if l.Username == m.ExistedStudent {
		role := user.RoleStudent
		suuid, _ := uuid.NewUUID()
		return &ent.User{
			Username: l.Username,
			ID:       uuuid,
			Password: hashed,
			Role:     &role,
			Edges: ent.UserEdges{
				Student: &ent.Student{
					ID:    suuid,
					Edges: ent.StudentEdges{},
				},
			},
		}, nil
	}

	if l.Username == m.ExistedTutor {
		role := user.RoleTutor
		tuuid, _ := uuid.NewUUID()
		return &ent.User{
			Username: l.Username,
			ID:       uuuid,
			Password: hashed,
			Role:     &role,
			Edges: ent.UserEdges{
				Tutor: &ent.Tutor{
					ID:    tuuid,
					Edges: ent.TutorEdges{},
				},
			},
		}, nil
	}
	return nil, errors.New("the username isn't match")
}

var (
	hashed, _      = utils.HashPassword(pass)
	mock_RepoLogin = mockup_RepoLogin{
		ExistedStudent: st1,
		Password:       pass,
		ExistedTutor:   tu1,
	}
	login_service = NewServiceLogin(mock_RepoLogin)
)

// test struct
type login_results struct {
	Response *schemas.SchemaLoginResponses
	Error    error
}

type login_testcase struct {
	Description string
	Input       *schemas.SchemaLogin
	Expected    login_results
}

var ltestcases = []login_testcase{
	//TC1
	{
		Description: "TC1 username doesn't exist",
		Input: &schemas.SchemaLogin{
			Username: St1,
			Password: pass,
		},
		Expected: login_results{
			Response: nil,
			Error:    errors.New("the username isn't match"),
		},
	},
	{
		Description: "TC2 password not correct",
		Input: &schemas.SchemaLogin{
			Username: st1,
			Password: wpass,
		},
		Expected: login_results{
			Response: nil,
			Error:    errors.New("the password isn't match"),
		},
	},
}

func TestServiceLogin_LoginService(t *testing.T) {
	for _, tc := range ltestcases {
		res, err := login_service.LoginService(tc.Input)
		result := login_results{
			Response: res,
			Error:    err,
		}
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed Test Case: %s", tc.Description))
	}
}

func TestStudentLoginSuccess(t *testing.T) {
	Description := "TC3 student login successful"
	ls := &schemas.SchemaLogin{
		Username: st1,
		Password: pass,
	}
	res, err := login_service.LoginService(ls)
	result := login_results{
		Response: res,
		Error:    err,
	}

	expected := login_results{
		Response: &schemas.SchemaLoginResponses{
			Username: st1,
			Token:    result.Response.Token,
			Role:     "student",
		},
		Error: nil,
	}
	assert.Equal(t, expected, result, fmt.Sprintf("Failed Test Case: %s", Description))
}

func TestTutorLoginSuccess(t *testing.T) {
	Description := "TC4 Tutor login successful"
	ls := &schemas.SchemaLogin{
		Username: tu1,
		Password: pass,
	}
	res, err := login_service.LoginService(ls)
	result := login_results{
		Response: res,
		Error:    err,
	}

	expected := login_results{
		Response: &schemas.SchemaLoginResponses{
			Username: tu1,
			Token:    result.Response.Token,
			Role:     "tutor",
		},
		Error: nil,
	}
	assert.Equal(t, expected, result, fmt.Sprintf("Failed Test Case: %s", Description))
}
