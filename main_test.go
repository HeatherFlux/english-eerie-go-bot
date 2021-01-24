package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_getBooks(t *testing.T) {
	ww := httptest.NewRecorder()

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				w: ww,
				r: httptest.NewRequest("GET", "/api/books", nil),
			},
			want: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getBooks(tt.args.w, tt.args.r)
			assert.Equal(t, tt.want, ww.Code)
		})
	}
}

func Test_getBook(t *testing.T) {
	books = append(books, Book{ID: "1", Isbn: "meh", Title: "B1", Author: &Author{Firstname: "John", Lastname: "Doe"}})

	req := httptest.NewRequest("GET", "/api/book/id", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	type args struct {
		w httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				w: *httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/api/book/bad", nil),
			},
			want: http.StatusNotFound,
		},
		{
			name: "test 1",
			args: args{
				w: *httptest.NewRecorder(),
				r: req,
			},
			want: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getBook(&tt.args.w, tt.args.r)
			assert.Equal(t, tt.want, tt.args.w.Code)
		})
	}
}

func Test_deleteBook(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteBook(tt.args.w, tt.args.r)
		})
	}
}

func Test_editBook(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			editBook(tt.args.w, tt.args.r)
		})
	}
}

func Test_createBook(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createBook(tt.args.w, tt.args.r)
		})
	}
}

func Test_main(t *testing.T) {
	type args struct {
		r mux.Router
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
