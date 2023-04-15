package repositorys

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	entPayment "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/payment"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

type RepositoryPayment interface {
	GetQRCodeForCoursePayment(sr *schema.SchemaGetQRCodeForCoursePayment) (*schemas.SchemaQRCode, error)
	GetQRCodeForTuitionFree(sr *schema.SchemaGetQRCodeForTuitionFree) (*schemas.SchemaQRCode, error)
	WebhookChargeHandler(sr *schema.SchemaChargeWebhook) (*ent.Payment, error)
}

type repositoryPayment struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryPayment(c *ent.Client) *repositoryPayment {
	return &repositoryPayment{client: c, ctx: context.Background()}
}

func (r *repositoryPayment) GetQRCodeForCoursePayment(sr *schema.SchemaGetQRCodeForCoursePayment) (*schemas.SchemaQRCode, error) {
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
		WithTutor(func(tq *ent.TutorQuery) {
			tq.WithUser()
		}).
		Only(r.ctx)

	if err != nil {
		return nil, err
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

	// Create Payment record
	payment, err := r.client.Payment.
		Create().
		SetAmount(sr.Amount).
		SetUserID(course.Edges.Tutor.Edges.User.ID).
		SetPaymentStatus("processing").
		SetQrPictureURL(qrCodePictureURL).
		SetCurrency("THB").
		SetChargeID(charge.ID).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return &schema.SchemaQRCode{
		Payment:   payment,
		QrCodeUrl: qrCodePictureURL,
		Charge_id: charge.ID,
	}, nil
}

func (r *repositoryPayment) GetQRCodeForTuitionFree(sr *schema.SchemaGetQRCodeForTuitionFree) (*schemas.SchemaQRCode, error) {
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
		WithUser().
		Only(r.ctx)

	if err != nil {
		return nil, err
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

	// Create Payment record
	payment, err := r.client.Payment.
		Create().
		SetAmount(sr.Amount).
		SetUserID(tutor.Edges.User.ID).
		SetPaymentStatus("processing").
		SetQrPictureURL(qrCodePictureURL).
		SetCurrency("THB").
		SetChargeID(charge.ID).
		SetAppointmentID(sr.Appointment_id).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return &schema.SchemaQRCode{
		Payment:   payment,
		QrCodeUrl: qrCodePictureURL,
		Charge_id: charge.ID,
	}, nil
}

func (r *repositoryPayment) WebhookChargeHandler(sr *schema.SchemaChargeWebhook) (*ent.Payment, error) {
	// Check if charge exists
	_, err := r.client.Payment.
		Query().
		Where(entPayment.ChargeID(sr.Data.ID)).
		Only(r.ctx)

	if err != nil {
		return nil, err
	}
	// Check if charge is pending
	if sr.Data.Status == "pending" {
		// Check if payment exists
		payment, err := r.client.Payment.
			Query().
			Where(entPayment.ChargeID(sr.Data.ID)).
			Only(r.ctx)

		if err != nil {
			return nil, err
		}

		// Update payment status to pending
		_, err = r.client.Payment.
			UpdateOne(payment).
			SetPaymentStatus("pending").
			Save(r.ctx)

		if err != nil {
			return nil, err
		}
		return payment, nil

	} else if sr.Data.Status == "successful" {
		// Check if payment exists
		payment, err := r.client.Payment.
			Query().
			Where(entPayment.ChargeID(sr.Data.ID)).
			Only(r.ctx)

		if err != nil {
			return nil, err
		}

		// Update payment status to successful
		_, err = r.client.Payment.
			UpdateOne(payment).
			SetPaymentStatus("successful").
			Save(r.ctx)

		if err != nil {
			return nil, err
		}
		return payment, nil

	} else if sr.Data.Status == "failed" {
		// Check if payment exists
		payment, err := r.client.Payment.
			Query().
			Where(entPayment.ChargeID(sr.Data.ID)).
			Only(r.ctx)

		if err != nil {
			return nil, err
		}

		// Update payment status to successful
		_, err = r.client.Payment.
			UpdateOne(payment).
			SetPaymentStatus("failed").
			Save(r.ctx)

		if err != nil {
			return nil, err
		}
		return payment, nil

	} else {
		return nil, errors.New("payment status is not valid")
	}
}
