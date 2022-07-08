#!/bin/bash
/usr/local/mongodb/bin/mongod --dbpath=/usr/local/mongodb/data/db --fork --logpath=/usr/local/mongodb/logs/mongod.log &&

sleep 5 &&

/usr/local/mongodb/bin/mongo r0Website /usr/local/mongodb/mongo-add-user.js &&

/usr/local/mongodb/bin/mongod -f /usr/local/mongodb/mongod.conf