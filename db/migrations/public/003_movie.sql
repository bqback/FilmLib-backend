CREATE TABLE IF NOT EXISTS public."movie"
(
    id serial NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    release date NOT NULL, 
    rating real NOT NULL,
    CONSTRAINT movie_pkey PRIMARY KEY (id),
    CONSTRAINT title_length_check CHECK (length(title) >= 1 AND length(title) <= 150),
    CONSTRAINT description_length_check CHECK (length(description) <= 150),
    CONSTRAINT rating_check CHECK (rating >= 0 AND rating <= 10)
);

---- create above / drop below ----

DROP TABLE IF EXISTS public."movie";
