CREATE SCHEMA kudesnik;

CREATE TABLE kudesnik.product_categories (
    product_category_id UUID PRIMARY KEY,
    product_category_name VARCHAR(60) NOT NULL CHECK(char_length(product_category_name) BETWEEN 3 AND 60),
    installation_price DECIMAL(10,2)
);

CREATE TABLE kudesnik.producers (
    producer_id UUID PRIMARY KEY,
    company_name VARCHAR(60) NOT NULL,
    website_link TEXT
);

CREATE TABLE kudesnik.products (
    product_id UUID PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,

    product_name VARCHAR(100) NOT NULL CHECK(char_length(product_name) BETWEEN 3 AND 100),
    price DECIMAL(10,2),
    description TEXT,
    is_visible BOOLEAN DEFAULT false, 
	category_id UUID NOT NULL,
	producer_id UUID,

	FOREIGN KEY (category_id) REFERENCES kudesnik.product_categories (product_category_id),
	FOREIGN KEY (producer_id) REFERENCES kudesnik.producers (producer_id)
);

CREATE TABLE kudesnik.doors (
    door_id UUID PRIMARY KEY,
    collection VARCHAR(60) NOT NULL,
    width INTEGER,
    height INTEGER,
    outside_material TEXT,
    outside_color TEXT,
    outside_picture TEXT,
    inside_material TEXT,
    inside_color TEXT,
    inside_picture TEXT,

    FOREIGN KEY (door_id) REFERENCES kudesnik.products (product_id)
);

CREATE TABLE kudesnik.entrance_doors (
    door_id UUID PRIMARY KEY,
    strength_class TEXT,
    sound_insulation TEXT,
    metal_thickness TEXT,
    box_thickness TEXT,
    leaf_thickness TEXT,
    leaf_description TEXT,
    filling_description TEXT,
    main_lock TEXT,
    additional_lock TEXT,
    insulation_description TEXT,
    hinges TEXT,

    FOREIGN KEY (door_id) REFERENCES kudesnik.doors (door_id)
);

CREATE TABLE kudesnik.interior_doors (
    door_id UUID PRIMARY KEY,
    opening_system TEXT,
    leaf_coating TEXT,
    handle TEXT,

    FOREIGN KEY (door_id) REFERENCES kudesnik.doors (door_id)
);

CREATE TABLE kudesnik.windows (
    window_id UUID PRIMARY KEY,
    purpose TEXT NOT NULL,
    width INTEGER NOT NULL,
    height INTEGER NOT NULL,
    material TEXT NOT NULL DEFAULT 'ПВХ',

    FOREIGN KEY (window_id) REFERENCES kudesnik.products (product_id)
);

CREATE TABLE kudesnik.balconies (
    balcony_id UUID PRIMARY KEY,
    purpose TEXT NOT NULL,
    material TEXT NOT NULL,

    FOREIGN KEY (balcony_id) REFERENCES kudesnik.products (product_id)
);

CREATE TABLE kudesnik.offices (
    office_id UUID PRIMARY KEY,
    address TEXT,
    phone_number VARCHAR(20)
);

CREATE TABLE kudesnik.product_locations (
    office_id UUID REFERENCES kudesnik.offices(office_id),
    product_id UUID REFERENCES kudesnik.products(product_id),
    PRIMARY KEY (office_id, product_id)
);

CREATE TABLE kudesnik.customers (
    customer_phone_number VARCHAR(15) PRIMARY KEY CHECK(
        customer_phone_number ~ '^\+[0-9]+$' AND 
        char_length(customer_phone_number) BETWEEN 10 AND 15
    ),
    customer_name VARCHAR(60)
);

CREATE TABLE kudesnik.admins (
    email VARCHAR(255) PRIMARY KEY,
    full_name VARCHAR(60) NOT NULL,
    password_hash TEXT NOT NULL,
    admin_type VARCHAR(20) NOT NULL CHECK(
        admin_type IN ('superadmin', 'manager')
    )
);

CREATE TABLE kudesnik.customer_requests (
    request_id UUID PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    desired_date DATE,
    desired_time TIME,
    extra_comment TEXT,
    customer_phone_number VARCHAR(15) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    handled_at TIMESTAMPTZ,
    handler_admin_email VARCHAR(255),
    product_id UUID,
    status VARCHAR(20) NOT NULL CHECK(
        status IN ('new', 'in_progress', 'completed', 'cancelled')
    ),

    CHECK (
        (status in ('completed', 'cancelled') AND handled_at IS NOT NULL AND handler_admin_email IS NOT NULL AND handled_at >= created_at) 
        OR
        (status = 'in_progress' AND handled_at IS NOT NULL) 
        OR 
        (status = 'new' AND handled_at IS NULL AND handler_admin_email IS NULL)
    ),

    FOREIGN KEY (product_id) REFERENCES kudesnik.products(product_id),
    FOREIGN KEY (handler_admin_email) REFERENCES kudesnik.admins(email),
    FOREIGN KEY (customer_phone_number) REFERENCES kudesnik.customers(customer_phone_number)
);