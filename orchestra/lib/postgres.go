package lib

import (
	"fmt"
	"log"
	"sync"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/stdlib"
)

var (
	Pg = &PostgresWithMutex{Mu: &sync.RWMutex{}} // Для глобального использования

	pgaddress = Getenv("POSTGRES_ADDRESS", "host=localhost port=5432")
	pgdb      = Getenv("POSTGRES_DB", "expressionist")
	pguser    = Getenv("POSTGRES_USER", "default")
	pgpass    = Getenv("POSTGRES_PASSWORD", "2S%23xYsW{")
)

// PostgresWithMutex Клиент Postgres с мьютексом для защиты от гонки
type PostgresWithMutex struct {
	Client *sqlx.DB
	Mu     *sync.RWMutex
}

// SetupPostgres подключает Postgres
func SetupPostgres() {
	// Инициализируем
	connectionString := fmt.Sprintf("%s dbname=%s user=%s password=%s sslmode=disable", pgaddress, pgdb, pguser, pgpass)
	client, err := sqlx.Open("pgx", connectionString)
	if err != nil {
		fmt.Println("ERR: Failed to connect to Postgres database")
		log.Fatal(err.Error())
	}

	err = client.Ping()
	if err != nil {
		fmt.Println("ERR: Failed to ping Postgres database")
		log.Fatal(err.Error())
	}

	Pg.Client = client

	// Если нет таблицы expressions, создаем
	_, err = Pg.Client.Exec("CREATE TABLE IF NOT EXISTS expressions (id uuid NOT NULL, expression character varying(255) COLLATE pg_catalog.\"default\" NOT NULL, calculation_time timestamp without time zone, result character varying(255) COLLATE pg_catalog.\"default\", dispatch_time timestamp without time zone, status character varying(255) COLLATE pg_catalog.\"default\" NOT NULL, is_sent boolean NOT NULL DEFAULT false, agent_id uuid, CONSTRAINT expressions_pk PRIMARY KEY (id))")
	if err != nil {
		fmt.Println("ERR: Failed to create if not exists Postgres table")
		log.Fatal(err.Error())
	}

	fmt.Println("SETUP: Postgres database has been connected")
}
