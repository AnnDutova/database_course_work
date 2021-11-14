CREATE TABLE users(
    id int not null primary key AUTO_INCREMENT,
    email varchar(30) not null unique,
    encrypted_password varchar(30) not null
)