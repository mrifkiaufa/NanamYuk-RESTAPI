package userplants

type Service interface {
	FindAll() ([]UserPlants, error)
	FindByID(ID int) (UserPlants, error)
	Create(userPlantsRequest UserPlantsRequestCreate) (UserPlants, error)
	Update(ID int, userPlantsRequest UserPlantsRequestUpdate) (UserPlants, error)
	Delete(ID int) (UserPlants, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]UserPlants, error) {
	userPlants, err := s.repository.FindAll()

	return userPlants, err
}

func (s *service) FindByID(ID int) (UserPlants, error) {
	userPlant, err := s.repository.FindByID(ID)

	return userPlant, err
}

func (s *service) Create(userPlantsRequest UserPlantsRequestCreate) (UserPlants, error) {
	userPlants := UserPlants{
		TagName:       userPlantsRequest.TagName,
		Date:          userPlantsRequest.Date,
		WateringState: userPlantsRequest.WateringState,
		DryState:      userPlantsRequest.DryState,
		HumidState:    userPlantsRequest.HumidState,
		PlantID:       userPlantsRequest.PlantID,
		UserID:        userPlantsRequest.UserID,
	}

	newUserPlant, err := s.repository.Create(userPlants)
	return newUserPlant, err
}

func (s *service) Update(ID int, userPlantsRequest UserPlantsRequestUpdate) (UserPlants, error) {
	userPlants, _ := s.repository.FindByID(ID)

	userPlants.TagName = userPlantsRequest.TagName
	userPlants.Date = userPlantsRequest.Date
	userPlants.WateringState = userPlantsRequest.WateringState
	userPlants.DryState = userPlantsRequest.DryState
	userPlants.HumidState = userPlantsRequest.HumidState
	userPlants.PlantID = userPlantsRequest.PlantID
	userPlants.UserID = userPlantsRequest.UserID

	newUserPlant, err := s.repository.Update(userPlants)
	return newUserPlant, err
}

func (s *service) Delete(ID int) (UserPlants, error) {
	userPlants, _ := s.repository.FindByID(ID)
	newUserPlant, err := s.repository.Delete(userPlants)
	return newUserPlant, err
}