package main

import ("encoding/json"
        "fmt")

type Patron struct {
  Name string
  CardNumber string `json:"card_number"`
  Branch string
  Loans Loans
  Holds []Hold
}

type Loans struct {
  Books []Book
  Videos []Video
}

type Book struct {
  Title string
  Author string
  CallNumber  string `json:"call_number"`
  DueDate string `json:"due_date"`
}

type Video struct {
  Title string
  CallNumber string `json:"call_number"`
  DueDate string `json:"due_date"`
}

type Hold struct {
  Title string
  Author string
  CallNumber string `json:"call_number"`
  Placed string `json:"placed_at"`
}


func main() {
  patron_data := `{"name": "Christine",
                   "card_number": "21111000001111",
                   "branch": "Bookton",
                   "holds": [ { "title": "Educated: a memoir",
                                "author": "Tara Westover",
                                "call_number": "921 WEST",
                                "placed_at": "2018-11-10" } ],
                   "loans": {
                     "books": [ { "title": "Nomadland: Surviving America in the Twenty-First Century",
                                  "author": "Jessica Bruder",
                                  "call_number": "331.398 BRU",
                                  "due_date": "2018-11-25" },
                                { "title": "Born a Crime",
                                  "author": "Trevor Noah",
                                  "call_number": "921 NOAH",
                                  "due_date": "2018-11-25" } ],
                     "videos": [ { "title": "Yosemite",
                                   "call_number": "DVD YOS",
                                   "due_date": "2018-11-25" } ]
                       } }`



  var patron Patron
  json.Unmarshal([]byte(patron_data), &patron)

  fmt.Println("Patron Name:", patron.Name)
  fmt.Println("Home Library:", patron.Branch)
  fmt.Println("Patron Library Card Number", patron.CardNumber)
  printLoans(patron.Loans)
  printHolds(patron.Holds)
}

func printLoans(loans Loans) {
  fmt.Printf("***Library Loans:***\n")
  fmt.Printf("***Books:***\n")
  for _, book := range loans.Books {
    fmt.Printf("Title: %s \nAuthor: %s \nCall Number: %s \nDue Date: %s \n\n", book.Title, book.Author, book.CallNumber, book.DueDate)
  }

  fmt.Printf("***Videos:***\n")
  for _, video := range loans.Videos {
    fmt.Printf("Title: %s \nCall Number: %s \nDueDate %s\n", video.Title, video.CallNumber, video.DueDate)
  }
}

func printHolds(holds []Hold) {
  fmt.Printf("\n***Library Materials on Hold:***\n")

  for _, hold_item := range holds {
    fmt.Printf("Title: %s \nAuthor: %s \nCall Number: %s \nHold placed on: %s\n", hold_item.Title, hold_item.Author, hold_item.CallNumber, hold_item.Placed)
  }
}
