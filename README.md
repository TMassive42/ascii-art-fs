# ASCII-ART-FS

ASCII-ART-FS is a program written in Go that converts input strings into graphic representations using ASCII characters and allow specification of banner files if need to be printed in a different font.

## Features

- Converts strings into ASCII art
- Supports numbers, letters, spaces, special characters, and newline characters ('\n')
- Utilizes specific graphical templates for ASCII representation
- Access the burner files in the third argument

## Installation

1. Clone the repository:

    ```bash
    git clone https://learn.zone01kisumu.ke/git/togondol/ascii-art-fs
    ```

2. Navigate to the project directory:

    ```bash
    cd ascii-art-fs
    ```

## Usage

To generate ASCII art for a string with a default font type (standard.txt), run the following command:

```go
$ go run . "input string"
```

Example:

```go
$ go run . "Hello\n" | cat -E
```

Output:

```bash
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$                                                  
```
If you want to generate with font in the thinkertoy.txt file, use this command :

```go
 go run . "input string" thinkertoy
```

Example:

```go
 go run . "Hello" thinkertoy | cat -E
```

Output:

```
                 $
o  o     o o     $
|  |     | |     $
O--O o-o | | o-o $
|  | |-' | | | | $
o  o o-o o o o-o $
                 $
                 $
```
If you want to generate with font in the  shadow.txt file, using this command :

```go
 go run . "input string" shadow
```

Example:

```go
$ go run . "Hello" shadow | cat -E
```

Output:

```
                                 $
_|    _|          _| _|          $
_|    _|   _|_|   _| _|   _|_|   $
_|_|_|_| _|_|_|_| _| _| _|    _| $
_|    _| _|       _| _| _|    _| $
_|    _|   _|_|_| _| _|   _|_|   $
                                 $
                                 $
```
## Testing 
To run the tests present do the following:

1. Open the inner folder using this command:

```
cd ascii/
```
2. Run the test using this command:

```
go test -v
```

## File Formats

- `standard.txt`: Standard ASCII character set
- `shadow.txt`: Shadowed ASCII character set
- `thinkertoy.txt`: ASCII character set with thinkertoy style

## Contributing

If you have suggestions for improvements, bug fixes, or new features, feel free to open an issue or submit a pull request.

## Author

This project was build and maintained by:

[Thadeus Ogondola](https://learn.zone01kisumu.ke/git/togondol)

[Kherld Hussein](https://learn.zone01kisumu.ke/git/khahussein)

[Bernad Okumu](https://learn.zone01kisumu.ke/git/bernaotieno)

