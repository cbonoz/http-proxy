http-proxy
---

A basic HTTP proxy written with the golang gin framework.

See `proxy_controller.go`


### Deployment
`fly deploy`


### Example call (axios):

PROXY_URL: `YOUR_SERVICE_DOMAIN/proxy`

<pre>
    return axios.post(PROXY_URL, {
        url, // target url
        type: 'GET', // can be GET,POST,PATCH,DELETE
        hash: MD5(window.location.origin).toString() // passed in checksum
        // body: // optional - json object
    });
</pre>