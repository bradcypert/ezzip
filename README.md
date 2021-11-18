# ezzip
Simple, easy zip/unzip cli with optional AES256 encryption

## Installation
### Homebrew
Coming soon.

### For Go Devs
```
go install github.com/bradcypert/ezzip
```

## Usage

### Zip a directory
```
ezzip my_dir
```

### Unzip a directory

```
ezzip my_dir.zip
```

### Optional encryption
```
ezzip my_dir --encrypt
...
use key: abc123 to decrypt
```

### Unzip and Decrypt
```
ezzip my_dir.zip --key=abc123
```
