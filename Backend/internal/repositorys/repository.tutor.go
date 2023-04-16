package repositorys

import (
	"context"
	"fmt"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	entCourse "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	entReview "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/review"
	entSchedule "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/schedule"
	entTutor "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"
	entUser "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
)

type RepositoryTutor interface {
	GetTutors() ([]*ent.Tutor, error)
	GetTutorByUsername(sr *schema.SchemaGetTutor) (*ent.Tutor, error)
	GetTutorScheduleByCourseId(sr *schema.SchemaGetTutorScheduleByCourseId) (*ent.Schedule, error)
	GetTutorByID(sr *schema.SchemaGetTutorByID) (*ent.Tutor, error)
	CreateTutor(sr *schema.SchemaCreateTutor) (*ent.Tutor, error)
	UpdateTutor(sr *schema.SchemaUpdateTutor) (*ent.Tutor, error)
	DeleteTutor(sr *schema.SchemaDeleteTutor) error
	UpdateSchedule(sr *schema.SchemaUpdateSchedule) (*ent.Schedule, error)
	GetSchedule(sr *schema.SchemaGetSchedule) (*ent.Schedule, error)
	GetReviews(sr *schema.SchemaGetReviews) ([]*ent.Review, error)
	GetCourses(sr *schema.SchemaGetCourses) ([]*ent.Course, error)
}

type repositoryTutor struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryTutor(c *ent.Client) *repositoryTutor {
	return &repositoryTutor{client: c, ctx: context.Background()}
}

func (r *repositoryTutor) GetTutorByUsername(sr *schema.SchemaGetTutor) (*ent.Tutor, error) {

	tutor, err := r.client.Tutor.
		Query().
		Where(entTutor.HasUserWith(entUser.UsernameEQ(sr.Username))).
		WithUser().
		Only(r.ctx)
	if err != nil {
		return nil, err
	}
	return tutor, nil
}

func (r *repositoryTutor) GetTutorScheduleByCourseId(sr *schema.SchemaGetTutorScheduleByCourseId) (*ent.Schedule, error) {
	//Check if course
	course, err := r.client.Course.
		Query().
		Where(entCourse.IDEQ(sr.Course_id)).
		WithTutor().
		Only(r.ctx)

	if err != nil {
		return nil, err
	}

	//Get Schedules by tutor id
	schedule, err := r.client.Schedule.
		Query().
		Where(entSchedule.HasTutorWith(entTutor.IDEQ(course.Edges.Tutor.ID))).
		Only(r.ctx)

	if err != nil {
		return nil, err
	}
	return schedule, nil
}
func (r *repositoryTutor) GetTutorByID(sr *schema.SchemaGetTutorByID) (*ent.Tutor, error) {
	tutor, err := r.client.Tutor.
		Query().
		Where(entTutor.IDEQ(sr.ID)).
		WithUser().
		Only(r.ctx)
	if err != nil {
		fmt.Println("Broken here", err)
		return nil, err
	}
	return tutor, nil
}

func (r *repositoryTutor) GetTutors() ([]*ent.Tutor, error) {
	tutor, err := r.client.Tutor.
		Query().
		WithUser().
		All(r.ctx)
	if err != nil {
		return nil, err
	}
	return tutor, nil
}

func (r *repositoryTutor) CreateTutor(sr *schema.SchemaCreateTutor) (*ent.Tutor, error) {

	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}
	// wrap the client with the transaction
	txc := tx.Client()

	newUser, err := txc.User.Create().
		SetUsername(sr.Username).
		SetPassword(sr.Password).
		SetFirstName(sr.Firstname).
		SetLastName(sr.Lastname).
		SetPhone(sr.Phone).
		SetAddress(sr.Address).
		SetProfilePictureURL(sr.ProfilePictureURL).
		Save(r.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, err
	}

	tutor, err := txc.Tutor.
		Create().
		SetUserID(newUser.ID).
		SetDescription(sr.Description).
		SetOmiseBankToken(sr.OmiseBankToken).
		SetCitizenID(sr.CitizenId).
		Save(r.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, err
	}

	return tutor, tx.Commit()
}

