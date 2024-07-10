// init.js

db.createUser({
    user: "hung",
    pwd: "123456",
    roles: [
        {
            role: "dbAdmin",
            db: "kafkadb",
        },
        {
            role: "readWrite",
            db: "kafkadb",
        },
    ],
});
