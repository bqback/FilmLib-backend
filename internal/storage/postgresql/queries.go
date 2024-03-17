package postgresql

import "filmlib/internal/pkg/dto"

var (
	actorTable          = "public.actor"
	actorIDField        = "public.actor.id"
	actorNameField      = "public.actor.name"
	tlActorNameField    = "LOWER(" + actorNameField + ")"
	actorGenderField    = "public.actor.gender"
	actorBirthDateField = "public.actor.dob"
)

var (
	movieTable            = "public.movie"
	movieIDField          = "public.movie.id"
	movieTitleField       = "public.movie.title"
	tlMovieTitleField     = "LOWER(" + movieTitleField + ")"
	movieDescriptionField = "public.movie.description"
	movieReleaseField     = "public.movie.release_date"
	movieRatingField      = "public.movie.rating"
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
	movieGetAllFields    = []string{movieIDField, movieTitleField, movieDescriptionField, movieReleaseField, movieRatingField,
		"jsonb_agg(jsonb_build_object( " +
			"'id'," + actorIDField + ",'name'," + actorNameField + ")) actors",
	}
)

var (
	actorMovieOnActorID = actorMovieTable + " ON " + actorIDField + " = " + actorMovieActorIDField
	movieOnMovieID      = movieTable + " ON " + movieIDField + " = " + actorMovieMovieIDField
	actorOnActorID      = actorTable + " ON " + actorIDField + " = " + actorMovieActorIDField
	actorMovieOnMovieID = actorMovieTable + " ON " + actorMovieMovieIDField + "=" + movieIDField
)

var (
	actorMovieFields = []string{"id_actor", "id_movie"}
)

var SortOptionsMap = map[int]string{
	dto.TitleSort:   movieTitleField,
	dto.RatingSort:  movieRatingField,
	dto.ReleaseSort: movieReleaseField,
}

var SortOrderMap = map[int]string{
	dto.AscSort:  "ASC",
	dto.DescSort: "DESC",
}
