package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
)

// TODO: Please create a struct to include the information of a video

type YouTubeIndex struct {
	Title        string
	Id           string
	ChannelTitle string
	LikeCount    string
	ViewCount    string
	PublishedAt  string
	CommentCount string
}

func YouTubePage(w http.ResponseWriter, r *http.Request) {
	// TODO: Get API token from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		http.ServeFile(w, r, "./error.html")
		return
	}

	apiKey := os.Getenv("YOUTUBE_API_KEY")

	// TODO: Get video ID from URL query `v`
	youtubeId := r.URL.Query().Get("v")

	if youtubeId == "" {
		fmt.Println(err)
		http.ServeFile(w, r, "./error.html")
		return
	}

	// TODO: Get video information from YouTube API
	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?key=%s&id=%s&part=snippet,statistics", apiKey, youtubeId)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		http.ServeFile(w, r, "./error.html")
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		http.ServeFile(w, r, "./error.html")
		return
	}

	// TODO: Parse the JSON response and store the information into a struct
	// var data map[string]interface{}
	var jsonResponse struct {
		Items []struct {
			Id      string `json:"id"`
			Snippet struct {
				Title        string `json:"title"`
				ChannelTitle string `json:"channelTitle"`
				PublishedAt  string `json:"publishedAt"`
			} `json:"snippet"`

			Statistics struct {
				LikeCount    string `json:"likeCount"`
				ViewCount    string `json:"viewCount"`
				CommentCount string `json:"commentCount"`
			} `json:"statistics"`
		} `json:"items"`
	}

	if err = json.Unmarshal(body, &jsonResponse); err != nil {
		fmt.Println(err)
		http.ServeFile(w, r, "./error.html")
		return
	}

	if len(jsonResponse.Items) < 1 {
		http.ServeFile(w, r, "./error.html")
		return
	}

	// TODO: Display the information in an HTML page through `template`

	toHtml := YouTubeIndex{
		Title:        jsonResponse.Items[0].Snippet.Title,
		Id:           jsonResponse.Items[0].Id,
		ChannelTitle: jsonResponse.Items[0].Snippet.ChannelTitle,
		LikeCount:    jsonResponse.Items[0].Statistics.LikeCount,
		ViewCount:    jsonResponse.Items[0].Statistics.ViewCount,
		PublishedAt:  jsonResponse.Items[0].Snippet.PublishedAt,
		CommentCount: jsonResponse.Items[0].Statistics.CommentCount,
	}

	if err = template.Must(template.ParseFiles("./index.html")).Execute(w, toHtml); err != nil {
		fmt.Println(err)
		http.ServeFile(w, r, "./error.html")
		return
	}

}

func main() {
	http.HandleFunc("/", YouTubePage)
	log.Fatal(http.ListenAndServe(":8085", nil))
}
