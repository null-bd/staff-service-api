package staff

type Staff struct {
	ID             string
	BranchID       string
	OrganizationID string
	FirstName      string
	LastName       string
	Code           string
	StaffType      string
	Status         string
	Specialities   []string
	Departments    Departments
	Schedule       Schedule
	Email          string
	Phone          string
	DateOfBirth    string
	Gender         string
	Address        Address
	Metadata       map[string]interface{}
	CreatedAt      string
	UpdatedAt      string
}

type Departments struct {
	DepartmentID string
	Role         []string
	IsPrimary    bool
}

type Schedule struct {
	Type   string
	Shifts []string
}

type Address struct {
	Street  string
	City    string
	State   string
	Country string
	ZipCode string
}
