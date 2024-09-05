test:
	export ACTIVE_PROFILE=test && go test "./..." -coverprofile="coverage.out" -covermode=count

server:
	go run main.go

lint:
	golangci-lint run ./... --verbose --out-format checkstyle > golangci-lint.xml --timeout 10m

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

mock:
	mockgen -package mock -destination db/mock/store.go golang-boiler-plate/db Store

migrate:
	migrate create -ext sql -dir db/migration -seq ${FILE_NAME}

sonar:
	sonar-scanner -Dsonar.projectKey=${PROJECT_KEY} -Dsonar.sources=. -Dsonar.tests=. -Dsonar.host.url=${SONAR_HOST} -Dsonar.go.exclusions=**/vendor/**,**/*_mock.go,**/db/** -Dsonar.go.coverage.reportPaths=coverage.out -Dsonar.test.inclusions=**/*_test.go,**/test/** -Dsonar.login=${SONAR_LOGIN} -Dsonar.projectVersion=${GIT_COMMIT_ID}

.PHONY: lint test sonar server migratedown migrateup sqlc mock migrate
