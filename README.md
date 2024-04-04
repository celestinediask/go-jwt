# go-jwt
user signup
### stack used
golang  
jwt  
gorm  
postgresql  
gin?
### USER GUIDE
#### Linux Debian
Update your system
```
sudo apt update && sudo apt upgrade
```
Install postgresql
```
sudo apt install postgresql
```
Open psql shell
```
sudo -u postgres psql
```
In psql shell set password for postgres user if not already set
```
ALTER USER postgres WITH PASSWORD 'your_password';
```
From psql shell create database titled as "gojwt"
```
CREATE DATABASE gojwt
```
Navigate to your root of your project
```
cd ~/path/to/your/project/root
```
Create .env file at your root of your project and set appropriate values
```
cat >> .env << EOF
PORT=
DB="host=localhost user=postgres password= dbname=gojwt port=5432 sslmode=disable"
EOF
```
From root of your project run program
```
go run main.go
```
