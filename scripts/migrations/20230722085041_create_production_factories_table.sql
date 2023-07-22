-- migrate:up
CREATE TABLE production_factories(
	id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	name varchar
);

-- migrate:down
DROP TABLE production_factories;

