// Code generated by ogen, DO NOT EDIT.

package filmlibAPI

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"
	"github.com/google/uuid"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// DeleteActorByIDParams is parameters of deleteActorByID operation.
type DeleteActorByIDParams struct {
	// ID актёра.
	ActorID uuid.UUID
}

func unpackDeleteActorByIDParams(packed middleware.Parameters) (params DeleteActorByIDParams) {
	{
		key := middleware.ParameterKey{
			Name: "actor_id",
			In:   "path",
		}
		params.ActorID = packed[key].(uuid.UUID)
	}
	return params
}

func decodeDeleteActorByIDParams(args [1]string, argsEscaped bool, r *http.Request) (params DeleteActorByIDParams, _ error) {
	// Decode path: actor_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "actor_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUUID(val)
				if err != nil {
					return err
				}

				params.ActorID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "actor_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetMoviesParams is parameters of getMovies operation.
type GetMoviesParams struct {
	// Поле сортировки.
	SortBy OptGetMoviesSortBy
	// Порядок сортировки.
	Order OptGetMoviesOrder
}

func unpackGetMoviesParams(packed middleware.Parameters) (params GetMoviesParams) {
	{
		key := middleware.ParameterKey{
			Name: "sort_by",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.SortBy = v.(OptGetMoviesSortBy)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "order",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Order = v.(OptGetMoviesOrder)
		}
	}
	return params
}

func decodeGetMoviesParams(args [0]string, argsEscaped bool, r *http.Request) (params GetMoviesParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: sort_by.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "sort_by",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotSortByVal GetMoviesSortBy
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotSortByVal = GetMoviesSortBy(c)
					return nil
				}(); err != nil {
					return err
				}
				params.SortBy.SetTo(paramsDotSortByVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.SortBy.Get(); ok {
					if err := func() error {
						if err := value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "sort_by",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: order.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "order",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotOrderVal GetMoviesOrder
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotOrderVal = GetMoviesOrder(c)
					return nil
				}(); err != nil {
					return err
				}
				params.Order.SetTo(paramsDotOrderVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Order.Get(); ok {
					if err := func() error {
						if err := value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "order",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// MoviesMovieIDDeleteParams is parameters of DELETE /movies/{movie_id} operation.
type MoviesMovieIDDeleteParams struct {
	// ID фильма.
	MovieID uuid.UUID
}

func unpackMoviesMovieIDDeleteParams(packed middleware.Parameters) (params MoviesMovieIDDeleteParams) {
	{
		key := middleware.ParameterKey{
			Name: "movie_id",
			In:   "path",
		}
		params.MovieID = packed[key].(uuid.UUID)
	}
	return params
}

func decodeMoviesMovieIDDeleteParams(args [1]string, argsEscaped bool, r *http.Request) (params MoviesMovieIDDeleteParams, _ error) {
	// Decode path: movie_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "movie_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUUID(val)
				if err != nil {
					return err
				}

				params.MovieID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "movie_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// MoviesMovieIDPatchParams is parameters of PATCH /movies/{movie_id} operation.
type MoviesMovieIDPatchParams struct {
	// ID фильма.
	MovieID uuid.UUID
}

func unpackMoviesMovieIDPatchParams(packed middleware.Parameters) (params MoviesMovieIDPatchParams) {
	{
		key := middleware.ParameterKey{
			Name: "movie_id",
			In:   "path",
		}
		params.MovieID = packed[key].(uuid.UUID)
	}
	return params
}

func decodeMoviesMovieIDPatchParams(args [1]string, argsEscaped bool, r *http.Request) (params MoviesMovieIDPatchParams, _ error) {
	// Decode path: movie_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "movie_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUUID(val)
				if err != nil {
					return err
				}

				params.MovieID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "movie_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// PatchActorByIdParams is parameters of patchActorById operation.
type PatchActorByIdParams struct {
	// ID актёра.
	ActorID uuid.UUID
}

func unpackPatchActorByIdParams(packed middleware.Parameters) (params PatchActorByIdParams) {
	{
		key := middleware.ParameterKey{
			Name: "actor_id",
			In:   "path",
		}
		params.ActorID = packed[key].(uuid.UUID)
	}
	return params
}

func decodePatchActorByIdParams(args [1]string, argsEscaped bool, r *http.Request) (params PatchActorByIdParams, _ error) {
	// Decode path: actor_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "actor_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUUID(val)
				if err != nil {
					return err
				}

				params.ActorID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "actor_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}
