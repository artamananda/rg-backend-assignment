CREATE TABLE persons(
    id INT PRIMARY KEY,
    nik VARCHAR(255) UNIQUE NOT NULL,
    fullname VARCHAR(255) NOT NULL,
    gender VARCHAR(50) NOT NULL,
    birth_date DATE NOT NULL,
    is_married BOOLEAN,
    height FLOAT,
    weight FLOAT,
    address TEXT
);