# ts-demo-gqlgen

## Schema

```mermaid
flowchart LR
    dbhasura[(PgSQL)]
    dbdata[(MySQL)]
    gqlhasura[Hasura]
    gqlthing[Thing GQL]

    gqlhasura --> |metadata| dbhasura
    gqlhasura --> |trollies| dbdata
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

## graphql
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