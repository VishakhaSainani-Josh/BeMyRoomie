CREATE TABLE "users"(
    "user_id" SERIAL,
    "name" VARCHAR(255) NOT NULL,
    "phone" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "gender" VARCHAR(255) CHECK
        ("gender" IN('Male', 'Female')) NOT NULL,
        "city" VARCHAR(255) DEFAULT '',
        "role" VARCHAR(255)
    CHECK
        ("role" IN('finder', 'lister')) NOT NULL,
        "required_vacancy" BIGINT DEFAULT 0,
        "tags" JSON DEFAULT '[]',
        "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
        "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "users" ADD PRIMARY KEY("user_id");
ALTER TABLE
    "users" ADD CONSTRAINT "users_email_unique" UNIQUE("email");
CREATE TABLE "properties"(
    "property_id" SERIAL,
    "lister_id" BIGINT NOT NULL,
    "title" VARCHAR(255) NOT NULL,
    "description" TEXT NOT NULL,
    "city" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255) NOT NULL,
    "rent" BIGINT NOT NULL,
    "facilities" JSON NOT NULL,
    "images" JSON NOT NULL,
    "preferred_gender" VARCHAR(255) CHECK
        (
            "preferred_gender" IN('Male', 'Female')
        ) NOT NULL,
        "status" VARCHAR(255)
    CHECK
        ("status" IN('vacant', 'fulfilled')) NOT NULL,
        "vacancy" BIGINT NOT NULL,
        "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
        "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "properties" ADD PRIMARY KEY("property_id");
CREATE TABLE "interested"(
    "user_id" BIGINT NOT NULL,
    "property_id" BIGINT NOT NULL,
    "is_accepted" BOOLEAN NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "interested" ADD PRIMARY KEY("user_id");
ALTER TABLE
    "properties" ADD CONSTRAINT "propertys_lister_id_foreign" FOREIGN KEY("lister_id") REFERENCES "users"("user_id");
ALTER TABLE
    "interested" ADD CONSTRAINT "interested_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("user_id");
ALTER TABLE
    "interested" ADD CONSTRAINT "interested_property_id_foreign" FOREIGN KEY("property_id") REFERENCES "properties"("property_id");