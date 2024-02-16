const http = require("http");

benchmarkHttpStack();

function benchmarkHttpStack() {
    const host = 'localhost';
    const port = 8000;

    const server = http.createServer((req, res) => res.writeHead(200).end("Hello world"));

    server.listen(port, host, () => {
        console.log(`Server is running on http://${host}:${port}`);
    });
}