CREATE DATABASE getgoal;

\c getgoal;

CREATE TABLE "public"."program" (
  "program_id" SERIAL PRIMARY KEY,
  "program_name" VARCHAR(150) NOT NULL,
  "rating" NUMERIC(3,1) NOT NULL,
  "media_url" VARCHAR(255),
  "program_description" VARCHAR(250),
  "expected_time" VARCHAR(30),
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE TABLE "public"."user_account" (
  "user_id" SERIAL PRIMARY KEY,
  "first_name" VARCHAR(70) NOT NULL,
  "last_name" VARCHAR(70) NOT NULL,
  "email" VARCHAR(100) NOT NULL,
  "password_hash" VARCHAR(250) NOT NULL,
  "password_salt" VARCHAR(100) NOT NULL,
  "email_validation_status_id" INT,
  "confirmation_token" VARCHAR(100),
  "token_generation_time" TIMESTAMP,
  "password_recovery_token" VARCHAR(100),
  "recovery_token_time" TIMESTAMP,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE TABLE "public"."email_validation_status" (
  "email_validation_status_id" SERIAL PRIMARY KEY,
  "status_description" VARCHAR(45) NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE TABLE "public"."task" (
  "task_id" SERIAL PRIMARY KEY,
  "task_name" VARCHAR(150) NOT NULL,
  "task_status" INT NOT NULL,
  "user_account_id" INT NOT NULL,
  "is_set_noti" INT NOT NULL,
  "start_time" TIMESTAMP NOT NULL,
  "end_time" TIMESTAMP,
  "program_id" INT,
  "category" VARCHAR(50),
  "time_before_notify" INT,
  "task_description" VARCHAR(250),
  "link" VARCHAR(255),
  "media_url" VARCHAR(255),
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE TABLE "public"."label" (
  "label_id" SERIAL PRIMARY KEY,
  "label_name" VARCHAR(50) NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);


CREATE TABLE "public"."label_program" (
  "label_id" INT NOT NULL,
  "program_id" INT NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP,
  PRIMARY KEY ("label_id", "program_id")
);


CREATE TABLE "public"."action_type" (
  "action_id" SERIAL PRIMARY KEY,
  "action_name" VARCHAR(50) NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE TABLE "public"."user_program" (
  "user_program_id" SERIAL PRIMARY KEY,
  "user_account_id" INT NOT NULL,
  "program_id" INT NOT NULL,
  "action_id" INT NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE INDEX "fk_task_program_idx" ON "public"."task" ("program_id");
CREATE INDEX "fk_task_user_account1_idx" ON "public"."task" ("user_account_id");
CREATE INDEX "fk_category_has_program_program1_idx" ON "public"."label_program" ("program_id");
CREATE INDEX "fk_category_has_program_category1_idx" ON "public"."label_program" ("label_id");
CREATE INDEX "fk_user_account_has_program_program1_idx" ON "public"."user_program" ("program_id");
CREATE INDEX "fk_user_account_has_program_user_account1_idx" ON "public"."user_program" ("user_account_id");
CREATE INDEX "fk_user_program_action_type1_idx" ON "public"."user_program" ("action_id");
CREATE INDEX "fk_user_account_email_validation_status1_idx" ON "public"."user_account" ("email_validation_status_id");

ALTER TABLE "public"."task" ADD CONSTRAINT "fk_task_program" FOREIGN KEY ("program_id") REFERENCES "public"."program" ("program_id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."task" ADD CONSTRAINT "fk_task_user_account1" FOREIGN KEY ("user_account_id") REFERENCES "public"."user_account" ("user_id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."label_program" ADD CONSTRAINT "fk_category_has_program_category1" FOREIGN KEY ("label_id") REFERENCES "public"."label" ("label_id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."label_program" ADD CONSTRAINT "fk_category_has_program_program1" FOREIGN KEY ("program_id") REFERENCES "public"."program" ("program_id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."user_program" ADD CONSTRAINT "fk_user_account_has_program_user_account1" FOREIGN KEY ("user_account_id") REFERENCES "public"."user_account" ("user_id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."user_program" ADD CONSTRAINT "fk_user_account_has_program_program1" FOREIGN KEY ("program_id") REFERENCES "public"."program" ("program_id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."user_program" ADD CONSTRAINT "fk_user_program_action_type1" FOREIGN KEY ("action_id") REFERENCES "public"."action_type" ("action_id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."user_account" ADD CONSTRAINT "fk_user_account" FOREIGN KEY ("email_validation_status_id") REFERENCES "public"."email_validation_status" ("email_validation_status_id") ON DELETE NO ACTION ON UPDATE NO ACTION;
