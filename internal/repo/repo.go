package repo

import (
	"context"
	"fmt"
	"github.com/abobacode/endpoint/internal/models"
	"github.com/abobacode/endpoint/pkg/database"

	builder "github.com/doug-martin/goqu/v9"
)

type Repo struct {
	database database.Pool
}

const (
	StatusUploaded int = iota
	StatusInProgress
	StatusError
	msgURL       = "The link to the stream has been changed"
	msgBlock     = "There was any change to the video block"
	deletedBlock = "The video has been deleted"
)

func (r *Repo) Videos(ctx context.Context) ([]models.Video, error) {
	var videos []models.Video

	err := r.database.Builder().
		Select(builder.L("*")).
		From("video").
		ScanStructsContext(ctx, &videos)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (r *Repo) Ads(ctx context.Context) ([]models.AdsStruct, error) {
	var ads []models.AdsStruct

	err := r.database.Builder().
		Select(builder.L("*")).
		From("ads").
		ScanStructsContext(ctx, &ads)
	if err != nil {
		return nil, err
	}

	return ads, nil
}

func (r *Repo) Save(ctx context.Context, videos []models.Video) error {
	_, err := r.database.Builder().
		Insert("video").
		Rows(videos).
		Executor().
		ExecContext(ctx)
	if err != nil {
		return err
	}
	return err
}

func (r *Repo) SaveAds(ctx context.Context, ads []models.AdsStruct) error {
	_, err := r.database.Builder().
		Insert("ads").
		Rows(ads).
		Executor().
		ExecContext(ctx)
	if err != nil {
		return err
	}
	return err
}

func (r *Repo) SaveM3u8URL(ctx context.Context, m3u8URL string, fileID int) error {
	_, err := r.database.Builder().
		Update("video").
		Set(builder.Record{"m3u8_url": m3u8URL}).
		Where(builder.C("video_id").Eq(fileID)).
		Executor().
		ExecContext(ctx)
	if err != nil {
		return err
	}
	return err
}

func (r *Repo) UpdateStatus(ctx context.Context, id int, status int) error {
	var statusValue string

	switch status {
	case StatusUploaded:
		statusValue = "Uploaded"
	case StatusInProgress:
		statusValue = "InProgress"
	case StatusError:
		statusValue = "Error"
	default:
		return fmt.Errorf("недопустимый статус")
	}

	_, err := r.database.Builder().
		Update("video").
		Set(builder.Record{"status": statusValue}).
		Where(builder.C("video_id").Eq(id)).
		Executor().
		ExecContext(ctx)
	if err != nil {
		return err
	}

	return err
}

func (r *Repo) CheckEtagInDataBase(ctx context.Context, etag string) (bool, error) {
	var count int

	_, err := r.database.Builder().
		Select(builder.COUNT("*")).
		From("video").
		Where(builder.C("etag").Eq(etag)).
		ScanValContext(ctx, &count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *Repo) CheckEtagInDataBaseAds(ctx context.Context, etag string) (bool, error) {
	var count int

	_, err := r.database.Builder().
		Select(builder.COUNT("*")).
		From("ads").
		Where(builder.C("etag").Eq(etag)).
		ScanValContext(ctx, &count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *Repo) CheckVideoID(ctx context.Context, videoID int) (bool, error) {
	var count int

	_, err := r.database.Builder().
		Select(builder.COUNT("*")).
		From("video").
		Where(builder.C("video_id").Eq(videoID)).
		ScanValContext(ctx, &count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *Repo) DeletedBlock(ctx context.Context, block models.VideoDelete, videoID int) error {
	_, err := r.database.Builder().
		Update("video").
		Set(builder.Record{
			"status":             block.Status,
			"status_description": deletedBlock,
		}).
		Where(builder.C("video_id").Eq(videoID)).
		Executor().
		ExecContext(ctx)
	if err != nil {
		return err
	}
	return err
}

func (r *Repo) ModifyBlock(ctx context.Context, block models.VideoModifiedURL, videoId int) error {
	_, err := r.database.Builder().
		Update("video").
		Set(builder.Record{
			"url":                block.URL,
			"etag":               block.Etag,
			"status":             block.Status,
			"status_description": msgURL,
		}).
		Where(builder.C("video_id").Eq(videoId)).
		Executor().
		ExecContext(ctx)
	if err != nil {
		return err
	}
	return err
}

func (r *Repo) ModifyBlockAll(ctx context.Context, block models.VideoModifiedAll, videoId int) error {
	_, err := r.database.Builder().
		Update("video").
		Set(builder.Record{
			"project_id":          block.ProjectID,
			"title":               block.Title,
			"title_project":       block.TitleProject,
			"description":         block.Description,
			"description_project": block.DescriptionProject,
			"categories":          block.Categories,
			"duration":            block.Duration,
			"views_count":         block.ViewsCount,
			"poster_url":          block.PosterURL,
			"poster_sqr_url":      block.PosterSqrURL,
			"aired_at":            block.AiredAt,
			"video_type":          block.VideoType,
			"age_restriction":     block.AgeRestriction,
			"format":              block.Format,
			"src":                 block.Src,
			"type":                block.Type,
			"status":              block.Status,
			"status_description":  msgBlock,
			"m3u8_url":            block.M3u8URL,
			"updated_at":          block.UpdatedAt,
		}).
		Where(builder.C("video_id").Eq(videoId)).
		Executor().
		ExecContext(ctx)
	if err != nil {
		return err
	}
	return err
}

func (r *Repo) CheckDeletedVideos(ctx context.Context) ([]models.Video, error) {
	var deleted []models.Video

	d := "Deleted"

	err := r.database.Builder().
		Select(builder.L("*")).
		From("video").
		Where(builder.C("status").Eq(d)).
		ScanStructsContext(ctx, &deleted)
	if err != nil {
		return nil, err
	}
	return deleted, nil
}

func (r *Repo) FetchAdsBlock(ctx context.Context, id int) ([]models.AdsStruct, error) {
	var ads []models.AdsStruct

	err := r.database.Builder().
		Select(builder.L("*")).
		From("ads").
		Where(builder.C("video_id").Eq(id)).
		ScanStructsContext(ctx, &ads)
	if err != nil {
		return nil, err
	}

	return ads, nil
}

func New(db database.Pool) *Repo {
	return &Repo{
		database: db,
	}
}
