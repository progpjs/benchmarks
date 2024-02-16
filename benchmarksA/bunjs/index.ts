const server = Bun.serve({
    port: 8000,
    fetch(request) {
        return new Response("Hello world");
    },
});

console.log(`Listening on localhost:${server.port}`);