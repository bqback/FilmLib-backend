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
	actorInfoFields      = []string{"id", "name"}
)

var (
	allMovieInsertFields = []string{"title", "description", "release_date", "rating"}
	movieInfoFields      = []string{"id", "title"}
	allMovieSelectFields = []string{"id", "title", "description", "release_date", "rating"}
)

var (
	actorMovieFields = []string{"id_actor", "id_movie"}
)
