package handlers

import (
	"construction-organization-system/internal/model"
	"construction-organization-system/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

type ConstructionProjectHandlers struct {
	constructionProjectService *service.ConstructionProjectService
}

func NewConstructionProjectHandlers(service *service.ConstructionProjectService) *ConstructionProjectHandlers {
	return &ConstructionProjectHandlers{constructionProjectService: service}
}

func (h *ConstructionProjectHandlers) Get(w http.ResponseWriter, r *http.Request) {
	var workTypeID int
	var err error

	workType := r.URL.Query().Get("work_type")
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	if workType == "" {
		workTypeID = 0
	} else if workTypeID, err = strconv.Atoi(workType); err != nil {
		http.Error(w, "Invalid work type format. Use integer value.", http.StatusBadRequest)
		return
	}

	if _, err := time.Parse("2006-01-02", startDate); startDate != "" && err != nil {
		http.Error(w, "Invalid start date format. Use YYYY-MM-DD format.", http.StatusBadRequest)
		return
	}

	if _, err := time.Parse("2006-01-02", endDate); endDate != "" && err != nil {
		http.Error(w, "Invalid end date format. Use YYYY-MM-DD format.", http.StatusBadRequest)
		return
	}

	projects, err := h.constructionProjectService.GetProjects(workTypeID, startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(projects)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ConstructionProjectHandlers) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionProjectHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionProjectHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionProjectHandlers) GetList(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionProjectHandlers) GetWorkSchedules(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	workSchedules, err := h.constructionProjectService.GetWorkSchedules(projectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(workSchedules)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ConstructionProjectHandlers) GetConstructionTeams(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	constructionTeams, err := h.constructionProjectService.GetConstructionTeams(projectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(constructionTeams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ConstructionProjectHandlers) GetMachines(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	if _, err := time.Parse("2006-01-02", startDate); startDate != "" && err != nil {
		http.Error(w, "Invalid start date format. Use YYYY-MM-DD format.", http.StatusBadRequest)
		return
	}

	if _, err := time.Parse("2006-01-02", endDate); endDate != "" && err != nil {
		http.Error(w, "Invalid end date format. Use YYYY-MM-DD format.", http.StatusBadRequest)
		return
	}

	machines, err := h.constructionProjectService.GetMachines(projectId, startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(machines)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ConstructionProjectHandlers) GetEstimate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	estimate, err := h.constructionProjectService.GetEstimate(projectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(estimate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ConstructionProjectHandlers) GetReports(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reports, err := h.constructionProjectService.GetReports(projectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(reports)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ConstructionProjectHandlers) GetExceededDeadlinesWorks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	workTypes, err := h.constructionProjectService.GetExceededDeadlinesWorks(projectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(workTypes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ConstructionProjectHandlers) GetSchoolList(w http.ResponseWriter, r *http.Request) {
	school, err := h.constructionProjectService.GetSchoolList()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(school)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ConstructionProjectHandlers) GetSchool(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	school, err := h.constructionProjectService.GetSchool(projectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(school)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ConstructionProjectHandlers) CreateSchool(w http.ResponseWriter, r *http.Request) {
	var school model.School

	//var data map[string]interface{}
	//if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}

	err := json.NewDecoder(r.Body).Decode(&school)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.constructionProjectService.CreateSchool(school)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(school); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ConstructionProjectHandlers) UpdateSchool(w http.ResponseWriter, r *http.Request) {
	var school model.School

	if err := json.NewDecoder(r.Body).Decode(&school); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.constructionProjectService.UpdateSchool(school)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(school); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ConstructionProjectHandlers) DeleteSchool(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)

	projectId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.constructionProjectService.DeleteSchool(projectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ConstructionProjectHandlers) GetApartmentHouseList(w http.ResponseWriter, r *http.Request) {
	houses, err := h.constructionProjectService.GetApartmentHouseList()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(houses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ConstructionProjectHandlers) GetApartmentHouse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	house, err := h.constructionProjectService.GetApartmentHouse(projectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(house)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ConstructionProjectHandlers) CreateApartmentHouse(w http.ResponseWriter, r *http.Request) {
	var house model.ApartmentHouse
	err := json.NewDecoder(r.Body).Decode(&house)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.constructionProjectService.CreateApartmentHouse(house)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(house); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ConstructionProjectHandlers) UpdateApartmentHouse(w http.ResponseWriter, r *http.Request) {
	var house model.ApartmentHouse
	err := json.NewDecoder(r.Body).Decode(&house)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.constructionProjectService.UpdateApartmentHouse(house)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(house); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ConstructionProjectHandlers) DeleteApartmentHouse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.constructionProjectService.DeleteApartmentHouse(projectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ConstructionProjectHandlers) GetBridgeList(w http.ResponseWriter, r *http.Request) {
	bridges, err := h.constructionProjectService.GetBridgeList()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(bridges)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ConstructionProjectHandlers) GetBridge(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bridge, err := h.constructionProjectService.GetBridge(projectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(bridge)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ConstructionProjectHandlers) CreateBridge(w http.ResponseWriter, r *http.Request) {
	var bridge model.Bridge

	err := json.NewDecoder(r.Body).Decode(&bridge)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.constructionProjectService.CreateBridge(bridge)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(bridge); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ConstructionProjectHandlers) UpdateBridge(w http.ResponseWriter, r *http.Request) {
	var bridge model.Bridge
	err := json.NewDecoder(r.Body).Decode(&bridge)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.constructionProjectService.UpdateBridge(bridge)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(bridge); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ConstructionProjectHandlers) DeleteBridge(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.constructionProjectService.DeleteBridge(projectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
