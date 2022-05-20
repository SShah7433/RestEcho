# Rest Echo

A simple REST API that will accept any path/payload and echo the URL query parameters and payload to the console for troubleshooting.

The server (should) always reply with response code 200.

### Configuration

__Port__: Set envirotnment variable `PORT` to desired port number. Defaults to 8080.

__Sample Output__
```
--------------------------------------------
--- URL PARAMETERS --- 
urlp1: urlv1

--- REQUEST BODY ---
{
        "key1": "value1"
}
--- REQUEST BODY END ---
--------------------------------------------
[GIN] 2022/05/20 - 11:31:33 | 200 |     149.651µs |       127.0.0.1 | POST     "/?urlp1=urlv1"
```