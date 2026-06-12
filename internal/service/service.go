package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"sqlc.dev/app/db/sqlc"
	"sqlc.dev/app/internal/models"
	"sqlc.dev/app/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, req models.CreateUserRequest) (models.UserResponse, error)
	GetUserByID(ctx context.Context, id int32) (models.UserResponse, error)
	UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (models.UserResponse, error)
	DeleteUser(ctx context.Context, id int32) error
	ListUsers(ctx context.Context, page, limit int) ([]models.UserResponse, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func CalculateAge(dob time.Time, now time.Time) int {
	years := now.Year() - dob.Year()
	if now.Month() < dob.Month() || (now.Month() == dob.Month() && now.Day() < dob.Day()) {
		years--
	}
	if years < 0 {
		return 0
	}
	return years
}

func parseDate(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}

func (s *userService) CreateUser(ctx context.Context, req models.CreateUserRequest) (models.UserResponse, error) {
	dobTime, err := parseDate(req.DOB)
	if err != nil {
		return models.UserResponse{}, errors.New("invalid date of birth format")
	}

	user, err := s.repo.Queries().CreateUser(ctx, sqlc.CreateUserParams{
		Name: req.Name,
		Dob:  dobTime,
	})
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
	}, nil
}

func (s *userService) GetUserByID(ctx context.Context, id int32) (models.UserResponse, error) {
	user, err := s.repo.Queries().GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.UserResponse{}, errors.New("user not found")
		}
		return models.UserResponse{}, err
	}

	age := CalculateAge(user.Dob, time.Now())

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
		Age:  &age,
	}, nil
}

func (s *userService) UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (models.UserResponse, error) {
	dobTime, err := parseDate(req.DOB)
	if err != nil {
		return models.UserResponse{}, errors.New("invalid date of birth format")
	}

	user, err := s.repo.Queries().UpdateUser(ctx, sqlc.UpdateUserParams{
		Name: req.Name,
		Dob:  dobTime,
		ID:   id,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.UserResponse{}, errors.New("user not found")
		}
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
	}, nil
}

func (s *userService) DeleteUser(ctx context.Context, id int32) error {
	_, err := s.repo.Queries().GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("user not found")
		}
		return err
	}
	return s.repo.Queries().DeleteUser(ctx, id)
}

func (s *userService) ListUsers(ctx context.Context, page, limit int) ([]models.UserResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	users, err := s.repo.Queries().ListUsers(ctx, sqlc.ListUsersParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}

	now := time.Now()
	res := make([]models.UserResponse, len(users))
	for i, u := range users {
		age := CalculateAge(u.Dob, now)
		res[i] = models.UserResponse{
			ID:   u.ID,
			Name: u.Name,
			DOB:  u.Dob.Format("2006-01-02"),
			Age:  &age,
		}
	}

	return res, nil
}
