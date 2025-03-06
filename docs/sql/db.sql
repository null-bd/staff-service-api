CREATE TYPE staff_type AS ENUM ('doctor', 'nurse', 'technician', 'administrative', 'support');
CREATE TYPE staff_status AS ENUM ('active', 'inactive', 'on-leave', 'terminated');
 
CREATE TABLE staffs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    branch_id UUID NOT NULL,
    organization_id UUID NOT NULL,
	first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    staff_code VARCHAR(10) NOT NULL CHECK (code ~ '^[A-Z0-9]{3,10}$'),
    type staff_type NOT NULL,
    specialty TEXT[],
    status staff_status NOT NULL,
    metadata JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(branch_id, code),
    UNIQUE(branch_id, name)
);


CREATE TYPE staff_role AS ENUM ('doctor', 'nurse', 'technician', 'administrative', 'support');
CREATE TYPE schedule_type AS ENUM ('full_time', 'part_time', 'on_call', 'rotating');
 
CREATE TABLE staff_assignments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    department_id UUID NOT NULL REFERENCES departments(id),
    staff_id UUID NOT NULL,
    role staff_role NOT NULL,
    schedule_type schedule_type NOT NULL,
    primary_department BOOLEAN DEFAULT false,
    start_date DATE NOT NULL,
    end_date DATE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(department_id, staff_id)
);
 
CREATE INDEX idx_departments_branch_id ON departments(branch_id);
CREATE INDEX idx_departments_organization_id ON departments(organization_id);
CREATE INDEX idx_staff_assignments_staff_id ON staff_assignments(staff_id);