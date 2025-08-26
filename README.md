## Migrations
### Create a new migration
`goose create <migration_name> sql`

### Run up migrations
`goose -dir ./migrations postgres "$DB_STRING" up`

or

`goose -dir ./migrations postgres "$DB_STRING" up-by-one`


to run one migration