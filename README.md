# mysql-sakila-sample

## install

need to install docker-compose
https://docs.docker.com/compose/install/#master-builds

## setup

```
cd .docker
docker-compose up -d
```

※少しデータを load するのに時間がかかるので、`docker logs -f sakila-mysql`でロードが終了したかをチェックすると良い  
↓ こんな感じのがログに表れていれば終了している ↓

```
INFO
CREATING DATABASE STRUCTURE
INFO
storage engine: InnoDB
INFO
LOADING departments
INFO
LOADING employees
INFO
LOADING dept_emp
INFO
LOADING dept_manager
INFO
LOADING titles
INFO
LOADING salaries
data_load_time_diff
```

## login

```
docker exec -it sakila-mysql bash
mysql -uroot -proot
```

## try

```
use employees;
show tables;
```

※表示されれば OK！
