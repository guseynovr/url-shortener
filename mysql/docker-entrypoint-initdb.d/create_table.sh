#!/bin/bash

echo 'use urls; create table urls ( id int not null unique auto_increment, Full varchar(2048) not null, Short varchar(10) not null unique, constraint pk_short primary key (id));' > urls.sql
mysql -ugo -pgopass < urls.sql
