create SEQUENCE seq_user;

CREATE TABLE "public"."user" (
  "id" int8 NOT NULL DEFAULT nextval('seq_user'::regclass),
  "name" varchar(255),
  "password" varchar(500),
  "email" varchar(255),
  "phone_numb" varchar(255),
  "fb" varchar(255),
  "status" int4 DEFAULT 1,
  PRIMARY KEY ("id")
)
;

create SEQUENCE seq_resto;
CREATE TABLE "public"."resto" (
  "id" int8 NOT NULL DEFAULT nextval('seq_resto'::regclass),
  "name" varchar(255),
  "resto_code" varchar(255),
  "desc" varchar(500),
  "address" varchar(255),
  "city" varchar(255),
  "province" varchar(255),
  "status" int4 DEFAULT 1,
  PRIMARY KEY ("id")
)
;

create SEQUENCE seq_e_menu_item;
CREATE TABLE "public"."e_menu_item" (
  "id" int8 NOT NULL DEFAULT nextval('seq_e_menu_item'::regclass),
  "group_id" int8,
  "resto_id" int8,
  "name" varchar(255),
  "desc" varchar(500),
  "img_url" varchar(255),
  "price" decimal(10,2),
  "status" int4 DEFAULT 1,
  PRIMARY KEY ("id")
)
;

create SEQUENCE seq_e_menu_group;
CREATE TABLE "public"."e_menu_group" (
  "id" int8 NOT NULL DEFAULT nextval('seq_e_menu_group'::regclass),
  "name" varchar(255),
  "img_url" varchar(255),
  "status" int4 DEFAULT 1,
  PRIMARY KEY ("id")
)
;


INSERT INTO public.resto
("name", "desc", address, city, province)
VALUES('Resto 1', 'Indonesian Food', 'Jl. Kemanggisan', 'Jakarta', 'Jawa barat');

INSERT INTO public.resto
("name", "desc", address, city, province)
VALUES('Resto 2', 'Chinese Food', 'Jl. Sudirman', 'Jakarta', 'Jawa barat');

INSERT INTO public.resto
("name", "desc", address, city, province)
VALUES('Resto 3', 'Western Food', 'Jl. Tanjung Duren', 'Jakarta', 'Jawa barat');




INSERT INTO public.e_menu_group
("name", img_url)
VALUES('Nasi', 'https://i.picsum.photos/id/1025/200/300.jpg');

INSERT INTO public.e_menu_group
("name", img_url)
VALUES('Mie', 'https://i.picsum.photos/id/0/200/300.jpg');

INSERT INTO public.e_menu_group
("name", img_url)
VALUES('Ikan', 'https://i.picsum.photos/id/204/200/300.jpg');

INSERT INTO public.e_menu_group
("name", img_url)
VALUES('Paket', 'https://i.picsum.photos/id/250/400/400.jpg');



INSERT INTO public.e_menu_item
(group_id, resto_id, "name", "desc", img_url, price)
VALUES(1, 1, 'Nasi Goreng Ayam', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.
            Pellentesque at ipsum ut metus vestibulum pharetra eget at mauris.
            Nulla pellentesque molestie sem, vel vulputate ligula commodo quis.
            Aenean vestibulum id nisl ut iaculis. Cras feugiat felis in ex aliquam, vitae ullamcorper lorem pellentesque.
            Mauris hendrerit urna eget
            rhoncus sagittis. Donec eleifend, quam id commodo dignissim, risus orci pulvinar nibh, et commodo nisl metus at velit.
            Quisque ultrices augue quis condimentum fringilla.', 'https://i.picsum.photos/id/1025/400/400.jpg', 15000);
           

INSERT INTO public.e_menu_item
(group_id, resto_id, "name", "desc", img_url, price)
VALUES(1, 2, 'Nasi Goreng Sapi', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.
            Pellentesque at ipsum ut metus vestibulum pharetra eget at mauris..', 'https://i.picsum.photos/id/1003/400/400.jpg', 25000);
            
INSERT INTO public.e_menu_item
(group_id, resto_id, "name", "desc", img_url, price)
VALUES(1, 1, 'Nasi Goreng Ayam 2', 'Nasi Goreng', 'https://i.picsum.photos/id/104/400/400.jpg', 20000);

INSERT INTO public.e_menu_item
(group_id, resto_id, "name", "desc", img_url, price)
VALUES(1, 1, 'Nasi Goreng seafood 2', 'Nasi Goreng udang, cumi', 'https://i.picsum.photos/id/1050/400/400.jpg', 22200);

INSERT INTO public.e_menu_item
(group_id, resto_id, "name", "desc", img_url, price)
VALUES(2, 1, 'Mie Pangsit', 'Mie Pangsit udang', 'https://i.picsum.photos/id/1056/400/400.jpg', 27500);

INSERT INTO public.e_menu_item
(group_id, resto_id, "name", "desc", img_url, price)
VALUES(2, 2, 'Indomie goreng ', 'Indomie goreng telor', 'https://i.picsum.photos/id/1074/400/400.jpg', 2000);

INSERT INTO public.e_menu_item
(group_id, resto_id, "name", "desc", img_url, price)
VALUES(3, 1, 'Ikan Gurame asam manis ', 'Ikan Gurame asam manis - Ikan Gurame asam manis', 
'https://i.picsum.photos/id/113/400/400.jpg', 199000);

INSERT INTO public.e_menu_item
(group_id, resto_id, "name", "desc", img_url, price)
VALUES(4, 1, 'Pahe 1 ', 'Nasi ikan bumbu cabe ijo', 'https://i.picsum.photos/id/132/400/400.jpg', 35000);

INSERT INTO public.e_menu_item
(group_id, resto_id, "name", "desc", img_url, price)
VALUES(4, 1, 'Pahe 2 ', 'yam bumbu krispi geprek pake cabe merah', 'https://i.picsum.photos/id/124/400/400.jpg', 999999);

INSERT INTO public.e_menu_item
(group_id, resto_id, "name", "desc", img_url, price)
VALUES(1, 1, 'Nasi Goreng spesial ', 'Nasi Goreng spesial Telorrrr', 'https://i.picsum.photos/id/139/400/400.jpg', 12345);

INSERT INTO public.e_menu_item
(group_id, resto_id, "name", "desc", img_url, price)
VALUES(1, 1, 'Nasi Goreng putih ', 'Nasi Putih goreng Telorrrr', 'https://i.picsum.photos/id/146/400/400.jpg', 987987);

INSERT INTO public.e_menu_item
(group_id, resto_id, "name", "desc", img_url, price)
VALUES(1, 2, 'Nasi Goreng rendang ', 'Nasi Putih goreng rendang', 'https://i.picsum.photos/id/152/400/400.jpg', 123);

INSERT INTO public.e_menu_item
(group_id, resto_id, "name", "desc", img_url, price)
VALUES(1, 1, 'Nasi Goreng Nanas ', 'Nasi Putih Nanas Telorrrr', 'https://i.picsum.photos/id/159/400/400.jpg', 5555)
;

