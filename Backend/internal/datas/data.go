package datas
import (
	"context"
	"fmt"
	"log"
	"time"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"

)


func TestData(client *ent.Client){

	ctx := context.Background()
	
	client.Tutor.Delete().Exec(ctx)
	client.Student.Delete().Exec(ctx)
	client.User.Delete().Exec(ctx)
	
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
	Save(context.Background())
	
	client.Tutor.Create().
	SetUserID(user2.ID).
	SetCitizenID("1234567890123").
	Save(context.Background())
		
	}