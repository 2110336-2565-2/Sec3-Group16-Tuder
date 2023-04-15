package repositorys

import (
	"context"
	"log"
	"os"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

type RepositoryPayment interface {
	GetQRCodeForCoursePayment(sr *schema.SchemaGetQRCodeForCoursePayment) (string, error)
	GetQRCodeForTuitionFree(sr *schema.SchemaGetQRCodeForTuitionFree) (string, error)
}

type repositoryPayment struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryPayment(c *ent.Client) *repositoryPayment {
	return &repositoryPayment{client: c, ctx: context.Background()}
}

func (r *repositoryPayment) GetQRCodeForCoursePayment(sr *schema.SchemaGetQRCodeForCoursePayment) (string, error) {
	omisePublicKey := os.Getenv("OMISE_PUBLIC_KEY")
	omiseSecretKey := os.Getenv("OMISE_SECRET_KEY")

	client, err := omise.NewClient(omisePublicKey, omiseSecretKey)

	if err != nil {
		log.Fatal(err)
	}
	// Convert payment.Amount to int64
	amount := int64(sr.Amount)

	// Check if course exists and get tutor with omise bank token
	course, err := r.client.Course.
		Query().
		Where(course.IDEQ(sr.Course_id)).
		WithTutor().
		Only(r.ctx)

	if err != nil {
		return "", err
	}

	// Create a new source
	source, createSource := &omise.Source{}, &operations.CreateSource{
		Amount:   amount,
		Currency: "THB",
		Type:     "promptpay",
	}

	if e := client.Do(source, createSource); e != nil {
		log.Fatal(e)
	}

	// Create a new charge
	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   amount,
		Currency: "thb",
		Customer: *course.Edges.Tutor.OmiseCustomerID,
		Source:   source.ID,
	}

	if e := client.Do(charge, createCharge); e != nil {
		log.Fatal(e)
	}

	qrCodePictureURL := charge.Source.ScannableCode.Image.DownloadURI
	return qrCodePictureURL, nil
}

func (r *repositoryPayment) GetQRCodeForTuitionFree(sr *schema.SchemaGetQRCodeForTuitionFree) (string, error) {
	omisePublicKey := os.Getenv("OMISE_PUBLIC_KEY")
	omiseSecretKey := os.Getenv("OMISE_SECRET_KEY")

	client, err := omise.NewClient(omisePublicKey, omiseSecretKey)

	if err != nil {
		log.Fatal(err)
	}
	// Convert payment.Amount to int64
	amount := int64(sr.Amount)

	//Check if tutor exists and get omise bank token
	tutor, err := r.client.Tutor.
		Query().
		Where(tutor.IDEQ(sr.Tutor_id)).
		Only(r.ctx)

	if err != nil {
		return "", err
	}

	// Create a new source
	source, createSource := &omise.Source{}, &operations.CreateSource{
		Amount:   amount,
		Currency: "THB",
		Type:     "promptpay",
	}

	if e := client.Do(source, createSource); e != nil {
		log.Fatal(e)
	}

	// Create a new charge
	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   amount,
		Currency: "thb",
		Customer: *tutor.OmiseCustomerID,
		Source:   source.ID,
	}

	if e := client.Do(charge, createCharge); e != nil {
		log.Fatal(e)
	}

	qrCodePictureURL := charge.Source.ScannableCode.Image.DownloadURI
	return qrCodePictureURL, nil
}
