#!/bin/bash

for file in `cat gofile.txt`
do
 sed -i 's#github.com/gophercloud/gophercloud#github.com/DashuOps/gophercloud#g' $file
done