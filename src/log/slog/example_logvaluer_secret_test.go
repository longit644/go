// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog_test

import (
	"log/slog"
	"log/slog/internal/slogtest"
	"os"
)

// A token is a secret value that grants permissions.
type Token string

// LogValue implements slog.LogValuer.
// It avoids revealing the token.
func (Token) LogValue() slog.Value {
	return slog.StringValue("REDACTED_TOKEN")
}

// Tokens represents a list of tokens.
type Tokens []Token

// LogValue implements slog.LogValuer.
// It avoids revealing the actual tokens.
func (ts Tokens) LogValue() slog.Value {
	return slog.ListValue(ts...)
}

// This example demonstrates a Value that replaces itself
// with an alternative representation to avoid revealing secrets.
func ExampleLogValuer_secret() {
	t := Token("shhhh!")
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{ReplaceAttr: slogtest.RemoveTime}))
	logger.Info("permission granted", "user", "Perry", "token", t)

	ts1 := []Token{"shhhh!", "whissss!"}
	logger.Info("user session info", "user", "Perry", "tokens", slog.ListValue(ts1...))

	ts2 := Tokens{"zhhhh!", "whizzzz!"}
	logger.Info("user session info", "user", "Heinz", "tokens", ts2)

	// Output:
	// level=INFO msg="permission granted" user=Perry token=REDACTED_TOKEN
	// level=INFO msg="user session info" user=Perry tokens[0]=REDACTED_TOKEN tokens[1]=REDACTED_TOKEN
	// level=INFO msg="user session info" user=Heinz tokens[0]=REDACTED_TOKEN tokens[1]=REDACTED_TOKEN
}
