CREATE TABLE consignments
(
    id               integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    barcode          text COLLATE pg_catalog."default",
    link_to_supplier text COLLATE pg_catalog."default",
    returned_at      date,
    CONSTRAINT consignments_pkey PRIMARY KEY (id)
);

CREATE TABLE customers
(
    id           integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    name         text COLLATE pg_catalog."default",
    postcode     text COLLATE pg_catalog."default",
    address      text COLLATE pg_catalog."default",
    geo_location text COLLATE pg_catalog."default"
);

CREATE TABLE van
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    name text COLLATE pg_catalog."default",
    latitude double precision,
    longitude double precision,
    CONSTRAINT "Van_pkey" PRIMARY KEY (id)
)