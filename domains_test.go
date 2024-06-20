package domains_test

import (
	"domains"
	"slices"
	"testing"
)

func TestListEmptyStoreGivesNoResults(t *testing.T) {
	s := domains.NewStore()
	t.Parallel()

	if len(s.List()) != 0 {
		t.Fatalf("Domain list should be empty, but has: %v", s.List())	
	} 
}

func TestListStoreWithOneDomainGivesExpectedResult(t *testing.T) {
	s := domains.NewStore()
	t.Parallel()
	expected := []string{"example.com"}
	s.Add("example.com")
	if !slices.Equal(s.List(), expected) {
		t.Fatalf("Expected %v, but got: %v", expected, s.List())
	} 
}


