package datas

import (
	"context"
	"fmt"
	"log"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"

	//"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/google/uuid"
)

func InsertData(client *ent.Client) {

	ctx := context.Background()

	client.Tutor.Delete().Exec(ctx)
	client.Student.Delete().Exec(ctx)
	client.User.Delete().Exec(ctx)
	client.Course.Delete().Exec(ctx)

	var user []*ent.User

	user, err := InsertUser(client)

	client.Student.Create().
		SetUserID(user[0].ID).
		Save(ctx)
	fmt.Println(user[0])

	fmt.Println(user[1])

	tutorId := uuid.New()

	tutor1, err := client.Tutor.Query().Where(tutor.CitizenIDEQ("1")).Only(ctx)
	if err != nil {

		tutor1, err = client.Tutor.Create().
			SetID(tutorId).
			SetUserID(user[1].ID).
			SetCitizenID("1").
			SetOmiseBankToken("bank token").
			SetDescription("test description").
			Save(ctx)
		if err != nil {
			log.Fatalf("failed creating tutor: %v", err)
		}
	}

	// update tutor user id
	client.Tutor.UpdateOne(tutor1).SetUserID(user[1].ID).Exec(ctx)
	//---------------------------------------------------------------------

	fmt.Println(user[2])

	tutorId2 := uuid.New()

	tutor2, err := client.Tutor.Query().Where(tutor.CitizenIDEQ("2")).Only(ctx)

	if err != nil {

		tutor2, err = client.Tutor.Create().
			SetID(tutorId2).
			SetUserID(user[2].ID).
			SetCitizenID("2").
			SetOmiseBankToken("bank token").
			SetDescription("test description").
			Save(ctx)
		if err != nil {
			log.Fatalf("failed creating tutor: %v", err)
		}
	}

	// update tutor user id
	client.Tutor.UpdateOne(tutor2).SetUserID(user[2].ID).Exec(ctx)

	//-----------------------------------------------------------------------

	fmt.Println(user[2])

	tutorId3 := uuid.New()

	tutor3, err := client.Tutor.Query().Where(tutor.CitizenIDEQ("3")).Only(ctx)

	if err != nil {

		tutor3, err = client.Tutor.Create().
			SetID(tutorId3).
			SetUserID(user[3].ID).
			SetCitizenID("3").
			SetOmiseBankToken("bank token").
			SetDescription("test description").
			Save(ctx)
		if err != nil {
			log.Fatalf("failed creating tutor: %v", err)
		}
	}

	// update tutor user id
	client.Tutor.UpdateOne(tutor3).SetUserID(user[3].ID).Exec(ctx)

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
		SetTitle("Ye hom hom").
		SetSubject("sex with me").
		SetTopic("Life of big hee").
		SetEstimatedTime(102).
		SetDescription("test description").
		SetPricePerHour(5108).
		SetLevel(course.LevelGrade4).
		SetCoursePictureURL("https://kuy2").
		Save(ctx)

}
