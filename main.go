package main

import(
	"fmt"
	"net/http"
	"path"
	"html/template"
	"io/ioutil"
	"encoding/json"
)


func main(){
	http.HandleFunc("/",index_handler)
	fmt.Println("App started on :8001")
	http.ListenAndServe(":8001",nil)
	
}

func index_handler(w http.ResponseWriter, r *http.Request){
	fmt.Println("...")
	page := path.Join("template","index.html")
	tmpl,err := template.ParseFiles(page)
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	feeds := get_news_feeds()
    if err := tmpl.Execute(w, feeds); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}


type NewsResult struct{
	Status string `json:status`
	TotalResults int `json:totalResults`
	Articles []EachArticle `json:articles`
}

type EachArticle struct{
	Author string `json:aurthor`
	Title string `json:title`
	Description string `json:description`
	Url string `json:url`
	UrlToImage string `json:urlToImage`
	PublishedAt string `json:publishedAt`
}

func get_news_feeds() NewsResult{
	fmt.Println("Fetching news feed...")
	response, err := http.Get("http://newsapi.org/v2/top-headlines?country=in&apiKey=94d7d9fa87754ce2b9dab21093634222")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		var newsFeeds NewsResult
		return newsFeeds
    } else {
        data, _ := ioutil.ReadAll(response.Body)
		// resp := string(data)
		var newsFeeds NewsResult
		json.Unmarshal(data,&newsFeeds)
		// fmt.Println(len(newsFeeds.Articles))
		return newsFeeds
    }
}