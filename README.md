![Coverage](https://img.shields.io/badge/Coverage-3.4%25-red)

Для сборки необходим `.env`-файл в папке config. Возможные варианты:

1. Создать `.env` вручную и вставить в него следующий текст: 

```
MIGRATOR_PASSWORD="postgres"
POSTGRES_USER="postgres"
POSTGRES_PASSWORD="postgres"
POSTGRES_DB="FilmLibrary"
POSTGRES_HOST="filmlib-db"
JWT_SECRET="Секрет"
```

2. Воспользоваться шаблоном `.env.template`. Возможны проблемы с несовпадением имён, не успел сделать импорт этих переменных в не-гошные конфиги (например в compose.yml).
3. Воспользоваться скриптом `create_env.sh`. Он принимает 6 аргументов, соответствующих полям в примере или шаблоне.