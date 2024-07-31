package domains

import (
	"os"
	"slices"
	"strings"

	"golang.org/x/exp/maps"
)

var storefile string = "./store.txt"

type store struct {
	domains map[string]struct{}
}

func (s *store) Add(domain string) error {
	s.domains[domain] = struct{}{}
	return os.WriteFile(storefile, []byte(strings.Join(s.List(), "\n")), 0o644)
}

func (s store) List() []string {
	keys := maps.Keys(s.domains)
	slices.Sort(keys)
	return keys
}

func NewStore() *store {
	var s store
	s.domains = map[string]struct{}{}
	bytes, err := os.ReadFile(storefile)
	if err != nil {
		return &s
	}
	for _, domain := range strings.Split(string(bytes), "\n") {
		s.domains[domain] = struct{}{}
	}
	return &s
}

func (s *store) Remove(domain string) error {
	delete(s.domains, domain)
	return os.WriteFile(storefile, []byte(strings.Join(s.List(), "\n")), 0o644)
}
