package database

func (c *DatabaseClient) Migrate() {
	c.DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY,
			first_name TEXT NOT NULL,
			last_name TEXT NOT NULL,
			user_name TEXT NOT NULL,
			password TEXT NOT NULL
		);
	`)
}
