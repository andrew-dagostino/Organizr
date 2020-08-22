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

## post

```sql
CREATE TABLE post (
    id SERIAL NOT NULL PRIMARY KEY,
    author INTEGER NOT NULL REFERENCES site_user(id) ON DELETE CASCADE,
    title VARCHAR(64) NOT NULL,
    text VARCHAR(512),
    img BYTEA,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT chk_post_content CHECK (text IS NOT NULL OR img IS NOT NULL)
);
```
