-- migrate:up
CREATE TABLE finishing_factories(
	id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	code varchar(50),
	name varchar(50)
);

-- migrate:down
DROP TABLE finishing_factories;
