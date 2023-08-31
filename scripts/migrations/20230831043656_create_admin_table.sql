-- migrate:up
CREATE TABLE admins (
  id integer NOT NULL,
  name character varying,
  created_at timestamp without time zone,
  updated_at timestamp without time zone,
  email character varying,
  password character varying
);

-- migrate:down
DROP TABLE admins;
