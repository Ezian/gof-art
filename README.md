
# gof-art

A small web service that take an image from somewhere and create a nice ascii art version of it.

It is used to show some feature of Golang:

* Simple Native http framework
* `go test` with `-race`  and `-bench` options
* go routine
* channels

# Quick test

```bash
curl -H "Content-Type: application/json" \
--request POST \
--data '{"url":"http://localhost:9999/Gophers.jpg","width":250}' \
http://localhost:9999/naive > test.txt
```

# Quick tour

Default port is 9999.

Every route `/naive`, `/mutex` and `/channel` does the same stuff:

* `POST` allow to upload a Ascii-converted image from any image on the net
* `GET` return the last upload ascii art.

The only differency is the implementation, and the handle of the concurrency on the uploaded art :

* `/naive`  just use a plain variable, with no security. It's a wrong implementation (since it is open to dataraces)
* `/mutex`  use a mutex to secure the variable that contains the last posted art.
* `/channel` use a monitor with channels to secure the variable that contains the last posted art. It's really go-like.

There is also another route that contains a pretty image : `/Gophers.jpg`