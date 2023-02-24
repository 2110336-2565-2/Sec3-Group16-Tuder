package datas

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
	"github.com/google/uuid"
)

func InsertUser(client *ent.Client, ctx context.Context) []*ent.User {

	// Create users
	user1 := CreateUser(client, "jacky", "P@ssw0rd", "a", "a", "0", "Jacky no love", "Jukjeejid", "female", time.Now(), "profile url", user.RoleStudent)
	user2 := CreateUser(client, "bright", "P@ssw0rd", "b", "a", "00", "BrightMenMen", "Jukjeejid", "female", time.Now(), "profile url", user.RoleStudent)
	user3 := CreateUser(client, "moo", "P@ssw0rd", "a", "a", "000", "MooMee (The Best)", "Jukjeejid", "male", time.Now(), "profile url", user.RoleTutor)
	user4 := CreateUser(client, "ballhomhom", "P@ssw0rd", "a", "a", "0000", "BALLBY JUKJEEJID", "Datastructure", "male", time.Now(), "profile url", user.RoleTutor)

	user, err := client.User.CreateBulk(user1, user2, user3, user4).Save(ctx)

	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}

	for _, u := range user {
		fmt.Println("User created: ", u.ID)
	}

	return user
}

func CreateUser(
	client *ent.Client,
	username string,
	pw string,
	address string,
	email string,
	phone string,
	firstname string,
	lastname string,
	gender string,
	birthdate time.Time,
	profilepictureurl string,
	role user.Role,
) *ent.UserCreate {

	// Gen UUID and Hash Password
	userId := uuid.New()
	pw1, _ := utils.HashPassword(pw)

	userCreate := client.User.Create().
		SetID(userId).
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
		SetRole(role)

	return userCreate
}
