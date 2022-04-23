param(
  [Parameter(Mandatory=$true)]
  [String]$dbPassword
)

migrate -database "postgres://postgres:${dbPassword}@localhost:5432/pf-dev?sslmode=disable" -path migrations up