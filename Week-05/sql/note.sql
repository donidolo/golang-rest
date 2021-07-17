-- public.note definition

-- Drop table

-- DROP TABLE public.note;

CREATE TABLE public.note (
	id varchar(36) NOT NULL,
	title varchar(255) NULL,
	"content" text NULL,
	created timestamptz NULL,
	created_by uuid NULL
);