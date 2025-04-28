package service

import "fmt"

// ItineraryService defines the interface for reconstructing flight itineraries.
type ItineraryService interface {
	// Reconstruct takes a slice of [source,destination] pairs and returns
	// the ordered list of airport codes or an error.
	Reconstruct(tickets [][2]string) ([]string, error)
}

type itineraryService struct{}

// NewItineraryService constructs a concrete ItineraryService.
func NewItineraryService() ItineraryService {
	return &itineraryService{}
}

// Reconstruct implements the ticket‐to‐itinerary logic.
func (s *itineraryService) Reconstruct(tickets [][2]string) ([]string, error) {
	if len(tickets) == 0 {
		return nil, fmt.Errorf("no tickets provided")
	}

	destMap := make(map[string]string)
	destSet := make(map[string]bool)

	for _, t := range tickets {
		src, dst := t[0], t[1]
		if src == "" || dst == "" {
			return nil, fmt.Errorf("invalid ticket: empty source or destination")
		}
		if _, dup := destMap[src]; dup {
			return nil, fmt.Errorf("duplicate source: %s", src)
		}
		destMap[src] = dst
		destSet[dst] = true
	}

	// find starting point (source not in any destination)
	var start string
	for src := range destMap {
		if !destSet[src] {
			start = src
			break
		}
	}
	if start == "" {
		return nil, fmt.Errorf("no valid starting point found")
	}

	// build the itinerary
	itinerary := []string{start}
	for {
		next, ok := destMap[start]
		if !ok {
			break
		}
		itinerary = append(itinerary, next)
		start = next
	}

	return itinerary, nil
}
