/*
Filtr cleansubject zamienia tekst w wierszu 'Subject:' nagłówka maila.

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
	-h: help

Przykład użycia w procmail w celu usunięcia nazwy listy mailowej '[go-nuts]'
z wierszy 'Subject':

	:0
	* ^Sender: golang-nuts@googlegroups.com
	{
		# Usuń '[go-nuts] ' z Subject
		:0fw
		| $HOME/bin/cleansubject -from "[go-nuts] " -to ""

		:0e
		{
			EXITCODE=$?
		}

		:0
		$MAILDIR/incoming/golang-nuts/
	}

*/
package main
