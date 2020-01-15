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