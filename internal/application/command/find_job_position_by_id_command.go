package command

import "github.com/google/uuid"

type FindJobPositionByIdCommand struct {
	ID uuid.UUID
}
