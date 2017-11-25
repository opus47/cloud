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

CREATE TABLE movement_parts (
  movement  UUID  REFERENCES movements NOT NULL,
  part      UUID  REFERENCES parts NOT NULL,

  UNIQUE(movement, part)
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
  title         text,
  description   text
);

CREATE TABLE recordings (
  id            UUID  PRIMARY KEY DEFAULT gen_random_uuid(),
  performance   UUID  REFERENCES performances NOT NULL,
  movement      UUID  REFERENCES movements NOT NULL,
  file          text,

  UNIQUE(performance, movement)
);

CREATE TABLE performer (
  performance   UUID  REFERENCES performances NOT NULL, 
  musician      UUID  REFERENCES musicians NOT NULL,
  part          UUID  REFERENCES parts NOT NULL,

  UNIQUE(performance, musician, part)
);


