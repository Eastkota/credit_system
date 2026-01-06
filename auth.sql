CREATE TABLE auth.users (
    id UUID PRIMARY KEY, -- Primary key for the user record
    email VARCHAR(50) UNIQUE, -- Unique for preventing duplicate accounts by email, can be NULL
    mobile VARCHAR(20) UNIQUE, -- Unique for preventing duplicate accounts by mobile, can be NULL
    identifier VARCHAR(11) UNIQUE NOT NULL, -- Unique 11 char identifier, NOT NULL, for profile URLs
    password_hash TEXT NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'active', -- e.g., 'active', 'inactive', 'suspended'
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    -- Constraint to ensure at least one primary login method (email or mobile) is present
    CONSTRAINT chk_primary_login_present CHECK (email IS NOT NULL OR mobile IS NOT NULL)
);

CREATE INDEX idx_users_email ON auth.users(email);
CREATE INDEX idx_users_mobile ON auth.users(mobile);
CREATE INDEX idx_users_user_identifier ON auth.users(identifier);

CREATE TABLE auth.auth_clients (
    id UUID PRIMARY KEY, -- Primary key for the auth client record
    name VARCHAR(255) NOT NULL, -- Name of the auth client
    secret TEXT NOT NULL, -- Secret for the auth client
    provider VARCHAR(50) NOT NULL, -- e.g., 'google', 'facebook', 'local'
    redirect TEXT NOT NULL, -- Redirect URI for the auth client
    personal_access_client BOOLEAN NOT NULL DEFAULT FALSE, -- Indicates if this is a personal access client
    password_client BOOLEAN NOT NULL DEFAULT FALSE, -- Indicates if this is a password client
    revoked BOOLEAN NOT NULL DEFAULT FALSE, -- Indicates if the client is revoked
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE auth.access_tokens (
    id UUID PRIMARY KEY, -- Primary key for the access token record
    user_id UUID NOT NULL REFERENCES auth.users(id) ON DELETE CASCADE, -- Foreign key to users table
    client_id UUID NOT NULL REFERENCES auth.auth_clients(id) ON DELETE CASCADE, -- Foreign key to auth clients table
    token TEXT NOT NULL UNIQUE, -- Unique access token
    scopes TEXT[], -- Array of scopes granted to the access token
    revoked BOOLEAN NOT NULL DEFAULT FALSE, -- Indicates if the access token is revoked
    expires_at TIMESTAMP WITH TIME ZONE, -- Expiration time for the access token
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_access_tokens_user_id ON auth.access_tokens(user_id);
CREATE INDEX idx_access_tokens_client_id ON auth.access_tokens(client_id);

CREATE TABLE auth.refresh_tokens (
    id UUID PRIMARY KEY, -- Primary key for the refresh token record
    access_token TEXT NOT NULL REFERENCES auth.access_tokens(token) ON DELETE CASCADE, -- Foreign key to access tokens
    token TEXT NOT NULL UNIQUE, -- Unique refresh token
    revoked BOOLEAN NOT NULL DEFAULT FALSE, -- Indicates if the refresh token is revoked
    expires_at TIMESTAMP WITH TIME ZONE, -- Expiration time for the refresh token
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_refresh_tokens_access_token ON auth.refresh_tokens(access_token);



