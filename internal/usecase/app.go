package usecase

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/urfave/cli/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"

	"github.com/abobacode/endpoint/internal/models"
	"github.com/abobacode/endpoint/pkg/signal"
)

const title = "API For Example"

type Fetcher interface {
	FetchDataBlock(ctx context.Context, id int) ([]models.DataStruct, error)
	SaveDataBlock(ctx context.Context, data []models.DataStruct) error
}

type AppCase struct {
	fetcher Fetcher
}

func NewApp(fetcher Fetcher) *AppCase {
	return &AppCase{
		fetcher: fetcher,
	}
}

func (a *AppCase) App(
	point context.Context,
	appContext context.Context,
	cli *cli.Context,
	key *string,
) error {
	app := fiber.New(fiber.Config{
		ServerHeader: title,
	})

	app.Get("/your/endpoint", a.adsHandler(point, key))

	var ln net.Listener
	var err error
	if ln, err = signal.Listener(
		appContext,
		cli.Int("listener"),
		cli.String("bind-socket"),
		cli.String("bind-address"),
	); err != nil {
		return err
	}

	if err = app.Listener(ln); err != nil {
		return err
	}

	return nil
}

func (a *AppCase) adsHandler(point context.Context, key *string) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		adsBlock, err := a.adsHandle(point, key)
		if err != nil {
			return err
		}

		return ctx.JSON(adsBlock)
	}
}

func (a *AppCase) adsHandle(ctx context.Context, key *string) ([]models.DataStruct, error) {
	var (
		blocks []models.DataStruct
	)

	serv, err := youtube.NewService(context.Background(), option.WithAPIKey(*key))
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	call := serv.Videos.List([]string{"snippet"}).Chart("mostPopular").MaxResults(10)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error fetching videos: %v", err)
	}

	for _, item := range response.Items {
		id := item.Id
		name := item.Snippet.Title
		link := fmt.Sprintf("https://www.youtube.com/watch?v=%s", id)

		blocks = append(blocks, models.DataStruct{
			Title: name,
			URL:   link,
		})
	}

	if err := a.fetcher.SaveDataBlock(ctx, blocks); err != nil {
		//return nil, err
	}

	return blocks, nil
}
