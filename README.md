# ts-demo-gqlgen

## Schema

```mermaid
flowchart LR
    dbhasura[(PgSQL)]
    dbdata[(MySQL)]
    gqlhasura[Hasura]
    gqlthing[Thing GQL]

    gqlhasura --> |metadata| dbhasura
    gqlhasura --> |trollies| dbhasura
    gqlhasura --> |things| gqlthing
    gqlthing --> dbdata
```

## walkthrough
1. fetch schema from original hasura using `@apollo/rover` 
```
rover graph introspect http://hasura/graphql > graph/raw/instrospect_schema.graphql
```
2. copy graph SDL need to graph/schema.graphqls
3. generate golang graphql interface using `gqlgen`
```
go run github.com/99designs/gqlgen generate
```
4. write gorm based code on graph/query 
5. run server `go run server.go`

## hasura 2.41
### cli

download https://github.com/hasura/graphql-engine/releases/tag/v2.41.0

```
hasura metadata export --metadata-database-url=postgres://postgres:root@host.docker.internal:5432/hasura_2_41 --database-url=postgres://postgres:root@host.docker.internal:5432/trial
```

## graphql

### qglgen-query
```
=== query ===
query($lm: Int, $wh: things_bool_exp) {
  thing(limit:$lm,where:$wh){
    id
    name
    amount
    created
  }
}
=== variables ===
{
  "lm": 1, 
  "wh": {
    "name": {"_like": "838%"},
    "amount": {"_gt": 10}
  }
}
```

### hasura-query

