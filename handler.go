package ferror

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorHandler struct {
	devMode         bool
	onErrorHandling func(ctx *fiber.Ctx, err Error)
}

func NewErrorHandler(options ...Option) *ErrorHandler {
	opts := new(Options)

	for _, option := range options {
		option(opts)
	}

	return &ErrorHandler{
		devMode:         opts.DevMode,
		onErrorHandling: opts.OnErrorHandling,
	}
}

// Error returns an error in client response.
func (e *ErrorHandler) Error(ctx *fiber.Ctx, err Error) error {
	if e.onErrorHandling != nil {
		e.onErrorHandling(ctx, err)
	}

	data := fiber.Map{
		"status":  err.FiberError.Message,
		"message": err.Message,
		"error":   err.Cause.Error(),
	}

	// add extra values if extra is valid
	for key, value := range err.Extra {
		if isKeyInvalid(key) {
			continue
		}
		data[key] = value
	}

	if e.devMode {
		return ctx.Status(err.FiberError.Code).JSON(data)
	}

	// no error in production environment!
	delete(data, "error")

	return ctx.Status(err.FiberError.Code).JSON(data)
}

func (e *ErrorHandler) BadRequest(ctx *fiber.Ctx, err error, message string) error {
	return e.Error(ctx, Error{
		FiberError: fiber.ErrBadRequest,
		Cause:      err,
		Message:    message,
	})
}

func (e *ErrorHandler) InternalServerError(ctx *fiber.Ctx, err error, message string) error {
	return e.Error(ctx, Error{
		FiberError: fiber.ErrInternalServerError,
		Cause:      err,
		Message:    message,
	})
}

func (e *ErrorHandler) Unauthorized(ctx *fiber.Ctx, err error, message string) error {
	return e.Error(ctx, Error{
		FiberError: fiber.ErrUnauthorized,
		Cause:      err,
		Message:    message,
	})
}

func (e *ErrorHandler) Forbidden(ctx *fiber.Ctx, err error, message string) error {
	return e.Error(ctx, Error{
		FiberError: fiber.ErrForbidden,
		Cause:      err,
		Message:    message,
	})
}

func (e *ErrorHandler) NotFound(ctx *fiber.Ctx, err error, message string) error {
	return e.Error(ctx, Error{
		FiberError: fiber.ErrNotFound,
		Cause:      err,
		Message:    message,
	})
}

func (e *ErrorHandler) Conflict(ctx *fiber.Ctx, err error, message string) error {
	return e.Error(ctx, Error{
		FiberError: fiber.ErrConflict,
		Cause:      err,
		Message:    message,
	})
}

func (e *ErrorHandler) TooManyRequest(ctx *fiber.Ctx, err error, message string) error {
	return e.Error(ctx, Error{
		FiberError: fiber.ErrTooManyRequests,
		Cause:      err,
		Message:    message,
	})
}

func (e *ErrorHandler) BadGateway(ctx *fiber.Ctx, err error, message string) error {
	return e.Error(ctx, Error{
		FiberError: fiber.ErrBadGateway,
		Cause:      err,
		Message:    message,
	})
}

func (e *ErrorHandler) GatewayTimeout(ctx *fiber.Ctx, err error, message string) error {
	return e.Error(ctx, Error{
		FiberError: fiber.ErrGatewayTimeout,
		Cause:      err,
		Message:    message,
	})
}
