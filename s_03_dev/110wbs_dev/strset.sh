#!/bin/bash

#--- <Pattearn A> Manual Operation.----------------------------#
# npm install -D webpack webpack-cli
# npm install -g typescript
# npm install sass

# npm install @babel/core @babel/preset-env @babel/preset-react --save-dev 
# npm install webpack webpack-cli webpack-dev-server babel-loader css-loader style-loader html-webpack-plugin --save-dev 


#--- <Pattearn B> create-react-app Operation.---(5m45sec)------#

# Saveing samplecodes.
mv public ..
mv src ..
mv README.md ..

# React installation.
npm install -g create-react-app@latest
create-react-app . --typescript
npm install sass@latest

# Restore samplecodes.(If you want to see React sample page, comment out mv command.)
mv ../public/Assets ./public
# mv ../public/LandingPage ./public
mv ../public/index.html ./public
mv ../src/components ./src
mv ../src/hooks ./src
mv ../src/types ./src
mv ../src/App.js ./src
mv ../src/logo.svg ./src
mv ../src/Pager.scss ./src
mv ../src/Router.jsx ./src
mv ../README.md .
# mv ../src/Top.jsx ./src
npm i -S react-router-dom@6.4.2

# Transfer version control files.
# \cp -f ./package.json ../../s_01_src/110wbs_src/package.json

# Import MUI
npm install @mui/material @emotion/react @emotion/styled

# Import Font Awesome
npm i -S @fortawesome/fontawesome-svg-core @fortawesome/react-fontawesome @fortawesome/free-regular-svg-icons @fortawesome/free-solid-svg-icons @fortawesome/free-brands-svg-icons

# Add accesspoint of build directroy.
# npm install -g serve
npm -g install serve
npm run build

# Make a standby state.
yarn start

#--- .gitkeep removal -----------------------------------------#
# rm .gitkeep
# rm -rf ./.*
# rm *
# rm -rf .[^.] .?*
# rm -rf
# rm -rf tafgetDir/.*
# rm -rf .gitkeep
# rm .dotfile
# rm.dotfile
# rm .dot*
# rm.dot*
# rm -r .*
# rm -R .*
# rm ./.gitkeep