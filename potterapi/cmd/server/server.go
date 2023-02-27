package main

import (
	"context"

	"github.com/MickStanciu/potter-api/potterapi/internal/gen/potter"
)

type Server struct {
	potter.UnimplementedHogwartsServiceServer
}

func (s *Server) GetTheGoodGuys(context.Context, *potter.GoodGuysRequest) (*potter.GoodGuysResponse, error) {
	fakeResponse := potter.GoodGuysResponse{
		People: []*potter.Character{
			{
				Name:  "Harry Potter",
				House: potter.Character_HOUSE_GRYFFINDOR,
			},
			{
				Name:  "Nevile Longbottom",
				House: potter.Character_HOUSE_GRYFFINDOR,
			},
			{
				Name:  "Cho Chang",
				House: potter.Character_HOUSE_RAVENCLAW,
			},
			{
				Name:  "Luna Lovegood",
				House: potter.Character_HOUSE_RAVENCLAW,
			},
		},
	}
	return &fakeResponse, nil
}
