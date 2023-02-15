// Code generated by ent, DO NOT EDIT.

package tutor

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Tutor {
	return predicate.Tutor(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Tutor {
	return predicate.Tutor(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Tutor {
	return predicate.Tutor(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Tutor {
	return predicate.Tutor(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Tutor {
	return predicate.Tutor(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Tutor {
	return predicate.Tutor(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Tutor {
	return predicate.Tutor(sql.FieldLTE(FieldID, id))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldUsername, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldPassword, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldEmail, v))
}

// FirstName applies equality check predicate on the "first_name" field. It's identical to FirstNameEQ.
func FirstName(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldFirstName, v))
}

// LastName applies equality check predicate on the "last_name" field. It's identical to LastNameEQ.
func LastName(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldLastName, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldAddress, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldPhone, v))
}

// BirthDate applies equality check predicate on the "birth_date" field. It's identical to BirthDateEQ.
func BirthDate(v time.Time) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldBirthDate, v))
}

// Gender applies equality check predicate on the "gender" field. It's identical to GenderEQ.
func Gender(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldGender, v))
}

// ProfilePictureURL applies equality check predicate on the "profile_picture_URL" field. It's identical to ProfilePictureURLEQ.
func ProfilePictureURL(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldProfilePictureURL, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldDescription, v))
}

// OmiseBankToken applies equality check predicate on the "omise_bank_token" field. It's identical to OmiseBankTokenEQ.
func OmiseBankToken(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldOmiseBankToken, v))
}

// CitizenID applies equality check predicate on the "citizen_id" field. It's identical to CitizenIDEQ.
func CitizenID(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldCitizenID, v))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.Tutor {
	return predicate.Tutor(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.Tutor {
	return predicate.Tutor(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldContainsFold(FieldUsername, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.Tutor {
	return predicate.Tutor(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.Tutor {
	return predicate.Tutor(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldContainsFold(FieldPassword, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Tutor {
	return predicate.Tutor(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Tutor {
	return predicate.Tutor(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldContainsFold(FieldEmail, v))
}

// FirstNameEQ applies the EQ predicate on the "first_name" field.
func FirstNameEQ(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldFirstName, v))
}

// FirstNameNEQ applies the NEQ predicate on the "first_name" field.
func FirstNameNEQ(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldNEQ(FieldFirstName, v))
}

// FirstNameIn applies the In predicate on the "first_name" field.
func FirstNameIn(vs ...string) predicate.Tutor {
	return predicate.Tutor(sql.FieldIn(FieldFirstName, vs...))
}

// FirstNameNotIn applies the NotIn predicate on the "first_name" field.
func FirstNameNotIn(vs ...string) predicate.Tutor {
	return predicate.Tutor(sql.FieldNotIn(FieldFirstName, vs...))
}

// FirstNameGT applies the GT predicate on the "first_name" field.
func FirstNameGT(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldGT(FieldFirstName, v))
}

// FirstNameGTE applies the GTE predicate on the "first_name" field.
func FirstNameGTE(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldGTE(FieldFirstName, v))
}

// FirstNameLT applies the LT predicate on the "first_name" field.
func FirstNameLT(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldLT(FieldFirstName, v))
}

// FirstNameLTE applies the LTE predicate on the "first_name" field.
func FirstNameLTE(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldLTE(FieldFirstName, v))
}

// FirstNameContains applies the Contains predicate on the "first_name" field.
func FirstNameContains(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldContains(FieldFirstName, v))
}

// FirstNameHasPrefix applies the HasPrefix predicate on the "first_name" field.
func FirstNameHasPrefix(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldHasPrefix(FieldFirstName, v))
}

// FirstNameHasSuffix applies the HasSuffix predicate on the "first_name" field.
func FirstNameHasSuffix(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldHasSuffix(FieldFirstName, v))
}

// FirstNameEqualFold applies the EqualFold predicate on the "first_name" field.
func FirstNameEqualFold(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEqualFold(FieldFirstName, v))
}

// FirstNameContainsFold applies the ContainsFold predicate on the "first_name" field.
func FirstNameContainsFold(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldContainsFold(FieldFirstName, v))
}

