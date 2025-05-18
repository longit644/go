// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog_test

import (
	"log/slog"
	"log/slog/internal/slogtest"
	"net/http"
	"os"
	"time"
)

func ExampleGroup() {
	r, _ := http.NewRequest("GET", "localhost", nil)
	// ...

	logger := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			ReplaceAttr: slogtest.RemoveTime,
		}),
	)
	logger.Info("finished",
		slog.Group("req",
			slog.String("method", r.Method),
			slog.String("url", r.URL.String())),
		slog.Int("status", http.StatusOK),
		slog.Duration("duration", time.Second))

	// Output:
	// level=INFO msg=finished req.method=GET req.url=localhost status=200 duration=1s
}

// FieldViolation describes a single bad request field.
type FieldViolation struct {
	Field       string
	Description string
}

// LogValue implements slog.LogValuer.
// It returns a group containing the fields of
// the FieldViolation, so that they appear together in the log output.
func (b *FieldViolation) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("field", b.Field),
		slog.String("description", b.Description),
	)
}

func ExampleList() {
	logger := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			ReplaceAttr: slogtest.RemoveTime,
		}),
	)

	// Consider the following request:
	//
	// 	type SignUpRequest struct {
	// 		Email    string
	// 		Password string
	// 	}

	logger.Info("finished",
		slog.Group("req", slog.String("method", "SignUp")),
		slog.Group("status",
			slog.Int("code", 3),
			slog.String("message", "Validation failed"),
			slog.List(
				"details",
				&FieldViolation{
					Field:       "email",
					Description: "Email address is invalid",
				},
				&FieldViolation{
					Field:       "password",
					Description: "Password is too weak",
				},
			),
		),
		slog.Duration("duration", time.Second),
	)
	// Output:
	// level=INFO msg=finished req.method=SignUp status.code=3 status.message="Validation failed" status.details[0].field=email status.details[0].description="Email address is invalid" status.details[1].field=password status.details[1].description="Password is too weak" duration=1s
}
