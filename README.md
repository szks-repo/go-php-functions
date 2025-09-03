# go-php-functions
go-php-functions is a Go language project that re-implements widely used PHP functions. The goal is to provide Go developers with familiar functionalities, especially for those migrating from PHP environments or working on projects requiring similar string and number manipulations.

## 🧪 Running Tests
To run the tests for the project, navigate to the root directory and execute:

```sh
go test ./...
```

## Functions

| PHP function | Corresponding function |
|---|---|
| `nl2br` | `phpstring.Nl2br` |
| `number_fomrat` | `phpstring.NumberFormat` |
| `wordwrap` | `phpstring.WordWrap` |
| `implode` | Use Go stdlib `strings.Join` |
| `explode` | Use Go stdlib `strings.Split` |
| `explode` with limit | Use Go stdlib `strings.SplitN` |

## 📄 License
This project is licensed under the MIT License. See the LICENSE file for details.
