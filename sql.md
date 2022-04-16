docker run --name mysql -e MYSQL_ROOT_PASSWORD=rootpass -e MYSQL_DATABASE=urls -e MYSQL_USER=go -e MYSQL_PASSWORD=gopass -d -p3306:3306 mysql:debian

- [ ] Add volume to mysql container
- [ ] Move containers to compose
- [ ] Change error handling, rm fatal

create table urls ( id int not null unique auto_increment, Full varchar(2048) not null, Short varchar(10) not null unique, constraint pk_short primary key (id));

Insert into urls values (1, 'https://www.google.com/search?q=golang+log&ei=dlhAYur_Hcik3APNqZ_QCw&ved=0ahUKEwjqktiupeb2AhVIEncKHc3UB7oQ4dUDCA4&uact=5&oq=golang+log&gs_lcp=Cgdnd3Mtd2l6EAMyBQgAEIAEMgUIABCABDIFCAAQgAQyBQgAEIAEMgUIABCABDIFCAAQgAQyBQgAEIAEMgUIABCABDIFCAAQgAQyBQgAEIAEOgcIABBHELADOgcIABCwAxBDOgQIABBDSgQIQRgASgQIRhgAUMgGWP4IYOAJaAJwAXgAgAFIiAHAAZIBATOYAQCgAQHIAQrAAQE&sclient=gws-wiz', '3JNb6D9');

414 (Request-URI Too Long)
