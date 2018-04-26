var BnetStrategy = require('passport-bnet').Strategy;
var express = require('express');
var http = require('http');
var https = require('https');
var passport = require('passport');
var cookieParser = require('cookie-parser');
var session = require('express-session');
var BNET_ID = 'dtwpk6g4rjjjqcypnjdtms69x5jjpmm4';
var BNET_SECRET = 'QFkkRBDtchj3ZTsNAKAbZJbT5nM5gyBy';
var request = require('request');
var fs = require('fs');

// var privateKey = fs.readFileSync('./ssl/key.pem', 'utf8');
// var certificate = fs.readFileSync('./ssl/key.crt', 'utf8');
// var credentials = {
//   key: privateKey,
//   cert: certificate
// }

var app = express();
app.use(passport.initialize());
app.use(passport.session());
app.use(cookieParser());
app.use(session({ secret: 'blizzard',
  saveUninitialized: true,
  resave: true }));


passport.serializeUser(function(user, done) {
  done(null, user);
});

passport.deserializeUser(function(user, done) {
  done(null, user);
});

// Use the BnetStrategy within Passport.
passport.use(new BnetStrategy({
  clientID: BNET_ID,
  clientSecret: BNET_SECRET,
  region: "cn",
  scope: "sc2.profile",
  callbackURL: "https://sc2.darkjs.com/bnet/callback",
}, function(accessToken, refreshToken, profile, done) {
  console.log(accessToken, refreshToken, profile);
  if (accessToken != null) {
    request('http://localhost:9001/saveAccessToken?accessToken='+accessToken, function (error, response, body) {
    })
  }
  return done(null, profile);
}));


app.get('/login',
  passport.authenticate('bnet'));

app.get('/callback',
  passport.authenticate('bnet', { failureRedirect: '/' }),
  function(req, res){
    res.redirect('/callback.html');
  });

var httpServer = http.createServer(app);
// var httpsServer = https.createServer(credentials, app);

httpServer.listen(9002);
// httpsServer.listen(9443);