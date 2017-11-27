-- Opus47 seed data

--- ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
---
--- Composers
---
--- ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

INSERT INTO composers (first, middle, last) values
('Johannes', NULL, 'Brahms'),
('Antonin', 'Leopold', 'Dvořák'),
('Wolfgang', 'Amadeus', 'Mozart'),
('Ludwig', 'van', 'Beethoven')
;

--- ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
---
--- Keys
---
--- ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

INSERT INTO keys (name) values
('C Major'),
('G Major'),
('D Major'),
('A Major'),
('E Major'),
('B Major'),
('F Sharp Major'),
('C Sharp Major'),
('A Minor'),
('E Minor'),
('B Minor'),
('F Sharp Minor'),
('C Sharp Minor'),
('G Sharp Minor'),
('D Sharp Minor'),

('F Major'),
('B Flat Major'),
('E Flat Major'),
('A Flat Major'),
('D Flat Major'),
('G Flat Major'),
('C Flat Major'),
('D Minor'),
('G Minor'),
('C Minor'),
('F Minor'),
('B Flat Minor'),
('E Flat Minor'),
('A Flat minor')
;

--- ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
---
--- Parts
---
--- ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

INSERT INTO parts (name, rank) values
('Violin', 1),
('Violin 1', 2),
('Violin 2', 3),
('Violin 3', 4),
('Violin 4', 5),
('Viola', 6),
('Viola 1', 7),
('Viola 2', 8),
('Viola 3', 9),
('Viola 4', 10),
('Cello', 11),
('Cello 1', 12),
('Cello 2', 13),
('Piano', 14),
('Piano 1', 15),
('Piano 2', 16),
('Flute', 17),
('Harp', 18),
('Clarinet', 19),
('Oboe', 20),
('Horn', 21),
('Bass', 22)
;

--- ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
---
--- Pieces
---
--- ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

CREATE FUNCTION new_piece(
  pcomposer_first text, 
  pcomposer_last text, 
  ptitle text, 
  pkey text, 
  pnumber integer, 
  pcatalog text,
  pmovements text[],
  pparts text[]
) RETURNS void AS $$
DECLARE
  n_movements integer := array_length(pmovements, 1);
  n_parts integer := array_length(pparts, 1);
  composer_id UUID;
  key_id UUID;
  piece_id UUID;
BEGIN

  SELECT id INTO composer_id FROM composers 
  WHERE unaccent(first) = pcomposer_first AND unaccent(last) = pcomposer_last
  ;

  SELECT id INTO key_ID FROM keys where name = pkey
  ;

  INSERT INTO pieces (composer, title, key, number, catalog) values
  (composer_id, ptitle, key_id, pnumber, pcatalog) 
  RETURNING id INTO piece_id
  ;

  FOR i IN 1..n_movements
  LOOP
    INSERT INTO movements (piece, title, number) values
    (piece_id, pmovements[i], i);
  END LOOP;

  FOR i IN 1..n_parts
  LOOP
    INSERT INTO piece_parts (piece, part) values
    (piece_id, (select id from parts where name ilike pparts[i]));
  END LOOP;
END;
$$ LANGUAGE plpgsql;

SELECT new_piece(
  'Johannes', 'Brahms',
  'Piano Quartet',
  'A Major',
  2,
  'Opus 26',
  ARRAY[
    'Allegro non troppo',
    'Poco adagio',
    'Scherzo. Poco Allegro -Trio',
    'Finale. Allegro'
  ],
  Array[
    'Violin',
    'Viola',
    'Cello',
    'Piano'
  ]
);

SELECT new_piece(
  'Antonin', 'Dvorak',
  'String Quartet',
  'F Major',
  12,
  'Opus 96',
  ARRAY[
  'Allegro ma non troppo',
  'Lento',
  'Molto vivace',
  'Finale. Vivace ma non troppo'
  ],
  Array[
    'Violin 1',
    'Violin 2',
    'Viola',
    'Cello'
  ]
);

