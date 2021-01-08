-- Тестовые данные для таблицы студии
INSERT INTO "studio" (id,title) VALUES (1,'Mauris LLP'),(2,'Senectus Et Netus LLC'),(3,'Id Libero Donec Associates'),(4,'Fringilla Ornare Company'),(5,'Morbi Metus Corporation'),(6,'Suscipit Est Institute'),(7,'Purus Institute'),(8,'Nunc Interdum Feugiat LLC'),(9,'Sollicitudin Corp.'),(10,'Tincidunt Vehicula Limited');
INSERT INTO "studio" (id,title) VALUES (11,'Vitae Diam Company'),(12,'Duis Mi Ltd'),(13,'Arcu Associates'),(14,'Donec Est LLP'),(15,'Eu Lacus Industries'),(16,'Varius Et Euismod Limited'),(17,'Scelerisque Scelerisque Dui LLP'),(18,'Rutrum PC'),(19,'Pede Foundation'),(20,'Nunc Commodo Auctor Incorporated');
INSERT INTO "studio" (id,title) VALUES (21,'Non Ante Bibendum LLP'),(22,'Magna Duis Industries'),(23,'Nunc Lectus Pede Corp.'),(24,'Nec Urna Et Industries'),(25,'Velit Justo Nec Limited'),(26,'Nunc LLP'),(27,'Lacus Associates'),(28,'Nisl Nulla Foundation'),(29,'Fusce PC'),(30,'At Inc.');
INSERT INTO "studio" (id,title) VALUES (31,'Consectetuer Adipiscing Limited'),(32,'Sapien Company'),(33,'Ipsum Dolor Sit LLC'),(34,'Proin LLC'),(35,'Scelerisque Corp.'),(36,'Massa Suspendisse Corp.'),(37,'Eget Limited'),(38,'Cras Eget Nisi Company'),(39,'Metus PC'),(40,'Elit Foundation');
INSERT INTO "studio" (id,title) VALUES (41,'Vivamus Limited'),(42,'Praesent Industries'),(43,'Egestas Foundation'),(44,'Imperdiet Ornare Ltd'),(45,'Fusce Feugiat Foundation'),(46,'A Nunc In Corp.'),(47,'Rhoncus Consulting'),(48,'Metus In Nec Corporation'),(49,'Turpis Egestas Foundation'),(50,'Dui Quis Accumsan Limited');
INSERT INTO "studio" (id,title) VALUES (51,'Ultricies Dignissim PC'),(52,'Tempus Non Lacinia Corp.'),(53,'Eget Industries'),(54,'Molestie Institute'),(55,'Aliquet Inc.'),(56,'Adipiscing Elit Industries'),(57,'Eu Consulting'),(58,'Ut Institute'),(59,'Tempus Corporation'),(60,'Aliquam Company');
INSERT INTO "studio" (id,title) VALUES (61,'Morbi Foundation'),(62,'A Nunc Limited'),(63,'Lacus LLC'),(64,'Nam Interdum Enim LLC'),(65,'Et Magnis Dis Foundation'),(66,'Nec Euismod In Company'),(67,'Erat Corporation'),(68,'Sem Molestie LLC'),(69,'In Molestie LLC'),(70,'Nulla Facilisi Corporation');
INSERT INTO "studio" (id,title) VALUES (71,'Quis Turpis LLP'),(72,'Cursus Vestibulum Associates'),(73,'Eget Magna Company'),(74,'Phasellus Ornare LLC'),(75,'Neque Sed PC'),(76,'Faucibus Inc.'),(77,'Praesent Eu Dui Corporation'),(78,'Ante Ipsum Associates'),(79,'Arcu Morbi Sit Foundation'),(80,'Donec Tempor Est Ltd');
INSERT INTO "studio" (id,title) VALUES (81,'Magna Ltd'),(82,'Sem Elit Pharetra Inc.'),(83,'At Pede Cras Limited'),(84,'Commodo Incorporated'),(85,'Nulla Tempor Augue Corp.'),(86,'Neque Industries'),(87,'Amet LLP'),(88,'Lorem Ipsum Institute'),(89,'Non Egestas A Consulting'),(90,'Bibendum Ullamcorper Limited');
INSERT INTO "studio" (id,title) VALUES (91,'Senectus Et Netus Company'),(92,'Donec Ltd'),(93,'Velit Egestas Lacinia Consulting'),(94,'Orci Ut Semper LLC'),(95,'Commodo Ipsum Incorporated'),(96,'Sit Amet Consectetuer Institute'),(97,'Magna Malesuada Vel Company'),(98,'Netus Et Institute'),(99,'In Condimentum LLP'),(100,'Natoque Industries');
ALTER SEQUENCE studio_id_seq RESTART WITH 101;


