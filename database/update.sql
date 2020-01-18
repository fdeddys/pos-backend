create SEQUENCE seq_user;

CREATE TABLE "public"."user" (
  "id" int8 NOT NULL DEFAULT nextval('seq_user'::regclass),
  "name" varchar(255),
  "password" varchar(500),
  "email" varchar(255),
  "phone_numb" varchar(255),
  "fb" varchar(255),
  "resto_id" int8,
  "status" int4 DEFAULT 1,
  PRIMARY KEY ("id")
)
;

create SEQUENCE seq_resto;
CREATE TABLE "public"."resto" (
  "id" int8 NOT NULL DEFAULT nextval('seq_resto'::regclass),
  "name" varchar(255) COLLATE "pg_catalog"."default",
  "resto_code" varchar(255) COLLATE "pg_catalog"."default",
  "desc" varchar(500) COLLATE "pg_catalog"."default",
  "address" varchar(255) COLLATE "pg_catalog"."default",
  "city" varchar(255) COLLATE "pg_catalog"."default",
  "province" varchar(255) COLLATE "pg_catalog"."default",
  "img_url" varchar(1024) COLLATE "pg_catalog"."default",
  "status" int4 DEFAULT 1,
  PRIMARY KEY ("id")
)

create SEQUENCE seq_e_menu_item;
CREATE TABLE "public"."e_menu_item" (
  "id" int8 NOT NULL DEFAULT nextval('seq_e_menu_item'::regclass),
  "group_id" int8,
  "name" varchar(255),
  "desc" varchar(1024),
  "img_url" varchar(255),
  "price" decimal(10,2),
  "status" int4 DEFAULT 1,
  PRIMARY KEY ("id")
)
;

create SEQUENCE seq_e_menu_group;
CREATE TABLE "public"."e_menu_group" (
  "id" int8 NOT NULL DEFAULT nextval('seq_e_menu_group'::regclass),
  "resto_id" int8,
  "name" varchar(255),
  "img_url" varchar(1024),
  "status" int4 DEFAULT 1,
  PRIMARY KEY ("id")
)
;


INSERT INTO "public"."resto"("id", "name", "resto_code", "desc", "address", "city", "province", "status") VALUES (1, 'Resto 1', NULL, 'Indonesian Food', 'Jl. Kemanggisan', 'Jakarta', 'Jawa barat', 1);
INSERT INTO "public"."resto"("id", "name", "resto_code", "desc", "address", "city", "province", "status") VALUES (2, 'Resto 2', NULL, 'Chinese Food', 'Jl. Sudirman', 'Jakarta', 'Jawa barat', 1);
INSERT INTO "public"."resto"("id", "name", "resto_code", "desc", "address", "city", "province", "status") VALUES (3, 'Resto 3', NULL, 'Western Food', 'Jl. Tanjung Duren', 'Jakarta', 'Jawa barat', 1);




INSERT INTO "public"."e_menu_group"("id", "resto_id", "name", "img_url", "status") VALUES (1, 1, 'Nasi', 'https://i.picsum.photos/id/1025/200/300.jpg', 1);
INSERT INTO "public"."e_menu_group"("id", "resto_id", "name", "img_url", "status") VALUES (2, 1, 'Mie', 'https://i.picsum.photos/id/0/200/300.jpg', 1);
INSERT INTO "public"."e_menu_group"("id", "resto_id", "name", "img_url", "status") VALUES (3, 2, 'Ikan', 'https://i.picsum.photos/id/204/200/300.jpg', 1);
INSERT INTO "public"."e_menu_group"("id", "resto_id", "name", "img_url", "status") VALUES (4, 2, 'Paket', 'https://i.picsum.photos/id/250/400/400.jpg', 1);






