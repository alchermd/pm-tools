package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/alchermd/pm-tools/pkg/trello"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	trelloKey := flag.String("trelloKey", "", "Your Trello key")
	trelloToken := flag.String("trelloToken", "", "Your Trello token")
	boardId := flag.String("boardId", "", "The board ID of the board you want to use")

	flag.Parse()

	if *trelloKey == "" {
		fmt.Println("The trelloKey flag is required")
		os.Exit(1)
	}

	if *trelloToken == "" {
		fmt.Println("The trelloToken flag is required")
		os.Exit(1)
	}

	if *boardId == "" {
		fmt.Println("The boardId flag is required")
		os.Exit(1)
	}

	c := trello.NewTrelloClient(*trelloKey, *trelloToken)

	board, err := c.GetBoardById(*boardId)
	if err != nil {
		log.Fatal(err)
	}

	lists, err := c.GetLists(board)
	if err != nil {
		log.Fatal(err)
	}
	list := lists[len(lists)-1]

	card, err := c.CreateCard("test", "test", list.Id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Card (ID: %s) successfully added to %s\n", card.Id, list.Name)
}
