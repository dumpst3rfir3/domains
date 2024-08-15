package domains_test

import (
	"slices"
	"testing"

	"domains"
)

func TestListEmptyStoreGivesNoResults(t *testing.T) {
	t.Parallel()
	s := justGiveMeSomeTmpStore(t)
	if len(s.List()) != 0 {
		t.Fatalf("Domain list should be empty, but has: %v", s.List())
	}
}

func TestListStoreWithOneDomainGivesExpectedResult(t *testing.T) {
	t.Parallel()
	s := justGiveMeSomeTmpStore(t)
	expected := []string{"example.com"}
	s.Add("example.com")
	if !slices.Equal(s.List(), expected) {
		t.Fatalf("Expected %v, but got: %v", expected, s.List())
	}
}

func TestListReturnsListInSortedOrder(t *testing.T) {
	t.Parallel()
	s := justGiveMeSomeTmpStore(t)
	expected := []string{"example1.com", "example2.com"}
	s.Add("example2.com")
	s.Add("example1.com")
	if !slices.Equal(s.List(), expected) {
		t.Fatalf("Expected %v, but got: %v", expected, s.List())
	}
}

func TestRemoveExistingDomainReturnsExpectedList(t *testing.T) {
	t.Parallel()
	s := justGiveMeSomeTmpStore(t)
	s.Add("example1.com")
	s.Add("example2.com")
	s.Add("example3.com")
	s.Remove("example1.com")
	want := []string{"example2.com", "example3.com"}
	got := s.List()
	if !slices.Equal(got, want) {
		t.Fatalf("want %v, but got: %v", want, got)
	}
}

func TestRemoveNonexistentDomainHasNoEffect(t *testing.T) {
	t.Parallel()
	s := justGiveMeSomeTmpStore(t)
	s.Remove("doesntexist.com")
	if len(s.List()) != 0 {
		t.Errorf("want empty store, got %#v", s.List())
	}
}

func TestSavePersistsChangesToDisk(t *testing.T) {
	t.Parallel()
	storePath := t.TempDir() + "/tmp_store.txt"
	s := domains.OpenStore("/foo/bar/baz.txt")
	s.Add("example1.com")
	s.Add("example2.com")
	s.Add("example3.com")
	s.Save()
	s2 := domains.OpenStore(storePath)
	expected := []string{"example1.com", "example2.com", "example3.com"}
	if !slices.Equal(s2.List(), expected) {
		t.Fatalf("Expected %#v, but got: %#v", expected, s2.List())
	}
}

func justGiveMeSomeTmpStore(t *testing.T) *domains.Store {
	tmpDir := t.TempDir()
	s, err := domains.OpenStore(tmpDir + "/tmp_store.txt")
	if err != nil {
		t.Fatal(err)
	}
	return s
}
