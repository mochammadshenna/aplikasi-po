-- migrate:up
CREATE TABLE production_factories(
	id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	name varchar(50)
);

-- migrate:down
DROP TABLE production_factories;

