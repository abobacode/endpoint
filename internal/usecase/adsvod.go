package usecase

import (
	"context"
	"encoding/json"
	"log"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/urfave/cli/v2"

	"github.com/abobacode/endpoint/internal/models"
	"github.com/abobacode/endpoint/internal/service"
	"github.com/abobacode/endpoint/pkg/signal"
)

const title = "API Vod Ads"

type Response struct {
	Pre   []string
	Mid   []string
	Pause []string
	Post  []string
}

type FetcherAds interface {
	FetchAdsBlock(ctx context.Context, id int) ([]models.AdsStruct, error)
}

type AdsVodCase struct {
	fetcher FetcherAds
}

func NewAdsVod(fetcher FetcherAds) *AdsVodCase {
	return &AdsVodCase{
		fetcher: fetcher,
	}
}

func (a *AdsVodCase) App(appContext context.Context, cli *cli.Context, pudge *service.Pudge) error {
	app := fiber.New(fiber.Config{
		ServerHeader: title,
	})

	app.Get("/ads/block", a.adsHandler(pudge))

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

func (a *AdsVodCase) adsHandler(pudge *service.Pudge) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		adsBlock, err := a.adsHandle(pudge)
		if err != nil {
			return err
		}

		return ctx.JSON(adsBlock)
	}
}

func (a *AdsVodCase) adsHandle(pudge *service.Pudge) (Response, error) {
	block, err := a.fetcher.FetchAdsBlock(pudge.Context(), 243333)
	if err != nil {
		return Response{}, nil
	}

	var adsBlock Response

	for i := 0; i < len(block); i++ {
		adsBlock = Response{
			Pre:   bytesToStrings(block[i].Pre),
			Mid:   bytesToStrings(block[i].Mid),
			Pause: bytesToStrings(block[i].Pause),
			Post:  bytesToStrings(block[i].Post),
		}
	}

	return adsBlock, nil
}

func bytesToStrings(bytes []byte) []string {
	var strArray []string

	if err := json.Unmarshal(bytes, &strArray); err != nil {
		log.Println("Ошибка декодирования JSON:", err)
	}

	return strArray
}