-- Тестовые данные для таблицы актеров
INSERT INTO "actor" (id,name,born) VALUES (1,'Burks, Austin U.','1992-04-19'),(2,'Pope, Jasmine Y.','1996-02-11'),(3,'Casey, Farrah E.','2009-12-11'),(4,'Gilliam, Hayden P.','2007-11-09'),(5,'Cleveland, Daquan T.','2000-10-13'),(6,'Atkinson, McKenzie I.','2003-07-18'),(7,'Dyer, Dara T.','1985-02-12'),(8,'Atkins, Roth B.','2018-12-21'),(9,'Holt, Yoshi U.','1998-06-24'),(10,'Singleton, Marcia T.','2018-05-22');
INSERT INTO "actor" (id,name,born) VALUES (11,'Alford, Tobias Y.','2005-05-21'),(12,'Pruitt, Allistair E.','1989-07-10'),(13,'Morales, Dominic K.','2006-05-23'),(14,'Terrell, Knox S.','1985-03-13'),(15,'Berry, Anne I.','2012-07-18'),(16,'Mccall, Hunter B.','1995-02-24'),(17,'Guerra, Bethany J.','2000-01-08'),(18,'Dodson, Elton Y.','2018-06-08'),(19,'Tyson, Zachary L.','2011-04-30'),(20,'Acosta, Wayne Y.','2002-07-24');
INSERT INTO "actor" (id,name,born) VALUES (21,'Morrison, Thomas A.','1990-02-22'),(22,'Campbell, Cleo T.','2002-07-22'),(23,'Rosario, Wayne N.','2017-05-14'),(24,'Holt, Howard V.','2000-05-09'),(25,'Riley, Jenette W.','2003-03-14'),(26,'Ball, Serina P.','2018-05-18'),(27,'Fuentes, Marny X.','2018-06-04'),(28,'York, Acton X.','1991-04-11'),(29,'Hansen, Montana I.','2017-02-13'),(30,'Walters, Nora O.','1995-08-11');
INSERT INTO "actor" (id,name,born) VALUES (31,'Lester, Ina K.','1988-07-20'),(32,'Oneill, Gray G.','2014-11-30'),(33,'Bradley, Amelia W.','1999-06-14'),(34,'Hurley, Neil X.','1988-11-20'),(35,'Mccormick, Garrison J.','2021-10-23'),(36,'Mccray, Brandon I.','1992-04-07'),(37,'Wynn, Ferris Z.','1996-11-17'),(38,'Chase, Camden H.','2003-01-16'),(39,'Bird, Jerome L.','2015-04-28'),(40,'Burch, Byron A.','1991-02-23');
INSERT INTO "actor" (id,name,born) VALUES (41,'Callahan, Baxter S.','1989-11-02'),(42,'Fowler, Brianna N.','2009-09-03'),(43,'Dotson, Cheyenne F.','1993-03-09'),(44,'Soto, Zelda J.','1992-05-03'),(45,'Faulkner, Cedric Q.','2002-08-01'),(46,'Ferguson, Lana S.','2007-04-24'),(47,'Lloyd, Iliana N.','1984-04-26'),(48,'Dodson, Thaddeus R.','2008-09-07'),(49,'Chen, Cathleen J.','1992-08-04'),(50,'Kelley, Keane C.','2021-03-10');
INSERT INTO "actor" (id,name,born) VALUES (51,'Maynard, Mia T.','2009-11-02'),(52,'Graham, Yardley Z.','2003-10-07'),(53,'Vazquez, Kato K.','1995-05-07'),(54,'Wiggins, Lawrence C.','2005-07-06'),(55,'Walton, Byron J.','2002-06-12'),(56,'Gilbert, Emerson B.','1996-09-04'),(57,'Fuentes, Ivy G.','1985-01-02'),(58,'Potter, Kay Y.','1996-06-10'),(59,'Moon, Jillian T.','1988-08-15'),(60,'Guerrero, Ivana U.','2001-10-11');
INSERT INTO "actor" (id,name,born) VALUES (61,'Marquez, Colette C.','1998-03-31'),(62,'Head, Neve R.','1994-09-05'),(63,'Valenzuela, Colorado H.','2008-01-29'),(64,'Coleman, Laura S.','1990-05-07'),(65,'Whitney, Phillip N.','1994-02-22'),(66,'Guerra, Ruth U.','2020-03-01'),(67,'Merrill, Colby I.','2001-06-27'),(68,'Berry, Xandra D.','1994-04-23'),(69,'Mays, Geraldine Z.','2010-07-08'),(70,'Espinoza, Kaye B.','2003-05-27');
INSERT INTO "actor" (id,name,born) VALUES (71,'Rojas, Levi E.','2003-03-31'),(72,'Talley, Rogan O.','1992-07-26'),(73,'Allen, Justin E.','2015-09-07'),(74,'Barton, Fitzgerald Y.','1992-08-21'),(75,'Tyson, Wallace A.','2017-09-17'),(76,'Hart, Noelle P.','1996-10-19'),(77,'Jacobs, Fredericka F.','2000-03-27'),(78,'Pollard, Fatima B.','2000-01-14'),(79,'Landry, Ingrid M.','1988-05-28'),(80,'Carr, Shelley Y.','2011-01-30');
INSERT INTO "actor" (id,name,born) VALUES (81,'Little, David O.','1987-04-21'),(82,'Carpenter, Caryn E.','1992-03-10'),(83,'Carey, Xena G.','1985-05-27'),(84,'Mckenzie, Melyssa X.','2010-03-05'),(85,'Yates, Lani N.','2019-01-11'),(86,'Owens, Bryar Z.','2002-12-16'),(87,'Mills, Adria J.','2020-07-01'),(88,'Russell, Ulla O.','2018-12-27'),(89,'Kirby, Tanek F.','1996-09-10'),(90,'Nash, Zephania F.','1996-04-12');
INSERT INTO "actor" (id,name,born) VALUES (91,'Clark, Forrest F.','1988-07-31'),(92,'Baird, Avye S.','2015-05-22'),(93,'Shields, Eden Q.','1998-08-08'),(94,'Stafford, Merrill B.','2021-12-20'),(95,'Powers, Dane G.','2021-01-16'),(96,'Fuentes, Cleo D.','1986-12-13'),(97,'Ashley, Ariel Y.','2007-09-19'),(98,'Welch, Hasad L.','2019-05-13'),(99,'Mcmillan, Ciaran O.','2003-09-22'),(100,'Dickson, Lee U.','2021-07-29');
ALTER SEQUENCE actor_id_seq RESTART WITH 101;


