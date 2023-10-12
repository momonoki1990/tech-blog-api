リポジトリ構成の参考
https://blog.recruit.co.jp/rls/2018-03-16-go-ddd/

## Docker

```
$ docker-compose up
```

## sql-migrate

```
$ sql-migrate new -env="local" -config=sqlmigrate.yml create_articles_table
$ sql-migrate up -env="local" -config=sqlmigrate.yml
$ sql-migrate down -env="local" -config=sqlmigrate.yml
```

## sql-boiler

```
$ sqlboiler mysql -c sqlboiler.toml -o infra/database/model -p model --no-tests --wipe
```

## Test

```
$ docker-compose exec tech-blog-api go test -v ./...
```
