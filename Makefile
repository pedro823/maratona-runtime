
dev:
	CHALLENGE_MASTER_TOKEN=123 go run .

create-postgres-dev:
	docker run --name postgres-maratona -e POSTGRES_PASSWORD=secretpass -p 5432:5432 -d postgres

resume-postgres-dev:
	docker start -p 5432:5432 postgres-maratona

destroy-postgres-dev:
	docker kill postgres-maratona
	docker rm postgres-maratona

prod:
	CHALLENGE_MASTER_TOKEN=123 MARTINI_ENV=production go run .

build-prod:
	go build .

test:
	go test -count=5 ./...

docker_test: .docker_test_build
	docker run --rm maratona-runtime-test

docker_test_verbose: .docker_test_build
	docker run --rm maratona-runtime-test go test --count=5 -v ./...

.docker_test_build: Dockerfile.testing runtime/
	docker build -f Dockerfile.testing . -t maratona-runtime-test
	touch ".docker_test_build"
