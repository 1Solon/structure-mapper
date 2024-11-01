<div align="center">

# Structure Mapper

_Structure Mapper is a Go-based utility that prints a visual representation of a directory tree._

</div>

<div align="center">

![GitHub Repo stars](https://img.shields.io/github/stars/1Solon/structure-mapper?style=for-the-badge)
![GitHub forks](https://img.shields.io/github/forks/1Solon/structure-mapper?style=for-the-badge)

</div>

## ✨ Features

- Prints a directory tree with folder icons.
- Excludes specified directories from the tree.
- Simple and easy to use.

## 🚀 Installation

You can install the tool using Go:

```sh
go install github.com/1Solon/structure-mapper@latest
```

## 📚 Usage

To run the script, use the following command:

```sh
structure-mapper <target_directory>
```

You can also exclude directories from the tree by specifying them as arguments:

```sh
structure-mapper -e dir_to_exclude1, dir_to_exclude2 <target_directory>
```

To display the directory tree up to a certain depth, use the `-d` flag:

```sh
structure-mapper -d 2 <target_directory>
```

If you need a refresher on the available options, use the `-h` flag:

```sh
structure-mapper -h
```

## 📖 Example

Running the following command:

```sh
structure-mapper -e dir_to_exclude1, dir_to_exclude2 -d 2 /path/to/target_directory
```

Might produce an output like this:

```md
📁
├──📁 folder1
│ ├──📁 subfolder1
│ └──📁 subfolder2
└──📁 folder2
   └──📁 subfolder3
```

## 🤝 Contributing

Contributions are welcome! If you would like to contribute to the project, please open an issue or a pull request.

To set up the project locally, follow these steps:

1. Clone the repository:

   ```sh
   git clone https://github.com/1Solon/structure-mapper.git
   ```

2. Navigate to the project directory:

   ```sh
   cd structure-mapper
   ```

3. Install the tool

   Installing with Make:

   ```sh
   make install
   ```

   Installing with Go:

   ```sh
   go install github.com/1Solon/structure-mapper
   ```

## 📝 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## 📧 Contact

For any questions or suggestions, please open an issue.
