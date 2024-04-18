#!/bin/bash

help_msg=$(cat <<EOF
Usage:

listimages.sh [OPTIONS]

Options:
  -r <registry>
  -o <operator version>
  -h 
EOF
)

: "${REGISTRY:=registry.suse.com}"
: "${OPERATOR:=$(git describe --abbrev=0 | sed "s|v\(.*\)|\1|g")}"

while getopts "hr:o:" flag; do
 case $flag in
   h)
     echo "${help_msg}"
     exit 0
   ;;
   r)
     REGISTRY=${OPTARG}
   ;;
   o)
     OPERATOR=$OPTARG
   ;;
   \?)
     >&2 echo "Invalid options"
     echo "${help_msg}"
     exit 1
   ;;
 esac
done

echo "${REGISTRY}/rancher/elemental-operator:${OPERATOR}"
echo "${REGISTRY}/rancher/seedimage-builder:${OPERATOR}"

echo "${REGISTRY}/rancher/elemental-channel:${OPERATOR}"
echo "${REGISTRY}/rancher/elemental-rt-channel:${OPERATOR}"

# Assume if not registry.suse.com it should be some OBS path, hence
# charts require a different repository than containers
if [ ${REGISTRY} != "registry.suse.com" ]; then
  REGISTRY="$(dirname ${REGISTRY})/charts"
fi

echo "${REGISTRY}/rancher/elemental-operator-crds-chart:${OPERATOR}"
echo "${REGISTRY}/rancher/elemental-operator-chart:${OPERATOR}"

