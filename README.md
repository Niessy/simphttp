simphttp
========

LLike Python's SimpleHTTPServer but with a shorter name and written in Go

Just go install, which will install in your GOPATH then you can call it from
any directory like you would python's SimpleHTTPServer but instead of 

```
python -m SimpleHTTPServer
```

You can write

```
simphttp
```

Your fingers will thank you.

Usage of flags follow as such

The port flag

```
simphttp -port=9001 // defaults to 8000
simphttp -p=9001    // shortform
```

The dir flag

```
simphttp -dir=$HOME/Documents/blah // defaults to current directory
simphttp -d=$HOME/Documents/blah   // shortform
```

