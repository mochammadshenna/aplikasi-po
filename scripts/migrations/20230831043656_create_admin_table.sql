-- migrate:up
CREATE TABLE admins (
  id integer NOT NULL,
  name varchar(255) NOT NULL,
  created_at timestamp without time zone,
  updated_at timestamp without time zone,
  email varchar(255),
  password varchar(255)
);

-- migrate:down
DROP TABLE admins;
