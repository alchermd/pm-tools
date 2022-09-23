package trello

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TrelloBoard struct {
	Name string `json:"name"`
	Id   string `json:"id"`
	Url  string `json:"url"`
}

type TrelloList struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func (c *TrelloClient) GetBoards() ([]TrelloBoard, error) {
	url := fmt.Sprintf("%s/members/me/boards?key=%s&token=%s", c.baseUrl, c.key, c.token)

	resBody, err := httpRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var boards []TrelloBoard
	if err := json.Unmarshal(resBody, &boards); err != nil {
		return nil, err
	}

	return boards, nil
}

func (c *TrelloClient) GetBoardById(id string) (*TrelloBoard, error) {
	boards, err := c.GetBoards()

	if err != nil {
		return nil, err
	}

	for _, board := range boards {
		if board.Id == id {
			return &board, nil
		}
	}

	return nil, ErrTrelloBoardNotFound
}

func (c *TrelloClient) GetLists(b *TrelloBoard) ([]TrelloList, error) {
	url := fmt.Sprintf("%s/boards/%s/lists?key=%s&token=%s", c.baseUrl, b.Id, c.key, c.token)

	resBody, err := httpRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var lists []TrelloList
	if err := json.Unmarshal(resBody, &lists); err != nil {
		return nil, err
	}

	return lists, nil
}
