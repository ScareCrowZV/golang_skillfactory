package main

type pet struct {
	Name string
}

type cat struct {
	pet
	Meow_sound_level int64
}

func NewCat(name string, meow_sound_level int64) cat {
	return cat{pet: pet{Name: name}, Meow_sound_level: meow_sound_level}
}
