package storage

import (
	"fmt"

	"github.com/google/uuid"
)

type ItemWithID[T any] struct {
	Item T
	ID   string
}

type Store[T any] struct {
	Data map[string]ItemWithID[T]
}

func SetupStore[T any]() *Store[T] {
	return &Store[T]{
		Data: map[string]ItemWithID[T]{},
	}
}

func (s *Store[T]) GetByID(id string) (*ItemWithID[T], error) {
	if s == nil || s.Data == nil {
		return nil, fmt.Errorf("store is not setup correctly")
	}
	item, ok := s.Data[id]
	// If the key exists
	if !ok {
		return nil, fmt.Errorf("item not found")
	}

	return &item, nil
}

func (s *Store[T]) GetAll() ([]ItemWithID[T], error) {
	if s == nil || s.Data == nil {
		return nil, fmt.Errorf("store is not setup correctly")
	}
	data := []ItemWithID[T]{}

	for _, val := range s.Data {
		data = append(data, val)
	}

	return data, nil
}

func (s *Store[T]) Create(newItems []T) ([]ItemWithID[T], error) {
	if s == nil || s.Data == nil {
		return nil, fmt.Errorf("store is not setup correctly")
	}
	items := []ItemWithID[T]{}
	for _, newItem := range newItems {
		// create unique id
		newUUID := uuid.New().String()
		newItemWithID := ItemWithID[T]{
			Item: newItem,
			ID:   newUUID,
		}
		s.Data[newUUID] = newItemWithID
		items = append(items, newItemWithID)
	}

	return items, nil
}

func (s *Store[T]) Update(param ItemWithID[T]) (*ItemWithID[T], error) {
	if s == nil || s.Data == nil {
		return nil, fmt.Errorf("store is not setup correctly")
	}
	_, ok := s.Data[param.ID]
	// If the key exists
	if !ok {
		return nil, fmt.Errorf("item not found")
	}

	newItem := ItemWithID[T]{
		Item: param.Item,
		ID:   param.ID,
	}

	s.Data[param.ID] = newItem

	return &newItem, nil
}

func (s *Store[T]) Remove(id string) (*ItemWithID[T], error) {
	if s == nil || s.Data == nil {
		return nil, fmt.Errorf("store is not setup correctly")
	}
	value, ok := s.Data[id]
	// If the key exists
	if !ok {
		return nil, fmt.Errorf("item not found")
	}
	delete(s.Data, id)
	return &value, nil
}
