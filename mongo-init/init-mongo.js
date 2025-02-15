db = db.getSiblingDB('admin');
db.createUser({
    user: "root",
    pwd: "password",
    roles: [{ role: "root", db: "admin" }]
});

db = db.getSiblingDB('banking');
db.createUser({
    user: "user",
    pwd: "password",
    roles: [{ role: "readWrite", db: "banking" }]
});
