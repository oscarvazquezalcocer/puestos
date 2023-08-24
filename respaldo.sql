PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE `funcions` (`id` integer,`created_at` datetime,`updated_at` datetime,`deleted_at` datetime,`nombre` text UNIQUE,`clave` text UNIQUE,PRIMARY KEY (`id`));
INSERT INTO funcions VALUES(1,'2023-08-24 14:39:46.463112065-06:00','2023-08-24 14:39:46.463112065-06:00',NULL,'INGENIERO EN SISTEMAS','CF12027');
INSERT INTO funcions VALUES(2,'2023-08-24 14:40:38.767673015-06:00','2023-08-24 14:40:38.767673015-06:00',NULL,'JEFE DE DEPARTAMENTO','ITS0004');
INSERT INTO funcions VALUES(3,'2023-08-24 14:40:48.667156782-06:00','2023-08-24 14:40:48.667156782-06:00',NULL,'JEFE DE DIVISION','ITS0003');
INSERT INTO funcions VALUES(4,'2023-08-24 14:41:01.164326889-06:00','2023-08-24 14:41:01.164326889-06:00',NULL,'DIRECTOR GENERAL','ITS0001');
INSERT INTO funcions VALUES(5,'2023-08-24 14:41:19.2325374-06:00','2023-08-24 14:41:19.2325374-06:00',NULL,'TECNICO ESPECIALIZADO','CF33118');
INSERT INTO funcions VALUES(6,'2023-08-24 14:41:31.17296691-06:00','2023-08-24 14:41:31.17296691-06:00',NULL,'PSICOLOGO','P16004');
INSERT INTO funcions VALUES(7,'2023-08-24 14:42:06.530864849-06:00','2023-08-24 14:42:06.530864849-06:00',NULL,'ANALISTA ESPECIALIZADO','P01002');
INSERT INTO funcions VALUES(8,'2023-08-24 14:42:48.147287861-06:00','2023-08-24 14:42:48.147287861-06:00',NULL,'JEFE DE OFICINA','A01001');
INSERT INTO funcions VALUES(9,'2023-08-24 14:43:04.528866397-06:00','2023-08-24 14:43:04.528866397-06:00',NULL,'SECRETARIA DE DIRECTOR GENERAL','CF53455');
INSERT INTO funcions VALUES(10,'2023-08-24 14:43:16.290776124-06:00','2023-08-24 14:43:16.290776124-06:00',NULL,'SECRETARIO DE DIRECTOR','CF34006');
INSERT INTO funcions VALUES(11,'2023-08-24 14:43:26.705135609-06:00','2023-08-24 14:43:26.705135609-06:00',NULL,'ANALISTA TECNICO','P01001');
INSERT INTO funcions VALUES(12,'2023-08-24 14:43:37.630277502-06:00','2023-08-24 14:43:37.630277502-06:00',NULL,'SECRETARIA DE SUBDIRECTOR','CF34280');
INSERT INTO funcions VALUES(13,'2023-08-24 14:43:53.365151516-06:00','2023-08-24 14:43:53.365151516-06:00',NULL,'SECRETARIA DE JEFE DEPARTAMENTO','CF34004');
INSERT INTO funcions VALUES(14,'2023-08-24 14:44:13.196289853-06:00','2023-08-24 14:44:13.196289853-06:00',NULL,'BIBLIOTECARIO','T05003');
INSERT INTO funcions VALUES(15,'2023-08-24 14:44:26.917204192-06:00','2023-08-24 14:44:26.917204192-06:00',NULL,'TECNICO EN MANTENIMIENTO','S08011');
INSERT INTO funcions VALUES(16,'2023-08-24 14:44:34.956240997-06:00','2023-08-24 14:44:34.956240997-06:00',NULL,'INTENDENTE','S06002');
INSERT INTO funcions VALUES(17,'2023-08-24 14:44:46.379896817-06:00','2023-08-24 14:44:46.379896817-06:00',NULL,'VIGILANTE','S14001');
INSERT INTO funcions VALUES(18,'2023-08-24 14:45:16.598074217-06:00','2023-08-24 14:45:16.598074217-06:00',NULL,'CHOFER DE DIRECTOR','S13008');
INSERT INTO funcions VALUES(19,'2023-08-24 14:45:25.531688647-06:00','2023-08-24 14:45:25.531688647-06:00',NULL,'PROFESOR ASIGNATURA "A"','E13001');
INSERT INTO funcions VALUES(20,'2023-08-24 14:45:34.55279251-06:00','2023-08-24 14:45:34.55279251-06:00',NULL,'PROFESOR ASIGNATURA "B"','E13002');
INSERT INTO funcions VALUES(21,'2023-08-24 14:45:44.164789071-06:00','2023-08-24 14:45:44.164789071-06:00',NULL,'PROFESOR ASOCIADO "A"','E13010');
INSERT INTO funcions VALUES(22,'2023-08-24 14:46:13.439143943-06:00','2023-08-24 14:46:13.439143943-06:00',NULL,'PROFESOR ASOCIADO "B"','E13011');
INSERT INTO funcions VALUES(23,'2023-08-24 14:46:22.109261693-06:00','2023-08-24 14:46:22.109261693-06:00',NULL,'PROFESOR ASOCIADO "C"','E13012');
INSERT INTO funcions VALUES(24,'2023-08-24 14:46:30.721583337-06:00','2023-08-24 14:46:30.721583337-06:00',NULL,'PROFESOR TITULAR "A"','E13013');
INSERT INTO funcions VALUES(25,'2023-08-24 14:51:16.333171749-06:00','2023-08-24 14:51:16.333171749-06:00',NULL,'CAPTURISTA','T06027');
INSERT INTO funcions VALUES(26,'2023-08-24 14:51:25.919620403-06:00','2023-08-24 14:51:25.919620403-06:00',NULL,'LABORATORISTA','T16005');
INSERT INTO funcions VALUES(27,'2023-08-24 14:53:46.71891857-06:00','2023-08-24 14:53:46.71891857-06:00',NULL,'SUBDIRECTOR DE AREA','ITS0002');
COMMIT;