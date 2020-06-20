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
    geo_location text COLLATE pg_catalog."default",
    latitude     double precision,
    longitude    double precision
);

CREATE TABLE van
(
    id        integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    name      text COLLATE pg_catalog."default",
    latitude  double precision,
    longitude double precision,
    CONSTRAINT "Van_pkey" PRIMARY KEY (id)
);

CREATE TABLE van_run
(
    van_id  integer,
    cons_id integer,
    CONSTRAINT cons_id FOREIGN KEY (van_id)
        REFERENCES consignments (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT van_id FOREIGN KEY (van_id)
        REFERENCES van (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

CREATE TABLE "order"
(
    id          integer NOT NULL,
    customer_id integer,
    cons_id     integer,
    CONSTRAINT order_pkey PRIMARY KEY (id),
    CONSTRAINT cons_id FOREIGN KEY (cons_id)
        REFERENCES consignments (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT customer_id FOREIGN KEY (customer_id)
        REFERENCES customers (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)