SELECT new_piece(
  'Antonin', 'Dvorak',
  'String Quintet',
  'E Flat Major',
  3,
  'Opus 97',
  ARRAY[
    'Allegro non tanto',
    'Allegro vivo',
    'Larghetto',
    'Finale. Allegro giusto'
  ],
  Array[
    'Violin 1',
    'Violin 2',
    'Viola 1',
    'Viola 2',
    'Cello'
  ]
);

SELECT new_piece(
  'Ludwig', 'Beethoven',
  'String Quartet',
  'C Sharp Minor',
  14,
  'Opus 131',
  ARRAY[
    'Adagio ma non troppo e molto espressivo',
    'Allegro molto vivace',
    'Allegro moderato',
    'Andante ma non troppo e molto cantabile',
    'Presto',
    'Adagio quasi un poco andante',
    'Allegro'
  ],
  ARRAY[
    'Violin 1',
    'Violin 2',
    'Viola',
    'Cello'
  ]
);

  

--- ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
---
--- Musicians
---
--- ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

INSERT INTO musicians (first, middle, last) values
('Tien-Hsin', 'Cindy', 'Wu'),
('Orion', NULL, 'Weiss'),
('Jessica', NULL, 'Lee'),
('Che-Yen', 'Brian', 'Chen'),
('Nicholas', NULL, 'Canellakis'),
('Dimitri', NULL, 'Murrath'),
('Alan', NULL, 'Iglitzin'),
('Matthew', NULL, 'Zalkind')
;

REFRESH MATERIALIZED VIEW mv_movements;
REFRESH MATERIALIZED VIEW mv_pieces;

--- ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
---
--- Performances
---
--- ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~


-- Beethoven Opus 131 Oly ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

INSERT INTO performances (piece, title) values
(
  (SELECT id FROM pieces WHERE catalog = 'Opus 131'), 
  'Olympic Music Festival 2014'
);

INSERT INTO performers (performance, musician, part) values
(
  (SELECT id FROM performances WHERE title ilike '%olympic%'),
  (SELECT id FROM musicians WHERE first = 'Jessica'),
  (SELECT id FROM parts where name = 'Violin 1')
),
(
  (SELECT id FROM performances WHERE title ilike '%olympic%'),
  (SELECT id FROM musicians WHERE first = 'Tien-Hsin'),
  (SELECT id FROM parts where name = 'Violin 2')
),
(
  (SELECT id FROM performances WHERE title ilike '%olympic%'),
  (SELECT id FROM musicians WHERE first = 'Alan'),
  (SELECT id FROM parts where name = 'Viola')
),
(
  (SELECT id FROM performances WHERE title ilike '%olympic%'),
  (SELECT id FROM musicians WHERE first = 'Matthew'),
  (SELECT id FROM parts where name = 'Cello')
);

INSERT INTO recordings (performance, movement, file) values
(
  (SELECT id FROM performances WHERE title ilike '%olympic%'),
  (SELECT id FROM movements WHERE 
    piece = (SELECT id FROM pieces WHERE catalog = 'Opus 131') AND
    title = 'Adagio ma non troppo e molto espressivo'),
  'ce3a4dc-1.flac'
),
(
  (SELECT id FROM performances WHERE title ilike '%olympic%'),
  (SELECT id FROM movements WHERE 
    piece = (SELECT id FROM pieces WHERE catalog = 'Opus 131') AND
    title = 'Allegro molto vivace'),
  'ce3a4dc-2.flac'
),
(
  (SELECT id FROM performances WHERE title ilike '%olympic%'),
  (SELECT id FROM movements WHERE 
    piece = (SELECT id FROM pieces WHERE catalog = 'Opus 131') AND
    title = 'Allegro moderato'),
  'ce3a4dc-3.flac'
),
(
  (SELECT id FROM performances WHERE title ilike '%olympic%'),
  (SELECT id FROM movements WHERE 
    piece = (SELECT id FROM pieces WHERE catalog = 'Opus 131') AND
    title = 'Andante ma non troppo e molto cantabile'),
  'ce3a4dc-4.flac'
),
(
  (SELECT id FROM performances WHERE title ilike '%olympic%'),
  (SELECT id FROM movements WHERE 
    piece = (SELECT id FROM pieces WHERE catalog = 'Opus 131') AND
    title = 'Presto'),
  'ce3a4dc-5.flac'
),
(
  (SELECT id FROM performances WHERE title ilike '%olympic%'),
  (SELECT id FROM movements WHERE 
    piece = (SELECT id FROM pieces WHERE catalog = 'Opus 131') AND
    title = 'Adagio quasi un poco andante'),
  'ce3a4dc-6.flac'
),
(
  (SELECT id FROM performances WHERE title ilike '%olympic%'),
  (SELECT id FROM movements WHERE 
    piece = (SELECT id FROM pieces WHERE catalog = 'Opus 131') AND
    title = 'Allegro'),
  'ce3a4dc-7.flac'
);


