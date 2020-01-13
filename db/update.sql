create SEQUENCE seq_user;

CREATE TABLE "public"."user" (
  "id" int8 NOT NULL DEFAULT nextval('seq_user'::regclass),
  "name" varchar(255),
  "password" varchar(500),
  "email" varchar(255),
  "phone_numb" varchar(255),
  "fb" varchar(255),
  PRIMARY KEY ("id")
)
;