# Haproxy 2.0 HTX / `Expect` header / `http-buffer-request` issue

This is a minimal setup to replicate an issue introduced in Haproxy 2.0 with `Expect: 100-continue` and `option http-buffer-request`.

## Prerequisites

* `docker-compose` 1.17.1
* Curl (tested with 7.52.1)


## Running

Use `docker-compose up` to bring everything up. Then run `./tests.sh` to run the tests that expose the problem. The tests against port `20080` with an `Expect` header should fail with the following message:

```
<html><body><h1>502 Bad Gateway</h1>
The server returned an invalid or incomplete response.
</body></html>
```

And the following errors in the logs:

```
unix:1 [01/Jul/2019:20:35:35.743] http-in app/app 0/0/0/1000/11002 200 398 cL-- 2/1/0/0/0 0/0 "POST / HTTP/1.1"
192.168.176.1:37570 [01/Jul/2019:20:35:35.743] ssl_offload ssl_offload/http-in 0/0/0/-1/11003 502 273 SH-- 1/1/0/0/0 0/0 "POST / HTTP/1.1"
```

Note the 502 from the `ssl_offload` frontend with the `SH--` proxy termination status. That line also shows the client timeout of 10s has been exceeded.
