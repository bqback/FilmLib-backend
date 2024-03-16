package postgresql

var (
	actorTable   = "public.actor"
	actorIDField = "public.actor.id"
)

var (
	movieTable   = "public.movie"
	movieIDField = "public.movie.id"
)

var (
	actorMovieTable        = "public.actor_movie"
	actorMovieActorIDField = "public.actor_movie.id_actor"
	actorMovieMovieIDField = "public.actor_movie.id_movie"
)

var (
	allActorInsertFields = []string{"name", "gender", "dob"}
	allActorSelectFields = []string{"id", "name", "gender", "dob"}
)

var (
	movieInfoFields = []string{"id", "title"}
)
