## example of gopherjs and browserify

### First

```
$ go get -u github.com/gopherjs/gopherjs
$ gopherjs build main.go
```

### Second

```
$ npm i -g browserify
$ browserify index.js -o serve/index.js
```

### Third

```
$ go run serve/serve.go
```

then access to `localhost:8080`
