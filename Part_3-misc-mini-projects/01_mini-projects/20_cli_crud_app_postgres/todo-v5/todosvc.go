package todo

import "fmt"

// Service is the ToDo task service layer
type Service struct {
	repo *Repository
}

// NewService creates a new service to handle ToDo task related operations
func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// Create handles the creation
func (s *Service) Create(item ToDo) error {
	if item.Title == "" || item.Description == "" {
		return ErrEmptyTitleAndDescription
	}
	if err := s.repo.Save(&item); err != nil {
		return fmt.Errorf("could not create ToDo task:\n%v", err)
	}
	return nil
}

// RetrieveAll handles the retrieval of the existing ToDo tasks.
func (s *Service) RetrieveAll() ([]ToDo, error) {
	items, err := s.repo.All()
	if err != nil {
		return nil, fmt.Errorf("could not retrieve ToDo tasks:\n%v", err)
	}
	return items, nil



}

// Retrieve handles the retrieval of the existing ToDo tasks whose title matches the one given.
func (s *Service) Retrieve(title string) ([]ToDo, error) {
	items, err := s.repo.FindByTitle(title)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve ToDo tasks by title:\n%v", err)
	}
	return items, nil
}

// Update handles the update of an existing ToDo task given the id and patched data
func (s *Service) Update(id int, item *ToDo) error {
	if id == -1 {
		return ErrEmptyID
	}
	savedItem, err := s.repo.FindByID(int(id))
	if err != nil {
		return err
	}

	if (item.Title == "" && item.Description == "" && item.Done == savedItem.Done) ||
	 	(item.Title == savedItem.Title && item.Description == savedItem.Description && item.Done == savedItem.Done ) {
		return ErrAtLeastOneFieldShouldBeUpdated
	}

	if item.Title != "" {
		savedItem.Title = item.Title
	}
	if item.Description != "" {
		savedItem.Description = item.Description
	}
	savedItem.Done = item.Done
	if err := s.repo.Save(savedItem); err != nil {
		return fmt.Errorf("could not update ToDo task:\n%v", err)
	}
	return nil
}

// Delete handles the deletion of a given ToDo task given its ID
func (s *Service) Delete(id int) error {
	if id == -1 {
		return ErrEmptyID
	}

	savedItem, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if err := s.repo.Delete(savedItem); err != nil {
		return fmt.Errorf("could not delete ToDo task:\n%v", err)
	}
	return nil
}
