
# kim-guru   
A clone of [kim.guru](https://kim.guru/) in golang, It uses [apixu-api](https://www.apixu.com/api.aspx) and [apixu-go](https://github.com/mohan3d/apixu-go).    


## Installation
```sh
$ go get -u github.com/mohan3d/kim-guru
```

## Run Locally
Make sure you have [Go](http://golang.org/doc/install) and the [Heroku Toolbelt](https://toolbelt.heroku.com/) installed.

```sh
$ cd $GOPATH/src/github.com/mohan3d/kim-guru
$ export APIXU_KEY=<YOUR-APIXU-API-KEY>
$ heroku local
```

Your app should now be running on [localhost:5000](http://localhost:5000/).    
You should also install [govendor](https://github.com/kardianos/govendor).

## Deployment
[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/mohan3d/kim-guru)    

**Note:** Don't forget to set config var APIXU_KEY.

## Manual deployment
```sh
$ heroku create
$ heroku config:add APIXU_KEY=<YOUR-APIXU-API-KEY>
$ git push heroku master
$ heroku open
```