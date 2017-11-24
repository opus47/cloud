-- Opus47 seed data

INSERT INTO composers (first, middle, last) values
('Johannes', NULL, 'Brahms'),
('Antonin', 'Leopold', 'Dvořák'),
('Wolfgang', 'Amadeus', 'Mozart')
;

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

INSERT INTO parts (name) values
('Violin'),
('Violin 1'),
('Violin 2'),
('Violin 3'),
('Violin 4'),
('Viola'),
('Viola 1'),
('Viola 2'),
('Viola 3'),
('Viola 4'),
('Cello'),
('Cello 1'),
('Cello 2'),
('Piano'),
('Piano 1'),
('Piano 2'),
('Flute'),
('Harp'),
('Clarinet'),
('Oboe'),
('Horn'),
('Bass')
;

CREATE FUNCTION new_piece(
  pcomposer_first text, 
  pcomposer_last text, 
  ptitle text, 
  pkey text, 
  pnumber integer, 
  pmovements text[]
) RETURNS void AS $$
DECLARE
  n_movements integer := array_length(pmovements, 1);
  composer_id UUID;
  key_id UUID;
  piece_id UUID;
BEGIN

  SELECT id INTO composer_id FROM composers 
  WHERE first = pcomposer_first AND last = pcomposer_last
  ;

  SELECT id INTO key_ID FROM keys where name = pkey
  ;

  INSERT INTO pieces (composer, title, key, number) values
  (composer_id, ptitle, key_id, pnumber) RETURNING id INTO piece_id
  ;

  FOR i IN 1..n_movements
  LOOP
    INSERT INTO movements (piece, title, number) values
    (piece_id, pmovements[i], i);
  END LOOP;
END;
$$ LANGUAGE plpgsql;

SELECT new_piece(
  'Johannes', 'Brahms',
  'Piano Quartet',
  'A Major',
  2,
  ARRAY[
  'Allegro non troppo',
  'Poco adagio',
  'Scherzo. Poco Allegro -Trio',
  'Finale. Allegro'
  ]
);


