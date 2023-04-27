package services

import (
	"errors"
	"fmt"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	St1      = "Student1"
	st1      = "student1"
	tu1      = "tutor1"
	fname    = "Tom"
	lname    = "Kiattisevivee"
	email    = "bbbb@gmail.com"
	pass     = "P@ssw0rd"
	addr     = "999/99"
	ctn      = "0888888888"
	gen      = "male"
	bd, _    = time.Parse("01/01/2000", "01/01/2000")
	rs       = "student"
	rt       = "tutor"
	cid      = "1111111111111"
	desc     = "Hello World"
	cardname = "Tom Kiattisevivee"
	cardno   = "4242 4242 4242 4242"
	expir    = "12/2026"
	cvv      = "123"
	//======================= bad ===============
	wpass   = "P@ssw0rd2"
	spass   = "P@2s"
	lpass   = "P@ssw0rd123456789"
	xpass   = "P@ssw0rdก"
	nlpass  = "P@SSW0RD"
	nupass  = "p@ssw0rd"
	ndpass  = "P@ssword"
	simpass = "Kiattisevivee0"
	invRole = "emptystring"
)

type mockup_RepoRegister struct {
	ExistedUsername string
}

func (m mockup_RepoRegister) RegisterUser(sr *schemas.SchemaRegister) (*ent.User, error) {
	if sr.Username == m.ExistedUsername {
		return nil, errors.New("the username has already been used")
	}
	if sr.Role != "student" && sr.Role != "tutor" {
		err := fmt.Errorf("User Role is invalid : %s", sr.Role)
		return nil, fmt.Errorf("error creating role-specific user data: %w", err)
	}
	return nil, nil
}

// test struct
type results struct {
	Response *schemas.SchemaRegisterResponse
	Error    error
}

type testcase struct {
	Description string
	Input       *schemas.SchemaRegister
	Expected    results
}

