CREATE TABLE songs(
   song_id serial PRIMARY KEY,
   group_name VARCHAR(150) NOT NULL,
   song_name VARCHAR(150) NOT NULL,
   release_date DATE NOT NULL,
   lyrics VARCHAR(500) NOT NULL,
   link VARCHAR(100) NOT NULL
);