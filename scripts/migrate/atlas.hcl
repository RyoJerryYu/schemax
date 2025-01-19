// environment for local server
env "mysql" {
  src = "file://schema/mysql.sql"
  url = "mysql://schemaxuser:123456@127.0.0.1:13306/schemax?loc=Local&parseTime=true&charset=utf8mb4"
  dev = "docker://mysql/8/schema"
}

env "postgres" {
  src = "file://schema/postgres.sql"
  url = "postgres://schemaxuser:123456@127.0.0.1:15432/schemax?sslmode=disable"
  dev = "docker://postgres/15/schema"
}

env "sqlite3" {
  src = "file://schema/sqlite3.sql"
  url = "sqlite://./schemax.db"
  dev = "sqlite://file?mode=memory"
}