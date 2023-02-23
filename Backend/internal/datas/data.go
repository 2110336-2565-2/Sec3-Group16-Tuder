package datas

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
	"github.com/google/uuid"
)

func InsertData(client *ent.Client) {

	ctx := context.Background()

	client.Tutor.Delete().Exec(ctx)
	client.Student.Delete().Exec(ctx)
	client.User.Delete().Exec(ctx)
	client.Course.Delete().Exec(ctx)

	ps, _ := utils.HashPassword("brightHeemen")
	user1Id := uuid.New()
	user2Id := uuid.New()

	user1, err := client.User.Query().Where(user.Username("bighee")).Only(ctx)
	if err != nil {
		user1, err = client.User.Create().
			SetID(user1Id).
			SetUsername("bighee").
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

	}

	client.Student.Create().
		SetUserID(user1.ID).
		Save(ctx)
	fmt.Println(user1)

	user2, err := client.User.Query().Where(user.Username("hee")).Only(ctx)

	if err != nil {

		user2, err = client.User.Create().
			SetID(user2Id).
			SetUsername("hee").
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

	}

	fmt.Println(user2)

	tutorId := uuid.New()

	tutor1, err := client.Tutor.Query().Where(tutor.CitizenIDEQ("1234567890123")).Only(ctx)

	if err != nil {

		tutor1, err = client.Tutor.Create().
			SetID(tutorId).
			SetUserID(user2.ID).
			SetCitizenID("1234567890123").
			SetOmiseBankToken("bank token").
			SetDescription("test description").
			Save(ctx)
		if err != nil {
			log.Fatalf("failed creating tutor: %v", err)
		}
	}

	// update tutor user id
	client.Tutor.UpdateOne(tutor1).SetUserID(user2.ID).Exec(ctx)

	course1Id := uuid.New()
	course2Id := uuid.New()

	// course 1
	client.Course.Create().
		SetID(course1Id).
		SetTutorID(tutor1.ID).
		SetTitle("sexeducation").
		SetSubject("test").
		SetTopic("topichee").
		SetEstimatedTime(10).
		SetDescription("test description").
		SetPricePerHour(500).
		SetLevel(course.LevelGrade1).
		SetCoursePictureURL("https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png").
		Save(ctx)

	// course 2
	client.Course.Create().
		SetID(course2Id).
		SetTutorID(tutor1.ID).
		SetTitle("Ye hee hom hom").
		SetSubject("sex with me").
		SetTopic("Life of big hee").
		SetEstimatedTime(10).
		SetDescription("test description").
		SetPricePerHour(510).
		SetLevel(course.LevelGrade3).
		SetCoursePictureURL("https://media.istockphoto.com/id/475737947/th/%E0%B8%A3%E0%B8%B9%E0%B8%9B%E0%B8%96%E0%B9%88%E0%B8%B2%E0%B8%A2/%E0%B8%9A%E0%B8%B1%E0%B8%93%E0%B8%91%E0%B8%B4%E0%B8%95%E0%B8%A7%E0%B8%B4%E0%B8%97%E0%B8%A2%E0%B8%B2%E0%B8%A5%E0%B8%B1%E0%B8%A2%E0%B8%8A%E0%B8%B2%E0%B8%A2%E0%B9%81%E0%B8%A5%E0%B8%B0%E0%B8%9C%E0%B8%B9%E0%B9%89%E0%B8%9B%E0%B8%81%E0%B8%84%E0%B8%A3%E0%B8%AD%E0%B8%87.jpg?s=1024x1024&w=is&k=20&c=2-vUNDj7147LiiAE_syfPsfWhkY8FGtw7lbGq_S5X3A=").
		Save(ctx)

}
