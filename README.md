# progog

## Description

A simple tool that retrieves Bitcoin blockchain data from `blockchain.info` API and exports it as a basic SWI Prolog KB. Since a public API is used, a mechanism for a maximum number of retry attempts is implemented for each block request. Made because I was bored to surf 30 blocks through.

* Exporting is made with help of the `reflect` package
* No Go package dependencies

## Parallelism note

We retrieve next block hash from the HTTP request, hence, parallelism is very limited and I did not find an adequate applicaiton of parallelism here.