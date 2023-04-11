package teacher

import (
	"errors"
	"fmt"

	"teacher/pkg/validate"

	"github.com/google/uuid"
)

var (
	ErrInvalidTeacherData = errors.New("invalid teacher data")
)

type Teacher struct {
	id          uuid.UUID
	firstName   string
	lastName    string
	email       string
	phoneNumber string
	subjectID   uuid.UUID
}
type Factory struct{
	idGenerator IDGenerator
}

func (f Factory)New(
	firstName, lastname, email, phoneNumber string,
	subjectID uuid.UUID,
) (Teacher, error) {
	t := Teacher{
		id:          f.idGenerator.GenerateID(),
		firstName:   firstName,
		lastName:    lastname,
		email:       email,
		phoneNumber: phoneNumber,
		subjectID:   subjectID,
	}
	if err := t.validate(); err != nil {
		return Teacher{}, nil
	}
}

func (t Teacher) validate() error {
	if t.firstName == "" || t.lastName == "" {
		return fmt.Errorf("%w: empity firstname or lastname")
	}
	if err := validate.Email(t.email); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidTeacherData, err)
	}
	if err := validate.PhoneNumber(t.phoneNumber); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidTeacherData, err)
		return nil
	}
}
