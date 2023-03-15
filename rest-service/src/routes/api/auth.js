const express = require('express');

// Has prefix '/auth'
const authRoutes = express.Router();
// Get database connection
const dbConn = require('../../db/conn');

module.exports = authRoutes;