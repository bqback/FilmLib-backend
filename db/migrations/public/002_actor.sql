CREATE TABLE IF NOT EXISTS public."actor"
(
    id serial NOT NULL,
    name text NOT NULL,
    dob date NOT NULL, 
    CONSTRAINT actor_pkey PRIMARY KEY (id)
);

---- create above / drop below ----

DROP TABLE IF EXISTS public."actor";