// LastNameEQ applies the EQ predicate on the "last_name" field.
func LastNameEQ(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldLastName, v))
}

// LastNameNEQ applies the NEQ predicate on the "last_name" field.
func LastNameNEQ(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldNEQ(FieldLastName, v))
}

// LastNameIn applies the In predicate on the "last_name" field.
func LastNameIn(vs ...string) predicate.Tutor {
	return predicate.Tutor(sql.FieldIn(FieldLastName, vs...))
}

// LastNameNotIn applies the NotIn predicate on the "last_name" field.
func LastNameNotIn(vs ...string) predicate.Tutor {
	return predicate.Tutor(sql.FieldNotIn(FieldLastName, vs...))
}

// LastNameGT applies the GT predicate on the "last_name" field.
func LastNameGT(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldGT(FieldLastName, v))
}

// LastNameGTE applies the GTE predicate on the "last_name" field.
func LastNameGTE(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldGTE(FieldLastName, v))
}

// LastNameLT applies the LT predicate on the "last_name" field.
func LastNameLT(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldLT(FieldLastName, v))
}

// LastNameLTE applies the LTE predicate on the "last_name" field.
func LastNameLTE(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldLTE(FieldLastName, v))
}

// LastNameContains applies the Contains predicate on the "last_name" field.
func LastNameContains(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldContains(FieldLastName, v))
}

// LastNameHasPrefix applies the HasPrefix predicate on the "last_name" field.
func LastNameHasPrefix(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldHasPrefix(FieldLastName, v))
}

// LastNameHasSuffix applies the HasSuffix predicate on the "last_name" field.
func LastNameHasSuffix(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldHasSuffix(FieldLastName, v))
}

// LastNameEqualFold applies the EqualFold predicate on the "last_name" field.
func LastNameEqualFold(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEqualFold(FieldLastName, v))
}

// LastNameContainsFold applies the ContainsFold predicate on the "last_name" field.
func LastNameContainsFold(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldContainsFold(FieldLastName, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Tutor {
	return predicate.Tutor(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Tutor {
	return predicate.Tutor(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldContainsFold(FieldAddress, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.Tutor {
	return predicate.Tutor(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.Tutor {
	return predicate.Tutor(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldContainsFold(FieldPhone, v))
}

// BirthDateEQ applies the EQ predicate on the "birth_date" field.
func BirthDateEQ(v time.Time) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldBirthDate, v))
}

// BirthDateNEQ applies the NEQ predicate on the "birth_date" field.
func BirthDateNEQ(v time.Time) predicate.Tutor {
	return predicate.Tutor(sql.FieldNEQ(FieldBirthDate, v))
}

// BirthDateIn applies the In predicate on the "birth_date" field.
func BirthDateIn(vs ...time.Time) predicate.Tutor {
	return predicate.Tutor(sql.FieldIn(FieldBirthDate, vs...))
}

// BirthDateNotIn applies the NotIn predicate on the "birth_date" field.
func BirthDateNotIn(vs ...time.Time) predicate.Tutor {
	return predicate.Tutor(sql.FieldNotIn(FieldBirthDate, vs...))
}

// BirthDateGT applies the GT predicate on the "birth_date" field.
func BirthDateGT(v time.Time) predicate.Tutor {
	return predicate.Tutor(sql.FieldGT(FieldBirthDate, v))
}

// BirthDateGTE applies the GTE predicate on the "birth_date" field.
func BirthDateGTE(v time.Time) predicate.Tutor {
	return predicate.Tutor(sql.FieldGTE(FieldBirthDate, v))
}

// BirthDateLT applies the LT predicate on the "birth_date" field.
func BirthDateLT(v time.Time) predicate.Tutor {
	return predicate.Tutor(sql.FieldLT(FieldBirthDate, v))
}

// BirthDateLTE applies the LTE predicate on the "birth_date" field.
func BirthDateLTE(v time.Time) predicate.Tutor {
	return predicate.Tutor(sql.FieldLTE(FieldBirthDate, v))
}

// GenderEQ applies the EQ predicate on the "gender" field.
func GenderEQ(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldGender, v))
}

// GenderNEQ applies the NEQ predicate on the "gender" field.
func GenderNEQ(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldNEQ(FieldGender, v))
}

// GenderIn applies the In predicate on the "gender" field.
func GenderIn(vs ...string) predicate.Tutor {
	return predicate.Tutor(sql.FieldIn(FieldGender, vs...))
}

// GenderNotIn applies the NotIn predicate on the "gender" field.
func GenderNotIn(vs ...string) predicate.Tutor {
	return predicate.Tutor(sql.FieldNotIn(FieldGender, vs...))
}

// GenderGT applies the GT predicate on the "gender" field.
func GenderGT(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldGT(FieldGender, v))
}

// GenderGTE applies the GTE predicate on the "gender" field.
func GenderGTE(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldGTE(FieldGender, v))
}

// GenderLT applies the LT predicate on the "gender" field.
func GenderLT(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldLT(FieldGender, v))
}

