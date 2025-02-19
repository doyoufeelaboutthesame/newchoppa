package handlers

import (
	"TheRealOne/internal/userService"
	"TheRealOne/internal/web/users"
	"context"
)

type UserHandler struct {
	Service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}
	response := users.GetUsers200JSONResponse{}
	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.ID,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}
	return response, nil
}

func (h *UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	userToCreate := userService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	createdUser, err := h.Service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}
	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}
	return response, nil
}

func (h *UserHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := request.Id
	err := h.Service.DeleteUserByID(int64(id))
	if err != nil {
		return nil, err
	}
	return users.DeleteUsersId204Response{}, nil
}

func (h *UserHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	requestBody := request.Body
	requestId := request.Id
	updatedUser := userService.User{
		Email:    *requestBody.Email,
		Password: *requestBody.Password,
	}
	result, err := h.Service.UpdateUserByID(int64(requestId), updatedUser)
	if err != nil {
		return nil, err
	}
	response := users.PatchUsersId200JSONResponse{
		Id:       &result.ID,
		Email:    &result.Email,
		Password: &result.Password,
	}
	return response, nil
}

func (h *UserHandler) GetUsersUserIdTasks(ctx context.Context, request users.GetUsersUserIdTasksRequestObject) (users.GetUsersUserIdTasksResponseObject, error) {
	allTasks, err := h.Service.GetTasksForUser(request.UserId)
	if err != nil {
		return nil, err
	}
	response := users.GetUsersUserIdTasks200JSONResponse{}
	for _, tsk := range allTasks {
		task := users.Task{
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}
	return response, nil
}
