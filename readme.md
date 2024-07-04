# Pagination example

This is a server, written in Go, which serves paginated people data.

To start the application run: `go run pagination.go`.

## Functionality
Dynamically paginates data returned from Postgres. Should populate min/max buttons and provide convenient links to walk through the pages.

## Postgress

For the app to work correctly you must have an instance of postgres running locally and have updated the values in the .env file.
