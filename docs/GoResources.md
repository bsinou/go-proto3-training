# Resources for the Go Language

## HowTo-s

### Useful commands

- to start a server that locally expose the documentation of all libraries that are on the local machine, simply run:
 `godoc -http :6060` (or whatever free port)

### Upgrading go

On **CentOs 7**:

- Download the arm64 archive from [GoLang website](https://golang.org/dl/)
- Untar and move to `/usr/local/go1.10` (FI)
- Update the symbolic link `/usr/local/go`:

```bash
 cd /usr/local
 sudo rm go
 sudo ln -s go1.10 go
```

**Note:** to ease plugin install, make the `/usr/local/go/bin` subfolder writable for all users

```bash
 sudo chmod -R a+rw go/bin
```

### Post Upgrade Tasks

#### Visual Code IDE

To insure Visual Code IDE is up to date, restart the workstation, and start the IDE. You should accept "All Plugin" installation and also recompile suggested plugins

#### Protobuf

Protobuf is installed in the bin subfolder of the `/usr/local/go` folder, it must thus be reinstalled after update

## Baby steps

If you are already familliar with code programming, going first through the resources listed below in that order is a good way to be quickly up and running in GO:

- A [first tour](https://tour.golang.org/list) of Go
- Important [code conventions](https://golang.org/doc/code.html)
- The [specifications](https://golang.org/ref/spec): they are surprisingly easy to read.
- [Effective Go](https://golang.org/doc/effective_go.html)
- A list of [small examples](https://gobyexample.com): each article is covering a specific point

### The basics, again and again

A selection of good articles that focuses on a specific feature of the Go language, with another point of view.

- About [Mutexes](https://kylewbanks.com/blog/tutorial-synchronizing-state-with-mutexes-golang)

## Good resources

- Rob Pike videos:

## About testing

### Unit tests

- [5 Tips](https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742) to start with

## REST services

### HTTP basics

- [The reference](https://golang.org/pkg/net/http/) on golang website
- [A good recap](http://www.alexedwards.net/blog/a-recap-of-request-handling) of request handling

## Questions and answers

Some resources that gather Q&A about go. Always good to refresh  one's mind

- 30 rather easy [questions](http://www.golangpro.com/2015/08/golang-interview-questions-answers.html) 
- A [fun website](https://gophercises.com) with tutorial videos and exercices