-- Тестовые данные для таблицы режиссеров
INSERT INTO "producer" (id,name,born) VALUES (1,'Schroeder, Ray P.','2011-03-31'),(2,'Matthews, Cruz T.','1992-03-25'),(3,'Patrick, Mohammad P.','1991-04-01'),(4,'Welch, Barry L.','2003-03-14'),(5,'Frye, Ursa V.','1985-06-27'),(6,'Warner, Sasha R.','2021-04-20'),(7,'Gardner, Josiah N.','2016-11-22'),(8,'Harrell, Ivan Y.','2009-07-01'),(9,'Hawkins, Urielle T.','2014-12-28'),(10,'Carson, Rachel B.','2009-06-12');
INSERT INTO "producer" (id,name,born) VALUES (11,'Henry, Mari H.','1995-07-14'),(12,'Baldwin, Kylan K.','2011-04-10'),(13,'Lancaster, Galvin J.','1992-01-13'),(14,'Rodgers, Hedwig M.','1993-03-08'),(15,'Wolf, Ciara D.','2014-01-07'),(16,'Abbott, Trevor I.','2003-05-22'),(17,'Cervantes, Plato Q.','1983-12-26'),(18,'Peters, Drew H.','1987-11-02'),(19,'Harper, Jana J.','2021-08-16'),(20,'Beasley, Quemby Y.','2006-10-08');
INSERT INTO "producer" (id,name,born) VALUES (21,'Gutierrez, Abra J.','1983-01-01'),(22,'Pitts, Melodie K.','1994-08-24'),(23,'Levy, Beatrice O.','2005-12-29'),(24,'Arnold, Willa D.','2009-12-31'),(25,'Martinez, Daphne X.','2004-08-19'),(26,'Mcclain, Declan A.','2006-07-19'),(27,'Ross, Kibo Z.','2012-02-26'),(28,'Mathews, Hiroko V.','2017-08-02'),(29,'Welch, Cullen Y.','2010-11-01'),(30,'Baxter, Zia Y.','2012-08-24');
INSERT INTO "producer" (id,name,born) VALUES (31,'Gamble, Bradley Y.','2010-03-04'),(32,'Gentry, Demetria J.','2007-07-29'),(33,'Castaneda, Xena S.','2013-10-04'),(34,'Duran, Zoe I.','1984-03-10'),(35,'Wallace, Rahim B.','2018-11-17'),(36,'Harper, Christine S.','2003-02-15'),(37,'Dickson, Cade Y.','2000-03-01'),(38,'Jefferson, Damian H.','2002-05-24'),(39,'Stevenson, Hillary H.','1989-02-01'),(40,'Mayo, Thane V.','1998-04-16');
INSERT INTO "producer" (id,name,born) VALUES (41,'Abbott, Stone H.','2009-07-14'),(42,'Hernandez, Carlos Z.','2002-08-29'),(43,'Guy, Leo G.','2000-10-23'),(44,'Bowers, Wynter H.','1991-04-18'),(45,'Bowman, Tyrone X.','1996-10-02'),(46,'Massey, Rana R.','2006-02-23'),(47,'Medina, Cairo L.','2003-02-08'),(48,'Gomez, Cailin N.','1993-08-01'),(49,'Gilmore, Skyler W.','1997-03-30'),(50,'Phelps, Erich Y.','1983-10-16');
INSERT INTO "producer" (id,name,born) VALUES (51,'Douglas, Hasad R.','1991-04-21'),(52,'Velasquez, David P.','2014-06-29'),(53,'Steele, Preston T.','1999-09-23'),(54,'Estes, Amos S.','2002-05-04'),(55,'Hood, Raymond N.','1998-11-08'),(56,'Malone, Clare I.','1987-04-25'),(57,'Lowe, Adele Z.','2018-07-22'),(58,'Fitzgerald, Nell C.','2011-09-24'),(59,'Lynn, Rylee F.','2013-11-10'),(60,'Hunt, Karleigh V.','1997-06-30');
INSERT INTO "producer" (id,name,born) VALUES (61,'Kane, Herrod W.','2001-04-14'),(62,'Oneal, Cecilia V.','1999-12-18'),(63,'Garner, Jamal W.','2008-05-10'),(64,'Lott, Vladimir U.','2006-04-27'),(65,'Sherman, Ralph L.','2021-01-15'),(66,'Whitaker, Chancellor O.','1986-12-18'),(67,'Jordan, Jana D.','2004-11-02'),(68,'Mcintyre, Linus N.','1996-08-30'),(69,'Curry, Sara N.','2018-02-14'),(70,'Chavez, Judith P.','2002-03-10');
INSERT INTO "producer" (id,name,born) VALUES (71,'Russo, Macaulay I.','2011-10-09'),(72,'Quinn, Dacey V.','2012-06-28'),(73,'Fitzgerald, Sloane J.','1983-11-16'),(74,'Williamson, Chaim P.','2017-04-19'),(75,'Humphrey, Merritt N.','1991-12-02'),(76,'Powers, Merrill L.','1992-07-02'),(77,'Tran, Xerxes C.','1996-07-11'),(78,'Moss, Jessamine C.','2013-08-08'),(79,'Sanchez, Yasir Y.','2009-03-05'),(80,'Mcleod, Inez P.','1988-07-13');
INSERT INTO "producer" (id,name,born) VALUES (81,'Rosario, Carolyn I.','1993-01-12'),(82,'Guy, Driscoll H.','2017-05-19'),(83,'Byrd, Asher O.','1984-04-18'),(84,'English, Skyler H.','1991-08-11'),(85,'Armstrong, Basil W.','1996-11-05'),(86,'Gardner, Summer N.','2017-11-13'),(87,'Rosa, Sarah R.','1995-02-02'),(88,'Jordan, Colt W.','1984-03-25'),(89,'Albert, Grace J.','2019-10-30'),(90,'Dominguez, Shelby D.','1986-05-16');
INSERT INTO "producer" (id,name,born) VALUES (91,'Norton, Denton Y.','2008-07-20'),(92,'Wilson, Kadeem R.','1983-08-23'),(93,'Day, Clark M.','2013-02-17'),(94,'Morales, Lacey O.','1991-07-23'),(95,'Garrett, Chanda B.','1995-05-20'),(96,'Mccormick, Jin T.','2012-02-16'),(97,'Fields, Chastity J.','1992-07-15'),(98,'Hodge, Cally O.','2020-04-10'),(99,'Bowers, Dexter T.','2019-12-02'),(100,'Good, Eve Y.','1991-03-10');
ALTER SEQUENCE producer_id_seq RESTART WITH 101;


