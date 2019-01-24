CREATE TABLE IF NOT EXISTS plates (
    plate_id INT AUTO_INCREMENT,
    plate VARCHAR(10) NOT NULL,
    sticker VARCHAR(255) NOT NULL,
    counts INTEGER NOT NULL,
    timestamp TIMESTAMP,
    PRIMARY KEY (plate_id)
)  ENGINE=INNODB;