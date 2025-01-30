package event

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/gerry-sheva/tixmaster/pkg/database/sqlc"
	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/api/uploader"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/meilisearch/meilisearch-go"
)

func NewEvent(
	ctx context.Context,
	dbpool *pgxpool.Pool,
	meilisearch meilisearch.ServiceManager,
	ik *imagekit.ImageKit,
	thumbnail multipart.File,
	banner multipart.File,
	p *NewEventInput,
) (sqlc.NewEventRow, error) {
	thumbnailResp, err := ik.Uploader.Upload(ctx, thumbnail, uploader.UploadParam{
		FileName: fmt.Sprintf("%s.webp", p.Name),
	})
	if err != nil {
		return sqlc.NewEventRow{}, err
	}

	bannerResp, err := ik.Uploader.Upload(ctx, banner, uploader.UploadParam{
		FileName: fmt.Sprintf("%s.webp", p.Name),
	})
	if err != nil {
		return sqlc.NewEventRow{}, err
	}

	params := sqlc.NewEventParams{
		Name:            p.Name,
		Summary:         p.Summary,
		Description:     p.Description,
		AvailableTicket: p.Available_ticket,
		Thumbnail:       thumbnailResp.Data.Url,
		Banner:          bannerResp.Data.Url,
		StartingDate:    p.Starting_date,
		EndingDate:      p.Ending_date,
		VenueID:         p.Venue_id,
		HostID:          p.Host_id,
	}

	event, err := sqlc.New(dbpool).NewEvent(ctx, params)
	if err != nil {
		return sqlc.NewEventRow{}, err
	}

	println(event.Name)

	index := meilisearch.Index("event")
	_, err = index.AddDocuments(event)
	if err != nil {
		println("IS IT HERE?")
		return sqlc.NewEventRow{}, err
	}

	return event, nil
}