// GenderLTE applies the LTE predicate on the "gender" field.
func GenderLTE(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldLTE(FieldGender, v))
}

// GenderContains applies the Contains predicate on the "gender" field.
func GenderContains(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldContains(FieldGender, v))
}

// GenderHasPrefix applies the HasPrefix predicate on the "gender" field.
func GenderHasPrefix(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldHasPrefix(FieldGender, v))
}

// GenderHasSuffix applies the HasSuffix predicate on the "gender" field.
func GenderHasSuffix(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldHasSuffix(FieldGender, v))
}

// GenderEqualFold applies the EqualFold predicate on the "gender" field.
func GenderEqualFold(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEqualFold(FieldGender, v))
}

// GenderContainsFold applies the ContainsFold predicate on the "gender" field.
func GenderContainsFold(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldContainsFold(FieldGender, v))
}

// ProfilePictureURLEQ applies the EQ predicate on the "profile_picture_URL" field.
func ProfilePictureURLEQ(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldProfilePictureURL, v))
}

// ProfilePictureURLNEQ applies the NEQ predicate on the "profile_picture_URL" field.
func ProfilePictureURLNEQ(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldNEQ(FieldProfilePictureURL, v))
}

// ProfilePictureURLIn applies the In predicate on the "profile_picture_URL" field.
func ProfilePictureURLIn(vs ...string) predicate.Tutor {
	return predicate.Tutor(sql.FieldIn(FieldProfilePictureURL, vs...))
}

// ProfilePictureURLNotIn applies the NotIn predicate on the "profile_picture_URL" field.
func ProfilePictureURLNotIn(vs ...string) predicate.Tutor {
	return predicate.Tutor(sql.FieldNotIn(FieldProfilePictureURL, vs...))
}

// ProfilePictureURLGT applies the GT predicate on the "profile_picture_URL" field.
func ProfilePictureURLGT(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldGT(FieldProfilePictureURL, v))
}

// ProfilePictureURLGTE applies the GTE predicate on the "profile_picture_URL" field.
func ProfilePictureURLGTE(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldGTE(FieldProfilePictureURL, v))
}

// ProfilePictureURLLT applies the LT predicate on the "profile_picture_URL" field.
func ProfilePictureURLLT(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldLT(FieldProfilePictureURL, v))
}

// ProfilePictureURLLTE applies the LTE predicate on the "profile_picture_URL" field.
func ProfilePictureURLLTE(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldLTE(FieldProfilePictureURL, v))
}

// ProfilePictureURLContains applies the Contains predicate on the "profile_picture_URL" field.
func ProfilePictureURLContains(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldContains(FieldProfilePictureURL, v))
}

// ProfilePictureURLHasPrefix applies the HasPrefix predicate on the "profile_picture_URL" field.
func ProfilePictureURLHasPrefix(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldHasPrefix(FieldProfilePictureURL, v))
}

// ProfilePictureURLHasSuffix applies the HasSuffix predicate on the "profile_picture_URL" field.
func ProfilePictureURLHasSuffix(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldHasSuffix(FieldProfilePictureURL, v))
}

// ProfilePictureURLIsNil applies the IsNil predicate on the "profile_picture_URL" field.
func ProfilePictureURLIsNil() predicate.Tutor {
	return predicate.Tutor(sql.FieldIsNull(FieldProfilePictureURL))
}

