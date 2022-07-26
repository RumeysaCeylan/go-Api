-- Table: public.User

-- DROP TABLE IF EXISTS public."User";

CREATE TABLE IF NOT EXISTS public."User"
(
    "Id" bigint NOT NULL,
    "firstName" "char",
    "lastName" "char",
    "email" "char" NOT NULL,
    "Password" "char" NOT NULL,
    CONSTRAINT "User_pkey" PRIMARY KEY ("Id")
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."User"
    OWNER to postgres;