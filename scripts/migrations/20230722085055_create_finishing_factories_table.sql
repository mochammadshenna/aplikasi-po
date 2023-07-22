-- migrate:up
CREATE TABLE finishing_factories(
	id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	code varchar,
	name varchar
);

-- migrate:down
DROP TABLE finishing_factories;
