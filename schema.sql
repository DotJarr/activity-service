CREATE TABLE Users (
    User_Id SERIAL  PRIMARY KEY  NOT NULL,
    First_Name Varchar(100) NOT NULL ,
    Last_Name Varchar(100) NOT NULL,
    Email  Varchar(100) NOT NULL,
    Contact Int NOT NULL,
    Hash_Password Varchar(200) NOT NULL,
    Gender  Varchar(10),
    Created_At TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Colors (
    Category_Id SERIAL  PRIMARY KEY  NOT NULL,
    Color Varchar(100) NOT NULL ,
    Category Varchar(100) NOT NULL, 
     UNIQUE(Color)  
);

CREATE TABLE Activities (
    Activities_Id SERIAL  PRIMARY KEY  NOT NULL,
    Category Varchar(100) NOT NULL ,
    Category_Id INT NOT NULL,
    Activity  Varchar(100) NOT NULL,
    User_Id Int NOT NULL,
    Color Varchar(100) NOT NULL,
    Status Varchar(100),
    Created_At TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (Category_Id)
      REFERENCES Colors (Category_Id),
     FOREIGN KEY (User_Id)
      REFERENCES Users (User_Id)
);

INSERT INTO users
values (301,'arvind','pj', 'apj@gmail.com', 9934678, 'password', 'M');

INSERT INTO Colors
values (201, 'blue','going out');
INSERT INTO Activities
values (101,'going out', 201, 'go to beach', 301, 'blue','pending');
INSERT INTO Activities
values (102,'going out', 201, 'go to movies', 301, 'blue','pending');


CREATE TABLE account_roles (
  user_id INT NOT NULL,
  role_id INT NOT NULL,
  grant_date TIMESTAMP,
  PRIMARY KEY (user_id, role_id),
  FOREIGN KEY (role_id)
      REFERENCES roles (role_id),
  FOREIGN KEY (user_id)
      REFERENCES accounts (user_id)
);

-- // db.Joins("Company", db.Where(&Company{Alive: true})).Find(&users)
-- 		// db.Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "411111111111").Find(&user)
-- 		_ = db.Joins("JOIN users ON users.user_id = activities.user_id").Where("users.first_name=? ", "arvind").Find(&activities)
-- 		// _ = db.Find(&activities)