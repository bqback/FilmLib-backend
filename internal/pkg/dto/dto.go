package dto

type key int

type ActorMovie struct {
	Actor uint64 `db:"id_actor"`
	Movie uint64 `db:"id_movie"`
}
