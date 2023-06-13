#!/bin/bash

author=""
name=""
body=""
started=0

echo "[" > achievements.json

cat AchievementsList.md | while read line;
do
    if [ -z "$line" ]; then 
        continue;
    elif [[ $line == \#\#\#\#* ]]; then 
        name=$(echo $line | sed -e "s/^\#\#\#\# //");
    elif [[ $line == \#\#* ]]; then 
        author=$(echo $line | sed -e "s/^\#\# //" -e "s/\:$//");
    else 
        body="$body\n$line";
    fi

    if [ ! -z "$name" ] && [ ! -z "$body" ]; then
        body=$(echo $body | sed -e "s/^\\\n//");
        # echo "Name: $name"
        # echo "Body: '$body'"
        # echo "Author: $author"
        json=$( jq -n -r \
                  --arg nm "$name" \
                  --arg bd "$body" \
                  --arg auth "$author" \
                  '{name: $nm, body: $bd, author: $auth}' )
        if [ "$started" == "1" ]; then 
            echo "," >> achievements.json
        fi
        echo $json >> achievements.json
        body=""
        name=""
        started=1
    fi
done

echo "]" >> achievements.json
echo "Scanning Done"

json=$(cat achievements.json | jq .) 

echo "Formatting Done"

echo "$json" > achievements.json

echo "Written to file"
