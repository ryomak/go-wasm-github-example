package main

import(
 "net/http"
 "encoding/json"
 "example/common"
 trending "github.com/ryomak/go-trending"
)

var client = trending.NewClient()

func main() {
	http.HandleFunc("/trend", trendHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.ListenAndServe(":8080", nil)
}

func trendHandler(w http.ResponseWriter, r *http.Request) {
	// Show projects of today
	repos, err := client.GetRepository(trending.TIME_WEEK, "go")
	if err != nil {
		json.NewEncoder(w).Encode([]common.Repository{
			{
				Name: "cant find",
				User: "user",
				URL:  "/",
				Star: 0,
			},
		})
		return
	}
	list := []common.Repository{}
	for _, repo := range repos {
		list = append(list, common.Repository{
			Name: repo.Name,
			User: repo.Owner,
			URL:  repo.URL,
			Star: repo.Star,
			Description: repo.Description,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
	return
}
