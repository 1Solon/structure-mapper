<div align="center">

# Structure Mapper

_Structure Mapper is a Go-based utility that prints a visual representation of a directory tree._

</div>

<div align="center">

![GitHub Repo stars](https://img.shields.io/github/stars/1Solon/structure-mapper?style=for-the-badge)
![GitHub forks](https://img.shields.io/github/forks/1Solon/structure-mapper?style=for-the-badge)

</div>

## Features ✨

- Prints a directory tree with folder icons.
- Excludes specified directories from the tree.
- Simple and easy to use.

## Installation 🚀

1. Clone the repository:

   ```sh
   git clone https://github.com/1Solon/structure-mapper.git
   ```

2. Navigate to the project directory:

   ```sh
   cd structure-mapper
   ```

3. Build the project:

   ```sh
   make build
   ```

4. Run the script:

   ```sh
   make run <target_directory> [exceptions...]
   ```

## Usage 📚

To run the script, use the following command:

```sh
structure-mapper.exe <target_directory> [exceptions...]
```

## Example 📖

Running the following command:

```sh
structure-mapper.exe /path/to/target_directory dir_to_exclude1 dir_to_exclude2
```

Might produce an output like this:

```md
📁
├──📁 folder1
│   ├──📁 subfolder1
│   └──📁 subfolder2
└──📁 folder2
    └──📁 subfolder3
```

## Contributing 🤝

Contributions are welcome! If you would like to contribute to the project, please open an issue or a pull request.

## License 📝

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact 📧

For any questions or suggestions, please open an issue.
