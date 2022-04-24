# url-shortener

A simple containerized server to transofrm any URL to a standart length link.

<img width="409" alt="Screenshot 2022-04-24 at 22 14 55" src="https://user-images.githubusercontent.com/32015630/164992737-a5ebc5fb-3180-4cf2-a451-6aaafa713c1a.png">

URL's are stored in a MYSQL database.

Suffix for short URL is generated from a different representation of 'id' field of the URL in the db. 'id' field is transformed to the numeral system,
which consists of symbols [a-zA-Z0-9] and "-" (63 symbols in total).

Ten symbols long suffix allows to store 63^10 unique links of the same length.
