package venue_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/gerry-sheva/tixmaster/pkg/common"
	"github.com/gerry-sheva/tixmaster/pkg/database"
	"github.com/gerry-sheva/tixmaster/pkg/testhelper"
	"github.com/gerry-sheva/tixmaster/pkg/venue"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type VenueTestSuite struct {
	suite.Suite
	pgContainer *testhelper.PostgresContainer
	ctx         context.Context
	dbpool      *pgxpool.Pool
	ik          common.ImageKit
	img         *os.File
}

func (suite *VenueTestSuite) SetupSuite() {
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

func (suite *VenueTestSuite) TearDownSuite() {
	suite.dbpool.Close()
	suite.img.Close()
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func (suite *VenueTestSuite) TestCreateVenue() {
	t := suite.T()

	p := venue.NewVenueInput{
		Name:     "Test Venue",
		Capacity: 1000,
		City:     "City",
		State:    "State",
	}
	venue, err := venue.NewVenue(suite.ctx, suite.dbpool, suite.ik.ImageKit, suite.img, &p)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, p.Name, venue.Name)
	assert.Equal(t, p.Capacity, venue.Capacity)
	assert.Equal(t, p.City, venue.City)
	assert.Equal(t, p.State, venue.State)
}

func TestVenueSuite(t *testing.T) {
	suite.Run(t, new(VenueTestSuite))
}
