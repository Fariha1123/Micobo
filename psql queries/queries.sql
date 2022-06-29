CREATE DOMAIN GENDER CHAR(1)
    CHECK (value IN ( 'F' , 'M' ) )
    ;
    
CREATE DOMAIN YESNO CHAR(1)
    CHECK (value IN ('Y', 'N'))
    ;

CREATE TABLE employees (
    id INT GENERATED ALWAYS AS IDENTITY,
    fullname VARCHAR(25) NOT NULL,
    birthday DATE NOT NULL,
    gender GENDER NOT NULL,
    eventId INT,
    accomodation YESNO,
    constraint fk_event_employee
     foreign key (eventId) 
     REFERENCES events (id)
     ON DELETE SET NULL
);