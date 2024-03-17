package postgresql

var (
	actorTable          = "public.actor"
	actorIDField        = "public.actor.id"
	actorNameField      = "public.actor.name"
	actorGenderField    = "public.actor.gender"
	actorBirthDateField = "public.actor.dob"
)

var (
	movieTable      = "public.movie"
	movieIDField    = "public.movie.id"
	movieTitleField = "public.movie.title"
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
	actorGetAllFields    = []string{actorIDField, actorNameField, actorGenderField, actorBirthDateField,
		"jsonb_agg(jsonb_build_object( " +
			"'id'," + movieIDField + ",'title'," + movieTitleField + ")) movies",
	}
)

var (
	allMovieInsertFields = []string{"title", "description", "release_date", "rating"}
	movieInfoFields      = []string{"id", "title"}
	allMovieSelectFields = []string{"id", "title", "description", "release_date", "rating"}
)

var (
	actorJoinActorMovieOnActorID = actorMovieTable + " ON " + actorIDField + " = " + actorMovieActorIDField
	movieJoinActorMovieOnMovieID = movieTable + " ON " + movieIDField + " = " + actorMovieMovieIDField
	actorMovieJoinMovieOnMovieID = actorMovieTable + " ON " + actorMovieMovieIDField + "=" + movieIDField
)

var (
	actorMovieFields = []string{"id_actor", "id_movie"}
)
