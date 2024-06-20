package domains

type store []string

func (s *store) Add(domain string) {
	*s = append(*s, domain)
}

func (s store) List() []string {
	return s
}

func NewStore() *store {
	var s store
	return &s
}