func (r *repositoryTutor) UpdateTutor(sr *schema.SchemaUpdateTutor) (*ent.Tutor, error) {

	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}
	// wrap the client with the transaction
	txc := tx.Client()

	user, err := txc.User.Query().
		Where(entUser.UsernameEQ(sr.Username)).
		WithTutor().
		Only(r.ctx)
	if err != nil {
		// rollback
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
	}
	tutor := user.Edges.Tutor
	temp, _ := txc.Tutor.Query().Where(entTutor.IDEQ(tutor.ID)).WithSchedule().Only(r.ctx)
	schedule := temp.Edges.Schedule
	
	profilePictureURL := *user.ProfilePictureURL
	if sr.ProfilePicture != nil {
		profilePictureURL, _ = utils.GenerateProfilePictureURL(sr.ProfilePicture, sr.Username, "ProfilePicture")

	}
	user, err = txc.User.
		UpdateOne(user).
		SetFirstName(sr.Firstname).
		SetLastName(sr.Lastname).
		SetEmail(sr.Email).
		SetPhone(sr.Phone).
		SetAddress(sr.Address).
		SetBirthDate(sr.Birthdate).
		SetProfilePictureURL(profilePictureURL).
		Save(r.ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
	}

	tutor, err = txc.Tutor.
		UpdateOne(tutor).
		SetDescription(sr.Description).
		SetOmiseBankToken(sr.OmiseBankToken).
		Save(r.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, err
	}
	ata := sr.Schedule.AvailableTimeArrays()
	if sr.Schedule != nil {
		schedule, err = txc.Schedule.UpdateOne(schedule).
			SetDay0(ata[0]).
			SetDay1(ata[1]).
			SetDay2(ata[2]).
			SetDay3(ata[3]).
			SetDay4(ata[4]).
			SetDay5(ata[5]).
			SetDay6(ata[6]).
			Save(r.ctx)

	}

	// reload tutor->user
	tutor.Edges.User = user
	schedule.Edges.Tutor = tutor
	return tutor, tx.Commit()
}

func (r *repositoryTutor) DeleteTutor(sr *schema.SchemaDeleteTutor) error {
	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return fmt.Errorf("starting a transaction: %w", err)
	}
	// wrap the client with the transaction
	txc := tx.Client()

	err = txc.Tutor.
		DeleteOneID(sr.ID).
		Exec(r.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return err
	}

	err = txc.User.DeleteOneID(sr.ID).Exec(r.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return err
	}

	return nil
}

func (r *repositoryTutor) UpdateSchedule(sr *schema.SchemaUpdateSchedule) (*ent.Schedule, error) {
	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}
	// wrap the client with the transaction
	txc := tx.Client()

	user, err := txc.User.Query().
		Where(entUser.UsernameEQ(sr.Username)).
		WithTutor().
		Only(r.ctx)

	if err != nil {
		// rollback
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
	}


	// load schedule
	oldschedule, err := user.Edges.Tutor.QuerySchedule().Only(r.ctx)
	if err != nil {
		// rollback
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
	}
	// update schedule
	ata := sr.Schedule.AvailableTimeArrays()
	schedule, err := txc.Schedule.UpdateOne(oldschedule).
		SetDay0(ata[0]).
		SetDay1(ata[1]).
		SetDay2(ata[2]).
		SetDay3(ata[3]).
		SetDay4(ata[4]).
		SetDay5(ata[5]).
		SetDay6(ata[6]).
		Save(r.ctx)
	if err != nil {
		// rollback
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
	}
	return schedule, err
}

func (r *repositoryTutor) GetSchedule(sr *schema.SchemaGetSchedule) (*ent.Schedule, error) {
	user, err := r.client.User.Query().
		Where(entUser.UsernameEQ(sr.Username)).
		WithTutor().
		Only(r.ctx)
	if err != nil {
		return nil, err
	}
	schedule, err := user.Edges.Tutor.QuerySchedule().
		Only(r.ctx)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}

func (r *repositoryTutor) GetReviews(sr *schema.SchemaGetReviews) ([]*ent.Review, error) {
	// user, err := r.client.User.Query().
	// 	Where(entUser.IDEQ(sr.ID)).
	// 	WithTutor(func(q *ent.TutorQuery) {
	// 		q.WithCourse(func(q *ent.CourseQuery) {
	// 			q.WithReview()
	// 		})
	// 	}).
	// 	All(r.ctx)

	reviews, err := r.client.Review.Query().
		Where(
			entReview.HasCourseWith(
				entCourse.HasTutorWith(
					entTutor.HasUserWith(entUser.UsernameEQ(sr.Username)),
				),
			),
		).
		WithCourse(func(q *ent.CourseQuery) {
			q.WithTutor(func(q *ent.TutorQuery) {
				q.WithUser()
			},
			)
		}).
		All(r.ctx)

	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *repositoryTutor) GetCourses(sr *schema.SchemaGetCourses) ([]*ent.Course, error) {
	courses, err := r.client.Course.Query().
		Where(
			entCourse.HasTutorWith(
				entTutor.HasUserWith(entUser.UsernameEQ(sr.Username)),
			),
		).
		WithMatch().
		All(r.ctx)

	if err != nil {
		return nil, err
	}

	return courses, nil
}
