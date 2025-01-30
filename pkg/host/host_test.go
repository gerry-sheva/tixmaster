package host_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/gerry-sheva/tixmaster/pkg/common"
	"github.com/gerry-sheva/tixmaster/pkg/database"
	"github.com/gerry-sheva/tixmaster/pkg/host"
	"github.com/gerry-sheva/tixmaster/pkg/testhelper"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HostTestSuite struct {
	suite.Suite
	pgContainer *testhelper.PostgresContainer
	ctx         context.Context
	dbpool      *pgxpool.Pool
	ik          common.ImageKit
	img         *os.File
}

func (suite *HostTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	pgContainer, err := testhelper.CreateTestContainer(suite.ctx)
	if err != nil {
		log.Fatal(err)
	}
	suite.pgContainer = pgContainer

	dbpool := database.ConnectDB(suite.pgContainer.ConnectionString)
	suite.dbpool = dbpool

	ik, err := testhelper.CreateImageKit()
	if err != nil {
		log.Fatal(err)
	}
	suite.ik = ik

	img, err := testhelper.NewImage()
	if err != nil {
		log.Fatal(err)
	}
	suite.img = img
}

func (suite *HostTestSuite) TearDownSuite() {
	suite.dbpool.Close()
	suite.img.Close()
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func (suite *HostTestSuite) TestCreateHost() {
	t := suite.T()

	p := host.NewHostInput{
		Name: "Kivvvi",
		Bio:  "Heelo",
	}

	newHost, err := host.NewHost(suite.ctx, suite.dbpool, suite.ik, suite.img, &p)

	assert.NoError(t, err)
	assert.NotNil(t, newHost)

	assert.Equal(t, p.Name, newHost.Name)
	assert.Equal(t, p.Bio, newHost.Bio)

}

func TestHostSuite(t *testing.T) {
	suite.Run(t, new(HostTestSuite))
}
