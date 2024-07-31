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
	err := s.Add("example.com")
	if err != nil {
		t.Fatalf("Expected an error nil, but got: %v", err)
	}
	if !slices.Equal(s.List(), expected) {
		t.Fatalf("Expected %v, but got: %v", expected, s.List())
	}
}

func TestListReturnsListInSortedOrder(t *testing.T) {
	s := domains.NewStore()
	t.Parallel()
	expected := []string{"example1.com", "example2.com"}
	err := s.Add("example2.com")
	if err != nil {
		t.Fatalf("Expected an error nil, but got: %v", err)
	}
	err = s.Add("example1.com")
	if err != nil {
		t.Fatalf("Expected an error nil, but got: %v", err)
	}
	if !slices.Equal(s.List(), expected) {
		t.Fatalf("Expected %v, but got: %v", expected, s.List())
	}
}

func TestRemoveExistingDomainReturnsExpectedList(t *testing.T) {
	s := domains.NewStore()
	t.Parallel()
	for _, d := range []string{"example3.com", "example2.com", "example1.com"} {
		err := s.Add(d)
		if err != nil {
			t.Fatalf("Expected an error nil, but got: %v", err)
		}
	}

	err := s.Remove("ImNotOnTheList.com")
	if err != nil {
		t.Fatalf("Expected an error nil, but got: %v", err)
	}
	expected := []string{"example1.com", "example2.com", "example3.com"}

	if !slices.Equal(s.List(), expected) {
		t.Fatalf("Expected %v, but got: %v", expected, s.List())
	}

	err = s.Remove("example2.com")
	if err != nil {
		t.Fatalf("Expected an error nil, but got: %v", err)
	}
	expected = []string{"example1.com", "example3.com"}

	if !slices.Equal(s.List(), expected) {
		t.Fatalf("Expected %v, but got: %v", expected, s.List())
	}

	err = s.Remove("example3.com")
	if err != nil {
		t.Fatalf("Expected an error nil, but got: %v", err)
	}
	expected = []string{"example1.com"}

	if !slices.Equal(s.List(), expected) {
		t.Fatalf("Expected %v, but got: %v", expected, s.List())
	}

	err = s.Remove("example1.com")
	if err != nil {
		t.Fatalf("Expected an error nil, but got: %v", err)
	}
	expected = []string{}

	if !slices.Equal(s.List(), expected) {
		t.Fatalf("Expected %v, but got: %v", expected, s.List())
	}

	err = s.Remove("ListShouldBeEmpty.com")
	if err != nil {
		t.Fatalf("Expected an error nil, but got: %v", err)
	}
	expected = []string{}

	if !slices.Equal(s.List(), expected) {
		t.Fatalf("Expected %v, but got: %v", expected, s.List())
	}

}

func TestStoreChangesPersisted(t *testing.T) {
	s := domains.NewStore()
	t.Parallel()
	expected := []string{"example1.com", "example2.com", "example3.com"}

	for _, d := range []string{"example3.com", "example2.com", "example1.com"} {
		err := s.Add(d)
		if err != nil {
			t.Fatalf("Expected an error nil, but got: %v", err)
		}
	}

	s2 := domains.NewStore()
	if !slices.Equal(s2.List(), expected) {
		t.Fatalf("Expected %#v, but got: %#v", expected, s2.List())
	}
}
