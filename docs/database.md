# Database Tables

## site_user
```sql
CREATE TABLE site_user (
    id SERIAL PRIMARY KEY,
    username VARCHAR(32) NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    last_login TIMESTAMP,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```