package mocks

import (
	"github.com/Sirpyerre/payment-platform/internal/models"
	"github.com/stretchr/testify/mock"
)

type RefundRepositoryMock struct {
	mock.Mock
}

func (m *RefundRepositoryMock) Refund(*models.Refund) error {
	args := m.Called()
	return args.Error(0)
}
