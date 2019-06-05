# u-job api server

## Precondition

- Go >= 1.12
- docker
- docker-compose

## Setup project

```
docker-compose up
```

> Recommand using vistual-studio-code-insider with remote development extension, then you can connect to container and use the package in container, for more information https://code.visualstudio.com/docs/remote/containers

## ORM

we using gorm, here is the [document](https://gorm.io/docs/), please don't directly operate ORM in handler, all oprator should in repository, and the most important thing is don't use `Entity` for where condition, golang have [zero-value](https://golang.org/ref/spec#The_zero_value) problem, if we not assign the value, the value will be the zero-value.

for example
```go
type User struct {
    Name  string
    Age   int
    Email string
}

// bad case, in this case, the email will be empty string, there is no way to know the empty string is assign by us or assign by golang
db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)

// good case
db.Where("name = ? AND age = ?", "jinzhu", 20).First(&user)
```

## Migration

for create migration file

```bash
make create-migration NAME="migration file name"
```

for upgrade database

```bash
make migrate-up
```

for downgrade database

```bash
make migrate-down
```

## Contribute

please follow below commit message format

```
Short commit message

Connect to issue #{issue number}.
Anything you want to describe more
```
