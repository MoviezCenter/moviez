package model

type Movie struct {
	Id          uint32 `db:"id"`
	Title       string `db:"string"`
	Overview    string `db:"overview"`
	Releasedate string `db:"release_date"`
	PosterPath  string `db:"poster_path"`
	GenreId     string `db:"genre_id"`
}

type Genre struct {
	Id     uint32 `db:"id"`
	Name   string `db:"name"`
	TypeId int32  `db:"type_id"`
}
