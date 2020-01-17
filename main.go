package main

import (
	"encoding/json"
	"example/common"
	"fmt"
	"net/http"
	"syscall/js"
)

var document = js.Global().Get("document")
var app = js.Global().Get("document").Call("getElementById", "app")

func main() {
	done := make(chan bool)
	go loading(done)
	title()
	fetch(done)
	select {}
}

func getTrend() (common.Repositories, error) {
	r, err := http.Get("/trend")
	defer r.Body.Close()
	if err != nil {
		return common.Repositories{}, err
	}
	repos := common.Repositories{}
	err = json.NewDecoder(r.Body).Decode(&repos)
	return repos, err
}

func title() {
	t := createElement("h1")
	t.Set("textContent", "go-trending in github")
	app.Call("appendChild", t)
}

func fetch(done chan bool) {
	trendListDiv := createElement("div")
	class := trendListDiv.Get("classList")
	class.Call("add", "trend")
	app.Call("appendChild", trendListDiv)

	trends, err := getTrend()
	done <- true
	if err != nil {
		errorSentense := createElement("div")
		errorSentense.Set("textContent", "error")
		trendListDiv.Call("appendChild", errorSentense)
	} else {
		for _, trend := range trends {
			trendListDiv.Call("appendChild", createItemByTrend(trend))
		}
	}
}

func createItemByTrend(repo common.Repository) js.Value {
	item := createElement("div")
	class := item.Get("classList")
	class.Call("add", "item")

	itemLink := createElement("a")
	itemLink.Set("href", repo.URL)

	itemTitle := createElement("div")
	itemTitle.Set("textContent", fmt.Sprintf("%s/%s (☆ %d)", repo.User, repo.Name, repo.Star))
	class = itemTitle.Get("classList")
	class.Call("add", "item-title")

	itemDescription := createElement("div")
	itemDescription.Set("textContent", repo.Description)
	class = itemDescription.Get("classList")
	class.Call("add", "item-description")

	item.Call("appendChild", itemLink)
	itemLink.Call("appendChild", itemTitle)
	itemLink.Call("appendChild", itemDescription)
	return item
}

func loading(done chan bool) {
	chase := createElement("div")
	class := chase.Get("classList")
	class.Call("add", "sk-chase")
	app.Call("appendChild", chase)
	dots := []js.Value{}
	for i := 0; i < 6; i++ {
		chaseDot := createElement("div")
		class = chaseDot.Get("classList")
		class.Call("add", "sk-chase-dot")
		chase.Call("appendChild", chaseDot)
		dots = append(dots, chaseDot)
	}
	<-done
	style := chase.Get("style")
	style.Set("display", "none")

}

func createElement(elementName string) js.Value {
	return document.Call("createElement", elementName)
}
