	CURRENT_DIR=$(shell pwd)

proto-gen:
	./script/gen-proto.sh ${CURRENT_DIR}

exp:
	export DBURL='postgres://macbookpro:1111@localhost:5432/habittracker?sslmode=disable'

mig-up:
	migrate -path storage/migrations -database 'postgres://macbookpro:1111@localhost:5432/habittracker?sslmode=disable' -verbose up

mig-down:
	migrate -path storage/migrations -database ${DBURL} -verbose down

mig-create:
	migrate create -ext sql -dir storage/migrations -seq create_tables

mig-insert:
	migrate create -ext sql -dir storage/migrations -seq insert_table
