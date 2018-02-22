Manda API
================

## Running Locally

Make sure you have [Go](http://golang.org/doc/install) installed.

```sh
$ go get -u github.com/dafian47/manda-api
$ cd $GOPATH/src/github.com/dafian47/manda-api
$ govendor sync
$ go run server.go
```

Your app should now be running on [localhost:5000](http://localhost:5000/).

You should also install [GoVendor](https://github.com/kardianos/govendor) if you are going to add any dependencies

## Built With

* [Gin Gonic](https://github.com/gin-gonic/gin)
* [Gorm](https://github.com/jinzhu/gorm)
* [PostgreSQL Driver](https://github.com/lib/pq)

## Planning

* [X] Rest for Login & Register
* [ ] Rest for Forget Password & Email Verification
* [X] Rest for User, Channel & Thread
* [X] Rest for Comments
* [X] Rest for Edit & Delete Channel
* [X] Rest for Edit & Delete Thread
* [X] Rest for Data Master
* [ ] Rest for Statistic
* [ ] Using JWT Token
* [ ] Documentation

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
