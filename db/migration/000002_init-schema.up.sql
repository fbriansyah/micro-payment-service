CREATE TABLE "outlets" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "user" uuid UNIQUE NOT NULL,
  "deposit" bigint NOT NULL,
  "is_active" bool NOT NULL DEFAULT true
);

CREATE TABLE "products" (
  "product_code" varchar UNIQUE PRIMARY KEY,
  "product_name" varchar NOT NULL,
  "product_endpoint" varchar NOT NULL
);

CREATE TABLE "transactions" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "bill_number" varchar NOT NULL,
  "product" varchar NOT NULL,
  "transaction_datetime" timestamptz NOT NULL DEFAULT 'now()',
  "inquiry" uuid,
  "payment" uuid,
  "amount" bigint NOT NULL,
  "refference_number" varchar NOT NULL,
  "outlet" uuid NOT NULL,
  "status" int NOT NULL DEFAULT 0
);

CREATE TABLE "request_logs" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "mode" varchar NOT NULL,
  "product" varchar NOT NULL,
  "bill_number" bigint NOT NULL,
  "name" varchar NOT NULL,
  "total_amount" bigint NOT NULL,
  "biller_response" text NOT NULL DEFAULT '',
  "outlet" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

COMMENT ON COLUMN "outlets"."user" IS 'one user one outlet';

COMMENT ON COLUMN "transactions"."outlet" IS 'user who pay the bill';

COMMENT ON COLUMN "transactions"."status" IS '-1:failed;0:pending;1:success';

COMMENT ON COLUMN "request_logs"."mode" IS 'inq or pay';

COMMENT ON COLUMN "request_logs"."outlet" IS 'user who request the bill';

ALTER TABLE "transactions" ADD FOREIGN KEY ("product") REFERENCES "products" ("product_code");

ALTER TABLE "transactions" ADD FOREIGN KEY ("outlet") REFERENCES "outlets" ("id");

ALTER TABLE "request_logs" ADD FOREIGN KEY ("product") REFERENCES "products" ("product_code");

ALTER TABLE "request_logs" ADD FOREIGN KEY ("outlet") REFERENCES "outlets" ("id");
