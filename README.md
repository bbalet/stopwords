**stopwords** is a go package that removes stop words from a text content.
If instructed to do so, it will remove HTML tags and parse HTML entities.
The objective is to prepare a text in view to be used by natural processing algos
or text comparison algorithms such as SimHash.

It uses a curated list of the most frequent words used in these languages:
 * arabic
 * bulgarian
 * czech
 * dutch
 * english
 * french
 * finnish
 * german
 * greek
 * spanish
 * persian
 * hungarian
 * italian
 * norwegian
 * polish
 * portuguese
 * romanian
 * russian
 * swedish
 * turkish

## How to use this package

You can find an example here https://github.com/bbalet/gorelated where **stopwords**
package is used in conjunction with SimHash algorithm in order to find a list of
related content :

    import (
	      "github.com/bbalet/stopwords"
    )

    cleanContent := stopwords.CleanContent(content, "en", true)

Where *cleanContent* and *content* are strings containing an english text.
And *en* is the ISO 639-1 code for English.
https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes

## Credits

Most of the lists were built by IR Multilingual Resources at UniNE
http://members.unine.ch/jacques.savoy/clef/index.html

## license

**stopwords** is released under the BSD license.
