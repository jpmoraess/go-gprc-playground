package application

import "fmt"

type HelloService struct{}

func (a *HelloService) GenerateHello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
