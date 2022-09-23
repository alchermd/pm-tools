package trello

type TrelloClient struct {
	baseUrl string
	key     string
	token   string
}

func NewTrelloClient(key, token string) *TrelloClient {
	return &TrelloClient{
		baseUrl: "https://api.trello.com/1",
		key:     key,
		token:   token,
	}
}
