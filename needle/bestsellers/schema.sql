CREATE TABLE IF NOT EXISTS BestSellers (
   ID          INT GENERATED ALWAYS AS IDENTITY,
   Name        VARCHAR(255) NOT NULL,
   FirstBSTime TIMESTAMP DEFAULT NOW(),
   PRIMARY KEY(ID)
);