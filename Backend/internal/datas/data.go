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
	user3Id := uuid.New()
	user4Id := uuid.New()

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

	user2, err := client.User.Query().Where(user.Username("Hee")).Only(ctx)

	if err != nil {

		user2, err = client.User.Create().
			SetID(user2Id).
			SetUsername("Hee").
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

	tutor1, err := client.Tutor.Query().Where(tutor.CitizenIDEQ("1")).Only(ctx)
	if err != nil {

		tutor1, err = client.Tutor.Create().
			SetID(tutorId).
			SetUserID(user2.ID).
			SetCitizenID("1").
			SetOmiseBankToken("bank token").
			SetDescription("test description").
			Save(ctx)
		if err != nil {
			log.Fatalf("failed creating tutor: %v", err)
		}
	}

	// update tutor user id
	client.Tutor.UpdateOne(tutor1).SetUserID(user2.ID).Exec(ctx)
	//---------------------------------------------------------------------
	user3, err := client.User.Query().Where(user.Username("hee")).Only(ctx)

	if err != nil {

		user3, err = client.User.Create().
			SetID(user3Id).
			SetUsername("hee").
			SetPassword(ps).
			SetAddress("bc").
			SetEmail("bc").
			SetPhone("12").
			SetFirstName("Brightc").
			SetLastName("Jukjeejidc").
			SetGender("female").
			SetProfilePictureURL("profile url").
			SetBirthDate(time.Now()).
			SetRole(user.RoleTutor).
			Save(ctx)

		if err != nil {
			log.Fatalf("failed creating user: %v", err)
		}

	}

	fmt.Println(user3)

	tutorId2 := uuid.New()

	tutor2, err := client.Tutor.Query().Where(tutor.CitizenIDEQ("2")).Only(ctx)

	if err != nil {

		tutor2, err = client.Tutor.Create().
			SetID(tutorId2).
			SetUserID(user3.ID).
			SetCitizenID("2").
			SetOmiseBankToken("bank token").
			SetDescription("test description").
			Save(ctx)
		if err != nil {
			log.Fatalf("failed creating tutor: %v", err)
		}
	}

	// update tutor user id
	client.Tutor.UpdateOne(tutor2).SetUserID(user3.ID).Exec(ctx)

	//-----------------------------------------------------------------------
	user4, err := client.User.Query().Where(user.Username("pongkul")).Only(ctx)

	if err != nil {

		user4, err = client.User.Create().
			SetID(user4Id).
			SetUsername("pongkul").
			SetPassword(ps).
			SetAddress("bd").
			SetEmail("bd").
			SetPhone("1d").
			SetFirstName("Brightd").
			SetLastName("Jukjeejidd").
			SetGender("female").
			SetProfilePictureURL("profile url").
			SetBirthDate(time.Now()).
			SetRole(user.RoleTutor).
			Save(ctx)

		if err != nil {
			log.Fatalf("failed creating user: %v", err)
		}

	}

	fmt.Println(user3)

	tutorId3 := uuid.New()

	tutor3, err := client.Tutor.Query().Where(tutor.CitizenIDEQ("3")).Only(ctx)

	if err != nil {

		tutor3, err = client.Tutor.Create().
			SetID(tutorId3).
			SetUserID(user4.ID).
			SetCitizenID("3").
			SetOmiseBankToken("bank token").
			SetDescription("test description").
			Save(ctx)
		if err != nil {
			log.Fatalf("failed creating tutor: %v", err)
		}
	}

	// update tutor user id
	client.Tutor.UpdateOne(tutor3).SetUserID(user4.ID).Exec(ctx)

	//----------------------------------------------------------------
	course1Id := uuid.New()
	course2Id := uuid.New()
	course3Id := uuid.New()

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
		SetTutorID(tutor2.ID).
		SetTitle("Ye hee hom").
		SetSubject("sex with me").
		SetTopic("Life of big hee").
		SetEstimatedTime(106).
		SetDescription("test description").
		SetPricePerHour(5101).
		SetLevel(course.LevelGrade3).
		SetCoursePictureURL("https:kuy").
		Save(ctx)

	// course 2
	client.Course.Create().
		SetID(course3Id).
		SetTutorID(tutor3.ID).
		SetTitle("Ye  hom hom").
		SetSubject("sex with me").
		SetTopic("Life of big hee").
		SetEstimatedTime(102).
		SetDescription("test description").
		SetPricePerHour(5108).
		SetLevel(course.LevelGrade4).
		SetCoursePictureURL("https://kuy2").
		Save(ctx)

}
