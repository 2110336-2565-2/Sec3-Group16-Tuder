package repositorys

import (
	"bytes"
	"context"
	"encoding/base64"
	"image/png"
	"log"
	"os"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"github.com/skip2/go-qrcode"
)

type RepositoryPayment interface {
	GetQRCode(sr *schema.SchemaGetQRCode) (string, error)
}

type repositoryPayment struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryPayment(c *ent.Client) *repositoryPayment {
	return &repositoryPayment{client: c, ctx: context.Background()}
}

func (r *repositoryPayment) GetQRCode(sr *schema.SchemaGetQRCode) (string, error) {
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

	// Create a new charge
	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   amount,
		Currency: "thb",
		Card:     *course.Edges.Tutor.OmiseBankToken,
	}

	// // Set the expiration time to 5 minutes from now
	// charge.Metadata["expiration_time"] = time.Now().Add(time.Minute * 5).Format(time.RFC3339)

	if e := client.Do(charge, createCharge); e != nil {
		log.Fatal(e)
	}

	log.Printf("charge: %s  amount: %s %d\n", charge.ID, charge.Currency, charge.Amount)

	// Create a new QR code
	qrCode, err := qrcode.New(charge.Source.ScannableCode.Image.Object, qrcode.Medium)

	// Resize the image to 400x400 pixels
	qrImage := qrCode.Image(400)

	// Convert the image to PNG format
	var buf bytes.Buffer
	if err := png.Encode(&buf, qrImage); err != nil {
		return "", err
	}

	// Encode the PNG image to base64 string
	qrCodeBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())

	return qrCodeBase64, nil
}
