# My Go lang playground

## Hello world
To run the __hello world:__ 

    go run src/github.com/nvieirafelipe/hello/hello.go

## Github Issues
To run __issues__ add env var GITHUB_KEY with your github api key, exec `go get` to download [go-github](https://github.com/google/go-github) and [goauth2](https://code.google.com/p/goauth2/) and then:

    go run src/github.com/nvieirafelipe/issues/issues.go

## Web hello world
To run the __web hello world:__
    
    PORT=5000 go run src/github.com/nvieirafelipe/web/web.ho

And open your browser at `http://localhost:5000`