var (
	mock_Repo = mockup_RepoRegister{
		ExistedUsername: St1,
	}
	register_service = NewServiceRegister(mock_Repo)
)
var testcases = []testcase{
	//TC1
	{
		Description: "TC1-1 used username",
		Input: &schemas.SchemaRegister{
			Username:        St1,
			Password:        pass,
			Email:           email,
			ConfirmPassword: pass,
			Firstname:       fname,
			Lastname:        lname,
			Address:         addr,
			Phone:           ctn,
			Birthdate:       bd,
			Gender:          gen,
			Role:            rs,
		},
		Expected: results{
			Response: nil,
			Error:    errors.New("the username has already been used"),
		},
	},
	{
		Description: "TC1-2 confirm password invalid",
		Input: &schemas.SchemaRegister{
			Username:        st1,
			Password:        pass,
			Email:           email,
			ConfirmPassword: wpass,
			Firstname:       fname,
			Lastname:        lname,
			Address:         addr,
			Phone:           ctn,
			Birthdate:       bd,
			Gender:          gen,
			Role:            rs,
		},
		Expected: results{
			Response: nil,
			Error:    errors.New("the password isn't match"),
		},
	},
	{
		Description: "TC1-3 short password",
		Input: &schemas.SchemaRegister{
			Username:        st1,
			Password:        spass,
			Email:           email,
			ConfirmPassword: spass,
			Firstname:       fname,
			Lastname:        lname,
			Address:         addr,
			Phone:           ctn,
			Birthdate:       bd,
			Gender:          gen,
			Role:            rs,
		},
		Expected: results{
			Response: nil,
			Error:    errors.New("Password length must be not lower that 8 chars"),
		},
	},
	{
		Description: "TC1-4 long password",
		Input: &schemas.SchemaRegister{
			Username:        st1,
			Password:        lpass,
			Email:           email,
			ConfirmPassword: lpass,
			Firstname:       fname,
			Lastname:        lname,
			Address:         addr,
			Phone:           ctn,
			Birthdate:       bd,
			Gender:          gen,
			Role:            rs,
		},
		Expected: results{
			Response: nil,
			Error:    errors.New("Password length must be not greater that 16 chars"),
		},
	},
	{
		Description: "TC1-5 invalid char password",
		Input: &schemas.SchemaRegister{
			Username:        st1,
			Password:        xpass,
			Email:           email,
			ConfirmPassword: xpass,
			Firstname:       fname,
			Lastname:        lname,
			Address:         addr,
			Phone:           ctn,
			Birthdate:       bd,
			Gender:          gen,
			Role:            rs,
		},
		Expected: results{
			Response: nil,
			Error:    errors.New("the password must contains only abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890(~!@#$%^&*_-+=`|\\(){}[]:;\"'<>,.?/)"),
		},
	},
	{
		Description: "TC1-6 no lowercase",
		Input: &schemas.SchemaRegister{
			Username:        st1,
			Password:        nlpass,
			Email:           email,
			ConfirmPassword: nlpass,
			Firstname:       fname,
			Lastname:        lname,
			Address:         addr,
			Phone:           ctn,
			Birthdate:       bd,
			Gender:          gen,
			Role:            rs,
		},
		Expected: results{
			Response: nil,
			Error:    errors.New("Password must contains at least 1 chars from abcdefghijklmnopqrstuvwxyz"),
		},
	},
	{
		Description: "TC1-7 no uppercase",
		Input: &schemas.SchemaRegister{
			Username:        st1,
			Password:        nupass,
			Email:           email,
			ConfirmPassword: nupass,
			Firstname:       fname,
			Lastname:        lname,
			Address:         addr,
			Phone:           ctn,
			Birthdate:       bd,
			Gender:          gen,
			Role:            rs,
		},
		Expected: results{
			Response: nil,
			Error:    errors.New("Password must contains at least 1 chars from ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
		},
	},
	{
		Description: "TC1-8 no digit",
		Input: &schemas.SchemaRegister{
			Username:        st1,
			Password:        ndpass,
			Email:           email,
			ConfirmPassword: ndpass,
			Firstname:       fname,
			Lastname:        lname,
			Address:         addr,
			Phone:           ctn,
			Birthdate:       bd,
			Gender:          gen,
			Role:            rs,
		},
		Expected: results{
			Response: nil,
			Error:    errors.New("Password must contains at least 1 chars from 1234567890"),
		},
	},
	{
		Description: "TC1-9 too similar to an attribute [lastname]",
		Input: &schemas.SchemaRegister{
			Username:        st1,
			Password:        simpass,
			Email:           email,
			ConfirmPassword: simpass,
			Firstname:       fname,
			Lastname:        lname,
			Address:         addr,
			Phone:           ctn,
			Birthdate:       bd,
			Gender:          gen,
			Role:            rs,
		},
		Expected: results{
			Response: nil,
			Error:    errors.New(fmt.Sprintf("The password is too similar to the %s", lname)),
		},
	},
	{
		Description: "TC1-10 invalid user role",
		Input: &schemas.SchemaRegister{
			Username:        st1,
			Password:        pass,
			Email:           email,
			ConfirmPassword: pass,
			Firstname:       fname,
			Lastname:        lname,
			Address:         addr,
			Phone:           ctn,
			Birthdate:       bd,
			Gender:          gen,
			Role:            invRole,
		},
		Expected: results{
			Response: nil,
			Error:    fmt.Errorf("error creating role-specific user data: %w", fmt.Errorf("User Role is invalid : %s", invRole)),
		},
	},
	{
		Description: "TC1-11 valid student",
		Input: &schemas.SchemaRegister{
			Username:        st1,
			Password:        pass,
			Email:           email,
			ConfirmPassword: pass,
			Firstname:       fname,
			Lastname:        lname,
			Address:         addr,
			Phone:           ctn,
			Birthdate:       bd,
			Gender:          gen,
			Role:            rs,
		},
		Expected: results{
			Response: &schemas.SchemaRegisterResponse{
				Success: true,
				Message: "Register Success",
				Error:   nil,
			},
			Error: nil,
		},
	},
	{
		Description: "TC1-11 valid student",
		Input: &schemas.SchemaRegister{
			Username:        st1,
			Password:        pass,
			Email:           email,
			ConfirmPassword: pass,
			Firstname:       fname,
			Lastname:        lname,
			Address:         addr,
			Phone:           ctn,
			Birthdate:       bd,
			Gender:          gen,
			Role:            rt,
			CitizenID:       cid,
			Description:     desc,
			OmiseBankToken:  cardno,
		},
		Expected: results{
			Response: &schemas.SchemaRegisterResponse{
				Success: true,
				Message: "Register Success",
				Error:   nil,
			},
			Error: nil,
		},
	},
}

func TestServiceRegister_RegisterService(t *testing.T) {
	for _, tc := range testcases {
		res, err := register_service.RegisterService(tc.Input)
		Aresults := results{
			Response: res,
			Error:    err,
		}
		assert.Equal(t, tc.Expected, Aresults, fmt.Sprintf("Failed Test Case: %s", tc.Description))
	}
}
