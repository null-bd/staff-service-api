package rest

import "github.com/null-bd/staff-service-api/internal/staff"

func ToStaff(req *CreateStaffRequest) *staff.Staff {
	return &staff.Staff{
		BranchID:       req.BranchID,
		OrganizationID: req.OrganizationID,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Code:           req.Code,
		StaffType:      req.StaffType,
		Specialities:   req.Specialities,
		Departments: staff.Departments{
			DepartmentID: req.Departments.DepartmentID,
			Role:         req.Departments.Role,
			IsPrimary:    req.Departments.IsPrimary,
		},
		Schedule: staff.Schedule{
			Type:   req.Schedule.Type,
			Shifts: req.Schedule.Shifts,
		},
		Email:       req.Email,
		Phone:       req.Phone,
		DateOfBirth: req.DateOfBirth,
		Gender:      req.Gender,
		Address: staff.Address{
			Street:  req.Address.Street,
			City:    req.Address.City,
			State:   req.Address.State,
			Country: req.Address.Country,
			ZipCode: req.Address.ZipCode,
		},
		Metadata: req.Metadata,
	}
}

func ToStaffResponse(staff *staff.Staff) *StaffResponse {
	return &StaffResponse{
		ID:             staff.ID,
		BranchID:       staff.BranchID,
		OrganizationID: staff.OrganizationID,
		FirstName:      staff.FirstName,
		LastName:       staff.LastName,
		Code:           staff.Code,
		Status:         staff.StaffType,
		StaffType:      staff.StaffType,
		Specialities:   staff.Specialities,
		Departments: DepartmentDTO{
			DepartmentID: staff.Departments.DepartmentID,
			Role:         staff.Departments.Role,
			IsPrimary:    staff.Departments.IsPrimary,
		},
		Schedule: ScheduleDTO{
			Type:   staff.Schedule.Type,
			Shifts: staff.Schedule.Shifts,
		},
		Email:       staff.Email,
		Phone:       staff.Phone,
		DateOfBirth: staff.DateOfBirth,
		Gender:      staff.Gender,
		Address: AddressDTO{
			Street:  staff.Address.Street,
			City:    staff.Address.City,
			State:   staff.Address.State,
			Country: staff.Address.Country,
			ZipCode: staff.Address.ZipCode,
		},
		Metadata:  staff.Metadata,
		CreatedAt: staff.CreatedAt,
		UpdatedAt: staff.UpdatedAt,
		DeletedAt: staff.DeletedAt,
	}
}
