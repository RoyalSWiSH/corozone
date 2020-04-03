CREATE TYPE "order_status_enum" AS ENUM (
  'open',
  'accepted',
  'delayed',
  'payed',
  'received'
);

CREATE TABLE "user_profile" (
  "user_id" uuid PRIMARY KEY,
  "first_name" varchar,
  "last_name" varchar,
  "paypal_url" varchar,
  "profile_picture" bytea,
  "id_picture" bytea,
  "phone_nr" varchar,
  "phone_nr_verified" bool,
  "email" varchar,
  "email_verified" bool,
  "created_at" timestamp
);

CREATE TABLE "rate_helper" (
  "user_id" uuid,
  "rating" int,
  "order_id" int,
  "created_at" timestamp
);

CREATE TABLE "user_storage" (
  "user_id" uuid,
  "items" json
);

CREATE TABLE "user_self_quarantine" (
  "user_id" int,
  "started" timestamp
);

CREATE TABLE "hospitals" (
  "hospital_id" uuid PRIMARY KEY,
  "adress" json
);

CREATE TABLE "hospital_utilization" (
  "hospital_id" uuid,
  "utilization" float8,
  "reporting_time" timestamp
);

CREATE TABLE "contacts" (
  "user_id" uuid,
  "contact_names" json,
  "contact_id" uuid
);

CREATE TABLE "health_status" (
  "user_id" uuid,
  "status" varchar,
  "time_of_status" timestamp
);

CREATE TABLE "location_history" (
  "user_id" uuid,
  "location" json
);

CREATE TABLE "corona_locations" (
  "incident_id" uuid PRIMARY KEY,
  "location" json
);

CREATE TABLE "delivery_order" (
  "order_id" uuid PRIMARY KEY,
  "request_user_id" uuid,
  "delivery_location" json,
  "requested_items" json,
  "location" varchar,
  "longlat" geom,
  "budget" float8,
  "days_left" int,
  "elderly" boolean,
  "for_someone_else" boolean,
  "in_quarantine" boolean,
  "delivery_categories" json,
  "created_at" timestamp,
  "closed_at" timestamp
);

CREATE TABLE "delivery_status_table" (
  "order_id" uuid PRIMARY KEY,
  "helper_user_id" uuid,
  "status" order_status_enum,
  "created_at" timestamp,
  "closed_at" timestamp
);

CREATE TABLE "delivery_receipt" (
  "order_id" uuid,
  "bought_items" json,
  "price" float8,
  "receipt" bytea,
  "created_at" timestamp
);

CREATE TABLE "private_data_requests" (
  "request_id" uuid PRIMARY KEY,
  "author" json,
  "purpose" varchar,
  "requested_data" json,
  "request_time" timestamp,
  "status" varchar,
  "closed_at" timestamp,
  "user_ids" uuid
);

ALTER TABLE "delivery_status_table" ADD FOREIGN KEY ("order_id") REFERENCES "delivery_order" ("order_id");

ALTER TABLE "delivery_order" ADD FOREIGN KEY ("request_user_id") REFERENCES "user_profile" ("user_id");

ALTER TABLE "user_storage" ADD FOREIGN KEY ("user_id") REFERENCES "user_profile" ("user_id");

ALTER TABLE "user_profile" ADD FOREIGN KEY ("user_id") REFERENCES "user_self_quarantine" ("user_id");

ALTER TABLE "delivery_order" ADD FOREIGN KEY ("order_id") REFERENCES "delivery_receipt" ("order_id");

ALTER TABLE "hospital_utilization" ADD FOREIGN KEY ("hospital_id") REFERENCES "hospitals" ("hospital_id");

ALTER TABLE "health_status" ADD FOREIGN KEY ("user_id") REFERENCES "user_profile" ("user_id");

ALTER TABLE "private_data_requests" ADD FOREIGN KEY ("user_ids") REFERENCES "user_profile" ("user_id");

ALTER TABLE "rate_helper" ADD FOREIGN KEY ("user_id") REFERENCES "user_profile" ("user_id");

ALTER TABLE "delivery_order" ADD FOREIGN KEY ("order_id") REFERENCES "rate_helper" ("order_id");