INSERT INTO "public"."e_menu_item"("id", "group_id", "name", "desc", "img_url", "price", "status") VALUES (1, 1, 'Nasi Goreng Ayam', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.
            Pellentesque at ipsum ut metus vestibulum pharetra eget at mauris.
            Nulla pellentesque molestie sem, vel vulputate ligula commodo quis.
            Aenean vestibulum id nisl ut iaculis. Cras feugiat felis in ex aliquam, vitae ullamcorper lorem pellentesque.
            Mauris hendrerit urna eget
            rhoncus sagittis. Donec eleifend, quam id commodo dignissim, risus orci pulvinar nibh, et commodo nisl metus at velit.
            Quisque ultrices augue quis condimentum fringilla.', 'https://i.picsum.photos/id/1025/400/400.jpg', 15000.00, 1);
INSERT INTO "public"."e_menu_item"("id", "group_id", "name", "desc", "img_url", "price", "status") VALUES (2, 1, 'Nasi Goreng Sapi', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.
            Pellentesque at ipsum ut metus vestibulum pharetra eget at mauris..', 'https://i.picsum.photos/id/1003/400/400.jpg', 25000.00, 1);
INSERT INTO "public"."e_menu_item"("id", "group_id", "name", "desc", "img_url", "price", "status") VALUES (3, 1, 'Nasi Goreng Ayam 2', 'Nasi Goreng', 'https://i.picsum.photos/id/104/400/400.jpg', 20000.00, 1);
INSERT INTO "public"."e_menu_item"("id", "group_id", "name", "desc", "img_url", "price", "status") VALUES (4, 1, 'Nasi Goreng seafood 2', 'Nasi Goreng udang, cumi', 'https://i.picsum.photos/id/1050/400/400.jpg', 22200.00, 1);
INSERT INTO "public"."e_menu_item"("id", "group_id", "name", "desc", "img_url", "price", "status") VALUES (5, 2, 'Mie Pangsit', 'Mie Pangsit udang', 'https://i.picsum.photos/id/1056/400/400.jpg', 27500.00, 1);
INSERT INTO "public"."e_menu_item"("id", "group_id", "name", "desc", "img_url", "price", "status") VALUES (6, 2, 'Indomie goreng ', 'Indomie goreng telor', 'https://i.picsum.photos/id/1074/400/400.jpg', 2000.00, 1);
INSERT INTO "public"."e_menu_item"("id", "group_id", "name", "desc", "img_url", "price", "status") VALUES (7, 3, 'Ikan Gurame asam manis ', 'Ikan Gurame asam manis - Ikan Gurame asam manis', 'https://i.picsum.photos/id/113/400/400.jpg', 199000.00, 1);
INSERT INTO "public"."e_menu_item"("id", "group_id", "name", "desc", "img_url", "price", "status") VALUES (8, 4, 'Pahe 1 ', 'Nasi ikan bumbu cabe ijo', 'https://i.picsum.photos/id/132/400/400.jpg', 35000.00, 1);
INSERT INTO "public"."e_menu_item"("id", "group_id", "name", "desc", "img_url", "price", "status") VALUES (9, 4, 'Pahe 2 ', 'yam bumbu krispi geprek pake cabe merah', 'https://i.picsum.photos/id/124/400/400.jpg', 999999.00, 1);



INSERT INTO "public"."e_menu_item"("id", "group_id", "name", "desc", "img_url", "price", "status") VALUES (10, 1, 'Nasi Goreng spesial ', 'Nasi Goreng spesial Telorrrr', 'https://i.picsum.photos/id/139/400/400.jpg', 12345.00, 1);
INSERT INTO "public"."e_menu_item"("id", "group_id", "name", "desc", "img_url", "price", "status") VALUES (11, 1, 'Nasi Goreng putih ', 'Nasi Putih goreng Telorrrr', 'https://i.picsum.photos/id/146/400/400.jpg', 987987.00, 1);
INSERT INTO "public"."e_menu_item"("id", "group_id", "name", "desc", "img_url", "price", "status") VALUES (12, 1, 'Nasi Goreng rendang ', 'Nasi Putih goreng rendang', 'https://i.picsum.photos/id/152/400/400.jpg', 123.00, 1);
INSERT INTO "public"."e_menu_item"("id", "group_id", "name", "desc", "img_url", "price", "status") VALUES (13, 1, 'Nasi Goreng Nanas ', 'Nasi Putih Nanas Telorrrr', 'https://i.picsum.photos/id/159/400/400.jpg', 5555.00, 1);