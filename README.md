# Gator

Gator is a command-line tool for managing RSS feeds. It provides a simple interface to register users, add and follow feeds, and browse aggregated posts, all backed by a PostgreSQL database.

---

## Requirements

- **Go** (version 1.24.4 or newer)
- **PostgreSQL** (running locally or accessible remotely)

---

## Installation

1. **Install Go**: [Download and install Go](https://golang.org/dl/)
2. **Install Gator CLI**:

   ```sh
   go install github.com/frozendolphin/Gator@latest
   ```
   This will install the `gator` binary to your `$GOPATH/bin` (usually `~/go/bin`). Make sure this directory is in your `PATH`.

---

## Database Setup

Follow these steps to set up your PostgreSQL database and run the migrations required by Gator.

### 1. Install PostgreSQL

#### Linux:
```sh
sudo apt update
sudo apt install postgresql postgresql-contrib
```

#### macOS (using Homebrew):
```sh
brew install postgresql@15
```

### 2. Start the PostgreSQL Server

#### Linux:
```sh
sudo service postgresql start
```

#### macOS:
```sh
brew services start postgresql@15
```

### 3. Set the Postgres User Password (Linux only)
```sh
sudo passwd postgres
```

### 4. Access the PostgreSQL Shell

#### Linux:
```sh
sudo -u postgres psql
```

#### macOS:
```sh
psql postgres
```

### 5. Create the Database and Set Password

In the PostgreSQL shell, run:
```sql
CREATE DATABASE gator;
\c gator
-- (Linux only) Set the user password:
ALTER USER postgres PASSWORD 'postgres';
```

### 6. Run Database Migrations

Install [goose](https://github.com/pressly/goose):
```sh
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Navigate to the schema directory:
```sh
cd sql/schema
```

Run the migrations:
```sh
goose postgres "postgres://postgres:password@localhost:5432/gator?sslmode=disable" up
```

> **Note:** Adjust the connection string as needed for your environment (username, password, host, port).

---

## Configuration

Gator uses a config file located at `~/.gatorconfig.json`.

**Example config:**
```json
{
  "db_url": "postgres://postgres:password@localhost:5432/gator?sslmode=disable",
  "username": "any"
}
```
- `db_url`: Your PostgreSQL connection string.
- `username`: Your Gator username (set automatically when you use the `register` or `login` command).

---

## Usage

Run the CLI with:
```sh
gator <command> [arguments]
```

If you installed with `go install`, you may need to run it as:
```sh
~/go/bin/gator <command> [arguments]
```

### First Steps
1. **Register a user:**
   ```sh
   gator register <username>
   ```
2. **Login as a user:**
   ```sh
   gator login <username>
   ```
   This sets your username in the config file.

---

## Commands & Features

- **login <username>**: Set your username for the session (updates config file).
- **register <username>**: Create a new user and set as current.
- **reset**: Delete all users (dangerous, use with caution).
- **users**: List all users, marking the current one.
- **agg <duration>**: Aggregate (scrape) feeds at the given interval (e.g., `10m`, `1h`).
- **addfeed <feedname> <url>**: Add a new feed and follow it.
- **feeds**: List all feeds in the system.
- **follow <feed_url>**: Follow a feed by its URL.
- **following**: List feeds you are currently following.
- **unfollow <feed_url>**: Unfollow a feed by its URL.
- **browse [limit]**: Browse posts from feeds you follow (optionally limit the number shown, default is 2).

---

## Example Workflow

1. **Register and login:**
   ```sh
   gator register alice
   gator login alice
   ```
2. **Add and follow a feed:**
   ```sh
   gator addfeed "My Blog" https://example.com/rss
   ```
3. **List all feeds:**
   ```sh
   gator feeds
   ```
4. **Follow another feed:**
   ```sh
   gator follow https://another.com/rss
   ```
5. **See what you are following:**
   ```sh
   gator following
   ```
6. **Browse posts:**
   ```sh
   gator browse 5
   ```
7. **Unfollow a feed:**
   ```sh
   gator unfollow https://another.com/rss
   ```

---

## Development & Contributing

- Database queries are generated using [sqlc](https://github.com/kyleconroy/sqlc). See `sqlc.yaml` for configuration.
- To regenerate database code after editing SQL files:
  ```sh
  sqlc generate
  ```
- Database schema is in `sql/schema/`.

---

## License

This project is open source and available under the MIT License.
