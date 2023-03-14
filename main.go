package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type anime struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Link  string `json:"uri"`
	Image string `json:"imguri"`
}

var animes = []anime{
	{ID: "1", Title: "Paprika", Link: "https://anilist.co/anime/1943/Paprika/", Image: "https://s4.anilist.co/file/anilistcdn/media/anime/cover/large/b1943-YyEsHWhYKukN.png"},
	{ID: "2", Title: "ODDTAXI", Link: "https://anilist.co/anime/128547/ODDTAXI/", Image: "https://s4.anilist.co/file/anilistcdn/media/anime/cover/large/bx128547-TWRVIu5zRTYx.jpg"},
	{ID: "3", Title: "Perfect Blue", Link: "https://anilist.co/anime/437/Perfect-Blue/", Image: "https://s4.anilist.co/file/anilistcdn/media/anime/cover/large/bx437-dvq9hcIyojVX.jpg"},
	{ID: "4", Title: "Your Name.", Link: "https://anilist.co/anime/21519/Your-Name/", Image: "https://s4.anilist.co/file/anilistcdn/media/anime/cover/large/bx21519-XIr3PeczUjjF.png"},
	{ID: "5", Title: "Death Note", Link: "https://anilist.co/anime/1535/Death-Note/", Image: "https://s4.anilist.co/file/anilistcdn/media/anime/cover/large/bx1535-lawCwhzhi96X.jpg"},
}

func getAnime(p *gin.Context) { // carries request details, validates and serializes JSON
	p.IndentedJSON(http.StatusOK, animes) // serialize the struct into JSON and add it to the response
}

func getAnimeByID(p *gin.Context) {
	id := p.Param("id")
	for _, a := range animes {
		if a.ID == id {
			p.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	p.IndentedJSON(http.StatusNotFound, gin.H{"message": "anime not found"})
}

func main() {
	router := gin.Default()
	router.GET("/anime", getAnime)
	router.GET("/anime/:id", getAnimeByID)
	router.Run("localhost:8080")
}
