// mongo r0Website /usr/local/mongodb/mongo-add-user.js
db.createUser(
  {
    user: "admin",
    pwd: "cherilee233",
    roles: [{ role: "dbOwner", db: "r0Website" }]
  }
)
// 如果在一个空Database下执行此脚本，那么一个新的Collection是必须的。
db.createCollection("user");