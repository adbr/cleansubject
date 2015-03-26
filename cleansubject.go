// 2015-03-25 Adam Bryt

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const usageStr = `usage: cleansubject -from <old> [-to <new> -n <num>]
	-from <old>: zamieniany fragment tekstu
	-to <new>: na jaki tekst zamienić (domyślnie "")
	-n <num>: ile wystąpień tekstu <old> zamienić (domyślnie 1)
	-h: help
`

const helpStr = `Filtr cleansubject zamienia tekst w wierszu 'Subject:' nagłówka maila.

Czyta wiersze tekstu maila z stdin i drukuje wiersze na stdout, zamieniając w
wierszu nagłówka zaczynającym się od 'Subject:', fragment <old> na <new>.
Tekst <old> nie może być fragmentem napisu 'Subject:' ponieważ spowoduje
uszkodzenie nagłówka. Zmieniany jest tylko wiersz w nagłówku maila - wiersze w
treści maila nie są zmieniane.

Sposób użycia:
	cleansubject -from <old> [-to <new> -n <num>]

Opcje:
	-from <old>: zamieniany fragment tekstu
	-to <new>: na jaki tekst zamienić (domyślnie "")
	-n <num>: ile wystąpień tekstu <old> zamienić (domyślnie 1)
	-h help
`

var (
	from     = flag.String("from", "", "stary string")
	to       = flag.String("to", "", "nowy string")
	num      = flag.Int("n", 1, "liczba zastąpień")
	helpFlag = flag.Bool("h", false, "help")
)

func main() {
	flag.Usage = usage
	flag.Parse()

	if *helpFlag {
		help()
	}

	if *from == "" {
		usage()
	}

	err := cleansubject(os.Stdout, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
}

func usage() {
	fmt.Fprint(os.Stderr, usageStr)
	os.Exit(1)
}

func help() {
	fmt.Print(helpStr)
	os.Exit(0)
}

func cleansubject(w io.Writer, r io.Reader) error {
	scanner := bufio.NewScanner(r)
	header := true
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			// header od treści oddziela pusty wiersz
			header = false
		}
		if header {
			if strings.HasPrefix(s, "Subject:") {
				if strings.Index(s, *from) != -1 {
					s = strings.Replace(s, *from, *to, *num)
				}
			}
		}
		_, err := fmt.Fprintln(w, s)
		if err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
