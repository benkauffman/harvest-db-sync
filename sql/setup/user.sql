SET @pass := '';
DROP USER IF EXISTS 'harvest'@'%';
CREATE USER IF NOT EXISTS 'harvest'@'%' IDENTIFIED BY @pass;
GRANT ALL ON `harvest%`.* TO 'harvest'@'%';
