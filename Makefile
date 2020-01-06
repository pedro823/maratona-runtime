
dev:
	CHALLENGE_MASTER_TOKEN=123 go run .

prod:
	CHALLENGE_MASTER_TOKEN=123 MARTINI_ENV=production go run .

build-prod:
	go build .