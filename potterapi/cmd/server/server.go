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
				House: potter.Character_HOUSE_GRIFFINDOR,
			},
			{
				Name:  "Lilli Potter",
				House: potter.Character_HOUSE_GRIFFINDOR,
			},
			{
				Name:  "Nevil Longbottom",
				House: potter.Character_HOUSE_GRIFFINDOR,
			},
			{
				Name:  "George Weasly",
				House: potter.Character_HOUSE_GRIFFINDOR,
			},
		},
	}
	return &fakeResponse, nil
}
