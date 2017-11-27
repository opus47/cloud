-- The Opus47 PostgreSQL database

CREATE EXTENSION pgcrypto;
CREATE EXTENSION unaccent;

---
--- Composer / Piece Information
---

CREATE TABLE composers (
  id      UUID  PRIMARY KEY DEFAULT gen_random_uuid(),
  first   text  NOT NULL,
  middle  text,
  last    text  NOT NULL,

  UNIQUE (first, middle, last)
);

CREATE TABLE keys (
  id   UUID   PRIMARY KEY DEFAULT gen_random_uuid(),
  name text   NOT NULL,

  UNIQUE (name)
);

CREATE TABLE pieces (
  id        UUID      PRIMARY KEY DEFAULT gen_random_uuid(),
  composer  UUID      REFERENCES composers NOT NULL,
  title     text      NOT NULL,
  key       UUID      REFERENCES keys NOT NULL,
  number    integer,
  catalog   text,

  UNIQUE (composer, title, key, number)
);


CREATE TABLE movements (
  id      UUID      PRIMARY KEY DEFAULT gen_random_uuid(),
  piece   UUID      REFERENCES pieces NOT NULL,
  title   text      NOT NULL,
  number  integer,

  UNIQUE(piece, title)
);

CREATE TABLE parts (
  id    UUID  PRIMARY KEY DEFAULT gen_random_uuid(),
  name  text  NOT NULL,

  UNIQUE(name)
);

CREATE TABLE piece_parts (
  piece     UUID  REFERENCES pieces NOT NULL,
  part      UUID  REFERENCES parts NOT NULL,

  UNIQUE(piece, part)
);

---
--- Musician / Performance Information
---

CREATE TABLE musicians (
  id      UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  first   text NOT NULL,
  middle  text,
  last    text NOT NULL,

  UNIQUE(first, middle, last)
);

CREATE TABLE performances (
  id            UUID  PRIMARY KEY DEFAULT gen_random_uuid(),
  piece         UUID  REFERENCES pieces NOT NULL,
  title         text, -- TODO change to veneu
  description   text
);

CREATE TABLE recordings (
  id            UUID  PRIMARY KEY DEFAULT gen_random_uuid(),
  performance   UUID  REFERENCES performances NOT NULL,
  movement      UUID  REFERENCES movements NOT NULL,
  file          text,

  UNIQUE(performance, movement)
);

CREATE TABLE performers (
  performance   UUID  REFERENCES performances NOT NULL, 
  musician      UUID  REFERENCES musicians NOT NULL,
  part          UUID  REFERENCES parts NOT NULL,

  UNIQUE(performance, musician, part)
);

-- Create a materialized view for seraching movements
-- To update this view use `REFRESH MATERIALIZED VIEW mv_movements`
--
-- searching:
-- select cname, ptitle, kname from mv_movements 
-- where document @@ to_tsquery('english', 'antonin')

CREATE MATERIALIZED VIEW mv_movements AS
SELECT 
  m.id AS mid,
  p.id AS pid,
  c.first AS cfirst,
  c.last AS clast,
  p.title AS ptitle,
  k.name AS kname,
  m.title AS mname,
  to_tsvector('english', c.first) ||
  to_tsvector('english', c.last) ||
  to_tsvector('english', m.title) ||
  to_tsvector('english', to_char(m.number, '999')) ||
  to_tsvector('english', p.title) || 
  to_tsvector('english', to_char(p.number, '999')) ||
  to_tsvector('english', p.catalog) ||
  to_tsvector('english', k.name) as document
FROM movements AS m
JOIN pieces AS p on m.piece = p.id
JOIN composers AS c on p.composer = c.id
JOIN keys AS k on p.key = k.id
;
CREATE INDEX idx_movement_search ON mv_movements USING gin(document);

-- Create a materialized view for seraching pieces
CREATE MATERIALIZED VIEW mv_pieces AS
SELECT 
  p.id AS pid,
  c.first cfirst,
  c.last clast,
  p.title AS ptitle,
  k.name AS kname,
  p.catalog AS pcatalog,
  to_tsvector('english', unaccent(c.first)) ||
  to_tsvector('english', unaccent(c.last)) ||
  to_tsvector('english', p.title) || 
  to_tsvector('english', to_char(p.number, '999')) ||
  to_tsvector('english', p.catalog) ||
  to_tsvector('english', k.name) AS document
FROM pieces AS p
JOIN composers AS c on p.composer = c.id
JOIN keys AS k on p.key = k.id
;
CREATE INDEX idx_piece_search ON mv_pieces USING gin(document);

