# beuys

`beuys` is a Go package for generating a minimalist RSS feed of the felt.com weblog.

## Motivation

Hopefully this package won't be necessary soon but the `felt.com` website has had a weblog without an RSS (or Atom) feed for over a year now so it still is.

## Tools

```
$> make cli
go build -mod vendor -ldflags="-s -w" -o bin/feed cmd/feed/main.go
go build -mod vendor -ldflags="-s -w" -o bin/feed-server cmd/feed-server/main.go
```

### feed

Generate an RSS feed and print to STDOUT.

```
$> ./bin/feed | xmllint - | grep title
    <title>An auto-generated syndication feed for the felt.com weblog</title>
      <title>May Spotlight: 10 Best Felt Community Maps</title>
      <title>Felt for investigative journalism: mapping Malaysian deforestation</title>
      <title>From SVG to Canvas – Part 2: A new way of building interactions</title>
      <title>Mapping Protected Territories: Westchester Land Trust Case Study</title>
      <title>From SVG to Canvas – Part 1: Making Felt faster</title>
      <title>Graceful Startup and Shutdown for Phoenix Applications</title>
      <title>April Spotlight: 10 Best Felt Community Maps</title>
      <title>GeoTIFFs, XYZ Urls, and Raster Imagery – Now in Felt</title>
      <title>Routing Patterns for Manageable Phoenix Channels</title>
      <title>March Spotlight: 10 Best Felt Community Maps</title>
      <title> Modern Client Collaboration: Alta Planning + Design Case Study</title>
      <title>Creating Maps on the Web with QGIS &amp; Felt</title
```

### feed-server

```
$> ./bin/feed-server
2023/06/06 19:06:27 Listen for requests at http://localhost:8080
```

```
$> curl -s localhost:8080/ | xmllint - | grep title
    <title>An auto-generated syndication feed for the felt.com weblog</title>
      <title>May Spotlight: 10 Best Felt Community Maps</title>
      <title>Felt for investigative journalism: mapping Malaysian deforestation</title>
      <title>From SVG to Canvas – Part 2: A new way of building interactions</title>
      <title>Mapping Protected Territories: Westchester Land Trust Case Study</title>
      <title>From SVG to Canvas – Part 1: Making Felt faster</title>
      <title>Graceful Startup and Shutdown for Phoenix Applications</title>
      <title>April Spotlight: 10 Best Felt Community Maps</title>
      <title>GeoTIFFs, XYZ Urls, and Raster Imagery – Now in Felt</title>
      <title>Routing Patterns for Manageable Phoenix Channels</title>
      <title>March Spotlight: 10 Best Felt Community Maps</title>
      <title> Modern Client Collaboration: Alta Planning + Design Case Study</title>
      <title>Creating Maps on the Web with QGIS &amp; Felt</title>
```

#### Lambda

The easiest way to deploy the `feed-server` tool is as a Lambda Function URL. First run the handy `lambda-server` Makefile target, like this:

```
$> make lambda-server
if test -f main; then rm -f main; fi
if test -f feed-server.zip; then rm -f feed-server.zip; fi
GOOS=linux go build -mod vendor -ldflags="-s -w" -o main cmd/feed-server/main.go
zip feed-server.zip main
  adding: main (deflated 58%)
rm -f main
```

Create a new (Go) Lambda function using the `feed-server.zip` file and  configure it as a Lambda Function URL. The function does not require any special permissions and the only environment variables you'll need to set are:

| Name | Value |
| --- | --- |
| FELT_SERVER_URI | `functionurl://` |

For example:

```
$> curl -s https://{FUNCTION_URL_ID}.lambda-url.{REGION}.on.aws | xmllint - | grep title
    <title>An auto-generated syndication feed for the felt.com weblog</title>
      <title>May Spotlight: 10 Best Felt Community Maps</title>
      <title>Felt for investigative journalism: mapping Malaysian deforestation</title>
      <title>From SVG to Canvas – Part 2: A new way of building interactions</title>
      <title>Mapping Protected Territories: Westchester Land Trust Case Study</title>
      <title>From SVG to Canvas – Part 1: Making Felt faster</title>
      <title>Graceful Startup and Shutdown for Phoenix Applications</title>
      <title>April Spotlight: 10 Best Felt Community Maps</title>
      <title>GeoTIFFs, XYZ Urls, and Raster Imagery – Now in Felt</title>
      <title>Routing Patterns for Manageable Phoenix Channels</title>
      <title>March Spotlight: 10 Best Felt Community Maps</title>
      <title> Modern Client Collaboration: Alta Planning + Design Case Study</title>
      <title>Creating Maps on the Web with QGIS &amp; Felt</title>
```

## See also

* https://felt.com/blog