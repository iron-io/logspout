# Custom Logspout Builds

Forking logspout to change modules is unnecessary! Instead, you can create an
empty Dockerfile based on `gliderlabs/logspout:master` and include a new `modules.go` file for the
build context that will override the standard one.

This uses the LogHandler interface to show chaining concept and the ability to abort processing chain. 

Add your handlers to modules go file. 

Then build it by building an image in this directory:

```sh
docker build -t mylogspouter .
```

And start it up with:

```sh
docker run --rm --name=logspout \
     -v=/var/run/docker.sock:/var/run/docker.sock \
     -p 8000:80 \
     mylogspouter
```
