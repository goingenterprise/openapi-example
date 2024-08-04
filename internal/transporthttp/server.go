package transporthttp

import (
	"context"
	api "open-api-example"
)

type Server struct{}

func (s Server) ListPets(ctx context.Context, request api.ListPetsRequestObject) (api.ListPetsResponseObject, error) {

	pets := make([]api.Pet, 0)
	p := append(pets, api.Pet{
		Id:   0,
		Name: "Bingo",
	})
	return api.ListPets200JSONResponse{
		Body: p,
	}, nil
}

func (s Server) CreatePets(ctx context.Context, request api.CreatePetsRequestObject) (api.CreatePetsResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) ShowPetById(ctx context.Context, request api.ShowPetByIdRequestObject) (api.ShowPetByIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func NewServer() *Server {
	return &Server{}
}
