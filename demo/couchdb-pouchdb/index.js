var PouchDB = require('pouchdb');

var localDB = new PouchDB('mylocaldb')
var remoteDB = new PouchDB('http://admin:password@localhost:5984/kitten');

localDB.sync(remoteDB, {
  live: true
}).on('change', function (change) {
  console.log("change")
  console.log(change)
}).on('error', function (err) {
  console.error("err")
});