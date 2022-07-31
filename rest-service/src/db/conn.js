const { MongoClient } = require('mongodb');
const connectionString = process.env.DB_CONNECT_URI;
const client = new MongoClient(connectionString, {
  useNewUrlParser: true,
  useUnifiedTopology: true,
});

let dbConnection;

module.exports = {
  connectToServer: function (callback) {
    client.connect(function (err, db) {
      if (err || !db) {
        return callback(err);
      }

      dbConnection = db.db(process.env.DB_NAME);
      console.log(`Successfully connected to ${process.env.DB_NAME} database.`);

      return callback();
    });
  },

  getDb: function () {
    return dbConnection;
  },
};