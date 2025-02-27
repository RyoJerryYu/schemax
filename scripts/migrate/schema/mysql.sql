-- table with manual primary key
-- generate insert only (no update, save, upsert, delete)
CREATE TABLE a_manual_table (
  a_text VARCHAR(255)
) ENGINE=InnoDB;

-- table with sequence
CREATE TABLE a_sequence (
  a_seq INTEGER AUTO_INCREMENT,
  CONSTRAINT a_sequence_pkey PRIMARY KEY (a_seq)
) ENGINE=InnoDB;

CREATE TABLE a_sequence_multi (
  a_seq INTEGER AUTO_INCREMENT,
  a_text VARCHAR(255),
  CONSTRAINT a_sequence_multi_pkey PRIMARY KEY (a_seq)
) ENGINE=InnoDB;

-- table with primary key
CREATE TABLE a_primary (
  a_key INTEGER PRIMARY KEY
) ENGINE=InnoDB;

CREATE TABLE a_primary_multi (
  a_key INTEGER PRIMARY KEY,
  a_text VARCHAR(255)
) ENGINE=InnoDB;

-- table with composite primary key
CREATE TABLE a_primary_composite (
  a_key1 INTEGER,
  a_key2 INTEGER,
  PRIMARY KEY (a_key1, a_key2)
) ENGINE=InnoDB;

-- table with index
CREATE TABLE a_index (
  a_key INTEGER
) ENGINE=InnoDB;

CREATE INDEX a_index_idx ON a_index (a_key);

-- table with composite index
CREATE TABLE a_index_composite (
  a_key1 INTEGER,
  a_key2 INTEGER
) ENGINE=InnoDB;

CREATE INDEX a_index_composite_idx ON a_index_composite (a_key1, a_key2);

-- table with unique index
CREATE TABLE a_unique_index (
  a_key INTEGER UNIQUE
) ENGINE=InnoDB;

-- table with composite unique index
CREATE TABLE a_unique_index_composite (
  a_key1 INTEGER,
  a_key2 INTEGER,
  UNIQUE (a_key1, a_key2)
) ENGINE=InnoDB;

/*

bigint
binary
bit
blob
char
date
datetime
dec
fixed
decimal
double precision
float
int
integer
json
longblob
longtext
mediumblob
mediumint
mediumtext
mediumtext
numeric
real
set
smallint
text
time
timestamp
tinyblob
tinyint
tinytext
varbinary
varchar
year

*/

-- table with all field types and all nullable field types
CREATE TABLE a_bit_of_everything (
  a_bigint BIGINT NOT NULL,
  a_bigint_nullable BIGINT,
  a_binary BINARY NOT NULL,
  a_binary_nullable BINARY,
  a_bit BIT NOT NULL,
  a_bit_nullable BIT,
  a_blob BLOB NOT NULL,
  a_blob_nullable BLOB,
  a_bool BOOL NOT NULL,
  a_bool_nullable BOOL,
  a_char CHAR NOT NULL,
  a_char_nullable CHAR,
  a_date DATE NOT NULL,
  a_date_nullable DATE,
  a_datetime DATETIME NOT NULL,
  a_datetime_nullable DATETIME,
  a_dec DEC NOT NULL,
  a_dec_nullable DEC,
  a_fixed FIXED NOT NULL,
  a_fixed_nullable FIXED,
  a_decimal DECIMAL NOT NULL,
  a_decimal_nullable DECIMAL,
  a_double_precision DOUBLE PRECISION NOT NULL,
  a_double_precision_nullable DOUBLE PRECISION,
  a_float FLOAT NOT NULL,
  a_float_nullable FLOAT,
  a_int INT NOT NULL,
  a_int_nullable INT,
  a_integer INTEGER NOT NULL,
  a_integer_nullable INTEGER,
  a_json JSON NOT NULL,
  a_json_nullable JSON,
  a_longblob LONGBLOB NOT NULL,
  a_longblob_nullable LONGBLOB,
  a_longtext LONGTEXT NOT NULL,
  a_longtext_nullable LONGTEXT,
  a_mediumblob MEDIUMBLOB NOT NULL,
  a_mediumblob_nullable MEDIUMBLOB,
  a_mediumint MEDIUMINT NOT NULL,
  a_mediumint_nullable MEDIUMINT,
  a_mediumtext MEDIUMTEXT NOT NULL,
  a_mediumtext_nullable MEDIUMTEXT,
  a_numeric NUMERIC NOT NULL,
  a_numeric_nullable NUMERIC,
  a_real REAL NOT NULL,
  a_real_nullable REAL,
  a_set SET('ONE', 'TWO') NOT NULL,
  a_set_nullable SET('ONE', 'TWO'),
  a_smallint SMALLINT NOT NULL,
  a_smallint_nullable SMALLINT,
  a_text TEXT NOT NULL,
  a_text_nullable TEXT,
  a_time TIME NOT NULL,
  a_time_nullable TIME,
  a_timestamp TIMESTAMP NOT NULL,
  a_timestamp_nullable TIMESTAMP,
  a_tinyblob TINYBLOB NOT NULL,
  a_tinyblob_nullable TINYBLOB,
  a_tinyint TINYINT NOT NULL,
  a_tinyint_nullable TINYINT,
  a_tinytext TINYTEXT NOT NULL,
  a_tinytext_nullable TINYTEXT,
  a_varbinary VARBINARY(255) NOT NULL,
  a_varbinary_nullable VARBINARY(255),
  a_varchar VARCHAR(255) NOT NULL,
  a_varchar_nullable VARCHAR(255),
  a_year YEAR NOT NULL,
  a_year_nullable YEAR
) ENGINE=InnoDB;
