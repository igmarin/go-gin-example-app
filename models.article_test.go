package main

import "testing"

func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	if len(alist) != len(articleList) {
		t.Fail()
	}

	for i, v := range alist {
		if v.Content != alist[i].Content ||
			v.ID != alist[i].ID ||
			v.Title != alist[i].Title {

			t.Fail()
			break
		}
	}
}
