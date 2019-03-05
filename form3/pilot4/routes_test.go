package main_test

import (
	"net/http"
	"testing"
)

func TestCreate(t *testing.T) {
	testCases := []struct {
		desc     string
		path     string
		verb     string
		expected int
	}{
		{
			desc:     "Create a new payment",
			path:     "/payments",
			verb:     http.MethodPost,
			expected: 201,
			/*
				## Notes

				- If your API uses POST to create a resource, be sure to include a Location header in the response that includes the URL of the newly-created resource, along with a 201 status code — that is part of the HTTP standard.
			*/
		},
		{
			desc:     "Create a new payment on a pre-existing ID",
			path:     "/payments/1234-5678-abcd",
			verb:     http.MethodPost,
			expected: 409,
		},
		{
			desc:     "Create a new payment on a non-existent valid ID",
			path:     "/payments/1234-5678-abcd",
			verb:     http.MethodPost,
			expected: 404,
		},
		{
			desc:     "Create a new payment on an invalid ID",
			path:     "/payments/my-payment-id",
			verb:     http.MethodPost,
			expected: 404,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}

func TestRead(t *testing.T) {
	testCases := []struct {
		desc     string
		path     string
		verb     string
		expected int
	}{
		{
			desc:     "Read the entire collection of existing payments",
			path:     "/payments",
			verb:     http.MethodGet,
			expected: 200,
		},
		{
			desc:     "Read a limited collection of existing payments",
			path:     "/payments?offset=2&limit=2",
			verb:     http.MethodGet,
			expected: 200,
		},
		{
			desc:     "Read a single existing payment",
			path:     "/payments/1234-5678-abcd",
			verb:     http.MethodGet,
			expected: 200,
		},
		{
			desc:     "Read a non-existent payment at a valid ID",
			path:     "/payments/1234-5678-abcd",
			verb:     http.MethodGet,
			expected: 404,
		},
		{
			desc:     "Read a non-existent payment at an invalid ID",
			path:     "/payments/my-payment-id",
			verb:     http.MethodGet,
			expected: 404,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}

func TestUpdate(t *testing.T) {
	testCases := []struct {
		desc     string
		path     string
		verb     string
		expected int
	}{
		{
			desc:     "Update all existing payments",
			path:     "/payments",
			verb:     http.MethodPut,
			expected: 405,
		},
		{
			desc:     "Update an existing payment",
			path:     "/payments/1234-5678-abcd",
			verb:     http.MethodPut,
			expected: 204, // update is OK, but response has no body/content
		},
		{
			desc:     "Update a non-existent payment at a valid ID",
			path:     "/payments/1234-5678-abcd",
			verb:     http.MethodPut,
			expected: 404,
		},
		{
			desc:     "Update a non-existent payment at an invalid ID",
			path:     "/payments/1234-5678-abcd",
			verb:     http.MethodPut,
			expected: 404,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}

func TestDelete(t *testing.T) {
	testCases := []struct {
		desc     string
		path     string
		verb     string
		expected int
	}{
		{
			desc:     "Delete all existing payments",
			path:     "/payments",
			verb:     http.MethodDelete,
			expected: 405,
		},
		{
			desc:     "Delete an existing payment at a valid ID",
			path:     "/payments/1234-5678-abcd",
			verb:     http.MethodDelete,
			expected: 200,
		},
		{
			desc:     "Delete a non-existent payment at a valid ID",
			path:     "/payments/1234-5678-abcd",
			verb:     http.MethodDelete,
			expected: 404,
		},
		{
			desc:     "Delete a non-existent payment at an invalid ID",
			path:     "/payments/my-payment-id",
			verb:     http.MethodDelete,
			expected: 404,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}
