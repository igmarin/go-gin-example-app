package main

func initializeRouter() {
	router.GET("/", showIndexPage)
	router.GET("/article/view/:article_id", getArticle)
}
