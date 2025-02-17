package handlers

import (
	"TheRealOne/internal/taskService"
	"TheRealOne/internal/web/tasks"
	"context"
)

type Handler struct {
	Service *taskService.TaskService
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

//====================================================================

func (h *Handler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}
	response := tasks.GetTasks200JSONResponse{}
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	return response, nil
}

func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := request.Id
	err := h.Service.DeleteTaskByID(int64(id))
	if err != nil {
		return nil, err
	}
	return tasks.DeleteTasksId204Response{}, nil
}

func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	requestBody := request.Body
	requestId := request.Id
	updatedTask := taskService.Task{
		Task:   *requestBody.Task,
		IsDone: *requestBody.IsDone,
	}
	result, err := h.Service.UpdateTaskByID(int64(requestId), updatedTask)
	if err != nil {
		return nil, err
	}
	response := tasks.PatchTasksId200JSONResponse{
		Id:     &result.ID,
		Task:   &result.Task,
		IsDone: &result.IsDone,
	}
	return response, nil
}

//func (h *Handler) PatchTaskHandler(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	idstr := vars["id"]
//	id, _ := strconv.Atoi(idstr)
//	var task taskService.Task
//	err := json.NewDecoder(r.Body).Decode(&task)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//	updatedTask, err := h.Service.UpdateTaskByID(int64(id), task)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//	w.Header().Set("Content-Type", "application/json")
//	err = json.NewEncoder(w).Encode(updatedTask)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
//func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	idstr := vars["id"]
//	id, _ := strconv.Atoi(idstr)
//	err := h.Service.DeleteTaskByID(int64(id))
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//	w.WriteHeader(http.StatusNoContent)
//}
