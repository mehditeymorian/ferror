package ferror

import "github.com/gofiber/fiber/v2"

type Option func(options *Options)

type Options struct {
	DevMode         bool
	OnErrorHandling func(ctx *fiber.Ctx, err Error)
}

func DevelopmentMode(devMode bool) Option {
	return func(options *Options) {
		options.DevMode = devMode
	}
}

func OnErrorHandling(handler func(ctx *fiber.Ctx, err Error)) Option {
	return func(options *Options) {
		options.OnErrorHandling = handler
	}
}
