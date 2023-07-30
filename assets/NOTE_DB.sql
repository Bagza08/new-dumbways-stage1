-- sintax untuk menggunakan psql , masukan nama user pg admin kita, lalu masukan nama db yang akan kita gunakan
psql -U postgres -d db_personal_web


-- sintax tampilkan seluruh table
select * from tb_user;

-- sintax untuk menginput data ke dalam table
INSERT INTO tb_project (name,description,image) 
	VALUES 	
			('jaya','HALLOWORLD','COBA.png'),
			('wiguna','HALLOWORLD','COBA.png');
			

INSERT INTO tb_user (name,email,psw)
    VALUES
        ('Jaya','jaya@gmail.com','12345');

INSERT INTO tb_project (name,start_date,end_date,description,technologies,image,author_id)
VALUES
        ('project 1','01-01-2023','02-01-2023','ini adalah project pertama horee','{"a","i","u","e"}','contoh.png','1');

INSERT INTO tb_project (name,start_date,end_date,description,technologies,image,author_id)
VALUES
        ('project 2','01-01-2023','02-01-2023','ini adalah project kedua huhu','{"a","","","e"}','contoh2.png','100');

-- sintax untuk meng update data 
UPDATE tb_project SET name='project 1 baru' WHERE id=1;

-- sintax untuk delete data 
DELETE FROM tb_user WHERE id=2;

'{"fa-brands fa-react fa-2xl","fa-brands fa-node fa-2xl","fa-brands fa-bootstrap fa-2xl","fa-brands fa-laravel fa-2xl"}'

'{"fa-brands fa-react fa-2xl","","fa-brands fa-bootstrap fa-2xl",""}'

'{"Front End Developer","Back End Developer","Fullstack Developer"}'

'{"2021","2022","2023"}'

'{"Digital Marketing","Marketing","Fullstack Developer"}'

'{"2010","2011","2012"}'