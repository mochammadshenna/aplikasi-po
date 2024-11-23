-- migrate:up
CREATE TABLE purchase_orders(
	id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	production_factory integer REFERENCES production_factories(id),
	pic_name varchar(50),
	quantity_po integer,
	quantity_production integer,
	product_item jsonb,
	payment_term integer,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP,
	expired_at timestamp,
	unit_item varchar(50),
	description varchar(255),
	status varchar(50),
	status_history jsonb,
	finishing_factory integer REFERENCES finishing_factories(id)	
);

-- migrate:down
DROP TABLE purchase_orders;

