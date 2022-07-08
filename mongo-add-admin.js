// mongo admin /usr/local/mongodb/mongo-add-user.js
db.createUser(
  {
    user: "super_admin",
    pwd: "crayfish233",
    roles: ["root"]
  }
)