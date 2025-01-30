package host

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/gerry-sheva/tixmaster/pkg/common"
	"github.com/gerry-sheva/tixmaster/pkg/database/sqlc"
	"github.com/imagekit-developer/imagekit-go/api/uploader"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewHost(ctx context.Context, dbpool *pgxpool.Pool, ik common.ImageKit, avatar multipart.File, p *NewHostInput) (sqlc.NewHostRow, error) {
	resp, err := ik.Uploader.Upload(ctx, avatar, uploader.UploadParam{
		FileName: fmt.Sprintf("%s.webp", p.Name),
		Folder:   ik.Dir,
	})
	if err != nil {
		return sqlc.NewHostRow{}, err
	}

	params := sqlc.NewHostParams{
		Name:   p.Name,
		Avatar: resp.Data.Url,
		Bio:    p.Bio,
	}

	host, err := sqlc.New(dbpool).NewHost(ctx, params)
	if err != nil {
		return sqlc.NewHostRow{}, err
	}

	return host, nil
}
