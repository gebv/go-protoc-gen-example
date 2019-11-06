При запуске main.go ошибка `cannot find package "api/bar"`. Что ожидаемое. Потому что в `api/foo/foo.proto` импортируется proto из `api/bar/models.proto`.

Как правильно указать путь к proto файлу в другой дириктории в рамках одного и того же проекта так что бы сгенеренный го-код имел валидный импорт?

Правильный ответ

Для каждого заинклуженного proto файла надо явно указать го пакет

В аргументах указываем явно ассоцицаию `api/bar/models.proto` с го пакетом `github.com/gebv/go-protoc-gen-example/api/bar` что в итоге приводит к верному значению в import в сгенеренном го коде.

```
protoc --proto_path=./ --gogofast_out=Mapi/bar/models.proto=github.com/gebv/go-protoc-gen-example/api/bar,plugins=grpc:. ./api/foo/*.proto
```

