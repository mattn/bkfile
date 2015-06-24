package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var (
	format = flag.String("f", "${base}/${name}_${date}.${ext}", "format")
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	for _, arg := range flag.Args() {
		if _, err := os.Stat(arg); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
			os.Exit(1)
		}
		from, err := filepath.Abs(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
			os.Exit(1)
		}
		base, name := filepath.Split(from)
		ext := filepath.Ext(name)
		if ext != "" {
			if len(name) == len(ext) {
				name = ext
				ext = ""
			} else {
				ext = ext[1:]
			}
		}

		var unk string
		to := os.Expand(*format, func(s string) string {
			switch s {
			case "base":
				return base
			case "file":
				return from
			case "name":
				if ext != "" {
					return name[:len(name)-len(ext)-1]
				}
				return name
			case "ext":
				return ext
			case "date":
				return time.Now().Format("20060102")
			case "datetime":
				return time.Now().Format("20060102030405")
			default:
				unk = s
			}
			return ""
		})
		if unk != "" {
			fmt.Fprintf(os.Stderr, "%s: unknown variable ${%s}\n", os.Args[0], unk)
			os.Exit(1)
		}
		to = filepath.ToSlash(to)

		n := 1
		for {
			var target string
			if n > 1 {
				ext = filepath.Ext(to)
				target = to[:len(to)-len(ext)] + fmt.Sprintf("(%d)", n) + ext
			} else {
				target = to
			}
			if _, err = os.Stat(target); !os.IsNotExist(err) {
				n++
				continue
			}
			if err = copyFile(from, target); err != nil {
				fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
				os.Exit(1)
			}
			break
		}
	}
}
