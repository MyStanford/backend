package database

import "mystanford/person"

type Person struct {
	ID          int    `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Prompt      string `json:"prompt"`
}

func PersonLoadDefault() {
	persons := person.GetDefaultPersons()
	for _, item := range persons {
		if !PersonExist(item.Name) {
			PersonAdd(item)
		}
	}
}

func PersonExist(name string) bool {
	var i int64
	DB.Model(&Person{}).Where("name = ?", name).Count(&i)
	return i != 0
}

func PersonAdd(person person.Person) (e error) {
	e = DB.Create(&Person{
		Name:        person.Name,
		Description: person.Description,
		Prompt:      person.Prompt,
	}).Error
	return
}

func PersonGetAll() (data []Person, e error) {
	e = DB.Find(&data).Error
	return
}

func PersonGetByName(name string) (data Person, e error) {
	e = DB.Where("name = ?", name).First(&data).Error
	return
}
