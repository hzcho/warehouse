CREATE TABLE IF NOT EXISTS roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    role_name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    login VARCHAR(255) UNIQUE NOT NULL,
    pass_hash TEXT NOT NULL,
    refresh_token TEXT,
    token_expire TIMESTAMP, 
    role_id UUID,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE SET NULL 
);

INSERT INTO roles(role_name) VALUES ('employee'), ('manager'), ('admin');
