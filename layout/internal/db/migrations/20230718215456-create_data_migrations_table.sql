-- +migrate Up
CREATE TABLE `data_migrations` (
   `id` varchar(255) NOT NULL,
   `applied_at` datetime NOT NULL,
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- +migrate Down
DROP TABLE `data_migrations`;
