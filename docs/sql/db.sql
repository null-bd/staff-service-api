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
