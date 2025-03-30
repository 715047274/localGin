#DB_URL=postgresql://root:password@localhost:5432/simple_bank?sslmode=disable
# sqlite local db
DB_URL=sqlite3://sqliteDemo.db

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

# migrate create -ext sql -dir database/migration/ -seq init_mg
# migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose up