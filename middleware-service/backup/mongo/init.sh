#! /bin/bash
mongorestore --db=sneakerResaleDB  ./backup/mongo/sneakerResaleDB

mongo admin -u admin -p admin --eval "db.getSiblingDB('dummydb').createUser({user: 'root', pwd: 'greenJordans', roles: ['readWrite']})"
