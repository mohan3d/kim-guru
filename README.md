
# kim-guru
A clone of [kim.guru](https://kim.guru/) in golang.   
It uses [apixu-api](https://www.apixu.com/api.aspx) and [apixu-go](https://github.com/mohan3d/apixu-go).    

[Deployed version](https://kim-guru-clone.herokuapp.com/)

## Running Locally
Make sure you have [Go](http://golang.org/doc/install) and the [Heroku Toolbelt](https://toolbelt.heroku.com/) installed.

```sh
$ go get -u github.com/mohan3d/kim-guru
$ cd $GOPATH/src/github.com/mohan3d/kim-guru
$ heroku local
```

Your app should now be running on [localhost:5000](http://localhost:5000/).
You should also install [govendor](https://github.com/kardianos/govendor).

## Deploying to Heroku
```sh
$ heroku create
$ heroku config:add APIXU_KEY=<YOUR-APIXU-API-KEY>
$ git push heroku master
$ heroku open
```