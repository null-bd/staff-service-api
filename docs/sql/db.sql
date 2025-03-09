CREATE TYPE staff_type AS ENUM ('doctor', 'nurse', 'technician', 'administrative', 'support');
CREATE TYPE staff_status AS ENUM ('active', 'inactive', 'on-leave', 'terminated');
 
CREATE TABLE staffs (
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

CREATE TYPE user_type AS ENUM ('regular', 'vip', 'corporate', 'senior citizen');
CREATE TYPE user_status AS ENUM ('active', 'inactive', 'blocked');
 
CREATE TABLE users (
    id UUID PRIMARY KEY,
    branch_id UUID NOT NULL,
    organization_id UUID NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    user_code VARCHAR(50) NOT NULL UNIQUE,
    status user_status NOT NULL,
    type user_type NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20) NOT NULL,
    date_of_birth DATE NOT NULL,
    gender VARCHAR(10) NOT NULL,
    address TEXT NOT NULL,  
    emergencycontact_name VARCHAR(100) NOT NULL,
    emergencycontact_phone VARCHAR(20) NOT NULL,
    emergencycontact_relationship VARCHAR(100) NOT NULL,
    healthhistory_condition TEXT[],
    healthhistory_diagnosis TEXT[],
    healthhistory_date DATE[],
    consultanthistory_staffId UUID[],
    consultanthistory_startdate DATE[], 
    consultanthistory_enddate DATE[],    
    metadata JSONB DEFAULT '{}'::JSONB,      
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE
    UNIQUE(branch_id, code)
);

CREATE INDEX staffs_branch_id_idx ON staffs(branch_id);
CREATE INDEX staffs_organization_id_idx ON staffs(organization_id);
CREATE INDEX users_branch_id_idx ON users(branch_id);
CREATE INDEX users_organization_id_idx ON users(organization_id);

