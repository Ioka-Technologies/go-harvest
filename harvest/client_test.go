package harvest

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClientService_List(t *testing.T) {
	service, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/clients", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{})
		testBody(t, r, "client/list/body_1.json")
		testWriteResponse(t, w, "client/list/response_1.json")
	})

	clientList, _, err := service.Client.List(context.Background(), &ClientListOptions{})
	assert.NoError(t, err)

	createdOne := time.Date(
		2018, 1, 31, 20, 34, 30, 0, time.UTC)
	updatedOne := time.Date(
		2018, 5, 31, 21, 34, 30, 0, time.UTC)
	createdTwo := time.Date(
		2018, 3, 2, 10, 12, 13, 0, time.UTC)
	updatedTwo := time.Date(
		2018, 4, 30, 12, 13, 14, 0, time.UTC)

	want := &ClientList{
		Clients: []*Client{
			{
				Id:        Int64(1),
				Name:      String("Client 1"),
				IsActive:  Bool(true),
				Address:   String("Address line 1"),
				Currency:  String("EUR"),
				CreatedAt: &createdOne,
				UpdatedAt: &updatedOne,
			}, {
				Id:        Int64(2),
				Name:      String("Client 2"),
				IsActive:  Bool(false),
				Address:   String("Address line 2"),
				Currency:  String("EUR"),
				CreatedAt: &createdTwo,
				UpdatedAt: &updatedTwo,
			}},
		Pagination: Pagination{
			PerPage:      Int(100),
			TotalPages:   Int(1),
			TotalEntries: Int(2),
			NextPage:     nil,
			PreviousPage: nil,
			Page:         Int(1),
			Links: &PageLinks{
				First:    String("https://api.harvestapp.com/v2/clients?page=1&per_page=100"),
				Next:     nil,
				Previous: nil,
				Last:     String("https://api.harvestapp.com/v2/clients?page=1&per_page=100"),
			},
		},
	}

	assert.Equal(t, want, clientList)
}

func TestClientService_Get(t *testing.T) {
	service, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/clients/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{})
		testBody(t, r, "client/get/body_1.json")
		testWriteResponse(t, w, "client/get/response_1.json")
	})

	client, _, err := service.Client.Get(context.Background(), 1)
	assert.NoError(t, err)

	createdOne := time.Date(
		2018, 1, 31, 20, 34, 30, 0, time.UTC)
	updatedOne := time.Date(
		2018, 5, 31, 21, 34, 30, 0, time.UTC)

	want := &Client{
		Id:        Int64(1),
		Name:      String("Client 1"),
		IsActive:  Bool(true),
		Address:   String("Address line 1"),
		Currency:  String("EUR"),
		CreatedAt: &createdOne,
		UpdatedAt: &updatedOne,
	}

	assert.Equal(t, want, client)
}

func TestClientService_CreateClient(t *testing.T) {
	service, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/clients", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{})
		testBody(t, r, "client/create/body_1.json")
		testWriteResponse(t, w, "client/create/response_1.json")
	})

	client, _, err := service.Client.Create(context.Background(), &ClientCreateRequest{
		Name:     String("Client new"),
		IsActive: Bool(true),
		Address:  String("Address line 1"),
		Currency: String("EUR"),
	})
	assert.NoError(t, err)

	createdOne := time.Date(
		2018, 1, 31, 20, 34, 30, 0, time.UTC)
	updatedOne := time.Date(
		2018, 5, 31, 21, 34, 30, 0, time.UTC)

	want := &Client{
		Id:        Int64(1),
		Name:      String("Client 1"),
		IsActive:  Bool(true),
		Address:   String("Address line 1"),
		Currency:  String("EUR"),
		CreatedAt: &createdOne,
		UpdatedAt: &updatedOne,
	}

	assert.Equal(t, want, client)
}

func TestClientService_UpdateClient(t *testing.T) {
	service, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/clients/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		testFormValues(t, r, values{})
		testBody(t, r, "client/update/body_1.json")
		testWriteResponse(t, w, "client/update/response_1.json")
	})

	client, _, err := service.Client.Update(context.Background(), 1, &ClientUpdateRequest{
		Name:     String("Client new"),
		IsActive: Bool(true),
		Address:  String("Address line 1"),
		Currency: String("EUR"),
	})
	if err != nil {
		t.Errorf("UpdateClient returned error: %v", err)
	}

	createdOne := time.Date(
		2018, 1, 31, 20, 34, 30, 0, time.UTC)
	updatedOne := time.Date(
		2018, 5, 31, 21, 34, 30, 0, time.UTC)

	want := &Client{
		Id:        Int64(1),
		Name:      String("Client 1"),
		IsActive:  Bool(true),
		Address:   String("Address line 1"),
		Currency:  String("EUR"),
		CreatedAt: &createdOne,
		UpdatedAt: &updatedOne,
	}

	assert.Equal(t, want, client)
}

func TestClientService_DeleteClient(t *testing.T) {
	service, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/clients/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testFormValues(t, r, values{})
		testBody(t, r, "client/delete/body_1.json")
		testWriteResponse(t, w, "client/delete/response_1.json")
	})

	_, err := service.Client.Delete(context.Background(), 1)
	assert.NoError(t, err)
}
