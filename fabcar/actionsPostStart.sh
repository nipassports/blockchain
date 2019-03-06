#!/bin/sh
cd javascript

npm install
npm build

node enrollAdmin
node registerUser
