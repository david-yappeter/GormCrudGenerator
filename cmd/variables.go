package cmd

const defaultConfig = `
database:
    type:
        # Only The First One Will Be Applied
        - mysql
        - postgre
    path: ./config
    name: databaseGorm
    apply: true
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
        apply: true
service:
    from:
        path: ./gormgenerator
        name: model
        # Ignore Model (Case-Sensitive)
        ignore:
            - 
    to:
        path: ./service
        postfix: "Generated"
    apply: true
queryTools:
    path: ./tools
    name: dbGenerator
    apply: true
`
