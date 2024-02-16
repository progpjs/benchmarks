const port = 8000;

const handler = (request: Request): Response => {
    return new Response("Hello world", { status: 200 });
};

console.log(`HTTP server running. Access it at: http://localhost:8000/`);
Deno.serve({ port }, handler);