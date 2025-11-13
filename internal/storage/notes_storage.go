package storage

import "notes-app/internal/models"

// Простое хранилище заметок в памяти
var Notes = map[int]models.Note{}
var lastID = 0

// Добавить заметку
func Add(note models.Note) models.Note {
	// найти минимальный свободный ID
	id := 1
	for {
		if _, exists := Notes[id]; !exists {
			break
		}
		id++
	}
	note.ID = id
	Notes[id] = note
	return note
}

// Получить все заметки
func GetAll() []models.Note {
	result := []models.Note{}
	for _, note := range Notes {
		result = append(result, note)
	}
	return result
}

// Удалить заметку по id
func Delete(id int) {
	delete(Notes, id)
}
