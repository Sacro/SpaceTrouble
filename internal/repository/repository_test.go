package repository

import (
	"context"
	"testing"

	"github.com/Sacro/SpaceTrouble/internal/ticket"
	"github.com/docker/go-connections/nat"
	"github.com/go-pg/pg/v10"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type RepositoryTestSuite struct {
	suite.Suite

	container testcontainers.Container
	repo      *TicketRepo
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (s *RepositoryTestSuite) SetupTest() {
	var (
		ctx      = context.Background()
		user     = "postgres"
		password = "password"
		port     = "5432/tcp"
		db       = "spacetrouble"
	)

	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:latest",
			ExposedPorts: []string{port},
			Cmd:          []string{"postgres", "-c", "fsync=off"},
			Env: map[string]string{
				"POSTGRES_USER":     user,
				"POSTGRES_PASSWORD": password,
				"POSTGRES_DB":       db,
			},
			WaitingFor: wait.ForListeningPort(nat.Port(port)),
		},
		Started: true,
	}
	container, err := testcontainers.GenericContainer(ctx, req)
	s.Require().Nil(err)

	// Get postgres port
	mappedPort, err := container.MappedPort(ctx, nat.Port(port))

	s.Require().Nil(err)

	// Create a connection to the db
	s.repo = New(pg.Connect(&pg.Options{
		Addr:     "127.0.0.1:" + mappedPort.Port(),
		Database: db,
		Password: password,
		User:     user,
	}))

	// Run the migration
	err = s.repo.CreateSchema()
	s.Require().Nil(err)

	s.container = container
}

func (s *RepositoryTestSuite) TearDownTest() {
	err := s.container.Terminate(context.Background())
	s.Assert().Nil(err)
}

func (s *RepositoryTestSuite) TestInsertSingleTicket() {
	f := faker.New()

	result, err := s.repo.db.Model(ticket.Fake(f)).Insert()
	s.Assert().Nil(err)

	s.Assert().Equal(result.RowsAffected(), 1)
}
