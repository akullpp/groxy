# Reverse Proxy

Simple reverse proxy to proxy requests based on the first path segment to the correct service in order to substitute services partially in local development.

## Usage

Create a `.env.local` of the following form:

```
FOO=http://localhost:1000
BAR=http://localhost:2000
```

* Requests which paths start with `FOO` (e.g. `foo/baz`) will be forwarded to `http://localhost:1000` (e.g `http://localhost:1000/foo/baz`).

* Requests which paths start with `BAR` (e.g. `bar/baz`) will be forwarded to `http://localhost:2000` (e.g `http://localhost:1000/bar/baz`).

* Requests which first path segment do not match will be forwarded to the DEFAULT value in `.env` (e.g. `https://api.dev-le.com`).

## Example

If you want to route requests that start with `user` to `http://localhost:8100`:

```
USER=http://localhost:8100
```

## Dependencies

I extracted the reading of the reading of environment files to another package which also only uses the standard library.

## TODO

* Watch changes in the dotenv files.
