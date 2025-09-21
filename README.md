# progog

A simple tool that retrieves Bitcoin blockchain data from `blockchain.info` API and exports it as a basic SWI Prolog KB. Since a public API is used, a mechanism for a maximum number of retry attempts is implemented for each block request. Made because I was bored to surf 30 blocks through.

* Exporting is made with help of the `reflect` package
* No Go package dependencies

