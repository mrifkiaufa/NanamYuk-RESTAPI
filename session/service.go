package session

type Service interface {
	FindAll() ([]Session, error)
	FindByID(ID int) (Session, error)
	Create(sessionRequest SessionRequestCreate) (Session, error)
	Update(ID int, sessionRequest SessionRequestUpdate) (Session, error)
	Delete(ID int) (Session, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Session, error) {
	session, err := s.repository.FindAll()

	return session, err
}

func (s *service) FindByID(ID int) (Session, error) {
	session, err := s.repository.FindByID(ID)

	return session, err
}

func (s *service) Create(sessionRequest SessionRequestCreate) (Session, error) {
	session := Session{
		Date:         sessionRequest.Date,
		UserPlantsID: sessionRequest.UserPlantsID,
	}

	newSession, err := s.repository.Create(session)
	return newSession, err
}

func (s *service) Update(ID int, sessionRequest SessionRequestUpdate) (Session, error) {
	session, _ := s.repository.FindByID(ID)

	session.Date = sessionRequest.Date
	session.UserPlantsID = sessionRequest.UserPlantsID

	newSession, err := s.repository.Update(session)
	return newSession, err
}

func (s *service) Delete(ID int) (Session, error) {
	session, _ := s.repository.FindByID(ID)
	newSession, err := s.repository.Delete(session)
	return newSession, err
}