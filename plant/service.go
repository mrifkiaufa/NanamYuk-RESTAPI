package plant

type Service interface {
	FindAll() ([]Plant, error)
	FindByID(ID int) (Plant, error)
	Create(plantRequest PlantRequestCreate) (Plant, error)
	Update(ID int, plantRequest PlantRequestUpdate) (Plant, error)
	Delete(ID int) (Plant, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Plant, error) {
	plants, err := s.repository.FindAll()

	return plants, err
}

func (s *service) FindByID(ID int) (Plant, error) {
	plant, err := s.repository.FindByID(ID)

	return plant, err
}

func (s *service) Create(plantRequest PlantRequestCreate) (Plant, error) {
	plant := Plant{
		Name:             plantRequest.Name,
		Image:            plantRequest.Image,
		Description:      plantRequest.Description,
		Temperature:      plantRequest.Temperature,
		WateringDuration: plantRequest.WateringDuration,
		Soil:             plantRequest.Soil,
		Light:            plantRequest.Light,
		Humidity:         plantRequest.Humidity,
		Rainfall:         plantRequest.Rainfall,
		Tutorial:         plantRequest.Tutorial,
	}

	newPlant, err := s.repository.Create(plant)
	return newPlant, err
}

func (s *service) Update(ID int, plantRequest PlantRequestUpdate) (Plant, error) {
	plant, _ := s.repository.FindByID(ID)

	plant.Name = plantRequest.Name
	plant.Image = plantRequest.Image
	plant.Description = plantRequest.Description
	plant.Temperature = plantRequest.Temperature
	plant.WateringDuration = plantRequest.WateringDuration
	plant.Soil = plantRequest.Soil
	plant.Light = plantRequest.Light
	plant.Humidity = plantRequest.Humidity
	plant.Rainfall = plantRequest.Rainfall
	plant.Tutorial = plantRequest.Tutorial

	newPlant, err := s.repository.Update(plant)
	return newPlant, err
}

func (s *service) Delete(ID int) (Plant, error) {
	plant, _ := s.repository.FindByID(ID)
	newPlant, err := s.repository.Delete(plant)
	return newPlant, err
}