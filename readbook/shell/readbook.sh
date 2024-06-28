#!/bin/bash

text=hello.txt
voice_folder=voice

[ ! -f $text ] && { echo "file $text does not exist"; exit 1;}

[ ! -d $voice_folder ] && { echo "folder $voice_folder does not exist"; exit 1;}

while read i
do 
	say -o ${voice_folder}/${i}.wav --data-format=alaw $i
done < $text
