package postgresql

var (
	actorTable   = "public.actor"
	actorIDField = "public.actor.id"
)

var (
	allActorInsertFields = []string{"name", "gender", "dob"}
	allActorSelectFields = []string{"id", "name", "gender", "dob"}
)
