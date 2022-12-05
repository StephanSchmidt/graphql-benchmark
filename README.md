
# Measuring Performance of plain SQL, SQL that generates JSON, an ORM and GraphQL

This repository is for measuring the preformance of different database connections in Golang.

For each test a web page is implemented that reads a list of Todos (tasks) from a database and rendered into HTML.

When starting the code generates *data.sql* that can be imported as the test data.

*The benchmarks were run on WSL/Windows 11, Postgres 15, Go 1.19.3, Ryzen 3900x/12c, 32gb/3600, WD SN850 SSD.*

The tests are run with [k6](https://k6.io/)

    k6 run --vus 10 --duration 10s k6/graphql.js

Each test is run with 1,5,10,15,20,25,30,35,40 VUs simulating more concurrent users each time.

## Results


