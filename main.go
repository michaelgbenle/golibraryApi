package main

import "golibraryApi/router"

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Rules of life", Author: "Jordan Peterson", Quantity: 12},
	{ID: "2", Title: "Rules of love", Author: "Kate Peterson", Quantity: 10},
	{ID: "3", Title: "Rules of wealth", Author: "Jordan Peterson", Quantity: 12},
	{ID: "4", Title: "Rules of health", Author: "Kate Peterson", Quantity: 11},
}

func main() {
	library := router.SetupRouter()

}
