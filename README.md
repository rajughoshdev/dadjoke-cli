# Dadjok-cli

Dadjok-cli is a command-line tool that provides you with a random dad joke to brighten up your day. It's a simple utility written in Go that fetches dad jokes from a api https://icanhazdadjoke.com/ and displays them on your terminal.

## Installation

To install `dadjok-cli`, you need to have Go installed on your system. If you don't have Go installed, you can download it from the [official Go website](https://golang.org/dl/).

Once you have Go installed, you can use the following steps to install `dadjok-cli`:

1. Clone this repository:
   ```sh
   git clone https://github.com/yourusername/dadjok-cli.git
   ```

2. Change into the project directory:
   ```sh
   cd dadjok-cli
   ```

3. Install the project's dependencies:
   ```sh
   go mod download
   ```

4. Build the `dadjok-cli` executable:
   ```sh
   go build -o dadjoke main.go
   ```

5. Add the executable to your PATH (optional):
   ```sh
   mv dadjoke /usr/local/bin  # Or any directory in your PATH
   ```

## Usage

Once you have `dadjok-cli` installed, you can use it to get random dad jokes. Here are some usage examples:

- To get a random dad joke:
  ```sh
  dadjoke get
  ```

- To see available commands and options:
  ```sh
  dadjoke --help
  ```

- To generate shell autocompletion scripts:
  ```sh
  dadjoke completion [shell]
  ```

## Contributing

If you'd like to contribute to the project, you're welcome to submit issues or pull requests on the [GitHub repository](https://github.com/yourusername/dadjok-cli). Please ensure that your code follows the project's coding standards and conventions.

## License

This project is licensed under the [MIT License](LICENSE).

---

Enjoy the dad jokes and happy coding!
```

Replace `yourusername` with your actual GitHub username in the links. Make sure to provide clear and accurate instructions for installing, using, and contributing to your project. Feel free to customize the README to match your project's specifics and your preferred style.