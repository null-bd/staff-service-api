CREATE TYPE staff_type AS ENUM ('doctor', 'nurse', 'technician', 'administrative', 'support');
CREATE TYPE staff_status AS ENUM ('active', 'inactive', 'on-leave', 'terminated');
 
CREATE TABLE staffs (
    id UUID PRIMARY KEY,
    branch_id UUID NOT NULL,
    organization_id UUID NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    code VARCHAR(50) NOT NULL UNIQUE,
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
    gender VARCHAR(50) NOT NULL,
    address_street VARCHAR(100) NOT NULL, 
    address_city VARCHAR(50) NOT NULL,
    address_state VARCHAR(50) NOT NULL,
    address_country VARCHAR(50) NOT NULL,
    address_zipcode VARCHAR(50) NOT NULL,      
    metadata JSONB DEFAULT '{}'::JSONB,      
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
    UNIQUE(branch_id, code)
);

CREATE TYPE user_status AS ENUM ('active', 'inactive', 'blocked');
 
CREATE TABLE users (
    id UUID PRIMARY KEY,
    branch_id UUID NOT NULL,
    organization_id UUID NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    code VARCHAR(50) NOT NULL UNIQUE,
    status user_status NOT NULL,
    user_type UUID NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20) NOT NULL,
    date_of_birth DATE NOT NULL,
    gender VARCHAR(10) NOT NULL,
    address TEXT NOT NULL,  
    emergencycontact_name VARCHAR(100) NOT NULL,
    emergencycontact_phone VARCHAR(20) NOT NULL,
    emergencycontact_relationship VARCHAR(100) NOT NULL,  
    metadata JSONB DEFAULT '{}'::JSONB,      
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
    UNIQUE(branch_id, code)
);

CREATE TABLE usertypes (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    metadata JSONB DEFAULT '{}'::JSONB,      
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX staffs_branch_id_idx ON staffs(branch_id);
CREATE INDEX staffs_organization_id_idx ON staffs(organization_id);
CREATE INDEX users_branch_id_idx ON users(branch_id);
CREATE INDEX users_organization_id_idx ON users(organization_id);
CREATE INDEX usertypes_id_idx ON usertypes(id);