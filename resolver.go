package graphql_go_pets_example

import (
	"context"
	"fmt"
	"strconv"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Pet() PetResolver {
	return &petResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) AddPet(ctx context.Context, pet PetInput) (*Pet, error) {
	var id string
	if pet.ID == nil {
		id = "new"
	} else {
		id = *pet.ID
	}
	p := Pet{
		ID:    id,
		Owner: pet.OwnerID,
		Name:  pet.Name,
	}
	if pet.TagIDs != nil {
		t := PetTags{}
		for _, pt := range pet.TagIDs {
			t = append(t, pt)
		}
		p.Tags = t
	}
	petStore[id] = p
	return &p, nil
}
func (r *mutationResolver) UpdatePet(ctx context.Context, pet PetInput) (*Pet, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeletePet(ctx context.Context, userID string, petID string) (*bool, error) {
	panic("not implemented")
}

type petResolver struct{ *Resolver }

func (r *petResolver) Owner(ctx context.Context, obj *Pet) (*User, error) {
	o, ok := userStore[obj.Owner]
	if !ok {
		return nil, fmt.Errorf("pet %s does not have an owner", obj.Name)
	}
	return &o, nil
}
func (r *petResolver) Tags(ctx context.Context, obj *Pet) ([]*Tag, error) {
	t := make([]*Tag, 0)
	for _, pt := range obj.Tags {
		tt := tagStore[strconv.Itoa(pt)]
		t = append(t, &tt)
	}
	if len(t) == 0 {
		return nil, nil
	}
	return t, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetPets(ctx context.Context, ids []*string) ([]*Pet, error) {
	p := make([]*Pet, len(petStore))
	iter := 0
	for _, value := range petStore {
		p[iter] = &value
		iter++
	}
	return p, nil
}
func (r *queryResolver) GetUser(ctx context.Context, id string) (*User, error) {
	u := userStore[id]
	return &u, nil
}
func (r *queryResolver) GetPet(ctx context.Context, id string) (*Pet, error) {
	p := petStore[id]
	return &p, nil
}
func (r *queryResolver) GetTag(ctx context.Context, title string) (*Tag, error) {
	for _, tag := range tagStore {
		if tag.Title == title {
			return &tag, nil
		}
	}
	return nil, fmt.Errorf("no tag named %s found", title)
}
