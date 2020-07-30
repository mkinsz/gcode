#!/bin/sh
srcPath=".."
pkgFile="server.go"
outputPath="build"
app="gcode"
output="$outputPath/$app.exe"
src="$srcPath/$app/$pkgFile"

printf "\nBuilding: $app\n"
time go build -o $output $src
printf "\nBuilt: $app size:"
ls -lah $output | awk '{print $5}'
printf "\nDone building: $app\n\n"