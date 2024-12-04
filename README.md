# About

`pwg` or Password Generator is a command line tool to generate cryptographically secure passwords.

## Installation

Clone the repository and run the following commands to manually install

```bash
cd pwd
go mod install && go mod verify
mkdir bin && go build -o bin/pwg
chmod +x ./bin/pwg
./bin/pwg -h
```

Set your path variable accordingly to enable the tool at the command line.

## Usage

Generate a 32 character password string (default behavior).

```bash
pwg
```

Generate a 64 charactor password string.

```bash
pwg -s 64
```

Generate a password consisting of four random words separated by dashes.
Omit the `-d` flag to not use dashes. Use the `-s` flag to specify number of words (default is 4).

```bash
pwg -t -d
```

### Flags

- `-t` Create a random text/word-based password instead of random characters (default behavior).
- `-s` Specify size (either individual chars or words when using the `-t` flag).
- `-d` Separate words with dashes (for use only with the `-t` flag).
