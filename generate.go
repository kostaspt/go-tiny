package generate

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --target=./internal/sql/ent --template=./internal/sql/ent/template ./internal/sql/ent/schema
//go:generate wire ./...
