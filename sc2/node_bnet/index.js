var BnetStrategy = require('passport-bnet').Strategy;
var express = require('express');
var passport = require('passport');
var BNET_ID = 'dtwpk6g4rjjjqcypnjdtms69x5jjpmm4';
var BNET_SECRET = 'QFkkRBDtchj3ZTsNAKAbZJbT5nM5gyBy';

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

app.get('/login',function(){
  console.log(".login start")
  passport.authenticate('bnet');
});

app.get('/success',function(req, res){
  console.log(req,res);
  res.redirect('/#/profile');
});

var server = app.listen(9002, function() {
  console.log('Listening on port %d', server.address().port);
});