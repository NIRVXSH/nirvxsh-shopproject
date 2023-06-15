BEGIN;

--Set Tomezone
SET TIME ZONE 'Asia/Bangkok';


--Install uuid extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp"

--user_id -> U000001
--products_id -> U000001
--orders_id -> U000001

CREATE SEQUENCE users_id_seq START WITH 1 INCREMENT BY 1;
CREATE SEQUENCE products_id_seq START WITH 1 INCREMENT BY 1;
CREATE SEQUENCE orders_id_seq START WITH 1 INCREMENT BY 1;

--Auto update
CREATE OR REPLACE FUNCTION set_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;   
END;
$$ language 'plpgsql';

--Create enum
CREATE TYPE "order_status" AS ENUM (
    'waiting',
    'shipping',
    'completed',
    'canceled'
);

CREATE TABLE "users" (
  "id" VARCHAR(7) PRIMARY KEY DEFAULT CONCAT('U',LPAD(NEXTVAL('users_id_seq')::TEXT,6,'0')),
  "username" VARCHAR UNIQUE,
  "password" VARCHAR,
  "email" VARCHAR UNIQUE,
  "role_id" int,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "oauth" (
  "id" VARCHAR PRIMARY KEY,
  "user_id" VARCHAR,
  "access_token" VARCHAR,
  "refresh_token" VARCHAR,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "roles" (
  "id" int PRIMARY KEY,
  "title" VARCHAR
);

CREATE TABLE "products" (
  "id" VARCHAR PRIMARY KEY,
  "title" VARCHAR,
  "description" VARCHAR,
  "price" float,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "products_categories" (
  "id" VARCHAR PRIMARY KEY,
  "product_id" VARCHAR,
  "category_id" int
);

CREATE TABLE "categories" (
  "id" int PRIMARY KEY,
  "title" VARCHAR UNIQUE
);

CREATE TABLE "images" (
  "id" VARCHAR PRIMARY KEY,
  "filename" VARCHAR,
  "url" VARCHAR,
  "product_id" VARCHAR,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "orders" (
  "id" VARCHAR PRIMARY KEY,
  "user_id" VARCHAR,
  "contact" VARCHAR,
  "address" VARCHAR,
  "transfer_slip" jsonb,
  "status" VARCHAR,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "products_orders" (
  "id" VARCHAR PRIMARY KEY,
  "order_id" VARCHAR,
  "qty" int,
  "products" jsonb
);

ALTER TABLE "oauth" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "products_categories" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "products_categories" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "images" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "products_orders" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");
