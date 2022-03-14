package translator


type Translator interface {
	Client
}

type transaltor struct {
	Client
}

type Client interface {
	Do(string) (string, error)
}

func New(c Client) Translator {
	return &transaltor{c}
}
