const express = require('express');
// Has prefix '/users'
const userRoutes = express.Router();
// Get database connection
const dbConn = require('../../db/conn');
// bcrypt package for hashing passwords
const bcrypt = require('bcrypt');

//Number of salt rounds for bcrypt hashing
const SALT_ROUNDS = 5;

userRoutes.route('/create').post(async function ({body}, res) {

    if(!body.username || !body.password) {
        res.status(400).send("Error creating user, must have a username and password");
        return;
    }

    const db = dbConn.getDb();
    const usernameExists = await db.collection('Users').find({username: body.username}).hasNext();
    if(usernameExists) {
        res.status(409).send("Error creating user, username already exists.");
        return;
    }

    const hashedPassword = await bcrypt.hash(body.password, SALT_ROUNDS);
    const creationDate = new Date();
    const newUser = {
        username: body.username,
        password: hashedPassword,
        creationDate
    }

    db.collection('Users')
        .insertOne(newUser, function (err, result) {
            if (err) {
                res.status(400).send("Error creating user.");
            } else {
                console.log(`Added a new user with id ${result.insertedId}`);
                res.status(201).send('User created');
            }
        });
})

module.exports = userRoutes;