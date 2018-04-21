var BnetStrategy = require('passport-bnet').Strategy;
var express = require('express');
var http = require('http');
var https = require('https');
var passport = require('passport');
var BNET_ID = 'dtwpk6g4rjjjqcypnjdtms69x5jjpmm4';
var BNET_SECRET = 'QFkkRBDtchj3ZTsNAKAbZJbT5nM5gyBy';
var fs = require('fs');

// var privateKey = fs.readFileSync('./ssl/key.pem', 'utf8');
// var certificate = fs.readFileSync('./ssl/key.crt', 'utf8');
// var credentials = {
//   key: privateKey,
//   cert: certificate
// }

// Use the BnetStrategy within Passport.
passport.use(new BnetStrategy({
  clientID: BNET_ID,
  clientSecret: BNET_SECRET,
  region: "cn",
  scope: "sc2.profile",
  callbackURL: "https://sc2.darkjs.com/bnet/success",
}, function(accessToken, refreshToken, profile, done) {
  return done(null, profile);
}));

var app = express();

app.get('/login',
  passport.authenticate('bnet'));

app.get('/success',function(req, res){
  console.log(req,res);
  res.redirect('/#/profile');
});
var httpServer = http.createServer(app);
// var httpsServer = https.createServer(credentials, app);

httpServer.listen(9002);
// httpsServer.listen(9443);