# chainsaw

Execute SQL on JSONL.

## Installation

```bash
$ go get -u --tags=sqlite_json1 github.com/ariarijp/chainsaw
```

## Usage

```bash
$ curl -s https://jsonplaceholder.typicode.com/posts \
| jq -c ".[]" \
| chainsaw "SELECT JSON_EXTRACT(json, '$.userId') userId, JSON_EXTRACT(json, '$.title') title FROM _ ORDER BY JSON_EXTRACT(json, '$.id') DESC LIMIT 5"
```

## License

MIT

## Author

[Takuya Arita](https://github.com/ariarijp)
