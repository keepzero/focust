package focust

type Module struct {
	Name     string
	handlers map[int]func(unique string, requestStr string) (string, error)
	ids      [2]int
}

type ModuleInterface interface {
	GetName() string
	GetHandler(int) func(string, string) (string, error)
}
