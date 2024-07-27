<h1 align='center'>
  Welcome! Go Oportunities-api README
</h1>

<p align='center'>
  This documentation describes the features of the Go Oportunities-api project, such as the project structure, requirements, installation and use.
</p>

<div align="center">

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![MySQL](https://img.shields.io/badge/mysql-4479A1.svg?style=for-the-badge&logo=mysql&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)

</div>

## About Go Oportunities

### Description

- Go Oportunities it's a personal project i'm doing to solidify my skills with golang, it's a restful api with a idea of managing jo opportunities, with a good description of the vacancy.

### Requirements 

#### To run go oportunities api on your machine, it must meet the following requirements:


- [GOLANG 1.22.5](https://go.dev/dl/)
- [MySQL 8](https://www.mysql.com/downloads/)
- [Docker](https://www.docker.com/) (optional)

### How to install

- First, clone this repository on your machine:

```
git clone https://github.com/CLucasrodrigues22/gooportunities.git
```

- Open a new terminal windows at the root of the project and create an .env based on the .env.example:

```
cp .env .env.example
```

- Specify th MySQL credentials, change the http port, through the .env created:

```
APP_PORT=8000

DB_HOST=yourhost
DB_PORT=yourport
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=yourdbname
```

- Run the commando to install golang dependencies used:

```
go mod tidy
```



- To run on a local machine, just use the make command (command specified in the ./makefile)

```
make
```
