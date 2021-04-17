# To indicate that our commands don't create files
.PHONY: postgres adminer migrate

# This makefile command runs a docker container and removes it afterwards
# Create a database with user postgres and password secret, db name is postgres
# If you want to run it in production, change production
postgres:
	docker run --rm -ti --network host -e POSTGRES_PASSWORD=secret postgres

adminer:
	docker run --rm -ti --network host adminer

migrate:
	migrate -source file://migrations \
	 		-database postgres://postgres:secret@localhost/postgres?sslmode=disable up
			
migrate-down:
	migrate -source file://migrations \
	 		-database postgres://postgres:secret@localhost/postgres?sslmode=disable down