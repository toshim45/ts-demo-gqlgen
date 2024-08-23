CREATE TABLE `things` (
  `id` char(36) NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000',
  `name` varchar(16) DEFAULT NULL,
  `amount` smallint DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

INSERT INTO `things` (`id`, `name`, `amount`) VALUES
('01917255-aabd-7908-a5e0-f4912fc98f03', '838a6609d657', 11),
('01917256-e6f2-76ca-8c08-db07e9ff72e8', 'b147d14f913a', 22);