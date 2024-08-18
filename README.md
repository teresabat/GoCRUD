## UTILIZANDO COMO COLA PARA FUTUROS PROJETOS

**PARA CRIAR O PROJETO**

```
go mod init crud
```

**INSTALANDO MYSQL E GORILLA**
```
go get -u github.com/go-sql-driver/mysql
go get -u github.com/gorilla/mux
```

**ARQUIVO DE CONFIGURAÇÃO DO DATABASE**
```
// database/database.go
package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
    var err error
    DB, err = sql.Open("mysql", "USER:SENHA@tcp(127.0.0.1:3306)/SEU_BANCO")
    if err != nil {
        panic(err.Error())
    }
}
```

