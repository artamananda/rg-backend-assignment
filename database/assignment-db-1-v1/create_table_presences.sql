CREATE TABLE presences(
    id INT NOT NULL,
    user_id INT NOT NULL,
    presence_date DATE NOT NULL,
    status VARCHAR(50) NOT NULL,
    location VARCHAR(255),
    description VARCHAR(255),
    image_presence VARCHAR(255),
    image_location VARCHAR(255)
);