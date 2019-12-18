CREATE TABLE Persons (
    PersonID bigserial NOT NULL,
    LastName varchar(255),
    FirstName varchar(255),
    Address varchar(255),
    City varchar(255)
);

INSERT INTO Persons (LastName, FirstName, Address, City)
VALUES ('Cardinal','Tom B. Erichsen','Skagen 21','Stavanger');

INSERT INTO Persons (LastName, FirstName, Address, City)
VALUES ('Sam','sung B','Skagen 21','Stavanger');