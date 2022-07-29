-- Table: public.User

-- DROP TABLE IF EXISTS public."User";

CREATE TABLE Userr (
    id SERIAL PRIMARY KEY,
    firstName VARCHAR(100) NULL,
    lastName VARCHAR(100) NULL,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL
);