package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"string"`
	Content string `json:"content"`
}

var articleList = []article{
	article{ID: 1, Title: "My dear Billie", Content: "Billie is a cute dog"},
	article{ID: 2, Title: "Chester is the best", Content: "The original an unique dog"},
}

func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}

	return nil, errors.New("Article not found")
}
