CREATE TYPE staff_type AS ENUM ('doctor', 'nurse', 'technician', 'administrative', 'support');
CREATE TYPE staff_status AS ENUM ('active', 'inactive', 'on-leave', 'terminated');
 
CREATE TABLE staff (
    id UUID PRIMARY KEY,
    branch_id UUID NOT NULL,
    organization_id UUID NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    staff_code VARCHAR(50) NOT NULL UNIQUE,
    status staff_status NOT NULL,
    type staff_type NOT NULL,
    specialties TEXT[],  
    departments_Id UUID NOT NULL,
    departments_role TEXT[],
    departments_isprimary BOOLEAN,   
    schedule_type VARCHAR(50) NOT NULL,
    schedule_shifts TEXT[],     
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20) NOT NULL,
    date_of_birth DATE NOT NULL,
    gender VARCHAR(10) NOT NULL,
    address TEXT NOT NULL,       
    metadata JSONB DEFAULT '{}'::JSONB,      
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE
    UNIQUE(branch_id, code)
);


