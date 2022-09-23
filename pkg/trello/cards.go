package trello

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TrelloCard struct {
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	Desc          string    `json:"desc"`
	LastActivitiy time.Time `json:"dateLastActivity"`
}

func (c *TrelloClient) CreateCard(name, description, listId string) (*TrelloCard, error) {
	cardsUrl := fmt.Sprintf("%s/cards?key=%s&token=%s&name=%s&desc=%s&idList=%s", c.baseUrl, c.key, c.token, name, description, listId)

	resBody, err := httpRequest(http.MethodPost, cardsUrl, nil)

	var card TrelloCard
	if err := json.Unmarshal(resBody, &card); err != nil {
		return nil, err
	}

	return &card, err
}
