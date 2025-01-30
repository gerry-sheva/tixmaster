package venue

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/gerry-sheva/tixmaster/pkg/database/sqlc"
	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/api/uploader"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewVenue(ctx context.Context, dbpool *pgxpool.Pool, ik *imagekit.ImageKit, img multipart.File, p *NewVenueInput) (sqlc.NewVenueRow, error) {
	resp, err := ik.Uploader.Upload(ctx, img, uploader.UploadParam{
		FileName: fmt.Sprintf("%s.webp", p.Name),
	})

	if err != nil {
		println("Helloo!")
		return sqlc.NewVenueRow{}, err
	}

	params := sqlc.NewVenueParams{
		Name:     p.Name,
		Capacity: p.Capacity,
		City:     p.City,
		State:    p.State,
		Img:      resp.Data.Url,
	}

	venue, err := sqlc.New(dbpool).NewVenue(ctx, params)
	if err != nil {
		return sqlc.NewVenueRow{}, err
	}

	return venue, nil
}
