package special

func GetSomething() string {
	return "something"
}

type Specila struct {
	Name string
}

func (s *Specila) GetName() string {
	return s.Name
}
