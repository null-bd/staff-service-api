package staff

import (
	"context"

	"github.com/null-bd/logger"
	"github.com/stretchr/testify/mock"
)

type mockLogger struct {
	mock.Mock
}

func (m *mockLogger) Debug(msg string, fields logger.Fields) {
	m.Called(msg, fields)
}

func (m *mockLogger) Info(msg string, fields logger.Fields) {
	m.Called(msg, fields)
}

func (m *mockLogger) Warn(msg string, fields logger.Fields) {
	m.Called(msg, fields)
}

func (m *mockLogger) Error(msg string, fields logger.Fields) {
	m.Called(msg, fields)
}

func (m *mockLogger) Fatal(msg string, fields logger.Fields) {
	m.Called(msg, fields)
}

func (m *mockLogger) WithFields(fields logger.Fields) logger.Logger {
	m.Called(fields)
	return nil
}

// Mock Repository
type mockRepository struct {
	mock.Mock
}

func (m *mockRepository) Create(ctx context.Context, staff *Staff) (*Staff, error) {
	args := m.Called(ctx, staff)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Staff), args.Error(1)
}

func (m *mockRepository) GetByCode(ctx context.Context, code string) (*Staff, error) {
	args := m.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Staff), args.Error(1)
}

func (m *mockRepository) GetByID(ctx context.Context, id string) (*Staff, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Staff), args.Error(1)
}
