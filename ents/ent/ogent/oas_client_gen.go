// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = bits.LeadingZeros64
	_ = big.Rat{}
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = attribute.KeyValue{}
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
	_ = codes.Unset
)

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	cfg       config
	requests  metric.Int64Counter
	errors    metric.Int64Counter
	duration  metric.Int64Histogram
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...Option) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	c := &Client{
		cfg:       newConfig(opts...),
		serverURL: u,
	}
	if c.requests, err = c.cfg.Meter.NewInt64Counter(otelogen.ClientRequestCount); err != nil {
		return nil, err
	}
	if c.errors, err = c.cfg.Meter.NewInt64Counter(otelogen.ClientErrorsCount); err != nil {
		return nil, err
	}
	if c.duration, err = c.cfg.Meter.NewInt64Histogram(otelogen.ClientDuration); err != nil {
		return nil, err
	}
	return c, nil
}

// CreateTodo invokes createTodo operation.
//
// Creates a new Todo and persists it to storage.
//
// POST /todos
func (c *Client) CreateTodo(ctx context.Context, request CreateTodoReq) (res CreateTodoRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createTodo"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "CreateTodo",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1, otelAttrs...)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
		}
		span.End()
	}()
	c.requests.Add(ctx, 1, otelAttrs...)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodeCreateTodoRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer jx.PutWriter(buf)
	reqBody = bytes.NewReader(buf.Buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/todos"

	r := ht.NewRequest(ctx, "POST", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeCreateTodoResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DeleteTodo invokes deleteTodo operation.
//
// Deletes the Todo with the requested ID.
//
// DELETE /todos/{id}
func (c *Client) DeleteTodo(ctx context.Context, params DeleteTodoParams) (res DeleteTodoRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteTodo"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "DeleteTodo",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1, otelAttrs...)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
		}
		span.End()
	}()
	c.requests.Add(ctx, 1, otelAttrs...)
	u := uri.Clone(c.serverURL)
	u.Path += "/todos/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}

	r := ht.NewRequest(ctx, "DELETE", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeDeleteTodoResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListTodo invokes listTodo operation.
//
// List Todos.
//
// GET /todos
func (c *Client) ListTodo(ctx context.Context, params ListTodoParams) (res ListTodoRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listTodo"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "ListTodo",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1, otelAttrs...)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
		}
		span.End()
	}()
	c.requests.Add(ctx, 1, otelAttrs...)
	u := uri.Clone(c.serverURL)
	u.Path += "/todos"

	q := u.Query()
	{
		// Encode "page" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			if val, ok := params.Page.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["page"] = e.Result()
	}
	{
		// Encode "itemsPerPage" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			if val, ok := params.ItemsPerPage.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["itemsPerPage"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeListTodoResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ReadTodo invokes readTodo operation.
//
// Finds the Todo with the requested ID and returns it.
//
// GET /todos/{id}
func (c *Client) ReadTodo(ctx context.Context, params ReadTodoParams) (res ReadTodoRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("readTodo"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "ReadTodo",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1, otelAttrs...)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
		}
		span.End()
	}()
	c.requests.Add(ctx, 1, otelAttrs...)
	u := uri.Clone(c.serverURL)
	u.Path += "/todos/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeReadTodoResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UpdateTodo invokes updateTodo operation.
//
// Updates a Todo and persists changes to storage.
//
// PATCH /todos/{id}
func (c *Client) UpdateTodo(ctx context.Context, request UpdateTodoReq, params UpdateTodoParams) (res UpdateTodoRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateTodo"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "UpdateTodo",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1, otelAttrs...)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
		}
		span.End()
	}()
	c.requests.Add(ctx, 1, otelAttrs...)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodeUpdateTodoRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer jx.PutWriter(buf)
	reqBody = bytes.NewReader(buf.Buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/todos/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}

	r := ht.NewRequest(ctx, "PATCH", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeUpdateTodoResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
