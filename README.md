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
  +--------+--------------------------------+
  | USERID |             TITLE              |
  +--------+--------------------------------+
  |     10 | at nam consequatur ea labore   |
  |        | ea harum                       |
  |     10 | temporibus sit alias delectus  |
  |        | eligendi possimus magni        |
  |     10 | laboriosam dolor voluptates    |
  |     10 | quas fugiat ut perspiciatis    |
  |        | vero provident                 |
  |     10 | quaerat velit veniam amet      |
  |        | cupiditate aut numquam ut      |
  |        | sequi                          |
  +--------+--------------------------------+
```

## License

MIT

## Author

[Takuya Arita](https://github.com/ariarijp)
