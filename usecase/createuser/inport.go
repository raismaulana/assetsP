package createuser

import (
	"context"
)

// Inport of CreateUser
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase CreateUser
type InportRequest struct {
	Name     string `json:"name"`
	Email    string `gorm:"email"`
	Password string `gorm:"password"`
}

// InportResponse is response payload after running the usecase CreateUser
type InportResponse struct {
}
