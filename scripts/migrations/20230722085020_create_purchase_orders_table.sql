-- migrate:up
CREATE TABLE purchase_orders(
	id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	production_factory integer REFERENCES production_factories(id),
	pic_name varchar,
	quantity_po integer,
	quantity_production integer,
	product_item jsonb,
	payment_term integer,
	created_at timestamp,
	expired_at timestamp,
	unit_item varchar,
	description varchar,
	status varchar,
	status_history jsonb,
	finishing_factory integer REFERENCES finishing_factories(id)	
);

-- migrate:down
DROP TABLE purchase_orders;

