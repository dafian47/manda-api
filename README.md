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

* [ ] Rest for Login & Register
* [ ] Rest for Forget Password & Email Verification
* [ ] Rest for User, Channel & Post
* [ ] Rest for Comments
* [ ] Rest for Edit & Delete Channel
* [ ] Rest for Edit & Delete Posts
* [ ] Using JWT Token
* [ ] Documentation

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
