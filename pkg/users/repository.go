package users

import (
	"github.com/google/uuid"
	"github.com/rishimalgwa/FamPay-Backend-Task/pkg/models"
)

type Repository interface {
	Find(id *uuid.UUID) (*models.Users, error)
}