-- Beethoven Opus 131 Oly ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

INSERT INTO performances (piece, title) values
(
  (SELECT id FROM pieces WHERE catalog = 'Opus 131'), 
  'Santa Fe Chamber Music Festival 2007'
);

INSERT INTO performers (performance, musician, part) values
(
  (SELECT id FROM performances WHERE title ilike '%santa%'),
  (SELECT id FROM musicians WHERE first = 'Tien-Hsin'),
  (SELECT id FROM parts where name = 'Violin 1')
),
(
  (SELECT id FROM performances WHERE title ilike '%santa%'),
  (SELECT id FROM musicians WHERE first = 'Jessica'),
  (SELECT id FROM parts where name = 'Violin 2')
),
(
  (SELECT id FROM performances WHERE title ilike '%santa%'),
  (SELECT id FROM musicians WHERE first = 'Che-Yen'),
  (SELECT id FROM parts where name = 'Viola')
),
(
  (SELECT id FROM performances WHERE title ilike '%santa%'),
  (SELECT id FROM musicians WHERE first = 'Nicholas'),
  (SELECT id FROM parts where name = 'Cello')
);

INSERT INTO recordings (performance, movement, file) values
(
  (SELECT id FROM performances WHERE title ilike '%santa%'),
  (SELECT id FROM movements WHERE 
    piece = (SELECT id FROM pieces WHERE catalog = 'Opus 131') AND
    title = 'Adagio ma non troppo e molto espressivo'),
  'be3a4dc-1.flac'
),
(
  (SELECT id FROM performances WHERE title ilike '%santa%'),
  (SELECT id FROM movements WHERE 
    piece = (SELECT id FROM pieces WHERE catalog = 'Opus 131') AND
    title = 'Allegro molto vivace'),
  'be3a4dc-2.flac'
),
(
  (SELECT id FROM performances WHERE title ilike '%santa%'),
  (SELECT id FROM movements WHERE 
    piece = (SELECT id FROM pieces WHERE catalog = 'Opus 131') AND
    title = 'Allegro moderato'),
  'be3a4dc-3.flac'
),
(
  (SELECT id FROM performances WHERE title ilike '%santa%'),
  (SELECT id FROM movements WHERE 
    piece = (SELECT id FROM pieces WHERE catalog = 'Opus 131') AND
    title = 'Andante ma non troppo e molto cantabile'),
  'be3a4dc-4.flac'
),
(
  (SELECT id FROM performances WHERE title ilike '%santa%'),
  (SELECT id FROM movements WHERE 
    piece = (SELECT id FROM pieces WHERE catalog = 'Opus 131') AND
    title = 'Presto'),
  'be3a4dc-5.flac'
),
(
  (SELECT id FROM performances WHERE title ilike '%santa%'),
  (SELECT id FROM movements WHERE 
    piece = (SELECT id FROM pieces WHERE catalog = 'Opus 131') AND
    title = 'Adagio quasi un poco andante'),
  'be3a4dc-6.flac'
),
(
  (SELECT id FROM performances WHERE title ilike '%santa%'),
  (SELECT id FROM movements WHERE 
    piece = (SELECT id FROM pieces WHERE catalog = 'Opus 131') AND
    title = 'Allegro'),
  'be3a4dc-7.flac'
);
