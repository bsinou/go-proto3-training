# Integrated Development Environment for GO


## Visual studio code.

The [opens source IDE}(https://code.visualstudio.com) from Microsoft based on electron.


### How to install

- On centOS, simply follow [these instructions](https://code.visualstudio.com/docs/setup/linux#_rhel-fedora-and-centos-based-distributions)


## Goclipse - the go plugin for Eclipse

[The main site project](https://github.com/GoClipse/goclipse)
Notes: 
- this project is no longer maintained.
- it has a limited set of features 
- does not work very well

### How to install

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