// ProfilePictureURLNotNil applies the NotNil predicate on the "profile_picture_URL" field.
func ProfilePictureURLNotNil() predicate.Tutor {
	return predicate.Tutor(sql.FieldNotNull(FieldProfilePictureURL))
}

// ProfilePictureURLEqualFold applies the EqualFold predicate on the "profile_picture_URL" field.
func ProfilePictureURLEqualFold(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEqualFold(FieldProfilePictureURL, v))
}

// ProfilePictureURLContainsFold applies the ContainsFold predicate on the "profile_picture_URL" field.
func ProfilePictureURLContainsFold(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldContainsFold(FieldProfilePictureURL, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Tutor {
	return predicate.Tutor(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Tutor {
	return predicate.Tutor(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Tutor {
	return predicate.Tutor(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Tutor {
	return predicate.Tutor(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldContainsFold(FieldDescription, v))
}

// OmiseBankTokenEQ applies the EQ predicate on the "omise_bank_token" field.
func OmiseBankTokenEQ(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldOmiseBankToken, v))
}

// OmiseBankTokenNEQ applies the NEQ predicate on the "omise_bank_token" field.
func OmiseBankTokenNEQ(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldNEQ(FieldOmiseBankToken, v))
}

// OmiseBankTokenIn applies the In predicate on the "omise_bank_token" field.
func OmiseBankTokenIn(vs ...string) predicate.Tutor {
	return predicate.Tutor(sql.FieldIn(FieldOmiseBankToken, vs...))
}

// OmiseBankTokenNotIn applies the NotIn predicate on the "omise_bank_token" field.
func OmiseBankTokenNotIn(vs ...string) predicate.Tutor {
	return predicate.Tutor(sql.FieldNotIn(FieldOmiseBankToken, vs...))
}

// OmiseBankTokenGT applies the GT predicate on the "omise_bank_token" field.
func OmiseBankTokenGT(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldGT(FieldOmiseBankToken, v))
}

// OmiseBankTokenGTE applies the GTE predicate on the "omise_bank_token" field.
func OmiseBankTokenGTE(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldGTE(FieldOmiseBankToken, v))
}

// OmiseBankTokenLT applies the LT predicate on the "omise_bank_token" field.
func OmiseBankTokenLT(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldLT(FieldOmiseBankToken, v))
}

// OmiseBankTokenLTE applies the LTE predicate on the "omise_bank_token" field.
func OmiseBankTokenLTE(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldLTE(FieldOmiseBankToken, v))
}

// OmiseBankTokenContains applies the Contains predicate on the "omise_bank_token" field.
func OmiseBankTokenContains(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldContains(FieldOmiseBankToken, v))
}

// OmiseBankTokenHasPrefix applies the HasPrefix predicate on the "omise_bank_token" field.
func OmiseBankTokenHasPrefix(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldHasPrefix(FieldOmiseBankToken, v))
}

// OmiseBankTokenHasSuffix applies the HasSuffix predicate on the "omise_bank_token" field.
func OmiseBankTokenHasSuffix(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldHasSuffix(FieldOmiseBankToken, v))
}

// OmiseBankTokenIsNil applies the IsNil predicate on the "omise_bank_token" field.
func OmiseBankTokenIsNil() predicate.Tutor {
	return predicate.Tutor(sql.FieldIsNull(FieldOmiseBankToken))
}

// OmiseBankTokenNotNil applies the NotNil predicate on the "omise_bank_token" field.
func OmiseBankTokenNotNil() predicate.Tutor {
	return predicate.Tutor(sql.FieldNotNull(FieldOmiseBankToken))
}

// OmiseBankTokenEqualFold applies the EqualFold predicate on the "omise_bank_token" field.
func OmiseBankTokenEqualFold(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEqualFold(FieldOmiseBankToken, v))
}

// OmiseBankTokenContainsFold applies the ContainsFold predicate on the "omise_bank_token" field.
func OmiseBankTokenContainsFold(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldContainsFold(FieldOmiseBankToken, v))
}

// CitizenIDEQ applies the EQ predicate on the "citizen_id" field.
func CitizenIDEQ(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldCitizenID, v))
}

