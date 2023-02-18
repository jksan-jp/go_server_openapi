## install
```
brew install openapi-generator
```

## gen
```
openapi-generator generate -i ./schema/openapi.yaml -g go-echo-server -o ./api
```

## run
```

```
go run main.go
```

## 所管

例えばラベルだけ張り替えた際に
```
handlers/api_pets.go:17:21: Container.GetPetsPetID redeclared in this block
        handlers/api_default.go:9:21: other declaration of GetPetsPetID
```
みたいなことになることがあるので（タグつけていないコードのものが残ってダブる）、
一度handlers周りは消してからって感じの運用になりそう。
で、中に処理入れていたりすると、これもこれで多分対応的にNGになるから多分これ使えない
