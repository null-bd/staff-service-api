package domain

import (
	"testing"

	"github.com/null-bd/microservice-name/internal/testutil"
	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuite struct {
	suite.Suite
	tc   *testutil.TestContainer
	repo IDomainRepository
}

func TestRepositorySuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	suite.Run(t, new(RepositoryTestSuite))
}