-- Тестовые данные для таблицы фильмов
INSERT INTO "film" (id,title,release_date,box_office,rating,studio_id) VALUES (1,'interdum ligula eu','2015-07-25',430503320,'PG-13',65),(2,'Donec elementum,','2008-05-14',735608529,'PG-10',19),(3,'elit sed consequat','2017-01-14',500215181,'PG-13',10),(4,'Aliquam tincidunt, nunc','2004-02-23',711408877,'PG-13',39),(5,'eu,','2014-11-13',328580339,'PG-13',62),(6,'amet,','2009-01-02',941099367,'PG-18',54),(7,'lacus. Aliquam rutrum','1982-10-04',484533856,'PG-13',3),(8,'lorem. Donec','2003-08-15',107006971,'PG-10',94),(9,'odio sagittis semper.','1983-12-08',770396898,'PG-10',43),(10,'velit.','1994-08-18',820137022,'PG-18',40);
INSERT INTO "film" (id,title,release_date,box_office,rating,studio_id) VALUES (11,'parturient','1987-05-16',841879718,'PG-10',64),(12,'nec','1994-04-26',943965889,'PG-13',75),(13,'Quisque','2021-03-11',674001473,'PG-10',12),(14,'a odio','2008-12-18',877835531,'PG-18',53),(15,'magna. Phasellus','1999-05-30',162409923,'PG-13',42),(16,'tincidunt dui augue','1996-12-22',102493403,'PG-13',53),(17,'libero. Integer in','2006-06-25',381099531,'PG-10',4),(18,'consectetuer adipiscing','1985-05-07',771301118,'PG-13',67),(19,'lacus,','2012-08-06',584139527,'PG-13',16),(20,'lobortis quam a','2002-08-19',682923858,'PG-18',81);
INSERT INTO "film" (id,title,release_date,box_office,rating,studio_id) VALUES (21,'ullamcorper','2007-04-07',585067053,'PG-13',94),(22,'sit amet,','1985-10-13',688523712,'PG-13',81),(23,'vitae, erat.','2018-05-27',396297825,'PG-10',92),(24,'montes,','1987-03-31',272905832,'PG-13',91),(25,'libero nec ligula','2007-12-02',831918950,'PG-18',27),(26,'Aliquam','1992-05-19',801861516,'PG-13',20),(27,'Sed neque.','1990-11-27',777351455,'PG-13',45),(28,'vulputate eu,','1982-02-10',776025392,'PG-18',63),(29,'ac mattis semper,','2015-06-11',544811276,'PG-10',41),(30,'erat.','2016-10-17',589670889,'PG-18',23);
INSERT INTO "film" (id,title,release_date,box_office,rating,studio_id) VALUES (31,'nisi magna sed','2011-08-07',799289524,'PG-18',84),(32,'tempor erat','1996-08-31',306592604,'PG-13',45),(33,'Quisque','2006-10-09',588681835,'PG-10',50),(34,'litora torquent','1987-10-26',346147752,'PG-13',45),(35,'nibh.','1992-01-06',258001762,'PG-10',3),(36,'aliquam arcu.','1988-06-14',391404540,'PG-13',38),(37,'elit, dictum eu,','1995-01-14',33248832,'PG-10',24),(38,'vitae,','1996-04-23',113741098,'PG-10',81),(39,'elit sed consequat','2001-09-17',551525728,'PG-10',84),(40,'blandit','1989-08-10',834332021,'PG-13',82);
INSERT INTO "film" (id,title,release_date,box_office,rating,studio_id) VALUES (41,'dictum magna.','1993-05-16',517392852,'PG-10',42),(42,'eros. Nam','1996-07-13',874348089,'PG-10',20),(43,'iaculis aliquet diam.','1982-09-09',995096137,'PG-13',90),(44,'ut ipsum ac','1996-11-20',280329150,'PG-13',45),(45,'aliquet.','1986-05-07',832992193,'PG-18',17),(46,'Aliquam erat','2014-08-24',945339867,'PG-10',25),(47,'arcu. Nunc','2018-12-12',803156396,'PG-18',18),(48,'tempor arcu.','2000-12-02',487951160,'PG-13',25),(49,'risus.','2019-08-06',625763770,'PG-13',72),(50,'Cras interdum.','1997-05-12',570945153,'PG-10',71);
INSERT INTO "film" (id,title,release_date,box_office,rating,studio_id) VALUES (51,'et netus','2009-12-21',819425138,'PG-13',49),(52,'odio, auctor vitae,','2016-04-30',315426723,'PG-10',32),(53,'orci tincidunt adipiscing.','1986-05-13',976188188,'PG-13',61),(54,'pede,','2010-11-05',976794490,'PG-18',56),(55,'odio semper cursus.','1993-11-09',42659200,'PG-13',20),(56,'magna','1998-08-14',928937178,'PG-13',38),(57,'tincidunt. Donec vitae','2016-10-19',677319745,'PG-18',8),(58,'dui, semper','2019-12-17',762172506,'PG-18',24),(59,'nec urna et','2001-12-11',2944636,'PG-18',41),(60,'ultrices,','1982-12-23',799386815,'PG-18',26);
INSERT INTO "film" (id,title,release_date,box_office,rating,studio_id) VALUES (61,'enim,','1997-06-21',129217135,'PG-13',91),(62,'quis','2018-03-28',457979214,'PG-10',71),(63,'rhoncus. Nullam','1987-10-15',94897431,'PG-18',31),(64,'cursus','1998-03-21',928311527,'PG-10',58),(65,'Proin non','1991-11-12',732193264,'PG-18',74),(66,'aliquet molestie tellus.','1994-03-21',643661491,'PG-10',59),(67,'congue.','1999-10-15',44395676,'PG-13',85),(68,'placerat','1989-04-13',567814092,'PG-13',3),(69,'mi fringilla','2014-12-24',104804872,'PG-13',4),(70,'a','2016-02-29',715406135,'PG-18',59);
INSERT INTO "film" (id,title,release_date,box_office,rating,studio_id) VALUES (71,'risus a ultricies','1996-06-05',196959676,'PG-13',58),(72,'porttitor eros nec','1998-07-22',326682570,'PG-18',56),(73,'semper','2002-11-29',921709640,'PG-13',17),(74,'Sed eu','2008-05-24',35887501,'PG-10',46),(75,'vulputate','1999-06-02',871788798,'PG-10',16),(76,'velit justo nec','2000-06-09',784784814,'PG-13',12),(77,'tortor,','2009-06-24',576810471,'PG-10',68),(78,'augue scelerisque','1991-11-01',133837393,'PG-18',32),(79,'elementum, dui quis','2020-10-02',231855670,'PG-10',75),(80,'erat,','2005-03-03',524158112,'PG-13',85);
INSERT INTO "film" (id,title,release_date,box_office,rating,studio_id) VALUES (81,'augue id ante','1992-03-02',777713730,'PG-13',66),(82,'amet orci.','1989-03-16',516695990,'PG-18',60),(83,'hendrerit consectetuer, cursus','1989-01-15',338785446,'PG-13',52),(84,'hymenaeos. Mauris ut','2006-11-09',629403965,'PG-13',63),(85,'eu neque','2005-10-16',487766417,'PG-18',91),(86,'eleifend','1992-03-24',266038940,'PG-13',1),(87,'adipiscing','2021-11-25',105252793,'PG-10',58),(88,'tempus risus. Donec','1990-10-25',312369827,'PG-13',89),(89,'nunc risus','2019-02-17',806180205,'PG-10',36),(90,'Mauris blandit enim','2003-09-07',575899040,'PG-18',87);
INSERT INTO "film" (id,title,release_date,box_office,rating,studio_id) VALUES (91,'et','2010-04-21',992804626,'PG-18',7),(92,'mauris sagittis placerat.','1990-01-03',986135887,'PG-18',68),(93,'Nunc','2009-03-10',877045287,'PG-10',41),(94,'Mauris ut','1997-12-21',265599857,'PG-13',50),(95,'Praesent eu','2014-07-30',908354417,'PG-10',17),(96,'dictum mi, ac','1985-07-14',505289327,'PG-10',59),(97,'Aliquam auctor, velit','2006-11-21',938218378,'PG-18',41),(98,'lorem vitae odio','2011-12-22',966232845,'PG-10',52),(99,'turpis','2004-06-27',726336,'PG-18',68),(100,'ligula eu','2008-08-09',889990627,'PG-18',15);
ALTER SEQUENCE film_id_seq RESTART WITH 101;


