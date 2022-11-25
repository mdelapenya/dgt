CREATE TABLE IF NOT EXISTS stickers (
    sticker_id INT AUTO_INCREMENT,
    description VARCHAR(255) NOT NULL,
    emoji VARCHAR(255) NOT NULL,
    PRIMARY KEY (sticker_id)
)  ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci;

INSERT INTO stickers (description, emoji) VALUES ('Sin distintivo', '⚪️');
INSERT INTO stickers (description, emoji) VALUES ('Etiqueta Ambiental B Amarilla', '🟡');
INSERT INTO stickers (description, emoji) VALUES ('Etiqueta Ambiental C Verde', '🟢');
INSERT INTO stickers (description, emoji) VALUES ('Etiqueta Ambiental Eco', '🟣');
INSERT INTO stickers (description, emoji) VALUES ('Etiqueta Ambiental 0', '🔵');
INSERT INTO stickers (description, emoji) VALUES ('No se ha encontrado ningún resultado para la matrícula introducida', '❌');

CREATE TABLE IF NOT EXISTS plates (
    plate_id INT AUTO_INCREMENT,
    plate VARCHAR(10) NOT NULL,
    sticker_id INT NOT NULL,
    counts INTEGER NOT NULL,
    timestamp TIMESTAMP,
    PRIMARY KEY (plate_id),
    FOREIGN KEY (sticker_id) REFERENCES stickers(sticker_id)
)  ENGINE=INNODB;

CREATE INDEX idx_plates ON plates (plate);
