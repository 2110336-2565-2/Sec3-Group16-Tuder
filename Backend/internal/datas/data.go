package datas

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
)

func TestData(client *ent.Client) {

	ctx := context.Background()

	client.Tutor.Delete().Exec(ctx)
	client.Student.Delete().Exec(ctx)
	client.User.Delete().Exec(ctx)
	client.Course.Delete().Exec(ctx)

	ps, _ := utils.HashPassword("brightHeemen")
	user1, err := client.User.Create().
		SetUsername("hee").
		SetPassword(ps).
		SetAddress("a").
		SetEmail("a").
		SetPhone("0").
		SetFirstName("Bright").
		SetLastName("Jukjeejid").
		SetGender("male").
		SetBirthDate(time.Now()).
		SetProfilePictureURL("profile url").
		SetRole(user.RoleStudent).
		Save(ctx)

	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}

	user2, err := client.User.Create().
		SetUsername("bighee").
		SetPassword(ps).
		SetAddress("b").
		SetEmail("b").
		SetPhone("1").
		SetFirstName("Bright").
		SetLastName("Jukjeejid").
		SetGender("female").
		SetProfilePictureURL("profile url").
		SetBirthDate(time.Now()).
		SetRole(user.RoleTutor).
		Save(ctx)

	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}

	fmt.Println(user1)
	fmt.Println(user2)

	client.Student.Create().
		SetUserID(user1.ID).
		Save(ctx)

	tutor1, _ := client.Tutor.Create().
		SetUserID(user2.ID).
		SetCitizenID("1234567890123").
		SetOmiseBankToken("bank token").
		SetDescription("test description").
		Save(ctx)

	client.Course.Create().
		SetTutorID(tutor1.ID).
		SetTitle("sexeducation").
		SetSubject("test").
		SetTopic("topichee").
		SetEstimatedTime(10).
		SetDescription("test description").
		SetPricePerHour(500).
		SetLevel(course.LevelGrade1).
		SetCoursePictureURL("picture url").
		Save(ctx)

}
