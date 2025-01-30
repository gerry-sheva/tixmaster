package event_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/gerry-sheva/tixmaster/pkg/common"
	"github.com/gerry-sheva/tixmaster/pkg/database"
	"github.com/gerry-sheva/tixmaster/pkg/event"
	"github.com/gerry-sheva/tixmaster/pkg/host"
	"github.com/gerry-sheva/tixmaster/pkg/testhelper"
	"github.com/gerry-sheva/tixmaster/pkg/venue"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/meilisearch/meilisearch-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EventTestSuite struct {
	suite.Suite
	pgContainer    *testhelper.PostgresContainer
	meiliContainer *testhelper.MeilisearchContainer
	ctx            context.Context
	dbpool         *pgxpool.Pool
	meili          *meilisearch.ServiceManager
	ik             common.ImageKit
	hostImg        *os.File
	venueImg       *os.File
	thumbnail      *os.File
	banner         *os.File
}

func (suite *EventTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	pgContainer, err := testhelper.CreateTestContainer(suite.ctx)
	if err != nil {
		log.Fatalf("Failed to start pg container: %s", err)
	}
	suite.pgContainer = pgContainer

	dbpool := database.ConnectDB(suite.pgContainer.ConnectionString)
	suite.dbpool = dbpool

	meili, err := testhelper.CreateMeilisearchContainer(suite.ctx)
	if err != nil {
		log.Fatalf("Failed to start meili container: %s", err)
	}

	meiliClient := meilisearch.New(fmt.Sprintf("http://%s:%s", meili.Host, meili.Port), meilisearch.WithAPIKey("MASTER_KEY"))
	suite.meili = &meiliClient

	ik, err := testhelper.CreateImageKit()
	if err != nil {
		log.Fatalf("Failed to start ImageKit: %s", err)
	}
	suite.ik = ik

	hostImg, err := testhelper.NewImage()
	if err != nil {
		log.Fatalf("Failed to load img: %s", err)
	}
	suite.hostImg = hostImg

	venueImg, err := testhelper.NewImage()
	if err != nil {
		log.Fatalf("Failed to load img: %s", err)
	}
	suite.venueImg = venueImg

	thumbnail, err := testhelper.NewImage()
	if err != nil {
		log.Fatalf("Failed to load img: %s", err)
	}
	suite.thumbnail = thumbnail

	banner, err := testhelper.NewImage()
	if err != nil {
		log.Fatalf("Failed to load img: %s", err)
	}
	suite.banner = banner
}

func (suite *EventTestSuite) TearDownSuite() {
	suite.dbpool.Close()
	suite.hostImg.Close()
	suite.venueImg.Close()
	suite.thumbnail.Close()
	suite.banner.Close()
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func (suite *EventTestSuite) TestCreateEvent() {
	t := suite.T()

	venueParams := venue.NewVenueInput{
		Name:     "Test Venue",
		Capacity: 1000,
		City:     "City",
		State:    "State",
	}
	venue, err := venue.NewVenue(suite.ctx, suite.dbpool, suite.ik.ImageKit, suite.venueImg, &venueParams)
	if err != nil {
		log.Fatalf("Failed to create new venue: %s", err)
	}

	assert.Equal(t, venueParams.Name, venue.Name)
	assert.Equal(t, venueParams.Capacity, venue.Capacity)
	assert.Equal(t, venueParams.City, venue.City)
	assert.Equal(t, venueParams.State, venue.State)

	hostParams := host.NewHostInput{
		Name: "Kivvvi",
		Bio:  "Heelo",
	}

	newHost, err := host.NewHost(suite.ctx, suite.dbpool, suite.ik, suite.hostImg, &hostParams)
	if err != nil {
		log.Fatalf("Failed to create new host: %s", err)
	}

	assert.NoError(t, err)
	assert.NotNil(t, newHost)

	assert.Equal(t, hostParams.Name, newHost.Name)
	assert.Equal(t, hostParams.Bio, newHost.Bio)

	eventParams := event.NewEventInput{
		Name:             "Event A",
		Summary:          "Summary",
		Description:      "Description",
		Available_ticket: 1000,
		Starting_date:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
		Ending_date:      pgtype.Timestamptz{Time: time.Now().Add(5 * time.Hour), Valid: true},
		Venue_id:         venue.VenueID,
		Host_id:          newHost.HostID,
	}

	newEvent, err := event.NewEvent(suite.ctx, suite.dbpool, *suite.meili, suite.ik.ImageKit, suite.thumbnail, suite.banner, &eventParams)
	if err != nil {
		log.Fatalf("Failed to create new event: %s", err)
	}

	assert.Equal(t, newEvent.Name, eventParams.Name)
}

func TestEventSuite(t *testing.T) {
	suite.Run(t, new(EventTestSuite))
}
