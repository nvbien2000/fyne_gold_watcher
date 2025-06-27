BINARY_NAME="Gold Watcher.app"

build-darwin:
	rm -rf ${BINARY_NAME}
	fyne package -os darwin

tidy:
	go mod tidy

run:
	env DB_PATH="./sql.db" air

clean:
	@echo "Cleaning..."
	go clean
	rm -rf ${BINARY_NAME}
	@echo "Cleaned!"

test:
	go test -v ./...