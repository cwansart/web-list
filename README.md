# Web List

Web-List is a Go application that starts a web server on port 3000 and returns a list of files and directories in the specified path. If no path is provided, it defaults to the current directory.

## Installation

To install and run the project, follow these steps:

1. Clone the repository:

    ```bash
    git clone https://github.com/cwansart/web-list.git
    ```

1. Navigate to the project directory:

    ```bash
    cd web-list
    ```

1. Build the project:

    ```bash
    go build
    ```

1. Run the project:

    ```bash
    ./web-list
    ```

## Usage

To use the program, run the following command:

```bash
./web-list /path/to/directory
```

If no path is provided, the program will use the current directory.
