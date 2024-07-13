package port

type HelloServicePort interface {
	GenerateHello(name string) string
}

type BankServicePort interface {
	GetCurrentBalance(acc string) float64
}
