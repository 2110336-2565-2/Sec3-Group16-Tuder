package datas

import (
	"context"
	"log"
	"time"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
	"github.com/google/uuid"
)

func InsertUser(client *ent.Client) ([]*ent.User, error) {
	user1, err := genUser(client, "password1", "bighee1", "a", "a", "0", "Bright", "Jukjeejid", "male", time.Now(), "profile url", user.RoleStudent)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}

	user2, err := genUser(client, "password2", "bighee2", "a", "a", "0", "Bright", "Jukjeejid", "male", time.Now(), "profile url", user.RoleStudent)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}

	user3, err := genUser(client, "password3", "bighee3", "a", "a", "0", "Bright", "Jukjeejid", "male", time.Now(), "profile url", user.RoleStudent)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}
	user4, err := genUser(client, "password4", "bighee4", "a", "a", "0", "Bright", "Jukjeejid", "male", time.Now(), "profile url", user.RoleStudent)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}
	return []*ent.User{user1, user2, user3, user4}, err
}

func genUser(client *ent.Client,
	pw string,
	username string,
	address string,
	email string,
	phone string,
	firstname string,
	lastname string,
	gender string,
	birthdate time.Time,
	profilepictureurl string,
	role user.Role,
) (*ent.User, error) {

	ctx := context.Background()

	user1Id := uuid.New()

	pw1, _ := utils.HashPassword(pw)

	user1, err := client.User.Query().Where(user.Username(username)).Only(ctx)
	if err != nil {
		user1, err = client.User.Create().
			SetID(user1Id).
			SetUsername(username).
			SetPassword(pw1).
			SetAddress(address).
			SetEmail(email).
			SetPhone(phone).
			SetFirstName(firstname).
			SetLastName(lastname).
			SetGender(gender).
			SetBirthDate(birthdate).
			SetProfilePictureURL(profilepictureurl).
			SetRole(role).
			Save(ctx)
		if err != nil {
			log.Fatalf("failed creating user: %v", err)
		}
	}
	return user1, err
}

// user5, err := InsertUser(client, "password5", "bighee5", "a", "a", "0", "Bright", "Jukjeejid", "male", time.Now(), "profile url", user.RoleStudent)
// if err != nil {
//     log.Fatalf("failed creating user: %v", err)
// }

// user6, err := InsertUser(client, "password6", "bighee6", "a", "a", "0", "Bright", "Jukjeejid", "male", time.Now(), "profile url", user.RoleStudent)
// if err != nil {
//     log.Fatalf("failed creating user: %v", err)
// }

// user7, err := InsertUser(client, "password7", "bighee7", "a", "a", "0", "Bright", "Jukjeejid", "male", time.Now(), "profile url", user.RoleStudent)
// if err != nil {
//     log.Fatalf("failed creating user: %v", err)
// }

// user8, err := InsertUser(client, "password8", "bighee8", "a", "a", "0", "Bright", "Jukjeejid", "male", time.Now(), "profile url", user.RoleStudent)
// if err != nil {
//     log.Fatalf("failed creating user: %v", err)
// }

// user9, err := InsertUser(client, "password9", "bighee9", "a", "a", "0", "Bright", "Jukjeejid", "male", time.Now(), "profile url", user.RoleStudent)
// if err != nil {
//     log.Fatalf("failed creating user: %v", err)
// }

// user10, err := InsertUser(client, "password10", "bighee10", "a", "a", "0", "Bright", "Jukjeejid", "male", time.Now(), "profile url", user.RoleStudent)
// if err != nil {
//     log.Fatalf("failed creating user: %v", err)
// }
