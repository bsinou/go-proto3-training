# Goclipse - the go plugin for Eclipse

[The main site project](https://github.com/GoClipse/goclipse)
Note: this project is no longer maintained.

## How to install

- Get the lastest version of the plugin in the Eclipse Market place
- Add go dependencies:

``` 
  # get the sources 
  go get -u github.com/nsf/gocode
  go get golang.org/x/tools/cmd/guru
  go get -u github.com/rogpeppe/godef
  
  
  # build them 
  go build github.com/nsf/gocode
  go build golang.org/x/tools/cmd/guru
  go build github.com/rogpeppe/godef 

``` 