-- Тестовые данные для таблицы связи фильмов и актеров
INSERT INTO "film_actor" (id,film_id,actor_ID) VALUES (1,45,30),(2,8,96),(3,65,10),(4,49,57),(5,41,23),(6,53,60),(7,19,58),(8,90,10),(9,84,35),(10,76,58);
INSERT INTO "film_actor" (id,film_id,actor_ID) VALUES (11,51,37),(12,6,47),(13,45,35),(14,41,92),(15,56,76),(16,77,11),(17,11,72),(18,27,91),(19,67,68),(20,89,40);
INSERT INTO "film_actor" (id,film_id,actor_ID) VALUES (21,51,75),(22,85,32),(23,80,76),(24,57,7),(25,5,30),(26,6,5),(27,81,6),(28,48,34),(29,4,49),(30,85,14);
INSERT INTO "film_actor" (id,film_id,actor_ID) VALUES (31,29,8),(32,72,49),(33,88,91),(34,6,16),(35,86,59),(36,95,38),(37,62,85),(38,17,48),(39,100,64),(40,90,21);
INSERT INTO "film_actor" (id,film_id,actor_ID) VALUES (41,68,35),(42,34,72),(43,50,12),(44,76,23),(45,2,84),(46,44,60),(47,17,19),(48,44,92),(49,63,30),(50,85,24);
INSERT INTO "film_actor" (id,film_id,actor_ID) VALUES (51,7,26),(52,87,20),(53,25,94),(54,75,32),(55,15,85),(56,90,10),(57,2,75),(58,42,22),(59,27,4),(60,1,30);
INSERT INTO "film_actor" (id,film_id,actor_ID) VALUES (61,23,86),(62,70,67),(63,100,83),(64,33,86),(65,82,93),(66,61,53),(67,94,87),(68,17,21),(69,93,68),(70,14,21);
INSERT INTO "film_actor" (id,film_id,actor_ID) VALUES (71,62,24),(72,37,6),(73,6,33),(74,15,36),(75,20,97),(76,77,95),(77,63,36),(78,77,43),(79,74,80),(80,91,70);
INSERT INTO "film_actor" (id,film_id,actor_ID) VALUES (81,74,59),(82,100,29),(83,55,8),(84,31,51),(85,64,96),(86,34,14),(87,65,63),(88,64,97),(89,87,15),(90,2,86);
INSERT INTO "film_actor" (id,film_id,actor_ID) VALUES (91,1,27),(92,26,74),(93,4,38),(94,90,94),(95,60,91),(96,47,57),(97,100,10),(98,53,55),(99,3,90),(100,14,87);
ALTER SEQUENCE film_actor_id_seq RESTART WITH 101;


