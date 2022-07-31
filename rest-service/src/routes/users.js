const express = require('express');

// Has prefix '/users'
const userRoutes = express.Router();

// Get database connection
const dbConn = require('../db/conn');

userRoutes.route('/create').post(function (req, res) {
    const db = dbConn.getDb();
    const creationDate = new Date();
    const newUser = {
        username: req.body.username,
        password: req.body.password,
        creationDate
    }

    db
        .collection('Users')
        .insertOne(newUser, function (err, result) {
            if (err) {
                res.status(400).send("Error inserting matches!");
            } else {
                console.log(`Added a new match with id ${result.insertedId}`);
                res.status(201).send('User created');
            }
        });
})

module.exports = userRoutes;