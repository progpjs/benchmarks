# ProgpJS Benchmarks

This project allows benchmarking ProgpJS. This benchmarks are done using bombardier which is a Go tools
allowing to send a lot of request to a server and stress this server.

Here we are benchmarking NodeJS, BunJS, DenoJS, ProgpJS and Progp http stack with pure Go.

All the test are executed on the same computer (bombardier + the server) which is limiting the performance by a lot.
But the main goal is to detect drop in performance, while comparing to others solutions. This comparison allows to
know if the matter is in ProgpJS or the computer. For exemple we detect a drop in performance while using last
version of GoLang IDE which is constantly using 6% of the process ... (why?)

## When testing

* Close your IDE, since they can consume a lot of CPU resources even when IDLE.
* When benchmarking ProgpJS, use compiled mode and this without plugin. Others execution modes are very slow! 

## Benchmarks A

Here the servers must respond "hello world" to each request. It's a simple test allowing to detect the performance
of the HTTP stack and the internal mechanisms inside ProgpJS.

### Round 1: 10 clients bombarding