CREATE DATABASE AMACHAY
GO

USE AMACHAY
GO

CREATE TABLE PERSON
(
    Id       INT          NOT NULL IDENTITY,
    Name     VARCHAR(100) NOT NULL,
    LastName VARCHAR(100) NOT NULL,
    Cel      VARCHAR(20),
    Phone    VARCHAR(20),
    Address  VARCHAR(200),
    Dni      VARCHAR(11),
    Mail     VARCHAR(100),
    CONSTRAINT Pk_Person
        PRIMARY KEY (Id)
)
GO

CREATE TABLE CLIENT
(
    Id       INT NOT NULL IDENTITY,
    IdPerson INT NOT NULL,
    Type     VARCHAR(100),
    CONSTRAINT Pk_Client
        PRIMARY KEY (Id),
    CONSTRAINT Fk_Client_Person
        FOREIGN KEY (IdPerson) REFERENCES PERSON (Id)
)
GO

CREATE TABLE WAREHOUSE
(
    Id      INT          NOT NULL IDENTITY,
    Name    VARCHAR(100) NOT NULL,
    Address VARCHAR(200) NOT NULL,
    State   bit          NOT NULL DEFAULT 1,
    CONSTRAINT Pk_WareHouse
        PRIMARY KEY (Id)
)
GO

CREATE TABLE USERS
(
    Id          INT         NOT NULL IDENTITY,
    IdPerson    INT         NOT NULL,
    IdWarehouse INT         NULL,
    UserName    VARCHAR(20) NOT NULL,
    Password    TEXT        NOT NULL,
    Rol         VARCHAR(100),
    CONSTRAINT Pk_User
        PRIMARY KEY (Id),
    CONSTRAINT Fk_User_Person
        FOREIGN KEY (IdPerson) REFERENCES PERSON (Id),
    FOREIGN KEY (IdWarehouse) REFERENCES WAREHOUSE (Id)
)
GO

CREATE TABLE CATEGORY
(
    Id   INT          NOT NULL IDENTITY,
    Name VARCHAR(100) NOT NULL,
    CONSTRAINT Pk_Category
        PRIMARY KEY (Id)
)
GO


CREATE TABLE PRODUCT
(
    Id          INT            NOT NULL IDENTITY,
    IdCategory  INT            NOT NULL,
    Name        VARCHAR(100)   NOT NULL,
    Description VARCHAR(200)   NOT NULL,
    Price       NUMERIC(15, 4) NOT NULL,
    Stock       NUMERIC(15, 4) NOT NULL,
    Perishable  BIT            NULL,
    CONSTRAINT Pk_Product
        PRIMARY KEY (Id),
    CONSTRAINT Fk_Product_Category
        FOREIGN KEY (IdCategory) REFERENCES CATEGORY (Id),
)
GO

CREATE TABLE MEASURE
(
    Id   INT          NOT NULL IDENTITY,
    Name VARCHAR(100) NOT NULL,
    CONSTRAINT Pk_Measure
        PRIMARY KEY (Id),
)
GO

CREATE TABLE PRODUCT_MEASURE
(
    Id        INT NOT NULL IDENTITY,
    IdProduct INT NOT NULL,
    IdMeasure INT NOT NULL,
    Unity     INT NOT NULL,
    MinAlert  INT NULL,
    CONSTRAINT Pk_Product_Measure
        PRIMARY KEY (Id),
    CONSTRAINT Fk_Product_MeasureProduct
        FOREIGN KEY (IdProduct) REFERENCES PRODUCT (Id),
    CONSTRAINT Fk_Product_MeasureMeasure
        FOREIGN KEY (IdMeasure) REFERENCES MEASURE (Id),
)
GO

CREATE TABLE BUSINESS
(
    Id      INT          NOT NULL IDENTITY,
    Name    VARCHAR(150) NOT NULL,
    RUC     VARCHAR(50)  NOT NULL,
    Address VARCHAR(200) NOT NULL,
    Cel     VARCHAR(20)  NULL,
    Phone   VARCHAR(20)  NULL,
    Mail    VARCHAR(50)  NOT NULL,
    CONSTRAINT Pk_Business
        PRIMARY KEY (Id),
)
GO

CREATE TABLE PROVIDER
(
    Id         INT         NOT NULL IDENTITY,
    IdBusiness INT         NOT NULL,
    Type       VARCHAR(50) NOT NULL,
    CONSTRAINT Pk_Provider
        PRIMARY KEY (Id),
    CONSTRAINT Fk_ProviderBusiness
        FOREIGN KEY (IdBusiness) REFERENCES BUSINESS (Id),
)
GO

CREATE TABLE INVOICE
(
    Id   INT          NOT NULL IDENTITY,
    Name VARCHAR(250) NOT NULL,
    Code VARCHAR(50)  NOT NULL,
    Date DATETIME     NOT NULL,
    CONSTRAINT Pk_Provider
        PRIMARY KEY (Id),
)
GO

CREATE TABLE MOVEMENT
(
    Id          INT            NOT NULL IDENTITY,
    IdProduct   INT            NOT NULL,
    IdWareHouse INT            NOT NULL,
    DateTime    DATETIME       NOT NULL,
    Quantity    NUMERIC(15, 2) NOT NULL,
    Type        VARCHAR(100),
    IdUser      INT            NOT NULL,
    IdClient    INT            NULL,
    IdProvider  INT            NULL,
    LogDateTime DATETIME       NULL,

    Lot         VARCHAR(50)    NULL,
    DueDate     DATETIME       NULL,
    State       BIT            NULL,
    IdInvoice   INT            NULL,
    CONSTRAINT Pk_Movement
        PRIMARY KEY (Id),
    CONSTRAINT Fk_Movement_Product
        FOREIGN KEY (IdProduct) REFERENCES PRODUCT (Id),
    CONSTRAINT Fk_Movement_WareHouse
        FOREIGN KEY (IdWareHouse) REFERENCES WAREHOUSE (Id),
    CONSTRAINT Fk_Movement_User
        FOREIGN KEY (IdUser) REFERENCES USERS (Id),
    CONSTRAINT Fk_Movement_Client
        FOREIGN KEY (IdClient) REFERENCES CLIENT (Id),
    CONSTRAINT Fk_Movement_Provider
        FOREIGN KEY (IdProvider) REFERENCES PROVIDER (Id),
    CONSTRAINT Fk_Movement_Invoice
        FOREIGN KEY (IdInvoice) REFERENCES INVOICE (Id)
)
GO
