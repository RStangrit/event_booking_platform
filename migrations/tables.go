package migrations

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Table struct {
	Name  string
	Query string
}

var tables = []Table{
	{
		Name: "UsersTable",
		Query: `
            CREATE TABLE IF NOT EXISTS users(
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                name TEXT NOT NULL,
                email TEXT NOT NULL UNIQUE,
                password TEXT NOT NULL,
                role TEXT CHECK(role IN ('admin', 'user')) NOT NULL,
                created_at DATETIME NOT NULL,
                updated_at DATETIME
            )
        `,
	},
	{
		Name: "EventsTable",
		Query: `
            CREATE TABLE IF NOT EXISTS events(
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                title TEXT NOT NULL UNIQUE,
                description TEXT NOT NULL,
                date TEXT NOT NULL,
                location TEXT NOT NULL,
                capacity INTEGER NOT NULL,
                price FLOAT,
                created_by INTEGER NOT NULL,
                created_at DATETIME NOT NULL,
                updated_at DATETIME,
                FOREIGN KEY(created_by) REFERENCES users(id)
            )
        `,
	},
	{
		Name: "BookingsTable",
		Query: `
            CREATE TABLE IF NOT EXISTS bookings(
                id INTEGER PRIMARY KEY AUTOINCREMENT,
				user_id INTEGER NOT NULL,
				event_id INTEGER NOT NULL,
				booking_date DATETIME NOT NULL,
				status TEXT CHECK(status IN ('pending', 'confirmed', 'cancelled')) NOT NULL,
				created_at DATETIME NOT NULL,
                FOREIGN KEY(user_id) REFERENCES users(id),
				FOREIGN KEY(event_id) REFERENCES events(id)
            )
        `,
	},
	{
		Name: "CategoriesTable",
		Query: `
            CREATE TABLE IF NOT EXISTS categories(
                id INTEGER PRIMARY KEY AUTOINCREMENT,
				name TEXT NOT NULL UNIQUE,
				created_at DATETIME NOT NULL,
                updated_at DATETIME
            )
        `,
	},
	{
		Name: "EventCategoriesTable",
		Query: `
            CREATE TABLE IF NOT EXISTS event_categories(
                id INTEGER PRIMARY KEY AUTOINCREMENT,
				event_id INTEGER NOT NULL,
				category_id INTEGER NOT NULL,
				FOREIGN KEY(event_id) REFERENCES events(id),
				FOREIGN KEY(category_id) REFERENCES categories(id)
            )
        `,
	},
	{
		Name: "ReviewsTable",
		Query: `
            CREATE TABLE IF NOT EXISTS reviews(
                id INTEGER PRIMARY KEY AUTOINCREMENT,
				event_id INTEGER NOT NULL,
				user_id INTEGER NOT NULL,
				rating INTEGER CHECK(rating IN ('1', '2', '3', '4', '5')) NOT NULL,
				comment TEXT,
				created_at DATETIME NOT NULL,
				FOREIGN KEY(event_id) REFERENCES events(id),
				FOREIGN KEY(user_id) REFERENCES users(id)
            )
        `,
	},
	{
		Name: "TicketsTable",
		Query: `
            CREATE TABLE IF NOT EXISTS tickets(
                id INTEGER PRIMARY KEY AUTOINCREMENT,
				booking_id INTEGER NOT NULL,
				ticket_number TEXT UNIQUE,
				issued_at DATETIME NOT NULL,
				FOREIGN KEY(booking_id) REFERENCES bookings(id)
            )
        `,
	},
	{
		Name: "PaymentsTable",
		Query: `
            CREATE TABLE IF NOT EXISTS payments(
                id INTEGER PRIMARY KEY AUTOINCREMENT,
				user_id INTEGER NOT NULL,
				event_id INTEGER NOT NULL,
				amount FLOAT NOT NULL,
				payment_date DATETIME NOT NULL,
				payment_method TEXT CHECK(payment_method IN('card', 'PayPal', 'cash', 'on_word_of_honor')),
				status TEXT CHECK(status IN('success', 'failed')),
				FOREIGN KEY(user_id) REFERENCES users(id),
				FOREIGN KEY(event_id) REFERENCES events(id)
            )
        `,
	},
	{
		Name: "NotificationsTable",
		Query: `
            CREATE TABLE IF NOT EXISTS notifications(
                id INTEGER PRIMARY KEY AUTOINCREMENT,
				user_id INTEGER NOT NULL,
				message TEXT NOT NULL,
				is_read INTEGER NOT NULL CHECK(is_read IN (0, 1)),
				created_at DATETIME NOT NULL,
				FOREIGN KEY(user_id) REFERENCES users(id)
            )
        `,
	},
	{
		Name: "OrganizationsTable",
		Query: `
            CREATE TABLE IF NOT EXISTS organizations(
                id INTEGER PRIMARY KEY AUTOINCREMENT,
				name TEXT NOT NULL,
				email TEXT NOT NULL,
				phone TEXT NOT NULL,
				created_at DATETIME NOT NULL
            )
        `,
	},
}

func CreateTables(DB *sql.DB) {
	for i := range tables {
		createTable(tables[i], DB)
	}
}

func createTable(table Table, DB *sql.DB) {
	_, err := DB.Exec(table.Query)
	if err != nil {
		log.Fatalf("could not create table %s: %v", table.Name, err)
	}
}
