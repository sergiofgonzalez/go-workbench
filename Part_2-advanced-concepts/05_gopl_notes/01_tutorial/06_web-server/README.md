# A simple HTTP web server

## v0: initial version

Write a minimal HTTP echo server that establishes a handler function to incoming URLs that begin with `/`.

In the handler, you should write as response the path from where the request is received. For example, when receiving a request from http://localhost:8080/hello the response should be something like:


```
URL.Path = "/hello"
```

## v1: counting the requests

In this version, you should create an additional endpoint and modify the behavior of the handler so that:

+ all requests that begin with `/` should be counted. Note that you should use a mutex to ensure there is no race condition when updating the index.

+ the `/count` endpoint should report the current count. You should also use a mutex to ensure you're not reading the counter while it is being updated by the other handler.

## v2: displaying more information about the request

In this version, the handler function is improved to report on the headers and form data.

## v3: blending the lissajous example with the web server

