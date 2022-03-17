# Playground for Protobuf quick experiments.

## Token Encoding options. main.go output:

```
$ go run main.go
=> Basic base58-proto encoding
27ot46j87D5pqfZTzLqGRA8Wcw7hkj2uv9RV9Zay4VXD4fo3jpZQWNxvcvaSfNecT56tzmWmwjyR2TgoHNN6EZy4tiZ4cgg6cUxDeJGQUuv8eH7wuvGNsnbCgQ2Ket8xtwv8xRnMtq1dgf37imv6XpmVzo7AkiwpmcvLybAnTqnmYHQP21kbZfaCB26H3ZaYJyDKR5JipGjzCKWPDrGkFfy3BiYinB9kZvtvTZUiaC9n2ZoCzRZeiN (246 bytes)

=> base58-proto encoding with Oauth decoded to raw bytes
3n4e6cVf49aiS8UnWpYRQWzj4vCJxi7AJzy3JueZ4A45iZgPZHfe1vDAbpfwXxzUH4NoaGvhhL3oLcqZFnnChdXyBtznntVsV52N7GmLtw2cWXMesxr56X7qSdHhdLK18RhgZWm4YBwafBAxSg9yDabL2jkZ1Wyo9no3ZkefkWBHLFfduxL1pgFMZNFNxhuxpkztjqYedTHB4UvBatUjM (213 bytes)

=> base58-proto encoding with raw oauth and raw location
2j9avAyvcTfuMwCGed5tv2eomTELPt4o3r4btZYML4skBmdhvEQNhRSdKcCo4F6uHmrbGt8xbXtyrQxBASJZFipJv52mncgZHTpWv1Vn5dHgU5SbYH14hbF1WuTrAG2iVJU4g5Wawd (138 bytes)
```