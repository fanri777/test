package models

type User struct {
	ID    string `json: "id"    bson:"user_id"`
	Nama  string `json: "nama"  bson:"user_nama"`
	Umur  int64  `json: "umur"  bson:"user_umur"`
	Kelas string `json: "kelas" bson:"user_kelas"`
}
