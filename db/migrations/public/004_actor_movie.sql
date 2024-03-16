CREATE TABLE IF NOT EXISTS public.actor_movie
(
    id_actor serial NOT NULL,
    id_movie serial NOT NULL,
    CONSTRAINT actor_movie_pkey PRIMARY KEY (id_actor, id_movie),
    CONSTRAINT actor_movie_id_actor_fkey FOREIGN KEY (id_actor)
        REFERENCES public."actor" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT actor_movie_id_movie_fkey FOREIGN KEY (id_movie)
        REFERENCES public."movie" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.actor_movie;
