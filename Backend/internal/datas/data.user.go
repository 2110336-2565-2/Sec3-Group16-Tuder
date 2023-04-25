package datas

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
	"github.com/google/uuid"
)

func InsertUser(client *ent.Client, ctx context.Context) []*ent.User {
	// Seed the random number generator with the current time
    rand.Seed(time.Now().UnixNano())

    // Create a slice to hold the birthdates
    var birthdates []time.Time

    // Generate 6 birthdates and append them to the slice
    for i := 1; i <= 6; i++ {
        // Generate a random year between 1900 and 2023
        year := rand.Intn(123) + 1900

        // Generate a random month between 1 and 12
        month := rand.Intn(12) + 1

        // Generate a random day between 1 and the maximum number of days in the month
        maxDays := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()
        day := rand.Intn(maxDays) + 1

        // Create a time.Time object from the year, month, and day
        birthdate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

        // Append the birthdate to the slice
        birthdates = append(birthdates, birthdate)
    }

	// Create users
	user1 := CreateUser(client, "student1", "P@ssw0rd", "487 ถ.ไม้แดง พัทลุง 85930", "student1@hotmail.com", "0998702095", "STUDENT1", "JUKJEEJID", "female", birthdates[0], "profile url", user.RoleStudent)
	user2 := CreateUser(client, "student2", "P@ssw0rd", "1 ซ.นกทอง อ.ห้วยขวาง อุตรดิตถ์ 34110", "student2@hotmail.com", "0872445831", "STUDENT2", "JUKJEEJID", "female", birthdates[1], "profile url", user.RoleStudent)
	user3 := CreateUser(client, "tutor1", "P@ssw0rd", "8/1 ถนนนาถะพินธุ ต.เว่อเล็ก อ.เรณูนคร ลำพูน", "tutor1@hotmail.com", "025659384", "TUTOR1", "JUKJEEJID", "male", birthdates[2], "profile url", user.RoleTutor)
	user4 := CreateUser(client, "tutor2", "P@ssw0rd", "801 ถ.พงศ์ฉบับนภา กระบี่ 85350", "tutor2@hotmail.com", "052194896", "TUTOR2", "Data", "male", birthdates[3], "profile url", user.RoleTutor)
	user5 := CreateUser(client, "admin1", "P@ssw0rd", "933/2 ถนนถนิมมาศ บ้านสุขเดือนห้า จ.สุราษฎร์ธานี 51850", "admin1@hotmail.com", "057593829", "ADMIN1", "JUKJEEJID", "male", birthdates[4], "profile url", user.RoleAdmin)
	user6 := CreateUser(client, "admin2", "P@ssw0rd", "35/123 รังสิต กรุงเทพฯ 10420", "admin2@hotmail.com", "076048766", "ADMIN2", "JUKJEEJID", "male", birthdates[5], "profile url", user.RoleAdmin)

	user, err := client.User.CreateBulk(user1, user2, user3, user4, user5, user6).Save(ctx)

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
