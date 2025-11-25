package common

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestPotentialOptionGradeFromString(t *testing.T) {
	tests := map[string]PotentialOptionGrade{
		"레어":   PotentialOptionGradeRare,
		"에픽":   PotentialOptionGradeEpic,
		"유니크":  PotentialOptionGradeUnique,
		"레전드리": PotentialOptionGradeLegendary,
	}

	for input, expected := range tests {
		got, err := PotentialOptionGradeFromString(input)
		if err != nil {
			t.Fatalf("unexpected error for %q: %v", input, err)
		}
		if got != expected {
			t.Fatalf("PotentialOptionGradeFromString(%q) = %v, want %v", input, got, expected)
		}
	}

	if _, err := PotentialOptionGradeFromString("invalid"); err == nil {
		t.Fatal("expected error for invalid grade")
	}
}

func TestRemoveQuery(t *testing.T) {
	const raw = "https://example.com/path?query=123&test=abc"
	const expected = "https://example.com/path"

	if got := RemoveQuery(raw); got != expected {
		t.Fatalf("RemoveQuery(%q) = %q, want %q", raw, got, expected)
	}
}

func TestIsEmptyResponse(t *testing.T) {
	body := map[string]any{
		"date":       "2024-01-01",
		"popularity": nil,
	}
	if !IsEmptyResponse(body) {
		t.Fatal("expected empty response to be detected")
	}

	body["popularity"] = 10
	if IsEmptyResponse(body) {
		t.Fatal("expected non-empty response to be detected")
	}
}

func TestToDateStringRespectsMin(t *testing.T) {
	loc := FixedOffsetLocation(540)
	date := Date{Year: 2024, Month: 1, Day: 2}
	min := Date{Year: 2024, Month: 1, Day: 1}

	got, err := ToDateString(date, loc, &min)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "2024-01-02" {
		t.Fatalf("ToDateString returned %q, want %q", got, "2024-01-02")
	}

	beforeMin := Date{Year: 2023, Month: 12, Day: 31}
	if _, err := ToDateString(beforeMin, loc, &min); err == nil {
		t.Fatal("expected error when date is before min")
	}
}

func TestGetProperDefaultDate(t *testing.T) {
	loc := FixedOffsetLocation(540)

	nowAfterUpdate := time.Date(2024, 1, 2, 3, 0, 0, 0, loc)
	latest := GetProperDefaultDate(nowAfterUpdate, loc, LatestAPIUpdateTimeOptions{Hour: 2, Minute: 0})
	if latest.Year != 2024 || latest.Month != 1 || latest.Day != 2 {
		t.Fatalf("unexpected date after update: %+v", latest)
	}

	nowBeforeUpdate := time.Date(2024, 1, 2, 1, 59, 0, 0, loc)
	latest = GetProperDefaultDate(nowBeforeUpdate, loc, LatestAPIUpdateTimeOptions{Hour: 2, Minute: 0})
	if latest.Year != 2024 || latest.Month != 1 || latest.Day != 1 {
		t.Fatalf("unexpected date before update: %+v", latest)
	}
}

func TestClientBuildURL(t *testing.T) {
	client := NewClient("key", "maplestory", FixedOffsetLocation(540))

	got, err := client.buildURL("v1/id", map[string]string{"character_name": "foo", "empty": ""})
	if err != nil {
		t.Fatalf("buildURL returned error: %v", err)
	}

	expectedPrefix := BaseURL + "maplestory/v1/id?"
	if !strings.HasPrefix(got, expectedPrefix) {
		t.Fatalf("url %q does not start with %q", got, expectedPrefix)
	}
	if !strings.Contains(got, "character_name=foo") {
		t.Fatalf("url %q missing query parameter", got)
	}
	if strings.Contains(got, "empty=") {
		t.Fatalf("url %q should not include empty query param", got)
	}
}

func TestClientFetchHandlesErrorEnvelope(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":{"name":"OPENAPI00001","message":"invalid"}}`))
	}))
	defer server.Close()

	client := NewClient("key", "", time.UTC)
	client.SetBase(server.URL + "/")

	err := client.Fetch(context.Background(), "v1/test", nil, nil)
	if err == nil {
		t.Fatal("expected error from API")
	}

	apiErr, ok := err.(*MapleStoryAPIError)
	if !ok {
		t.Fatalf("expected MapleStoryAPIError, got %T", err)
	}
	if apiErr.Code != OpenAPI00001 {
		t.Fatalf("expected code OPENAPI00001, got %s", apiErr.Code)
	}
}
