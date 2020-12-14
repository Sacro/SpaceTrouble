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
	repo      *Repository
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (suite *RepositoryTestSuite) SetupTest() {
	var (
		ctx      = context.Background()
		user     = "postgres"
		password = "password"
		port     = "5432/tcp"
		db       = "spacetrouble"

		// dbURL = func(port nat.Port) string {
		// 	return fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", user, password, port, db)
		// }
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
	suite.Require().Nil(err)

	// Get postgres port
	mappedPort, err := container.MappedPort(ctx, nat.Port(port))

	suite.Require().Nil(err)

	// Create a connection to the db
	suite.repo = New(pg.Connect(&pg.Options{
		Addr:     "127.0.0.1:" + mappedPort.Port(),
		Database: db,
		Password: password,
		User:     user,
	}))

	// Run the migration
	err = suite.repo.CreateSchema()
	suite.Require().Nil(err)

	suite.container = container
}

func (suite *RepositoryTestSuite) TearDownTest() {
	suite.container.Terminate(context.Background())
}

func (suite *RepositoryTestSuite) TestInsertSingleTicket() {
	f := faker.New()

	result, err := suite.repo.db.Model(ticket.Fake(f)).Insert()
	suite.Assert().Nil(err)

	suite.Assert().Equal(result.RowsAffected(), 1)
}
