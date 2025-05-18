// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog_test

import (
	"log/slog"
)

// Book represents a book with a title and author.
type Book struct {
	Title  string
	Author string
}

// LogValue implements slog.LogValuer.
// It returns a group containing the fields of
// the Book, so that they appear together in the log output.
func (b *Book) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("title", b.Title),
		slog.String("author", b.Author),
	)
}

// Books is a slice of pointers to Book.
type Books []*Book

// LogValue implements slog.LogValuer for a slice of Books.
// It returns a list of Books, where each Book is logged using
// their respective LogValue method, providing structured
// output for each.
func (bs Books) LogValue() slog.Value {
	return slog.ListValue(bs...)
}

func ExampleLogValuer_books() {
	books := Books{
		{
			Title:  "Principia Mathematica",
			Author: "Isaac Newton",
		},
		{
			Title:  "Relativity: The Special and the General Theory",
			Author: "Albert Einstein",
		},
	}

	slog.Info("book collection", "books", books)

	// JSON Output would look in part like:
	// {
	//     ...
	//     "msg": "book collection",
	//     "books": [
	//         {
	//             "title": "Principia Mathematica",
	//             "author": "Isaac Newton"
	//         },
	//         {
	//             "title": "Relativity: The Special and the General Theory",
	//             "author": "Albert Einstein"
	//         }
	//     ]
	// }
}
