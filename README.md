# gator
Blog Aggregator in Go—Boot.dev Project

Gator is a command-line RSS feed aggregator. It lets you subscribe to RSS feeds from your favorite blogs and news sites, then browse all the posts in one place.

## Prerequisites

Before you can use Gator, you'll need two things installed on your computer:

### 1. Go

Go is the programming language Gator is written in. You'll need it to install and run the program.

**Mac:**
```bash
brew install go
```

**Linux:**
```bash
sudo apt install golang-go
```

Verify it's installed by running:
```bash
go version
```

### 2. PostgreSQL

PostgreSQL is the database where Gator stores users, feeds, and posts.

**Mac:**
```bash
brew install postgresql@15
brew services start postgresql@15
```

You may also need to add PostgreSQL to your PATH:
```bash
echo 'export PATH="/opt/homebrew/opt/postgresql@15/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

**Linux:**
```bash
sudo apt install postgresql
sudo service postgresql start
```

Verify it's installed by running:
```bash
psql --version
```

### 3. Create the database

Connect to PostgreSQL and create the Gator database:
```bash
psql postgres
```

Then run:
```sql
CREATE DATABASE gator;
\q
```

## Installation

Once you have Go and PostgreSQL set up, install Gator with:
```bash
go install github.com/ankamason/gator@latest
```

This downloads the code, compiles it, and puts the `gator` binary in your Go bin directory.

If you get "command not found" when trying to run `gator`, add Go's bin directory to your PATH:
```bash
echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

## Configuration

Gator needs a config file to know how to connect to your database. Create a file called `.gatorconfig.json` in your home directory:
```bash
cat > ~/.gatorconfig.json << 'CONFIG'
{
  "db_url": "postgres://YOUR_USERNAME:@localhost:5432/gator?sslmode=disable"
}
CONFIG
```

Replace `YOUR_USERNAME` with your computer username (the one you see in your terminal prompt).

## Usage

Here are the commands you can run:

### Register a new user
```bash
gator register alice
```

This creates a new user and logs you in as that user.

### Log in as an existing user
```bash
gator login alice
```

### List all users
```bash
gator users
```

Shows all registered users, with `(current)` next to the one you're logged in as.

### Add a feed
```bash
gator addfeed "Boot.dev Blog" "https://blog.boot.dev/index.xml"
```

Adds an RSS feed to your account.

### Reset the database
```bash
gator reset
```

Deletes all users and feeds. Useful for testing.

## Development

If you want to work on Gator yourself:

1. Clone the repo
2. Run database migrations with Goose
3. Use `go run .` instead of `gator` to run your local code
```bash
git clone https://github.com/ankamason/gator.git
cd gator
go run . register yourname
```
