CREATE TABLE image (
	id INTEGER PRIMARY KEY,
	created_at TEXT NOT NULL,
	deleted_at TEXT,
	url TEXT NOT NULL,
);

CREATE TABLE containers (
	id INTEGER PRIMARY KEY,
	created_at TEXT NOT NULL,
	deleted_at TEXT,
	image_id INTEGER NOT NULL,
	name TEXT NOT NULL,
	FOREIGN KEY (image_id)
		REFERENCES image (id)
			ON DELETE CASCADE
			ON UPDATE NO ACTION,
);

CREATE TABLE exposure (
	id INTEGER PRIMARY KEY,
	address TEXT NOT NULL,
	container_id INTEGER,
	container_port INTEGER,
	created_at TEXT NOT NULL,
	deleted_at TEXT,
	host_port INTEGER,
	protocol TEXT CHECK( protocol IN ('tcp', 'udp') NOT NULL,
	FOREIGN KEY (container_id)
		REFERENCES containers (id)
			ON DELETE CASCADE
			ON UPDATE NO ACTION,
	);
);

CREATE TABLE environment_value (
	id INTEGER PRIMARY KEY,
	container_id INTEGER NOT NULL,
	created_at TEXT NOT NULL,
	deleted_at TEXT,
	key TEXT NOT NULL,
	value TEXT NOT NULL,
	created_at TEXT NOT NULL,
	deleted_at TEXT,
	FOREIGN KEY (container_id)
		REFERENCES containers (id)
			ON DELETE CASCADE
			ON UPDATE NO ACTION,
	);

CREATE TABLE mount (
	id INTEGER PRIMARY KEY,
	container_id INTEGER NOT NULL,
	created_at TEXT NOT NULL,
	deleted_at TEXT,
	destination_path TEXT NOT NULL,
	source_path TEXT NOT NULL,
	FOREIGN KEY (container_id)
		REFERENCES containers (id)
			ON DELETE CASCADE
			ON UPDATE NO ACTION,
	);
)
