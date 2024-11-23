-- migrate:up
CREATE TABLE admins (
  id integer NOT NULL,
  name varchar(255) NOT NULL,
  created_at timestamp without time zone,
  updated_at timestamp without time zone,
  email varchar(255),
  password varchar(255),
);

-- migrate:down
DROP TABLE admins;


ALTER TABLE admins
ADD COLUMN role varchar(50) DEFAULT 'admin',
ADD COLUMN status varchar(50) DEFAULT 'active',
ADD COLUMN provider varchar(50) DEFAULT 'email',
ADD COLUMN picture text,
ALTER COLUMN id SET DEFAULT nextval('admins_id_seq'),
ALTER COLUMN created_at SET DEFAULT CURRENT_TIMESTAMP,
ALTER COLUMN updated_at SET DEFAULT CURRENT_TIMESTAMP;

-- Add unique constraint to email
ALTER TABLE admins
ADD CONSTRAINT admins_email_unique UNIQUE (email);


-- migrate:down
ALTER TABLE admins
DROP COLUMN role,
DROP COLUMN status,
DROP COLUMN provider,
DROP COLUMN picture,
ALTER COLUMN id DROP DEFAULT,
ALTER COLUMN created_at DROP DEFAULT,
ALTER COLUMN updated_at DROP DEFAULT;

ALTER TABLE admins
DROP CONSTRAINT IF EXISTS admins_email_unique;
