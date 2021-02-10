# Graphql + GORM SQL CRUD Generator
```
$ go run github.com/david-yappeter/GormCrudGenerator init
$ go run github.com/david-yappeter/GormCrudGenerator
```
Run The Normal Command After you run Init
<br>
Init command will initialize a .yaml file (gormCrud.yaml) :
```
database:
    type:
        # Only The First One Will Be Applied
        - mysql
        - postgre
    path: ./config
    name: databaseGorm
    setting:
        path: ./logger
        name: logMode
        singularTable: true
        tablePrefix: ""
        logLevel:
            # Only The First One Will Be Applied
            - Info
            - Silent
            - Warn
            - Error
        slowThreshold: 1
service:
    from:
        path: ./gormgenerator
        name: model
        # Ignore Model (Case-Sensitive)
        ignore:
            - User
    to:
        path: ./service
        postfix: "Generated"
queryTools:
    path: ./tools
    name: dbGenerator
```
#### Path
Path is the path to the directory from where you run your command
###### Example:
```
$USER/myapp$ go run github.com/david-yappeter/GormCrudGenerator
```
This will generate path relative from 'myapp' folder

## Database
-   ##### Type 
    It will take first array of yaml as Database Connection type
    <br>
    ###### Example :
    ```
    database:
    type:
        # Only The First One Will Be Applied
        - mysql
        - postgre
    ```
    This will take 'mysql' as Connection Type
    ```
    database:
    type:
        # Only The First One Will Be Applied
        - postgre
        - mysql
    ```
    This will take 'postgre' as Connection Type
-   ##### Setting
    Contain Logger Config for [gorm.io v2](https://gorm.io/)

## Service
-   ##### From
    This is the origin path of the model that want to be generate
    <br>
    You can ignore some of the model by declaring it inside 'ignore'