// CitizenIDNEQ applies the NEQ predicate on the "citizen_id" field.
func CitizenIDNEQ(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldNEQ(FieldCitizenID, v))
}

// CitizenIDIn applies the In predicate on the "citizen_id" field.
func CitizenIDIn(vs ...string) predicate.Tutor {
	return predicate.Tutor(sql.FieldIn(FieldCitizenID, vs...))
}

// CitizenIDNotIn applies the NotIn predicate on the "citizen_id" field.
func CitizenIDNotIn(vs ...string) predicate.Tutor {
	return predicate.Tutor(sql.FieldNotIn(FieldCitizenID, vs...))
}

// CitizenIDGT applies the GT predicate on the "citizen_id" field.
func CitizenIDGT(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldGT(FieldCitizenID, v))
}

// CitizenIDGTE applies the GTE predicate on the "citizen_id" field.
func CitizenIDGTE(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldGTE(FieldCitizenID, v))
}

// CitizenIDLT applies the LT predicate on the "citizen_id" field.
func CitizenIDLT(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldLT(FieldCitizenID, v))
}

// CitizenIDLTE applies the LTE predicate on the "citizen_id" field.
func CitizenIDLTE(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldLTE(FieldCitizenID, v))
}

// CitizenIDContains applies the Contains predicate on the "citizen_id" field.
func CitizenIDContains(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldContains(FieldCitizenID, v))
}

// CitizenIDHasPrefix applies the HasPrefix predicate on the "citizen_id" field.
func CitizenIDHasPrefix(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldHasPrefix(FieldCitizenID, v))
}

// CitizenIDHasSuffix applies the HasSuffix predicate on the "citizen_id" field.
func CitizenIDHasSuffix(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldHasSuffix(FieldCitizenID, v))
}

// CitizenIDEqualFold applies the EqualFold predicate on the "citizen_id" field.
func CitizenIDEqualFold(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldEqualFold(FieldCitizenID, v))
}

// CitizenIDContainsFold applies the ContainsFold predicate on the "citizen_id" field.
func CitizenIDContainsFold(v string) predicate.Tutor {
	return predicate.Tutor(sql.FieldContainsFold(FieldCitizenID, v))
}

// HasIssueReport applies the HasEdge predicate on the "issue_report" edge.
func HasIssueReport() predicate.Tutor {
	return predicate.Tutor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, IssueReportTable, IssueReportColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIssueReportWith applies the HasEdge predicate on the "issue_report" edge with a given conditions (other predicates).
func HasIssueReportWith(preds ...predicate.IssueReport) predicate.Tutor {
	return predicate.Tutor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(IssueReportInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, IssueReportTable, IssueReportColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCourse applies the HasEdge predicate on the "course" edge.
func HasCourse() predicate.Tutor {
	return predicate.Tutor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CourseTable, CourseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCourseWith applies the HasEdge predicate on the "course" edge with a given conditions (other predicates).
func HasCourseWith(preds ...predicate.Course) predicate.Tutor {
	return predicate.Tutor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CourseInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CourseTable, CourseColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReviewTutor applies the HasEdge predicate on the "review_tutor" edge.
func HasReviewTutor() predicate.Tutor {
	return predicate.Tutor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReviewTutorTable, ReviewTutorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReviewTutorWith applies the HasEdge predicate on the "review_tutor" edge with a given conditions (other predicates).
func HasReviewTutorWith(preds ...predicate.ReviewTutor) predicate.Tutor {
	return predicate.Tutor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReviewTutorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReviewTutorTable, ReviewTutorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSchedule applies the HasEdge predicate on the "schedule" edge.
func HasSchedule() predicate.Tutor {
	return predicate.Tutor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ScheduleTable, ScheduleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasScheduleWith applies the HasEdge predicate on the "schedule" edge with a given conditions (other predicates).
func HasScheduleWith(preds ...predicate.Schedule) predicate.Tutor {
	return predicate.Tutor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ScheduleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ScheduleTable, ScheduleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Tutor) predicate.Tutor {
	return predicate.Tutor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Tutor) predicate.Tutor {
	return predicate.Tutor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Tutor) predicate.Tutor {
	return predicate.Tutor(func(s *sql.Selector) {
		p(s.Not())
	})
}
