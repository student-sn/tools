# Student SN Tools
Инструменты для **Student SN Services**. Здесь есть свой мини-логгер и код для остановки сервера.

Как использовать логгер:

```go
package main

import (
	"github.com/student-sn/tools/logs"
	"github.com/student-sn/tools/stopper"
)

func main() {
	// Инициализация стоппера - требуется *mongo.Client, context.Context, *http.Server
	stop := stopper.Init(client, ctx, server)
	// Инициализация логгера - требуется название сервиса и инициализированный стоппер
	logger := logs.Init("НАЗВАНИЕ СЕРВИСА", stop)
	// Дебаг (работает если глобальная переменная ST_DEBUG не равна 0
	logger.Debug("Это дебаг")
	// Лог
	logger.Log("Это лог")
	// Предупреждение
	logger.Warn("Это предупреждение")
	// Fatal - вывод ошибки и полная остановка
	logger.Fatal("Это fatal")
}
```

Как использовать стоппер:
```go
package main

import "github.com/student-sn/tools/stopper"

func main() {
	// Инициализация стоппера - требуется *mongo.Client, context.Context, *http.Server
	stop := stopper.Init(client, ctx, server)
	// Стандартное завершение
	stop.Stop(0)
	// Завершение с ошибкой
	stop.Stop(-1, "Произошла ошибка")
}
```