# rss-agg

Lightweight RSS Feed aggregator.

## Motivation

It's always been a pain to keep up with my favourite dev blogs like [DHH's](https://world.hey.com/dhh), [Dave Cheney's](https://dave.cheney.net/), etc. So I don't need to remember 10 more URLs in my daily life this creation was born!

## Installation

Compiling from source:

```bash
go get github.com/dimadudin/rss-agg

export $PORT=<port>
export $CONN=<dbconn>

go build .
```

Where port is the `port` the server will run on and `dbconn` is the connection string to the PostgreSQL database.

## Usage

There are multiple endpoints available on the server

```text
├── "v1/users"
├── "v1/feeds"
├── "v1/feed_follows"
├── "v1/posts"
```

The users endpoint can be used to create  and view users
The feeds endpoint can be used to add feeds and view all feeds
The feed_follows endpoint can be used to follow, unfollow feeds and view all follows
The posts endpoint can be used to get the posts from a feed

This server uses API Keys for authentication
This server runs a scraping worker that updates the feeds, that haven't been fetched in a while.
