#/bin/bash

# include demo-magic
. ./tools/demo-magic.sh -n

# Will wait max 1 seconds until user presses
PROMPT_TIMEOUT=1

TYPE_SPEED=10

# hide the evidence
clear

p "ls -al"
cat cache/ls-al-cache.txt

p "brew install pv"
wait
cat cache/install-pv-cache.txt

# Will wait until user presses enter
PROMPT_TIMEOUT=0
p "ssh 123.123.111.666"
echo password\(ENTER\):
wait
echo login success
