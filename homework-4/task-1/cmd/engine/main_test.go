package main

import (
	"testing"

	"go-core.course/homework-4/task-1/pkg/dummyScan"
)

func TestEngine(t *testing.T) {
	sites := []string{"https://yandex.ru", "https://ya.ru"}
	depth := 2

	d := new(dummyScan.Dummy)
	hash := scanResults(d, sites, depth)

	want := "Блог Яндекса"
	got := hash["https://yandex.ru"]["https://yandex.ru/blog/company/"]

	if got != want {
		t.Fatalf("заголовок https://yandex.ru/blog/company/ - %s, а ожидалось %s", got, want)
	}
}
