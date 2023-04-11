package teacher

import "github.com/google/uuid"

type IDGenerator interface{
	GenerateID() uuid.UUID
}


type uuidGenerator struct{}

func(u uuidGenerator) GenerateID() uuid.UUID{
	return uuid.New()
 }