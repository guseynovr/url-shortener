#!/bin/sh

sleep 10
# while ! mysql -h mysql -ugo -pgopass  -e ";" 2>/dev/null ; do
while ! mysql -h mysql -ugo -pgopass  -e ";" ; do
		sleep 1
		echo "Waiting for db"
done

url_shortener
