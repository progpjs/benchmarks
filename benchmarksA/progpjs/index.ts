import "@progp/core"
import {HttpServer} from "@progp/http"

let server = new HttpServer(8000);
let host = server.getHost("localhost");

host.GET("/", async req => {
    //console.log("Request IP:" + req.requestIP());
    //console.log("Request path: ", req.requestPath());
    req.returnHtml(200, "Hello world")
});

server.start();

console.log("Server started at http://localhost:8000")