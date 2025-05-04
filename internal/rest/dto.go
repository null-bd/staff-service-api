package rest

type CreateStaffRequest struct {
	BranchID       string                 `json:"branchId"`
	OrganizationID string                 `json:"organizationId"`
	FirstName      string                 `json:"firstName" validate:"required"`
	LastName       string                 `json:"lastName" validate:"required"`
	Code           string                 `json:"code" validate:"required"`
	StaffType      string                 `json:"type" validate:"required,oneof=doctor nurse technician administrative support"`
	Specialities   []string               `json:"specialties" validate:"required,dive,required"`
	Departments    DepartmentDTO          `json:"departments"`
	Schedule       ScheduleDTO            `json:"schedule"`
	Email          string                 `json:"email" validate:"required"`
	Phone          string                 `json:"phone" validate:"required"`
	DateOfBirth    string                 `json:"dateOfBirth" validate:"required,datetime=2006-01-02"`
	Gender         string                 `json:"gender" validate:"required"`
	Address        AddressDTO             `json:"address" validate:"required"`
	Metadata       map[string]interface{} `json:"metadata"`
}

type DepartmentDTO struct {
	DepartmentID string   `json:"departmentId" validate:"required"`
	Role         []string `json:"role" validate:"required"`
	IsPrimary    bool     `json:"isPrimary"`
}

type ScheduleDTO struct {
	Type   string   `json:"type" validate:"required"`
	Shifts []string `json:"shifts"`
}

type AddressDTO struct {
	Street  string `json:"street" validate:"required"`
	City    string `json:"city" validate:"required"`
	State   string `json:"state" validate:"required"`
	Country string `json:"country" validate:"required"`
	ZipCode string `json:"zipCode" validate:"required"`
}

type StaffResponse struct {
	ID             string                 `json:"Id" validate:"required"`
	BranchID       string                 `json:"branchId" validate:"required"`
	OrganizationID string                 `json:"organizationId" validate:"required"`
	FirstName      string                 `json:"firstName" validate:"required"`
	LastName       string                 `json:"lastName" validate:"required"`
	Code           string                 `json:"code" validate:"required"`
	Status         string                 `json:"status" validate:"required"`
	StaffType      string                 `json:"type" validate:"required,oneof=doctor nurse technician administrative support"`
	Specialities   []string               `json:"specialties" validate:"required"`
	Departments    DepartmentDTO          `json:"departments" validate:"required"`
	Schedule       ScheduleDTO            `json:"schedule" validate:"required"`
	Email          string                 `json:"email" validate:"required"`
	Phone          string                 `json:"phone" validate:"required"`
	DateOfBirth    string                 `json:"dateOfBirth"`
	Gender         string                 `json:"gender" validate:"required"`
	Address        AddressDTO             `json:"address" validate:"required"`
	Metadata       map[string]interface{} `json:"metadata"`
	CreatedAt      string                 `json:"createdAt"`
	UpdatedAt      string                 `json:"updatedAt"`
	DeletedAt      string                 `json:"deletedAt"`
}
