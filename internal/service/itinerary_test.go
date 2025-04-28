package service

import (
	"reflect"
	"testing"
)

func TestReconstruct_Success(t *testing.T) {
	svc := NewItineraryService()
	tickets := [][2]string{
		{"LAX", "DXB"},
		{"JFK", "LAX"},
		{"SFO", "SJC"},
		{"DXB", "SFO"},
	}
	want := []string{"JFK", "LAX", "DXB", "SFO", "SJC"}

	got, err := svc.Reconstruct(tickets)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReconstruct_NoTickets(t *testing.T) {
	svc := NewItineraryService()
	if _, err := svc.Reconstruct([][2]string{}); err == nil {
		t.Error("expected error for empty tickets, got nil")
	}
}

func TestReconstruct_DuplicateSource(t *testing.T) {
	svc := NewItineraryService()
	tickets := [][2]string{{"A", "B"}, {"A", "C"}}
	if _, err := svc.Reconstruct(tickets); err == nil {
		t.Error("expected error for duplicate source, got nil")
	}
}

func TestReconstruct_Cycle(t *testing.T) {
	svc := NewItineraryService()
	tickets := [][2]string{{"A", "B"}, {"B", "A"}}
	if _, err := svc.Reconstruct(tickets); err == nil {
		t.Error("expected error for cycle, got nil")
	}
}
