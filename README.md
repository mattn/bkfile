# bkfile

copy files for backup

## Usage

```
Usage of bkfile:
  -f string
        format (default "${base}/${name}_${date}.${ext}")
```

```
$ ls
important-file.txt

$ bkfile important-file.txt
$ ls
important-file.txt
important-file_20150624.txt

```

Do you use Windows? Then make shortcut file to `bkfile.exe` and put it into your `%APPDATA%\Microsoft\Windows\SendTo`.

## Requirements

golang

## Installation

```
go get github.com/mattn/bkfile
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a mattn)
