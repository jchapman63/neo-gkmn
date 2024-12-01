package databasetest

import (
	"context"
	"fmt"
	"os"
	"testing"

	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type DatabaseTestSuite struct {
	suite.Suite
	postgres postgres

	pool *pgxpool.Pool
	tx   pgx.Tx
}

type postgres struct {
	container                            testcontainers.Container
	user, password, database, host, port string
}

func (s *DatabaseTestSuite) dbURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		s.postgres.user, s.postgres.password, s.postgres.host, s.postgres.port, s.postgres.database)
}

func (s *DatabaseTestSuite) SetupSuite() {
	s.postgres.user = "postgres"
	s.postgres.password = "password"
	s.postgres.database = "cloud"

	ctx := context.Background()

	ct, err := testcontainers.GenericContainer(ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: testcontainers.ContainerRequest{
				Image:        "postgres:14.2",
				ExposedPorts: []string{"5432/tcp"},
				WaitingFor:   wait.ForExposedPort(),
				Env: map[string]string{
					"POSTGRES_USER":     s.postgres.user,
					"POSTGRES_PASSWORD": s.postgres.password,
					"POSTGRES_DB":       s.postgres.database,
				},
				Tmpfs: map[string]string{"/var/lib/postgresql/data": "rw"},
			},
			Started: true,
		})
	s.NoError(err)

	s.postgres.container = ct

	ip, err := ct.Host(ctx)
	s.NoError(err)
	s.postgres.host = ip

	port, err := ct.MappedPort(ctx, "5432")
	s.NoError(err)
	s.postgres.port = port.Port()

	conn, err := pgx.Connect(ctx, s.dbURL())
	s.NoError(err)
	defer conn.Close(ctx)

	b, err := os.ReadFile("../../../db/schema.sql")
	s.NoError(err)
	_, err = conn.Exec(ctx, string(b))
	s.NoError(err)

	pool, err := pgxpool.New(ctx, s.dbURL())
	s.NoError(err)

	s.pool = pool
	// TODO - consider seeding
	//if err := Seed(ctx, conn); err != nil {
	//	s.FailNow("failed to seed database")
	//}
}

func (s *DatabaseTestSuite) SetupTest() {
	tx, err := s.pool.Begin(context.Background())
	if err != nil {
		s.FailNow("failed to begin transaction", err)
	}

	s.tx = tx
}

func (s *DatabaseTestSuite) TearDownTest() {
	if err := s.tx.Rollback(context.Background()); err != nil {
		s.FailNow("failed to rollback transaction", err)
	}
}

func (s *DatabaseTestSuite) TearDownSuite() {
	ctx := context.Background()
	s.NoError(s.postgres.container.Terminate(ctx))
}

func TestDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseTestSuite))
}
