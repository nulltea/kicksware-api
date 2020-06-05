#! /bin/bash
mongodump --uri="mongodb://root:greenJordans@mongodb:27017" --db=sneakerResaleDB  --out=/backup/mongo
