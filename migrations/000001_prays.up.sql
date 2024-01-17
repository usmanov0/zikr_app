CREATE TABLE "prays"(
    "id" SERIAL PRIMARY KEY,
    "language" VARCHAR NOT NULL,
    "text" VARCHAR NOT NULL,
    "count_pray" INTEGER
)