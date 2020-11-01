// Package dummyScan выдает подготовленные данные вместо сканера.
package dummyScan

// Scan осуществляет рекурсивный обход ссылок сайта, указанного в URL,
// с учётом глубины перехода по ссылкам, переданной в depth.
import "go-core.course/homework-3/task-1/pkg/scan"

type Dummy struct {
	scan.Scanner
}

func (d *Dummy) Scan(url string, depth int) (data map[string]string, err error) {
	data = map[string]string{
		"https://yandex.ru/collections/":  "Яндекс.Коллекции",
		"https://yandex.ru/all":           "Все сервисы Яндекса",
		"https://yandex.ru/blog/company/": "Блог Яндекса",
	}

	return data, nil
}
