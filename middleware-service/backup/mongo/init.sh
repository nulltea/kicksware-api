#! /bin/bash
mongorestore --uri="mongodb://root:greenJordans@mongodb:27017" --db=sneakerResaleDB  ./backup/mongo/sneakerResaleDB
