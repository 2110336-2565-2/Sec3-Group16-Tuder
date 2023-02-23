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

func InsertUser(client *ent.Client) []*ent.User {

	ctx := context.Background()

	user1 := CreateUser(client, ctx, "jacky", "P@ssw0rd", "a", "a", "0", "Jacky no love", "Jukjeejid", "female", time.Now(), "profile url", user.RoleStudent)
	user2 := CreateUser(client, ctx, "bright", "P@ssw0rd", "b", "a", "00", "BrightMenMen", "Jukjeejid", "female", time.Now(), "profile url", user.RoleStudent)
	user3 := CreateUser(client, ctx, "moo", "P@ssw0rd", "a", "a", "000", "MooMee (The Best)", "Jukjeejid", "male", time.Now(), "profile url", user.RoleTutor)
	user4 := CreateUser(client, ctx, "ballhomhom", "P@ssw0rd", "a", "a", "0000", "BALLBY JUKJEEJID", "Datastructure", "male", time.Now(), "profile url", user.RoleTutor)

	return []*ent.User{user1, user2, user3, user4}
}

func CreateUser(
	client *ent.Client,
	ctx context.Context,
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
) *ent.User {

	// Gen UUID and Hash Password
	userId := uuid.New()
	pw1, _ := utils.HashPassword(pw)

	
	user, err := client.User.Create().
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
		SetRole(role).
		Save(ctx)
		
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}
	

	fmt.Println("User created: ", user.ID)

	return user
}
