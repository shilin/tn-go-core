// Package dummySpider реализует сканер содержимого веб-сайтов.

// Пакет сделан для тестирования и позволяет получить предварительно заготовленый список ссылок и заголовков страниц внутри веб-сайта по его URL.

package dummySpider

import (
	"testing"
)

func TestScanSite(t *testing.T) {
	const url = "https://yandex.ru"
	const depth = 2
	var want = map[string]string{
		"https://yandex.ru/collections/":  "Яндекс.Коллекции",
		"https://yandex.ru/all":           "Все сервисы Яндекса",
		"https://yandex.ru/blog/company/": "Блог Яндекса",
	}
	got, err := Scan(url, depth)
	if err != nil {
		t.Fatal(err)
	}

	if got["https://yandex.ru/collections/"] != want["https://yandex.ru/collections/"] {
		t.Fatalf("Получено от сканера %s, а ожидалось %s", got, want)
	}
}