-- Тестовые данные для таблицы связи фильмов и режиссеров
INSERT INTO "film_producer" (id,film_id,producer_id) VALUES (1,80,95),(2,6,4),(3,31,39),(4,37,35),(5,44,48),(6,79,48),(7,21,45),(8,40,87),(9,69,30),(10,64,44);
INSERT INTO "film_producer" (id,film_id,producer_id) VALUES (11,53,36),(12,78,10),(13,46,21),(14,10,44),(15,7,95),(16,36,86),(17,99,47),(18,53,96),(19,27,26),(20,100,3);
INSERT INTO "film_producer" (id,film_id,producer_id) VALUES (21,14,27),(22,56,45),(23,15,2),(24,34,60),(25,89,61),(26,18,31),(27,96,4),(28,2,78),(29,25,89),(30,9,85);
INSERT INTO "film_producer" (id,film_id,producer_id) VALUES (31,46,60),(32,56,81),(33,53,20),(34,7,51),(35,70,78),(36,97,63),(37,12,57),(38,48,30),(39,54,44),(40,81,93);
INSERT INTO "film_producer" (id,film_id,producer_id) VALUES (41,35,23),(42,73,100),(43,72,66),(44,20,82),(45,92,96),(46,71,53),(47,80,30),(48,11,54),(49,32,70),(50,64,21);
INSERT INTO "film_producer" (id,film_id,producer_id) VALUES (51,79,15),(52,70,13),(53,60,24),(54,29,68),(55,1,59),(56,33,70),(57,60,15),(58,40,19),(59,59,64),(60,69,26);
INSERT INTO "film_producer" (id,film_id,producer_id) VALUES (61,12,78),(62,26,32),(63,26,82),(64,98,51),(65,30,58),(66,33,87),(67,80,75),(68,53,41),(69,93,4),(70,69,63);
INSERT INTO "film_producer" (id,film_id,producer_id) VALUES (71,20,78),(72,53,7),(73,79,70),(74,44,86),(75,62,72),(76,81,74),(77,6,48),(78,12,85),(79,12,93),(80,73,57);
INSERT INTO "film_producer" (id,film_id,producer_id) VALUES (81,86,68),(82,99,42),(83,44,95),(84,76,20),(85,63,57),(86,40,46),(87,51,98),(88,13,9),(89,16,9),(90,54,85);
INSERT INTO "film_producer" (id,film_id,producer_id) VALUES (91,27,74),(92,23,2),(93,73,85),(94,8,75),(95,36,67),(96,42,17),(97,65,8),(98,70,6),(99,6,35),(100,82,47);
ALTER SEQUENCE film_producer_id_seq RESTART WITH 101;


-- Проверка на ограничение по дате
-- INSERT INTO "film" (title,release_date,box_office,rating,studio_id) VALUES ('interdum ligula eu','1700-07-25',430503320,'PG-13',65);

-- Проверка на ограничение: один фильм с уникальным названием в год
-- INSERT INTO "film" (title,release_date,box_office,rating,studio_id) VALUES ('interdum ligula eu','2015-07-25',430503320,'PG-13',65);