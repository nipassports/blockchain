#!/bin/bash

nvm use 8
npm install

node enrollAdmin.js
node registeruser.js

node query.js

exit 0