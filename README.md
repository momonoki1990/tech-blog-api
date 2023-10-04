リポジトリ構成の参考
https://blog.recruit.co.jp/rls/2018-03-16-go-ddd/

## Docker

```
$ docker-compose up
```

## sql-migrate

```
$ sql-migrate up -env="local" -config=sqlmigrate.yml
$ sql-migrate down -env="local" -config=sqlmigrate.yml
```

## sql-boiler

```
$ sqlboiler mysql -c sqlboiler.toml -o infra/database/db_model -p db_model --no-tests --wipe
```
