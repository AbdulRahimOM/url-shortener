```markdown
# URL Shortener

This project provides a simple API to generate shortened URLs.

## Endpoint

**`POST /generate`**

### Request Body

The request body should be sent as `x-www-form-urlencoded`.

| Field     | Expected Format                                                              |
|-----------|------------------------------------------------------------------------------|
| `longUrl` | `https://www.postgresql.org/docs/current/app-psql.html`                      |


### Example Request
```http
POST /generate HTTP/1.1
Host: example.com
Content-Type: application/x-www-form-urlencoded

longUrl=https://www.postgresql.org/docs/current/app-psql.html
```

## Responses

### Successful URL Generation

```json
{
    "status": true,
    "message": "URL generated successfully",
    "shortened_url": "<base_url>/<path>"
}
```

### URL Already Registered

```json
{
    "status": true,
    "message": "This URL is already registered",
    "shortened_url": "<base_url>/<path>"
}
```

### Invalid input long-URL
```json
{
    "status": false,
    "message": "Invalid URL",
    "error": <reason/error>
}
```


## Usage

1. Send a POST request to `/generate` with the `longUrl` field containing the URL you wish to shorten.
2. The API will return the shortened URL or indicate if the URL is invalid or already registered.

## Note: Requests to this shortened url is automatically redirected to the original long url
### If shorturl is not registered
```json
{
    "status": false,
    "message": "URL not found",
    "error": "record not found"
}
```