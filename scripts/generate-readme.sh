#bin/bash

if ! command -v yq &> /dev/null
then
    echo "yq package could not be found. You can install it by running: brew install yq "
    exit
fi

DESTINATION=${PWD}/documentation/app-docs/docs/cli

rm -fr "$DESTINATION"
mkdir "$DESTINATION"
echo "{\"label\": \"CLI\",\"position\": 2}
" >> "$DESTINATION/_category_.json"

yq '.tasks[] | key' Taskfile.yml | while read -r commandName; do
  title=(${commandName//:/ })

  fileName=(${commandName//:/-})

  filePath="${DESTINATION}/${fileName}.md"

  {
    echo "# $(tr '[:lower:]' '[:upper:]' <<< "${title:0:1}")${title:1}" "$(tr '[:lower:]' '[:upper:]' <<< "${title[1]:0:1}")${title[1]:1}"

    description=$(yq ".tasks.\"${commandName}\".summary" Taskfile.yml);


    if [ "$description" != "null" ]; then
      echo "\n**Description:**  "
      echo $description
    fi

    echo "\n**How to execute:**  "
    echo "\`\`\`shell\n$ task $commandName\n\`\`\`"

    envsDescription=$(yq ".tasks.\"${commandName}\".env" Taskfile.yml;)

    if [ "$envsDescription" != "null" ]; then
      echo "\n**Environmental variables:**"
      echo "\`\`\`"
      echo "$envsDescription"
      echo "\`\`\`"
    fi

  } >>"$filePath"
done
