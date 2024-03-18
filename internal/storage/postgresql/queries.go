package postgresql

import (
	"filmlib/internal/pkg/dto"
	"slices"
	"strings"
)

var shortIDField = "id"

var (
	actorTable          = "public.actor"
	actorIDField        = "public.actor.id"
	actorNameField      = "public.actor.name"
	tlActorNameField    = "LOWER(" + actorNameField + ")"
	actorGenderField    = "public.actor.gender"
	actorBirthDateField = "public.actor.dob"
)

var (
	actorShortNameField      = "name"
	actorShortGenderField    = "gender"
	actorShortBirthDateField = "dob"
	actorShortFields         = []string{actorShortNameField, actorShortGenderField, actorShortBirthDateField}
)

func ValidateActorUpdate(values map[string]interface{}) bool {
	if len(values) == 0 {
		return false
	}
	for key := range values {
		if !slices.Contains(actorShortFields, key) {
			return false
		}
	}
	return true
}

var (
	movieShortTitleField       = "title"
	movieShortDescriptionField = "description"
	movieShortReleaseDateField = "release"
	movieShortRatingField      = "rating"
	movieShortFields           = []string{movieShortTitleField, movieShortDescriptionField, movieShortReleaseDateField, movieShortRatingField}
)

func ValidateMovieUpdate(values map[string]interface{}) bool {
	if len(values) == 0 {
		return false
	}
	for key := range values {
		if !slices.Contains(movieShortFields, key) {
			return false
		}
	}
	return true
}

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
	allActorInsertFields = []string{actorShortNameField, actorShortGenderField, actorShortBirthDateField}
	allActorSelectFields = []string{shortIDField, actorShortNameField, actorShortGenderField, actorShortBirthDateField}
	actorInfoFields      = []string{shortIDField, actorShortNameField}
	actorGetAllFields    = []string{actorIDField, actorNameField, actorGenderField, actorBirthDateField,
		"jsonb_agg(jsonb_build_object( " +
			"'id'," + movieIDField + ",'title'," + movieTitleField + ")) movies",
	}
)

var actorUpdateReturnSuffix = "RETURNING " + strings.Join(allActorSelectFields, ", ")

var (
	allMovieInsertFields = []string{"title", "description", "release_date", "rating"}
	movieInfoFields      = []string{shortIDField, "title"}
	allMovieSelectFields = []string{shortIDField, "title", "description", "release_date", "rating"}
	movieGetAllFields    = []string{movieIDField, movieTitleField, movieDescriptionField, movieReleaseField, movieRatingField,
		"jsonb_agg(jsonb_build_object( " +
			"'id'," + actorIDField + ",'name'," + actorNameField + ")) actors",
	}
)

var movieUpdateReturnSuffix = "RETURNING " + strings.Join(allMovieSelectFields, ", ")

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

// var actorUpdateStructToField = map[string]string{
// 	"Name":      actorShortNameField,
// 	"Gender":    actorShortGenderField,
// 	"BirthDate": actorBirthDateField,
// }
