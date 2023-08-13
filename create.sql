DROP DATABASE powerlifterplpus;
CREATE DATABASE powerlifterplpus;
USE powerlifterplpus;

CREATE TABLE lifts (
  id          INT AUTO_INCREMENT NOT NULL,
  user_id     INT NOT NULL,
  name        VARCHAR(255) NOT NULL,
  lift_date   DATE NOT NULL,
  weight      INT NOT NULL DEFAULT '0',
  reps        INT NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
);

CREATE TABLE users (
  id          INT AUTO_INCREMENT NOT NULL,
  name        VARCHAR(255) NOT NULL,
  email        VARCHAR(255) NOT NULL,
  password        VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);
