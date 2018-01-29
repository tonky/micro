package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/tonky/micro/mock"
)

func TestSiteInfoUrl(t *testing.T) {
	gc := gsmCode("test_gsm_code")
	si := siteId("test_site_id")

	siUrl := Url{gc}.SiteInfo(si)

	if siUrl != "/test_gsm_code/test_site_id/info" {
		t.Error("info uri mismatch: ", siUrl)
	}
}

func TestApiCall(t *testing.T) {
	gc := gsmCode("test_gsm_code")
	si := siteId("test_site_id")
	expected := "test response"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, expected)
	}))

	defer ts.Close()

	a := Api{ts.URL}

	siUrl := Url{gc}.SiteInfo(si)

	got, err := a.Fetch(siUrl)

	if err != nil {
		t.Error("Error calling Api")
	}

	if !reflect.DeepEqual(strings.TrimSpace(string(got)), expected) {
		t.Errorf("api call test: expected '%v', got '%v'", string(expected), string(got))
	}
}

func TestApiFetch(t *testing.T) {
	baseURL, mux, teardown := mock.ServerMock()
	defer teardown()

	expected := "Fine!"
	gc := gsmCode("test_gsm_code")

	mux.HandleFunc(fmt.Sprintf("/%s/health", gc), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(expected))
	})

	a := Api{baseURL}

	u := Url{gc}.Health()

	got, err := a.Fetch(u)

	if err != nil {
		t.Error("Error calling Api")
	}

	if !reflect.DeepEqual(strings.TrimSpace(string(got)), expected) {
		t.Errorf("api call test: expected '%v', got '%v'", string(expected), string(got))
	}
}
