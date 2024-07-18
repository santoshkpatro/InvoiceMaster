CREATE TABLE IF NOT EXISTS invoices (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    invoice_number VARCHAR(255) NOT NULL,
    invoice_type VARCHAR(50) NOT NULL,
    invoice_name VARCHAR(255) NOT NULL,
    order_id BIGINT,
    order_total DECIMAL(10, 2),
    order_name VARCHAR(255),
    issue_date DATE,
    due_date DATE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    archived_at TIMESTAMPTZ,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    UNIQUE (user_id, invoice_name)
);