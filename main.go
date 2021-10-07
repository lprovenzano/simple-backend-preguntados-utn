package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type question struct {
	ID       int64    `json:"id"`
	Title    string   `json:"title"`
	Options  []answer `json:"options"`
	UrlImage string   `json:"url_image"`
}

type answer struct {
	Title   string `json:"title"`
	IsRight bool   `json:"right"`
}

var questions = []question{
	{ID: 1,
		Title: "¿Cuántas zonas horarias tiene en Rusia?", Options: []answer{
		{
			Title:   "11",
			IsRight: true,
		}, {
			Title:   "3",
			IsRight: false,
		}, {
			Title:   "5",
			IsRight: false,
		}}, UrlImage: "https://images.unsplash.com/photo-1513326738677-b964603b136d?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=749&q=80",
	}, {
		ID:    2,
		Title: "¿Cuál es la flor nacional de Japón?", Options: []answer{
			{
				Title:   "Flor de Cerezo",
				IsRight: true,
			}, {
				Title:   "Flor de Lotto",
				IsRight: false,
			}, {
				Title:   "Flor de Koi",
				IsRight: false,
			}}, UrlImage: "https://images.unsplash.com/photo-1508610048659-a06b669e3321?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=735&q=80",
	}, {
		ID:    3,
		Title: "¿Cuál es el animal nacional de Australia?", Options: []answer{
			{
				Title:   "Canguro",
				IsRight: true,
			}, {
				Title:   "Leopardo",
				IsRight: false,
			}, {
				Title:   "Koala",
				IsRight: false,
			}}, UrlImage: "https://images.unsplash.com/photo-1497752531616-c3afd9760a11?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=1170&q=80",
	}, {
		ID:    4,
		Title: "¿Qué país tiene la mayor cantidad de islas en el mundo?", Options: []answer{
			{
				Title:   "Suecia",
				IsRight: true,
			}, {
				Title:   "EEUU",
				IsRight: false,
			}, {
				Title:   "Inglaterra",
				IsRight: false,
			}}, UrlImage: "https://images.unsplash.com/photo-1534201041980-ab6cb6c36cc3?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=687&q=80",
	}, {
		ID:    5,
		Title: "¿Cuál es el país más pequeño del mundo?", Options: []answer{
			{
				Title:   "El Vaticano",
				IsRight: true,
			}, {
				Title:   "Nueva Zelanda",
				IsRight: false,
			}, {
				Title:   "San Marino",
				IsRight: false,
			}}, UrlImage: "https://images.unsplash.com/photo-1534201041980-ab6cb6c36cc3?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=687&q=80",
	}, {
		ID:    6,
		Title: "¿Qué artista pinto el techo de la Capilla Sixtina en Roma?", Options: []answer{
			{
				Title:   "Miguel Ángel",
				IsRight: true,
			}, {
				Title:   "Pablo Picasso",
				IsRight: false,
			}, {
				Title:   "Vincent van Gogh",
				IsRight: false,
			}}, UrlImage: "https://images.unsplash.com/photo-1576016770956-debb63d92058?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=1026&q=80",
	}, {
		ID:    7,
		Title: "¿Cuál es la obra más famosa de Edvard Munch?", Options: []answer{
			{
				Title:   "La Gioconda",
				IsRight: false,
			}, {
				Title:   "El grito",
				IsRight: true,
			}, {
				Title:   "Guernica",
				IsRight: false,
			}}, UrlImage: "https://images.unsplash.com/photo-1513364776144-60967b0f800f?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=1171&q=80",
	}, {
		ID:    8,
		Title: "¿Cuál es la serie de libros mejor vendida del siglo 21?", Options: []answer{
			{
				Title:   "El alquimista, Paulo Coelho",
				IsRight: false,
			}, {
				Title:   "El señor de los anillos, J. R. R. Tolkien",
				IsRight: false,
			}, {
				Title:   "Harry Potter, J. K. Rowling",
				IsRight: true,
			}}, UrlImage: "https://images.unsplash.com/photo-1532012197267-da84d127e765?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=687&q=80",
	}, {
		ID:    9,
		Title: "¿Quién inventó la World Wide Web, y cuándo?", Options: []answer{
			{
				Title:   "Bjarne Stroustrup, 1987",
				IsRight: false,
			}, {
				Title:   "Tim Berners-Lee, 1990",
				IsRight: true,
			}, {
				Title:   "Dennis Ritchie, 1992",
				IsRight: false,
			}}, UrlImage: "https://images.unsplash.com/photo-1544197150-b99a580bb7a8?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=1170&q=80",
	}, {
		ID:    10,
		Title: "¿Cuál es la canción más reproducida en Spotify hasta el momento?", Options: []answer{
			{
				Title:   "Drake, One Dance",
				IsRight: false,
			}, {
				Title:   "Ed Sheeran, The Shape of You",
				IsRight: true,
			}, {
				Title:   "The Chainsmokers, Closer",
				IsRight: false,
			}}, UrlImage: "https://images.unsplash.com/photo-1470225620780-dba8ba36b745?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=1170&q=80",
	},
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200", "https://tp-sala-juegos-provenzano-luca.herokuapp.com/"},
		AllowMethods:     []string{"GET"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/", getQuestions)

	router.Run(fmt.Sprintf(":%s", port))
}

func getQuestions(c *gin.Context) {
	c.JSON(http.StatusOK, questions)
}
