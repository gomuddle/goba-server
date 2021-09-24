package goba_server

import (
	"github.com/gomuddle/goba"
	"github.com/valyala/fasthttp"
	"net/http"
)

// handlerFunc is the type of any endpoint handler func.
type handlerFunc func(ctx *fasthttp.RequestCtx) error

// handle wraps the given handler.
func (s Server) handle(handler handlerFunc) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		globalOptionsHandler()(ctx)
		if err := handler(ctx); err != nil {
			s.error(ctx, err, http.StatusInternalServerError)
		}
	}
}

// getImage retrieves the image matching the
// given parameters and sends it to the client.
func (s Server) getImage() handlerFunc {
	type response goba.Image
	return func(ctx *fasthttp.RequestCtx) error {
		typ, name := pathParameter(ctx, "type"), pathParameter(ctx, "name")
		image, err := s.goba.FindImage(goba.DatabaseType(typ), name)
		if err != nil {
			return err
		}
		return s.writeJSON(ctx, response(*image))
	}
}

// getAllImages retrieves all images with the
// given type and sends them to the client.
func (s Server) getAllImages() handlerFunc {
	type response []goba.Image
	return func(ctx *fasthttp.RequestCtx) error {
		typ := pathParameter(ctx, "type")
		images, err := s.goba.AllImages(goba.DatabaseType(typ))
		if err != nil {
			return err
		}
		return s.writeJSON(ctx, response(images))
	}
}

// applyImage applies the image matching the given
// parameters to the corresponding database.
func (s Server) applyImage() handlerFunc {
	return func(ctx *fasthttp.RequestCtx) error {
		typ, name := pathParameter(ctx, "type"), pathParameter(ctx, "name")
		return s.goba.ApplyImage(goba.DatabaseType(typ), name)
	}
}

// createImage creates an image of the database
// with the given type and sends it to the client.
func (s Server) createImage() handlerFunc {
	type response goba.Image
	return func(ctx *fasthttp.RequestCtx) error {
		typ := pathParameter(ctx, "type")
		image, err := s.goba.CreateImage(goba.DatabaseType(typ))
		if err != nil {
			return err
		}
		return s.writeJSON(ctx, response(*image))
	}
}

// deleteImage deletes the image with the given parameters.
func (s Server) deleteImage() handlerFunc {
	return func(ctx *fasthttp.RequestCtx) error {
		typ, name := pathParameter(ctx, "type"), pathParameter(ctx, "name")
		return s.goba.DeleteImage(goba.DatabaseType(typ), name)
	}
}

// pathParameter returns the value associated with the given key in the request's path.
func pathParameter(ctx *fasthttp.RequestCtx, key string) string {
	return ctx.UserValue(key).(string)
}
