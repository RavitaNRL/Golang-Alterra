-- SOAL :

-- Terdapat sebuah query pada SQL yaitu `SELECT * FROM users;`

-- Dengan tujuan yang sama, tuliskan dalam bentuk perintah:

-- 1. Redis
-- 2. Neo4j
-- 3. Cassandra


-- Jawaban :

-- 1. Redis
GET FROM users;

-- 2. Neo4j
MATCH (r:User) RETURN r;

-- 3. Cassandra
SELECT * FROM users;