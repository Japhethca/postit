CREATE TABLE ps_user (
  user_id SERIAL,
  firstname VARCHAR(255),
  lastname VARCHAR(255),
  email VARCHAR(500) UNIQUE NOT NULL,
  PRIMARY KEY (user_id)
);

CREATE TABLE ps_group (
  id SERIAL,
  name VARCHAR(1000),
  PRIMARY KEY (id)
);

CREATE TABLE postit (
  id SERIAL,
  content TEXT NOT NULL,
  color VARCHAR(200),
  user_id INTEGER REFERENCES ps_user (user_id) ON DELETE CASCADE,
  category_id INTEGER REFERENCES ps_group,
  updated_at TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL,
  PRIMARY KEY (id)
);

