package db

import "log"

func Migrate() {
	schema := `
		CREATE TABLE IF NOT EXISTS "users" (
			"id" uuid PRIMARY KEY NOT NULL,
			"created_at" timestamptz NOT NULL DEFAULT (now()),
			"name" varchar,
			"email" varchar,
			"password" varchar
		);

		CREATE TABLE IF NOT EXISTS "link_groups" (
			"id" uuid PRIMARY KEY NOT NULL,
			"created_at" timestamptz NOT NULL DEFAULT (now()),
			"url" varchar,
			"name" varchar,
			"status" bool,
			"user_id" uuid
		);

		CREATE TABLE IF NOT EXISTS "links" (
			"id" uuid PRIMARY KEY NOT NULL,
			"created_at" timestamptz NOT NULL DEFAULT (now()),
			"duration" timestamp,
			"url" varchar,
			"name" varchar,
			"qrcode" varchar,
			"status" bool,
			"group_id" uuid
		);

		ALTER TABLE "link_groups" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
		ALTER TABLE "links" ADD FOREIGN KEY ("group_id") REFERENCES "link_groups" ("id");

	`

	db, err := OpenConnection()
	if err != nil {
		log.Fatal(err)
	}

	db.MustExec(schema)
}

func Drop() {
	schema := `
		DROP TABLE IF EXISTS "links";	
		DROP TABLE IF EXISTS "link_groups";
		DROP TABLE IF EXISTS "users";
	`

	db, err := OpenConnection()
	if err != nil {
		log.Fatal(err)
	}

	db.MustExec(schema)
}
