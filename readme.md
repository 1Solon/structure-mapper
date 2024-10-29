<div align="center">

# Structure Mapper

_Structure Mapper is a Go-based utility that prints a visual representation of a directory tree._

</div>

<div align="center">

![GitHub Repo stars](https://img.shields.io/github/stars/1Solon/structure-mapper?style=for-the-badge)
![GitHub forks](https://img.shields.io/github/forks/1Solon/structure-mapper?style=for-the-badge)

</div>

## Features

- Prints a directory tree with folder icons.
- Excludes specified directories from the tree.
- Simple and easy to use.

## Installation

1. Ensure you have [Go](https://golang.org/dl/) installed.

2. Clone the repository:

```sh
git clone <repository-url>
```

1. Navigate to the project directory:

```sh
cd structure-mapper
```

## Usage

To run the script, use the following command:

```sh
go run repo-mapper.go <target_directory> [exceptions...]
```

### Example

```sh
go run repo-mapper.go /path/to/target_directory dir_to_exclude1 dir_to_exclude2
```

This will print the directory tree of `/path/to/target_directory`, excluding `dir_to_exclude1` and `dir_to_exclude2`.

## Contributing

Contributions are welcome! If you would like to contribute to the project, please open an issue or a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any questions or suggestions, please open an issue.